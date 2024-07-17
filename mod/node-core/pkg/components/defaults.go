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
	"cosmossdk.io/core/transaction"

	// bktransaction "github.com/berachain/beacon-kit/mod/node-core/pkg/components/transaction"
	"github.com/cosmos/cosmos-sdk/codec"
)

func DefaultComponentsWithStandardTypes() []any {
	return []any{
		ProvideABCIMiddleware,
		ProvideAttributesFactory,
		ProvideAvailabilityPruner,
		ProvideAvailibilityStore,
		ProvideBlsSigner,
		ProvideBlobFeed,
		ProvideBlockFeed,
		ProvideBlobProcessor,
		ProvideBlobProofVerifier,
		ProvideBlobVerifier,
		ProvideChainService,
		ProvideChainSpec,
		ProvideConfig,
		ProvideDAService,
		ProvideDBManager,
		ProvideDepositPruner,
		ProvideDepositService,
		ProvideDepositStore,
		ProvideBeaconDepositContract,
		ProvideEngineClient,
		ProvideExecutionEngine,
		ProvideGenesisBroker,
		ProvideJWTSecret,
		ProvideLocalBuilder,
		ProvideServiceRegistry,
		ProvideSidecarFactory,
		ProvideStateProcessor,
		ProvideSlotBroker,
		ProvideStatusBroker,
		ProvideStorageBackend,
		ProvideTelemetrySink,
		ProvideTrustedSetup,
		ProvideValidatorService,
		ProvideValidatorUpdateBroker,
		// ProvideNoopTxConfig,
		// ProvideTxCodec[*bktransaction.SSZTx],
		ProvideTxCodec[transaction.Tx],

		// server v2
		ProvideBlockStore,
		ProvideStateStore, // i pray this works otherwise i kms
		ProvideStateManager,
		ProvideStoreOptions,

		codec.ProvideInterfaceRegistry,
		codec.ProvideAddressCodec,
		codec.ProvideProtoCodec,
		codec.ProvideLegacyAmino,
	}
}
