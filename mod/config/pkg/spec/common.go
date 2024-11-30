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

package spec

import "github.com/berachain/beacon-kit/mod/primitives/pkg/version"

// Common is the underlying data structure for chain-specific parameters.
//
//nolint:lll // struct tags may create long lines.
type Common[
	DomainTypeT ~[4]byte,
	EpochT ~uint64,
	ExecutionAddressT ~[20]byte,
	SlotT ~uint64,
	CometBFTConfigT any,
] struct {
	// Gwei value constants.
	//
	// MinDepositAmount is the minimum deposit amount per deposit
	// transaction.
	MinDepositAmount uint64 `mapstructure:"min-deposit-amount"`
	// MaxEffectiveBalance is the maximum effective balance allowed for a
	// validator.
	MaxEffectiveBalance uint64 `mapstructure:"max-effective-balance"`
	// EjectionBalance is the balance at which a validator is ejected.
	EjectionBalance uint64 `mapstructure:"ejection-balance"`
	// EffectiveBalanceIncrement is the effective balance increment.
	EffectiveBalanceIncrement uint64 `mapstructure:"effective-balance-increment"`

	// HysteresisQuotient is the quotient used in effective balance calculations
	HysteresisQuotient uint64 `mapstructure:"hysteresis-quotient"`
	// HysteresisDownwardMultiplier is the multiplier for downward balance adjustments.
	HysteresisDownwardMultiplier uint64 `mapstructure:"hysteresis-downward-multiplier"`
	// HysteresisUpwardMultiplier is the multiplier for upward balance adjustments.
	HysteresisUpwardMultiplier uint64 `mapstructure:"hysteresis-upward-multiplier"`
	// Time parameters constants.
	//
	// SlotsPerEpoch is the number of slots per epoch.
	SlotsPerEpoch uint64 `mapstructure:"slots-per-epoch"`
	// SlotsPerHistoricalRoot is the number of slots per historical root.
	SlotsPerHistoricalRoot uint64 `mapstructure:"slots-per-historical-root"`
	// MinEpochsToInactivityPenalty is the minimum number of epochs before a
	// validator is penalized for inactivity.
	MinEpochsToInactivityPenalty uint64 `mapstructure:"min-epochs-to-inactivity-penalty"`

	// Signature domains.
	//
	// DomainDomainTypeProposerProposer is the domain for beacon proposer
	// signatures.
	DomainTypeProposer DomainTypeT `mapstructure:"domain-type-beacon-proposer"`
	// DomainTypeAttester is the domain for beacon attester signatures.
	DomainTypeAttester DomainTypeT `mapstructure:"domain-type-beacon-attester"`
	// DomainTypeRandao is the domain for RANDAO reveal signatures.
	DomainTypeRandao DomainTypeT `mapstructure:"domain-type-randao"`
	// DomainTypeDeposit is the domain for deposit contract signatures.
	DomainTypeDeposit DomainTypeT `mapstructure:"domain-type-deposit"`
	// DomainTypeVoluntaryExit is the domain for voluntary exit signatures.
	DomainTypeVoluntaryExit DomainTypeT `mapstructure:"domain-type-voluntary-exit"`
	// DomainTypeSelectionProof is the domain for selection proof signatures.
	DomainTypeSelectionProof DomainTypeT `mapstructure:"domain-type-selection-proof"`
	// DomainTypeAggregateAndProof is the domain for aggregate and proof
	// signatures.
	DomainTypeAggregateAndProof DomainTypeT `mapstructure:"domain-type-aggregate-and-proof"`
	// DomainTypeApplicationMask is the domain for the application mask.
	DomainTypeApplicationMask DomainTypeT `mapstructure:"domain-type-application-mask"`

	// Eth1-related values.
	//
	// DepositContractAddress is the address of the deposit contract.
	DepositContractAddress ExecutionAddressT `mapstructure:"deposit-contract-address"`
	// MaxDepositsPerBlock specifies the maximum number of deposit operations
	// allowed per block.
	MaxDepositsPerBlock uint64 `mapstructure:"max-deposits-per-block"`
	// DepositEth1ChainID is the chain ID of the execution client.
	DepositEth1ChainID uint64 `mapstructure:"deposit-eth1-chain-id"`
	// Eth1FollowDistance is the distance between the eth1 chain and the beacon
	// chain with respect to reading deposits.
	Eth1FollowDistance uint64 `mapstructure:"eth1-follow-distance"`
	// TargetSecondsPerEth1Block is the target time between eth1 blocks.
	TargetSecondsPerEth1Block uint64 `mapstructure:"target-seconds-per-eth1-block"`

	// Fork-related values.
	//
	// DenebPlus is the epoch at which the Deneb+ fork is activated.
	DenebPlusForkEpoch EpochT `mapstructure:"deneb-plus-fork-epoch"`
	// ElectraForkEpoch is the epoch at which the Electra fork is activated.
	ElectraForkEpoch EpochT `mapstructure:"electra-fork-epoch"`

	// State list lengths
	//
	// EpochsPerHistoricalVector is the number of epochs in the historical
	// vector.
	EpochsPerHistoricalVector uint64 `mapstructure:"epochs-per-historical-vector"`
	// EpochsPerSlashingsVector is the number of epochs in the slashings vector.
	EpochsPerSlashingsVector uint64 `mapstructure:"epochs-per-slashings-vector"`
	// HistoricalRootsLimit is the maximum number of historical roots.
	HistoricalRootsLimit uint64 `mapstructure:"historical-roots-limit"`
	// ValidatorRegistryLimit is the maximum number of validators in the
	// registry.
	ValidatorRegistryLimit uint64 `mapstructure:"validator-registry-limit"`

	// Rewards and penalties constants.
	//
	// InactivityPenaltyQuotient is the inactivity penalty quotient.
	InactivityPenaltyQuotient uint64 `mapstructure:"inactivity-penalty-quotient"`
	// ProportionalSlashingMultiplier is the slashing multiplier relative to the
	// base penalty.
	ProportionalSlashingMultiplier uint64 `mapstructure:"proportional-slashing-multiplier"`

	// Capella Values
	//
	// MaxWithdrawalsPerPayload indicates the maximum number of withdrawal
	// operations allowed in a single payload.
	MaxWithdrawalsPerPayload uint64 `mapstructure:"max-withdrawals-per-payload"`
	// MaxValidatorsPerWithdrawalsSweep specifies the maximum number of
	// validator
	// withdrawals allowed per sweep.
	MaxValidatorsPerWithdrawalsSweep uint64 `mapstructure:"max-validators-per-withdrawals-sweep"`

	// Deneb Values
	//
	// MinEpochsForBlobsSidecarsRequest is the minimum number of epochs the node
	// will keep the blobs for.
	MinEpochsForBlobsSidecarsRequest uint64 `mapstructure:"min-epochs-for-blobs-sidecars-request"`
	// MaxBlobCommitmentsPerBlock specifies the maximum number of blob
	// commitments allowed per block.
	MaxBlobCommitmentsPerBlock uint64 `mapstructure:"max-blob-commitments-per-block"`
	// MaxBlobsPerBlock specifies the maximum number of blobs allowed per block.
	MaxBlobsPerBlock uint64 `mapstructure:"max-blobs-per-block"`
	// FieldElementsPerBlob specifies the number of field elements per blob.
	FieldElementsPerBlob uint64 `mapstructure:"field-elements-per-blob"`
	// BytesPerBlob denotes the size of EIP-4844 blobs in bytes.
	BytesPerBlob uint64 `mapstructure:"bytes-per-blob"`
	// KZGCommitmentInclusionProofDepth is the depth of the KZG inclusion proof.
	KZGCommitmentInclusionProofDepth uint64 `mapstructure:"kzg-commitment-inclusion-proof-depth"`

	// Comet Values
	CometValues CometBFTConfigT `mapstructure:"comet-bft-config"`

	// Berachain Values
	//
	// EVMInflationAddress is the address on the EVM which will receive the
	// inflation amount of native EVM balance through a withdrawal every block.
	EVMInflationAddress ExecutionAddressT `mapstructure:"evm-inflation-address"`
	// EVMInflationPerBlock is the amount of native EVM balance (in Gwei) to be
	// minted to the EVMInflationAddress via a withdrawal every block.
	EVMInflationPerBlock uint64 `mapstructure:"evm-inflation-per-block"`
}

// Validate ensures that the chain spec is valid, returning error if it is not.
func (c *Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) Validate() error {
	if c.GetMaxWithdrawalsPerPayload() <= 1 {
		return ErrInsufficientMaxWithdrawalsPerPayload
	}

	// EVM Inflation values can be zero or non-zero, no validation needed.

	// TODO: Add more validation rules here.
	return nil
}

// MinDepositAmount returns the minimum deposit amount required.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetMinDepositAmount() uint64 {
	return c.MinDepositAmount
}

// MaxEffectiveBalance returns the maximum effective balance.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetMaxEffectiveBalance() uint64 {
	return c.MaxEffectiveBalance
}

// EjectionBalance returns the balance below which a validator is ejected.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetEjectionBalance() uint64 {
	return c.EjectionBalance
}

// EffectiveBalanceIncrement returns the increment of effective balance.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetEffectiveBalanceIncrement() uint64 {
	return c.EffectiveBalanceIncrement
}

func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetHysteresisQuotient() uint64 {
	return c.HysteresisQuotient
}

func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetHysteresisDownwardMultiplier() uint64 {
	return c.HysteresisDownwardMultiplier
}

func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetHysteresisUpwardMultiplier() uint64 {
	return c.HysteresisUpwardMultiplier
}

// SlotsPerEpoch returns the number of slots per epoch.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetSlotsPerEpoch() uint64 {
	return c.SlotsPerEpoch
}

// SlotsPerHistoricalRoot returns the number of slots per historical root.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetSlotsPerHistoricalRoot() uint64 {
	return c.SlotsPerHistoricalRoot
}

// MinEpochsToInactivityPenalty returns the minimum number of epochs before an
// inactivity penalty is applied.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetMinEpochsToInactivityPenalty() uint64 {
	return c.MinEpochsToInactivityPenalty
}

// DomainTypeProposer returns the domain for beacon proposer signatures.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetDomainTypeProposer() DomainTypeT {
	return c.DomainTypeProposer
}

// DomainTypeAttester returns the domain for beacon attester signatures.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetDomainTypeAttester() DomainTypeT {
	return c.DomainTypeAttester
}

// DomainTypeRandao returns the domain for RANDAO reveal signatures.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetDomainTypeRandao() DomainTypeT {
	return c.DomainTypeRandao
}

// DomainTypeDeposit returns the domain for deposit contract signatures.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetDomainTypeDeposit() DomainTypeT {
	return c.DomainTypeDeposit
}

// DomainTypeVoluntaryExit returns the domain for voluntary exit signatures.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetDomainTypeVoluntaryExit() DomainTypeT {
	return c.DomainTypeVoluntaryExit
}

// DomainTypeSelectionProof returns the domain for selection proof signatures.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetDomainTypeSelectionProof() DomainTypeT {
	return c.DomainTypeSelectionProof
}

// DomainTypeAggregateAndProof returns the domain for aggregate and proof
// signatures.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetDomainTypeAggregateAndProof() DomainTypeT {
	return c.DomainTypeAggregateAndProof
}

// DomainTypeApplicationMask returns the domain for the application mask.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetDomainTypeApplicationMask() DomainTypeT {
	return c.DomainTypeApplicationMask
}

// DepositContractAddress returns the address of the deposit contract.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetDepositContractAddress() ExecutionAddressT {
	return c.DepositContractAddress
}

// MaxDepositsPerBlock returns the maximum number of deposits per block.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetMaxDepositsPerBlock() uint64 {
	return c.MaxDepositsPerBlock
}

// DepositEth1ChainID returns the chain ID of the execution chain.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetDepositEth1ChainID() uint64 {
	return c.DepositEth1ChainID
}

// Eth1FollowDistance returns the distance between the eth1 chain and the beacon
// chain.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetEth1FollowDistance() uint64 {
	return c.Eth1FollowDistance
}

// TargetSecondsPerEth1Block returns the target time between eth1 blocks.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetTargetSecondsPerEth1Block() uint64 {
	return c.TargetSecondsPerEth1Block
}

// DenebPlusForEpoch returns the epoch of the Deneb+ fork.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetDenebPlusForkEpoch() EpochT {
	return c.DenebPlusForkEpoch
}

// ElectraForkEpoch returns the epoch of the Electra fork.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetElectraForkEpoch() EpochT {
	return c.ElectraForkEpoch
}

// EpochsPerHistoricalVector returns the number of epochs per historical vector.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetEpochsPerHistoricalVector() uint64 {
	return c.EpochsPerHistoricalVector
}

// EpochsPerSlashingsVector returns the number of epochs per slashings vector.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetEpochsPerSlashingsVector() uint64 {
	return c.EpochsPerSlashingsVector
}

// HistoricalRootsLimit returns the limit of historical roots.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetHistoricalRootsLimit() uint64 {
	return c.HistoricalRootsLimit
}

// ValidatorRegistryLimit returns the limit of the validator registry.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetValidatorRegistryLimit() uint64 {
	return c.ValidatorRegistryLimit
}

// InactivityPenaltyQuotient returns the inactivity penalty quotient.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetInactivityPenaltyQuotient() uint64 {
	return c.InactivityPenaltyQuotient
}

// ProportionalSlashingMultiplier returns the proportional slashing multiplier.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetProportionalSlashingMultiplier() uint64 {
	return c.ProportionalSlashingMultiplier
}

// MaxWithdrawalsPerPayload returns the maximum number of withdrawals per
// payload.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetMaxWithdrawalsPerPayload() uint64 {
	return c.MaxWithdrawalsPerPayload
}

// MaxValidatorsPerWithdrawalsSweep returns the maximum number of validators per
// withdrawals sweep.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetMaxValidatorsPerWithdrawalsSweep() uint64 {
	return c.MaxValidatorsPerWithdrawalsSweep
}

// MinEpochsForBlobsSidecarsRequest returns the minimum number of epochs for
// blobs sidecars request.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetMinEpochsForBlobsSidecarsRequest() uint64 {
	return c.MinEpochsForBlobsSidecarsRequest
}

// MaxBlobCommitmentsPerBlock returns the maximum number of blob commitments per
// block.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetMaxBlobCommitmentsPerBlock() uint64 {
	return c.MaxBlobCommitmentsPerBlock
}

// MaxBlobsPerBlock returns the maximum number of blobs per block.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetMaxBlobsPerBlock() uint64 {
	return c.MaxBlobsPerBlock
}

// FieldElementsPerBlob returns the number of field elements per blob.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetFieldElementsPerBlob() uint64 {
	return c.FieldElementsPerBlob
}

// BytesPerBlob returns the number of bytes per blob.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetBytesPerBlob() uint64 {
	return c.BytesPerBlob
}

// GetCometBFTConfigForSlot returns the CometBFT configuration for the given
// slot.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetCometBFTConfigForSlot(_ SlotT) CometBFTConfigT {
	return c.CometValues
}

// EVMInflationAddress returns the address on the EVM which will receive the
// inflation amount of native EVM balance through a withdrawal every block.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetEVMInflationAddress() ExecutionAddressT {
	return c.EVMInflationAddress
}

// EVMInflationPerBlock returns the amount of native EVM balance (in Gwei) to
// be minted to the EVMInflationAddress via a withdrawal every block.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetEVMInflationPerBlock() uint64 {
	return c.EVMInflationPerBlock
}

// ActiveForkVersionForSlot returns the active fork version for a given slot.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetActiveForkVersionForSlot(
	slot SlotT,
) uint32 {
	return c.GetActiveForkVersionForEpoch(c.GetSlotToEpoch(slot))
}

// ActiveForkVersionForEpoch returns the active fork version for a given epoch.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetActiveForkVersionForEpoch(
	epoch EpochT,
) uint32 {
	switch {
	case epoch >= c.ElectraForkEpoch:
		return version.Electra
	case epoch >= c.DenebPlusForkEpoch:
		return version.DenebPlus
	default:
		return version.Deneb
	}
}

// SlotToEpoch converts a slot to an epoch.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetSlotToEpoch(slot SlotT) EpochT {
	//#nosec:G701 // realistically fine in practice.
	return EpochT(uint64(slot) / c.GetSlotsPerEpoch())
}

// WithinDAPeriod checks if the block epoch is within
// MIN_EPOCHS_FOR_BLOB_SIDECARS_REQUESTS
// of the given current epoch.
func (c Common[
	DomainTypeT, EpochT, ExecutionAddressT, SlotT, CometBFTConfigT,
]) GetWithinDAPeriod(
	block, current SlotT,
) bool {
	return c.GetSlotToEpoch(block)+
		EpochT(c.GetMinEpochsForBlobsSidecarsRequest()) >=
		c.GetSlotToEpoch(current)
}
