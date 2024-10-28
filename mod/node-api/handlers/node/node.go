// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package node

import (
	"context"
	"fmt"

	nodetypes "github.com/berachain/beacon-kit/mod/node-api/handlers/node/types"
	"github.com/berachain/beacon-kit/mod/node-api/handlers/types"
	cmtclient "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
)

// Syncing returns the syncing status of the beacon node.
func (h *Handler[ContextT]) Syncing(_ ContextT) (any, error) {
	// Create a new RPC client
	rpcClient, err := cmtclient.New("tcp://localhost:26657")
	if err != nil {
		return nil, fmt.Errorf("failed to create RPC client: %w", err)
	}

	// Create client context with the RPC client
	clientCtx := client.Context{}
	clientCtx = clientCtx.WithClient(rpcClient)

	// Query CometBFT status
	status, err := cmtservice.GetNodeStatus(context.Background(), clientCtx)
	if err != nil {
		return nil, fmt.Errorf("err in getting node status %w", err)
	}
	response := nodetypes.SyncingData{}
	response.HeadSlot = status.SyncInfo.LatestBlockHeight

	// Calculate sync distance
	if status.SyncInfo.LatestBlockHeight < status.SyncInfo.EarliestBlockHeight {
		syncDistance := status.SyncInfo.EarliestBlockHeight -
			status.SyncInfo.LatestBlockHeight
		response.SyncDistance = syncDistance
		response.IsSyncing = status.SyncInfo.CatchingUp
	} else {
		response.SyncDistance = 0
		response.IsSyncing = false
	}

	// Keep existing values for these fields
	response.IsOptimistic = true
	response.ELOffline = false

	return types.Wrap(&response), nil
}

// Version returns the version of the beacon node.
func (h *Handler[ContextT]) Version(_ ContextT) (any, error) {
	version, err := h.backend.GetNodeVersion()
	if err != nil {
		return nil, err
	}
	return types.Wrap(nodetypes.VersionData{
		Version: version,
	}), nil
}
