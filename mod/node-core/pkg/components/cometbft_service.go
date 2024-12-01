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
	storetypes "cosmossdk.io/store/types"
	"github.com/berachain/beacon-kit/mod/config"
	"github.com/berachain/beacon-kit/mod/config/pkg/spec"
	cometbft "github.com/berachain/beacon-kit/mod/consensus/pkg/cometbft/service"
	"github.com/berachain/beacon-kit/mod/log"
	"github.com/berachain/beacon-kit/mod/node-core/pkg/builder"
	cmtcfg "github.com/cometbft/cometbft/config"
	dbm "github.com/cosmos/cosmos-db"
)

// ProvideCometBFTService provides the CometBFT service component.
func ProvideCometBFTService[
	LoggerT log.AdvancedLogger[LoggerT],
](
	logger LoggerT,
	storeKey *storetypes.KVStoreKey,
	abciMiddleware cometbft.MiddlewareI,
	db dbm.DB,
	cmtCfg *cmtcfg.Config,
	appOpts config.AppOptions,
	chainSpec spec.Chain[any],
) *cometbft.Service[LoggerT] {
	return cometbft.NewService(
		storeKey,
		logger,
		db,
		abciMiddleware,
		cmtCfg,
		chainSpec,
		builder.DefaultServiceOptions[LoggerT](appOpts)...,
	)
}
