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

package proof

import (
	"net/http"

	"github.com/berachain/beacon-kit/log"
	"github.com/berachain/beacon-kit/node-api/handlers"
)

func (h *Handler[_, _, ContextT]) RegisterRoutes(logger log.Logger) {
	h.SetLogger(logger)
	h.BaseHandler.AddRoutes([]*handlers.Route[ContextT]{
		{
			Method:  http.MethodGet,
			Path:    "bkit/v1/proof/block_proposer/:timestamp_id",
			Handler: h.GetBlockProposer,
		},
		{
			Method:  http.MethodGet,
			Path:    "bkit/v1/proof/execution_number/:timestamp_id",
			Handler: h.GetExecutionNumber,
		},
		{
			Method:  http.MethodGet,
			Path:    "bkit/v1/proof/execution_fee_recipient/:timestamp_id",
			Handler: h.GetExecutionFeeRecipient,
		},
	})
}
