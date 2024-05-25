// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package blockchain

import (
	"context"
	"time"

	"github.com/berachain/beacon-kit/mod/consensus-types/pkg/genesis"
	"github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
	engineprimitives "github.com/berachain/beacon-kit/mod/primitives-engine"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/transition"
	"golang.org/x/sync/errgroup"
)

// ProcessGenesisState processes the genesis state and initializes the beacon
// state.
func (s *Service[
	AvailabilityStoreT,
	ReadOnlyBeaconStateT,
	BlobSidecarsT,
	DepositStoreT,
]) ProcessGenesisState(
	ctx context.Context,
	genesisData *genesis.Genesis[
		*types.Deposit, *types.ExecutionPayloadHeaderDeneb,
	],
) ([]*transition.ValidatorUpdate, error) {
	return s.sp.InitializePreminedBeaconStateFromEth1(
		s.sb.StateFromContext(ctx),
		genesisData.Deposits,
		genesisData.ExecutionPayloadHeader,
		genesisData.ForkVersion,
	)
}

// ProcessStateTransition receives an incoming beacon block, it first validates
// and then processes the block.
//

func (s *Service[
	AvailabilityStoreT,
	ReadOnlyBeaconStateT,
	BlobSidecarsT,
	DepositStoreT,
]) ProcessStateTransition(
	ctx context.Context,
	blk types.BeaconBlock,
	sidecars BlobSidecarsT,
	optimisticEngine bool,
) ([]*transition.ValidatorUpdate, error) {
	// If the block is nil, exit early.
	if blk == nil || blk.IsNil() {
		return nil, ErrNilBlk
	}

	// Create a new errgroup with the provided context.
	g, gCtx := errgroup.WithContext(ctx)
	st := s.sb.StateFromContext(ctx)

	// Launch a goroutine to process the state transition.
	var valUpdates []*transition.ValidatorUpdate
	g.Go(func() error {
		var (
			err       error
			startTime = time.Now()
		)
		defer s.metrics.measureStateTransitionDuration(startTime)
		valUpdates, err = s.sp.Transition(
			// We set `OptimisticEngine` when this is called during
			// FinalizeBlock. We want to assume the payload is valid. If it
			// ends up not being valid later, the node will simply AppHash,
			// which is completely fine. This means we were syncing from a
			// bad peer, and we would likely AppHash anyways.
			//
			// TODO: Figure out why SkipPayloadIfExists being `true`
			// causes nodes to create gaps in their chain.
			&transition.Context{
				Context:          gCtx,
				OptimisticEngine: optimisticEngine,
			},
			st,
			blk,
		)
		return err
	})

	// Launch a goroutine to process the blob sidecars.
	g.Go(func() error {
		startTime := time.Now()
		defer s.metrics.measureBlobProcessingDuration(startTime)
		return s.bp.ProcessBlobs(
			blk.GetSlot(),
			s.sb.AvailabilityStore(ctx),
			sidecars,
		)
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	// If the blobs needed to process the block are not available, we
	// return an error. It is safe to use the slot off of the beacon block
	// since it has been verified as correct already.
	if !s.sb.AvailabilityStore(ctx).IsDataAvailable(
		ctx, blk.GetSlot(), blk.GetBody(),
	) {
		return nil, ErrDataNotAvailable
	}

	// No matter what happens we always want to forkchoice at the end of post
	// block processing.
	defer func() {
		go s.sendPostBlockFCU(ctx, st, blk)
	}()

	//
	//
	//
	//
	//
	// TODO: EVERYTHING BELOW THIS LINE SHOULD NOT PART OF THE
	//  MAIN BLOCK PROCESSING THREAD.
	//
	//
	//
	//
	//
	//

	// Prune deposits.
	// TODO: This should be moved into a go-routine in the background.
	// Watching for logs should be completely decoupled as well.
	idx, err := st.GetEth1DepositIndex()
	if err != nil {
		return nil, err
	}

	// TODO: pruner shouldn't be in main block processing thread.
	if err = s.PruneDepositEvents(ctx, idx); err != nil {
		return nil, err
	}

	var lph engineprimitives.ExecutionPayloadHeader
	lph, err = st.GetLatestExecutionPayloadHeader()
	if err != nil {
		return nil, err
	}

	// Process the logs from the previous blocks execution payload.
	// TODO: This should be moved out of the main block processing flow.
	// TODO: eth1FollowDistance should be done actually proper
	eth1FollowDistance := math.U64(1)
	if err = s.retrieveDepositsFromBlock(
		ctx, lph.GetNumber()-eth1FollowDistance,
	); err != nil {
		s.logger.Error("failed to process logs", "error", err)
		return nil, err
	}

	return valUpdates, nil
}

// VerifyPayload validates the execution payload on the block.
func (s *Service[
	AvailabilityStoreT,
	ReadOnlyBeaconStateT,
	BlobSidecarsT,
	DepositStoreT,
]) VerifyPayloadOnBlk(
	ctx context.Context,
	blk types.BeaconBlock,
) error {
	if blk == nil || blk.IsNil() {
		return ErrNilBlk
	}

	// We notify the engine of the new payload.
	var (
		parentBeaconBlockRoot = blk.GetParentBlockRoot()
		body                  = blk.GetBody()
		payload               = body.GetExecutionPayload()
	)

	if err := s.ee.VerifyAndNotifyNewPayload(
		ctx,
		engineprimitives.BuildNewPayloadRequest(
			payload,
			body.GetBlobKzgCommitments().ToVersionedHashes(),
			&parentBeaconBlockRoot,
			false,
			// We do not want to optimistically assume truth here, since
			// this is being called in process proposal.
			false,
		),
	); err != nil {
		return err
	}

	s.logger.Info(
		"successfully verified execution payload 💸",
		"payload-block-number", payload.GetNumber(),
		"num-txs", len(payload.GetTransactions()),
	)
	return nil
}
