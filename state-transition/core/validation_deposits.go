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

package core

import (
	"fmt"

	"github.com/berachain/beacon-kit/config/spec"
	ctypes "github.com/berachain/beacon-kit/consensus-types/types"
	"github.com/berachain/beacon-kit/errors"
	"github.com/berachain/beacon-kit/primitives/math"
)

func (sp *StateProcessor[
	_, BeaconStateT, _, _,
]) validateGenesisDeposits(
	st BeaconStateT,
	deposits []*ctypes.Deposit,
) error {
	switch {
	case sp.cs.DepositEth1ChainID() == spec.BartioChainID:
		// Bartio does not properly validate deposits index
		// We skip checks for backward compatibility
		return nil

	case sp.cs.DepositEth1ChainID() == spec.BoonetEth1ChainID:
		// Boonet inherited the bug from Bartio and it may have added some
		// validators before we activate the fork. So we skip all validations
		// but the validator set cap.
		//#nosec:G701 // can't overflow.
		if uint64(len(deposits)) > sp.cs.ValidatorSetCap() {
			return fmt.Errorf("validator set cap %d, deposits count %d: %w",
				sp.cs.ValidatorSetCap(),
				len(deposits),
				ErrValSetCapExceeded,
			)
		}
		return nil

	default:
		index, err := st.GetEth1DepositIndex()
		if err != nil {
			return err
		}
		if index != 0 {
			// there should not be Eth1DepositIndex stored before
			// genesis first deposit
			return errors.Wrap(
				ErrDepositMismatch,
				"Eth1DepositIndex should be unset at genesis",
			)
		}
		if len(deposits) == 0 {
			// there should be at least a validator in genesis
			return errors.Wrap(
				ErrDepositsLengthMismatch,
				"at least one validator should be in genesis",
			)
		}
		for i, deposit := range deposits {
			// deposit indices should be contiguous
			if deposit.GetIndex() != math.U64(i) {
				return errors.Wrapf(
					ErrDepositIndexOutOfOrder,
					"genesis deposit index: %d, expected index: %d",
					deposit.GetIndex().Unwrap(), i,
				)
			}
		}

		// BeaconKit enforces a cap on the validator set size.
		// If genesis deposits breaches the cap we return an error.
		//#nosec:G701 // can't overflow.
		if uint64(len(deposits)) > sp.cs.ValidatorSetCap() {
			return fmt.Errorf("validator set cap %d, deposits count %d: %w",
				sp.cs.ValidatorSetCap(),
				len(deposits),
				ErrValSetCapExceeded,
			)
		}
		return nil
	}
}

func (sp *StateProcessor[
	BeaconBlockT, BeaconStateT, _, _,
]) validateNonGenesisDeposits(
	st BeaconStateT,
	blk BeaconBlockT,
	deposits []*ctypes.Deposit,
) error {
	slot, err := st.GetSlot()
	if err != nil {
		return fmt.Errorf(
			"failed loading slot while processing deposits: %w", err,
		)
	}
	switch {
	case sp.cs.DepositEth1ChainID() == spec.BartioChainID:
		// Bartio does not properly validate deposits index
		// We skip checks for backward compatibility
		return nil

	case sp.cs.DepositEth1ChainID() == spec.BoonetEth1ChainID &&
		slot < math.U64(spec.BoonetFork2Height):
		// Boonet inherited the bug from Bartio and it may have added some
		// validators before we activate the fork. So we skip validation
		// before fork activation
		return nil

	default:
		// Verify that outstanding deposits match those listed by contract
		var depositIndex uint64
		depositIndex, err = st.GetEth1DepositIndex()
		if err != nil {
			return err
		}

		var localDeposits []*ctypes.Deposit
		localDeposits, err = sp.ds.GetDepositsByIndex(
			depositIndex, sp.cs.MaxDepositsPerBlock(),
		)
		if err != nil {
			return err
		}

		sp.logger.Info(
			"Processing deposits in range",
			"start_index", depositIndex,
			"range_length", len(localDeposits),
		)

		if len(localDeposits) != len(deposits) {
			return errors.Wrapf(
				ErrDepositsLengthMismatch,
				"local: %d, payload: %d", len(localDeposits), len(deposits),
			)
		}

		// Verify that the local deposits have the same root as the block deposits.
		eth1Data, err := st.GetEth1Data()
		if err != nil {
			return err
		}
		newEth1Data := blk.GetBody().GetEth1Data()
		localDepositsRoot := ctypes.Deposits(localDeposits).CombiHashTreeRoot(eth1Data.DepositRoot)
		if localDepositsRoot != newEth1Data.DepositRoot {
			return errors.Wrapf(
				ErrDepositMismatch, "deposits root mismatch, local: %s, received: %s",
				localDepositsRoot, newEth1Data.DepositRoot,
			)
		}
		return nil
	}
}
