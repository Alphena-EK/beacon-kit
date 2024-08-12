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

package components

import (
	"cosmossdk.io/depinject"
	"github.com/berachain/beacon-kit/mod/async/pkg/broker"
	asynctypes "github.com/berachain/beacon-kit/mod/async/pkg/types"
	"github.com/berachain/beacon-kit/mod/beacon/validator"
	"github.com/berachain/beacon-kit/mod/config"
	"github.com/berachain/beacon-kit/mod/log"
	"github.com/berachain/beacon-kit/mod/node-core/pkg/components/metrics"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/crypto"
)

// ValidatorServiceInput is the input for the validator service provider.
type ValidatorServiceInput[
	BeaconBlockT any,
	LoggerT log.AdvancedLogger[any, LoggerT],
] struct {
	depinject.In
	BeaconBlockFeed *broker.Broker[*asynctypes.Event[BeaconBlockT]]
	Cfg             *config.Config
	ChainSpec       common.ChainSpec
	LocalBuilder    *LocalBuilder
	Logger          LoggerT
	StateProcessor  *StateProcessor
	StorageBackend  *StorageBackend
	Signer          crypto.BLSSigner
	SidecarsFeed    *broker.Broker[*asynctypes.Event[BlobSidecarsT]]
	SidecarFactory  *SidecarFactory
	SlotBroker      *broker.Broker[*asynctypes.Event[SlotDataT]]
	TelemetrySink   *metrics.TelemetrySink
}

// ProvideValidatorService is a depinject provider for the validator service.
func ProvideValidatorService(
	in ValidatorServiceInput,
) (*ValidatorService, error) {
	slotSubscription, err := in.SlotBroker.Subscribe()
	if err != nil {
		in.Logger.Error("failed to subscribe to slot feed", "err", err)
		return nil, err
	}
	// Build the builder service.
	return validator.NewService[
		*AttestationData,
		*BeaconBlock,
		*BeaconBlockBody,
		*BeaconState,
		*SidecarFactory,
		*BlobSidecars,
		*Context,
		*Deposit,
		*DepositStore,
		*Eth1Data,
		*ExecutionPayload,
		*ExecutionPayloadHeader,
		*ForkData,
		log.Logger[any],
		*LocalBuilder,
		*SlashingInfo,
		*SlotData,
		*StateProcessor,
		*StorageBackend,
	](
		&in.Cfg.Validator,
		in.Logger.With("service", "validator"),
		in.ChainSpec,
		in.StorageBackend,
		in.StateProcessor,
		in.Signer,
		in.SidecarFactory,
		in.LocalBuilder,
		in.TelemetrySink,
		in.BeaconBlockFeed,
		in.SidecarsFeed,
		slotSubscription,
	), nil
}
