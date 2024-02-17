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

	"github.com/ethereum/go-ethereum/common"
	"github.com/itsdevbear/bolaris/beacon/execution"
	"github.com/itsdevbear/bolaris/types/consensus/interfaces"
)

// FinalizeBeaconBlock finalizes a beacon block by processing the logs, deposits,
// and voluntary exits. It also updates the finalized and safe eth1 block hashes
// on the beacon state.
func (s *Service) FinalizeBeaconBlock(
	ctx context.Context,
	blk interfaces.ReadOnlyBeaconKitBlock,
) error {
	payload, err := blk.ExecutionPayload()
	if err != nil {
		return err
	}

	// Process logs, including deposit and withdrawal events.
	err = s.en.ProcessLogs(ctx, payload.GetBlockNumber())
	if err != nil {
		return err
	}
	// TODO: PROCESS VOLUNTARY EXITS HERE

	if err = s.en.NotifyForkchoiceUpdate(ctx, &execution.FCUConfig{
		HeadEth1Hash: common.Hash(payload.GetBlockHash()),
		Attributes:   nil,
	}); err != nil {
		s.Logger().Error("failed to notify forkchoice update", "error", err)
	}

	state := s.BeaconState(ctx)
	eth1BlockHash := common.Hash(payload.GetBlockHash())
	state := s.BeaconState(ctx)
	state.SetFinalizedEth1BlockHash(eth1BlockHash)
	state.SetSafeEth1BlockHash(eth1BlockHash)
	state.SetLastValidHead(eth1BlockHash)

	return nil
}
