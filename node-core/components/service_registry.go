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
	"github.com/berachain/beacon-kit/beacon/blockchain"
	"github.com/berachain/beacon-kit/beacon/validator"
	cometbft "github.com/berachain/beacon-kit/consensus/cometbft/service"
	engineprimitives "github.com/berachain/beacon-kit/engine-primitives/engine-primitives"
	"github.com/berachain/beacon-kit/execution/client"
	"github.com/berachain/beacon-kit/log"
	"github.com/berachain/beacon-kit/node-api/server"
	"github.com/berachain/beacon-kit/node-core/components/metrics"
	service "github.com/berachain/beacon-kit/node-core/services/registry"
	"github.com/berachain/beacon-kit/node-core/services/version"
	"github.com/berachain/beacon-kit/observability/telemetry"
)

// ServiceRegistryInput is the input for the service registry provider.
type ServiceRegistryInput[
	AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT, BlobSidecarsT],
	ConsensusBlockT ConsensusBlock[BeaconBlockT],
	BeaconBlockT BeaconBlock[BeaconBlockT, BeaconBlockBodyT],
	BeaconBlockBodyT BeaconBlockBody[
		BeaconBlockBodyT,
		ExecutionPayloadT, *SlashingInfo,
	],
	BeaconBlockStoreT BlockStore[BeaconBlockT],
	BeaconStateT BeaconState[
		BeaconStateT, BeaconStateMarshallableT,
		ExecutionPayloadHeaderT, KVStoreT,
	],
	BeaconStateMarshallableT any,
	ConsensusSidecarsT ConsensusSidecars[BlobSidecarsT],
	BlobSidecarT any,
	BlobSidecarsT BlobSidecars[BlobSidecarsT, BlobSidecarT],
	DepositStoreT DepositStore,
	ExecutionPayloadT ExecutionPayload[
		ExecutionPayloadT, ExecutionPayloadHeaderT,
	],
	ExecutionPayloadHeaderT ExecutionPayloadHeader[ExecutionPayloadHeaderT],
	GenesisT Genesis[ExecutionPayloadHeaderT],
	KVStoreT any,
	LoggerT log.AdvancedLogger[LoggerT],
	NodeAPIContextT NodeAPIContext,
] struct {
	depinject.In
	ChainService *blockchain.Service[
		AvailabilityStoreT, DepositStoreT,
		ConsensusBlockT, BeaconBlockT, BeaconBlockBodyT,
		BeaconStateT, BeaconBlockStoreT,
		ExecutionPayloadT,
		ExecutionPayloadHeaderT, GenesisT,
		ConsensusSidecarsT, BlobSidecarsT,
		*engineprimitives.PayloadAttributes,
	]
	EngineClient *client.EngineClient[
		ExecutionPayloadT,
		*engineprimitives.PayloadAttributes,
	]
	Logger           LoggerT
	NodeAPIServer    *server.Server[NodeAPIContextT]
	ReportingService *version.ReportingService[
		ExecutionPayloadT,
		*engineprimitives.PayloadAttributes,
	]
	TelemetrySink    *metrics.TelemetrySink
	TelemetryService *telemetry.Service
	ValidatorService *validator.Service[
		BeaconBlockT, BeaconBlockBodyT,
		BeaconStateT, BlobSidecarT, BlobSidecarsT, DepositStoreT,
		ExecutionPayloadT, ExecutionPayloadHeaderT,
		*SlashingInfo, *SlotData,
	]
	CometBFTService *cometbft.Service[LoggerT]
}

// ProvideServiceRegistry is the depinject provider for the service registry.
func ProvideServiceRegistry[
	AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT, BlobSidecarsT],
	ConsensusBlockT ConsensusBlock[BeaconBlockT],
	BeaconBlockT BeaconBlock[BeaconBlockT, BeaconBlockBodyT],
	BeaconBlockBodyT BeaconBlockBody[
		BeaconBlockBodyT,
		ExecutionPayloadT, *SlashingInfo,
	],
	BeaconBlockStoreT BlockStore[BeaconBlockT],
	BeaconStateT BeaconState[
		BeaconStateT, BeaconStateMarshallableT,
		ExecutionPayloadHeaderT, KVStoreT,
	],
	BeaconStateMarshallableT any,
	ConsensusSidecarsT ConsensusSidecars[BlobSidecarsT],
	BlobSidecarT any,
	BlobSidecarsT BlobSidecars[BlobSidecarsT, BlobSidecarT],
	DepositStoreT DepositStore,
	ExecutionPayloadT ExecutionPayload[ExecutionPayloadT,
		ExecutionPayloadHeaderT],
	ExecutionPayloadHeaderT ExecutionPayloadHeader[ExecutionPayloadHeaderT],
	GenesisT Genesis[ExecutionPayloadHeaderT],
	KVStoreT any,
	LoggerT log.AdvancedLogger[LoggerT],
	NodeAPIContextT NodeAPIContext,
](
	in ServiceRegistryInput[
		AvailabilityStoreT,
		ConsensusBlockT, BeaconBlockT, BeaconBlockBodyT,
		BeaconBlockStoreT, BeaconStateT,
		BeaconStateMarshallableT,
		ConsensusSidecarsT, BlobSidecarT, BlobSidecarsT,
		DepositStoreT, ExecutionPayloadT, ExecutionPayloadHeaderT,
		GenesisT, KVStoreT, LoggerT, NodeAPIContextT,
	],
) *service.Registry {
	return service.NewRegistry(
		service.WithLogger(in.Logger),
		service.WithService(in.ValidatorService),
		service.WithService(in.NodeAPIServer),
		service.WithService(in.ReportingService),
		service.WithService(in.EngineClient),
		service.WithService(in.TelemetryService),
		service.WithService(in.ChainService),
		service.WithService(in.CometBFTService),
	)
}
