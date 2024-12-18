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

package genesis

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/berachain/beacon-kit/cli/context"
	"github.com/berachain/beacon-kit/consensus-types/types"
	"github.com/berachain/beacon-kit/errors"
	"github.com/berachain/beacon-kit/primitives/encoding/json"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

// CollectGenesisDepositsCmd - return the cobra command to
// collect genesis transactions.
func CollectGenesisDepositsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "collect-premined-deposits",
		Short: "adds a validator to the genesis file",
		RunE: func(cmd *cobra.Command, _ []string) error {
			config := context.GetConfigFromCmd(cmd)

			appGenesis, err := genutiltypes.AppGenesisFromFile(
				config.GenesisFile(),
			)
			if err != nil {
				return errors.Wrap(err, "failed to read genesis doc from file")
			}

			// create the app state
			appGenesisState, err := genutiltypes.GenesisStateFromAppGenesis(
				appGenesis,
			)
			if err != nil {
				return err
			}

			genesisInfo := &types.Genesis{}

			if genesisInfo.DepositDatas, err = CollectValidatorJSONFiles(
				filepath.Join(config.RootDir, "config", "premined-deposits"), appGenesis,
			); err != nil {
				return errors.Wrap(
					err,
					"failed to collect validator json files",
				)
			}

			if err = json.Unmarshal(appGenesisState["beacon"], genesisInfo); err != nil {
				return errors.Wrap(err, "failed to unmarshal beacon genesis")
			}

			appGenesisState["beacon"], err = json.Marshal(genesisInfo)
			if err != nil {
				return errors.Wrap(err, "failed to marshal beacon genesis")
			}

			if appGenesis.AppState, err = json.MarshalIndent(
				appGenesisState, "", "  ",
			); err != nil {
				return err
			}

			return genutil.ExportGenesisFile(appGenesis, config.GenesisFile())
		},
	}

	return cmd
}

// CollectValidatorJSONFiles collects JSON files from the specified directory
// and unmarshals them into a list of DepositData objects.
func CollectValidatorJSONFiles(
	genTxsDir string,
	genesis *genutiltypes.AppGenesis,
) ([]*types.DepositData, error) {
	// prepare a map of all balances in genesis state to then validate
	// against the validators addresses
	var appState map[string]json.RawMessage
	if err := json.Unmarshal(genesis.AppState, &appState); err != nil {
		return nil, err
	}

	// get the list of files in the genTxsDir
	fos, err := os.ReadDir(genTxsDir)
	if err != nil {
		return nil, err
	}

	// prepare the list of validators
	deposits := make([]*types.DepositData, 0)
	for _, fo := range fos {
		if fo.IsDir() {
			continue
		}
		if !strings.HasSuffix(fo.Name(), ".json") {
			continue
		}

		var bz []byte
		bz, err = afero.ReadFile(
			afero.NewOsFs(),
			filepath.Join(genTxsDir, fo.Name()),
		)
		if err != nil {
			return nil, err
		}

		val := &types.DepositData{}
		if err = json.Unmarshal(bz, val); err != nil {
			return nil, err
		}

		deposits = append(deposits, val)
	}

	return deposits, nil
}
