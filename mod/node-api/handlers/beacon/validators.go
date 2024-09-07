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

package beacon

import (
	"strconv"
	"strings"

	consensustypes "github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
	"github.com/berachain/beacon-kit/mod/errors"
	beacontypes "github.com/berachain/beacon-kit/mod/node-api/handlers/beacon/types"
	"github.com/berachain/beacon-kit/mod/node-api/handlers/types"
	"github.com/berachain/beacon-kit/mod/node-api/handlers/utils"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/crypto"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/encoding/json"
)

func (h *Handler[_, ContextT, _, _]) GetStateValidators(
	c ContextT,
) (any, error) {
	req, err := utils.BindAndValidate[beacontypes.GetStateValidatorsRequest](
		c, h.Logger(),
	)
	if err != nil {
		return nil, err
	}
	// TODO: remove this once status filter is implemented.
	if len(req.Statuses) > 0 {
		return nil, types.ErrNotImplemented
	}
	slot, err := utils.SlotFromStateID(req.StateID, h.backend)
	if err != nil {
		return nil, err
	}
	validators, err := h.backend.ValidatorsByIDs(
		slot,
		req.IDs,
		req.Statuses,
	)
	if err != nil {
		return nil, err
	}
	if len(validators) == 0 {
		return nil, types.ErrNotFound
	}
	return beacontypes.ValidatorResponse{
		ExecutionOptimistic: false, // stubbed
		Finalized:           false, // stubbed
		Data:                validators,
	}, nil
}

// CustomValidator to hold the validator data for BeaconAPIs.
//
//nolint:lll
type CustomValidator struct {
	Pubkey                     crypto.BLSPubkey                     `json:"pubkey"`
	WithdrawalCredentials      consensustypes.WithdrawalCredentials `json:"withdrawal_credentials"`
	EffectiveBalance           string                               `json:"effective_balance"`
	Slashed                    bool                                 `json:"slashed"`
	ActivationEligibilityEpoch string                               `json:"activation_eligibility_epoch"`
	ActivationEpoch            string                               `json:"activation_epoch"`
	ExitEpoch                  string                               `json:"exit_epoch"`
	WithdrawableEpoch          string                               `json:"withdrawable_epoch"`
}

func (h *Handler[_, ContextT, _, _]) PostStateValidators(
	c ContextT,
) (any, error) {
	req, err := utils.BindAndValidate[beacontypes.PostStateValidatorsRequest](
		c, h.Logger(),
	)
	if err != nil {
		return nil, err
	}
	// TODO: remove this once status filter is implemented.
	if len(req.Statuses) > 0 {
		return nil, types.ErrNotImplemented
	}
	slot, err := utils.SlotFromStateID(req.StateID, h.backend)
	if err != nil {
		return nil, err
	}
	validators, err := h.backend.ValidatorsByIDs(
		slot,
		req.IDs,
		req.Statuses,
	)
	if err != nil {
		return nil, err
	}

	convertedValidators := make(
		[]beacontypes.ValidatorData[CustomValidator],
		len(validators),
	)
	for i, validator := range validators {
		convertedValidator, convErr := convertValidator(validator)
		if convErr != nil {
			return nil, errors.Wrap(convErr, "failed to convert validator")
		}
		convertedValidators[i] = convertedValidator
	}
	return beacontypes.ValidatorResponse{
		ExecutionOptimistic: false, // stubbed
		Finalized:           false, // stubbed
		Data:                convertedValidators,
	}, nil
}

func convertValidator[ValidatorT any](
	validator *beacontypes.ValidatorData[ValidatorT],
) (beacontypes.ValidatorData[CustomValidator], error) {
	// Convert the original validator to JSON
	jsonData, err := json.Marshal(validator.Validator)
	if err != nil {
		return beacontypes.ValidatorData[CustomValidator]{},
			errors.Wrap(err, "failed to marshal validator")
	}

	// Unmarshal into our ConvertedValidator struct
	var customValidator CustomValidator
	if err = json.Unmarshal(jsonData, &customValidator); err != nil {
		return beacontypes.ValidatorData[CustomValidator]{},
			errors.Wrap(err, "failed to unmarshal validator")
	}

	if err = convertHexFields(&customValidator); err != nil {
		return beacontypes.ValidatorData[CustomValidator]{},
			errors.Wrap(err, "failed to convert hex fields")
	}

	return beacontypes.ValidatorData[CustomValidator]{
		ValidatorBalanceData: validator.ValidatorBalanceData,
		Status:               validator.Status,
		Validator:            customValidator,
	}, nil
}

func convertHexFields(v *CustomValidator) error {
	var err error
	v.EffectiveBalance, err = hexToDecimalString(v.EffectiveBalance)
	if err != nil {
		return errors.Wrap(err, "failed to convert effective balance")
	}
	v.ActivationEligibilityEpoch, err = hexToDecimalString(
		v.ActivationEligibilityEpoch,
	)
	if err != nil {
		return errors.Wrap(err, "failed to convert activation eligibility epoch")
	}
	v.ActivationEpoch, err = hexToDecimalString(v.ActivationEpoch)
	if err != nil {
		return errors.Wrap(err, "failed to convert activation epoch")
	}
	v.ExitEpoch, err = hexToDecimalString(v.ExitEpoch)
	if err != nil {
		return errors.Wrap(err, "failed to convert exit epoch")
	}
	v.WithdrawableEpoch, err = hexToDecimalString(v.WithdrawableEpoch)
	if err != nil {
		return errors.Wrap(err, "failed to convert withdrawable epoch")
	}
	return nil
}

func hexToDecimalString(hexStr string) (string, error) {
	hexStr = strings.TrimPrefix(hexStr, "0x")

	// Convert hex string to uint64
	value, err := strconv.ParseUint(hexStr, 16, 64)
	if err != nil {
		return "", err
	}

	return strconv.FormatUint(value, 10), nil
}

func (h *Handler[_, ContextT, _, _]) GetStateValidator(
	c ContextT,
) (any, error) {
	req, err := utils.BindAndValidate[beacontypes.GetStateValidatorRequest](
		c, h.Logger(),
	)
	if err != nil {
		return nil, err
	}
	slot, err := utils.SlotFromStateID(req.StateID, h.backend)
	if err != nil {
		return nil, err
	}
	validator, err := h.backend.ValidatorByID(
		slot,
		req.ValidatorID,
	)
	if err != nil {
		return nil, err
	}
	return validator, nil
}

func (h *Handler[_, ContextT, _, _]) GetStateValidatorBalances(
	c ContextT,
) (any, error) {
	req, err := utils.BindAndValidate[beacontypes.GetValidatorBalancesRequest](
		c, h.Logger(),
	)
	if err != nil {
		return nil, err
	}
	slot, err := utils.SlotFromStateID(req.StateID, h.backend)
	if err != nil {
		return nil, err
	}
	balances, err := h.backend.ValidatorBalancesByIDs(
		slot,
		req.IDs,
	)
	if err != nil {
		return nil, err
	}
	return beacontypes.ValidatorResponse{
		ExecutionOptimistic: false, // stubbed
		Finalized:           false, // stubbed
		Data:                balances,
	}, nil
}

func (h *Handler[_, ContextT, _, _]) PostStateValidatorBalances(
	c ContextT,
) (any, error) {
	var ids []string
	if err := c.Bind(&ids); err != nil {
		return nil, errors.Wrapf(errors.New("err in func bind"), "err %v", err)
	}

	req := beacontypes.PostValidatorBalancesRequest{
		StateIDRequest: types.StateIDRequest{StateID: "head"},
		IDs:            ids,
	}

	if err := c.Validate(&req); err != nil {
		return nil, errors.Wrapf(errors.New("err in validate"), "err %v", err)
	}

	slot, err := utils.SlotFromStateID(req.StateID, h.backend)
	if err != nil {
		return nil, errors.Wrapf(
			errors.New("err in getting slot "),
			" slot req err %v %v %v", req, slot, err)
	}

	h.Logger().Info("PostStateValidatorBalances", "slot", slot, "req", req)

	balances, err := h.backend.ValidatorBalancesByIDs(slot, req.IDs)
	if err != nil {
		return nil, errors.Wrapf(errors.New("err in backend "), "err %v", err)
	}
	return beacontypes.ValidatorResponse{
		ExecutionOptimistic: false, // stubbed
		Finalized:           false, // stubbed
		Data:                balances,
	}, nil
}
