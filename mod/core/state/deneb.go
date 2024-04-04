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

package state

import (
	"github.com/berachain/beacon-kit/mod/core/types"
	types0 "github.com/berachain/beacon-kit/mod/execution/types"
	"github.com/berachain/beacon-kit/mod/primitives"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// DefaultBeaconStateDeneb returns a default BeaconStateDeneb.
//
// TODO: take in BeaconConfig params to determine the
// default length of the arrays, which we are currently
// and INCORRECTLY setting to 0.
func DefaultBeaconStateDeneb() *BeaconStateDeneb {
	//nolint:gomnd // default allocs.
	return &BeaconStateDeneb{
		GenesisValidatorsRoot: primitives.Root{},

		Slot: 0,
		LatestBlockHeader: &primitives.BeaconBlockHeader{
			Slot:          0,
			ProposerIndex: 0,
			ParentRoot:    primitives.Root{},
			StateRoot:     primitives.Root{},
			BodyRoot:      primitives.Root{},
		},
		BlockRoots:             make([][32]byte, 1),
		StateRoots:             make([][32]byte, 1),
		LatestExecutionPayload: DefaultGenesisExecutionPayload(),
		Eth1BlockHash: common.HexToHash(
			"0xa63c365d92faa4de2a64a80ed4759c3e9dfa939065c10af08d2d8d017a29f5f4",
		),
		Eth1DepositIndex: 0,
		Validators:       make([]*types.Validator, 0),
		Balances:         make([]uint64, 0),
		RandaoMixes:      make([][32]byte, 8),
		Slashings:        make([]uint64, 1),
		TotalSlashing:    0,
	}
}

func DefaultGenesisExecutionPayload() *types0.ExecutableDataDeneb {
	return &types0.ExecutableDataDeneb{
		ParentHash:    primitives.ExecutionHash{},
		FeeRecipient:  primitives.ExecutionAddress{},
		StateRoot:     primitives.ExecutionHash{},
		ReceiptsRoot:  primitives.ExecutionHash{},
		LogsBloom:     make([]byte, 256), //nolint:gomnd // default allocs.
		Random:        primitives.ExecutionHash{},
		Number:        0,
		GasLimit:      0,
		GasUsed:       0,
		Timestamp:     0,
		ExtraData:     make([]byte, 32), //nolint:gomnd // default allocs.
		BaseFeePerGas: make([]byte, 32), //nolint:gomnd // default allocs.
		Transactions:  [][]byte{},
		Withdrawals:   []*primitives.Withdrawal{},
		BlobGasUsed:   0,
		ExcessBlobGas: 0,
	}
}

// TODO: should we replace ? in ssz-size with values to ensure we are hash tree
// rooting correctly?
//
//go:generate go run github.com/fjl/gencodec -type BeaconStateDeneb -field-override beaconStateDenebJSONMarshaling -out deneb.json.go
//nolint:lll // various json tags.
type BeaconStateDeneb struct {
	// Versioning
	//
	//nolint:lll
	GenesisValidatorsRoot primitives.Root `json:"genesisValidatorsRoot" ssz-size:"32"`
	Slot                  primitives.Slot `json:"slot"`

	// History
	LatestBlockHeader *primitives.BeaconBlockHeader `json:"latestBlockHeader"`
	BlockRoots        [][32]byte                    `json:"blockRoots"        ssz-size:"?,32" ssz-max:"8192"`
	StateRoots        [][32]byte                    `json:"stateRoots"        ssz-size:"?,32" ssz-max:"8192"`

	// Eth1
	LatestExecutionPayload *types0.ExecutableDataDeneb `json:"latestExecutionPayload"`
	Eth1BlockHash          primitives.ExecutionHash    `json:"eth1BlockHash"          ssz-size:"32"`
	Eth1DepositIndex       uint64                      `json:"eth1DepositIndex"`

	// Registry
	Validators []*types.Validator `json:"validators" ssz-max:"1099511627776"`
	Balances   []uint64           `json:"balances"   ssz-max:"1099511627776"`

	// Randomness
	RandaoMixes [][32]byte `json:"randaoMixes" ssz-size:"?,32" ssz-max:"65536"`

	// Withdrawals
	NextWithdrawalIndex          uint64 `json:"nextWithdrawalIndex"`
	NextWithdrawalValidatorIndex uint64 `json:"nextWithdrawalValidatorIndex"`

	// Slashing
	Slashings     []uint64        `json:"slashings"     ssz-max:"1099511627776"`
	TotalSlashing primitives.Gwei `json:"totalSlashing"`
}

// String returns a string representation of BeaconStateDeneb.
func (b *BeaconStateDeneb) String() string {
	return "TODO: BeaconStateDeneb"
}

// beaconStateDenebJSONMarshaling is a type used to marshal/unmarshal
// BeaconStateDeneb.
type beaconStateDenebJSONMarshaling struct {
	GenesisValidatorsRoot  hexutil.Bytes
	LatestExecutionPayload *types0.ExecutableDataDeneb
	BlockRoots             []primitives.Root
	StateRoots             []primitives.Root
	RandaoMixes            []primitives.Root
}
