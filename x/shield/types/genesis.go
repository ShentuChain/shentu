package types

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewGenesisState creates a new genesis state.
func NewGenesisState(shieldAdmin sdk.AccAddress, nextPoolID, nextPurchaseID uint64, poolParams PoolParams,
	claimProposalParams ClaimProposalParams, totalCollateral, totalWithdrawing, totalShield, totalClaimed sdk.Int,
	serviceFees sdk.DecCoins, pools []Pool, providers []Provider, withdraws []Withdraw,
	globalStakingPool sdk.Int, stakingPurchases []Purchase,
	donationPool DonationPool, pendingPayouts []PendingPayout, blockRewardParams BlockRewardParams) GenesisState {
	return GenesisState{
		ShieldAdmin:         shieldAdmin.String(),
		NextPoolId:          nextPoolID,
		NextPurchaseId:      nextPurchaseID,
		PoolParams:          poolParams,
		ClaimProposalParams: claimProposalParams,
		TotalCollateral:     totalCollateral,
		TotalWithdrawing:    totalWithdrawing,
		TotalShield:         totalShield,
		TotalClaimed:        totalClaimed,
		Fees:                serviceFees,
		Pools:               pools,
		Providers:           providers,
		Withdraws:           withdraws,
		GlobalStakingPool:   globalStakingPool,
		Purchases:           stakingPurchases,
		DonationPool:        donationPool,
		PendingPayouts:      pendingPayouts,
		BlockRewardParams:   blockRewardParams,
	}
}

// DefaultGenesisState returns a default genesis state.
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		NextPoolId:          uint64(1),
		NextPurchaseId:      uint64(1),
		PoolParams:          DefaultPoolParams(),
		ClaimProposalParams: DefaultClaimProposalParams(),
		TotalCollateral:     sdk.ZeroInt(),
		TotalWithdrawing:    sdk.ZeroInt(),
		TotalShield:         sdk.ZeroInt(),
		TotalClaimed:        sdk.ZeroInt(),
		Fees:                sdk.NewDecCoins(),
		DonationPool:        InitialDonationPool(),
		BlockRewardParams:   DefaultBlockRewardParams(),
	}
}

// ValidateGenesis validates shield genesis data.
func ValidateGenesis(data GenesisState) error {
	if data.NextPoolId < 1 {
		return fmt.Errorf("failed to validate %s genesis state: NextPoolID must be positive ", ModuleName)
	}
	if data.DonationPool.Amount.IsNegative() {
		return fmt.Errorf("donation pool amount is negative %v", data.DonationPool.Amount)
	}
	if err := validatePoolParams(data.PoolParams); err != nil {
		return fmt.Errorf("failed to validate %s pool params: %w", ModuleName, err)
	}
	if err := validateClaimProposalParams(data.ClaimProposalParams); err != nil {
		return fmt.Errorf("failed to validate %s claim proposal params: %w", ModuleName, err)
	}
	if err := validateBlockRewardParams(data.BlockRewardParams); err != nil {
		return fmt.Errorf("failed to validate %s block reward params: %w", ModuleName, err)
	}

	return nil
}

// GetGenesisStateFromAppState returns GenesisState given raw application genesis state.
func GetGenesisStateFromAppState(cdc codec.Codec, appState map[string]json.RawMessage) GenesisState {
	var genesisState GenesisState
	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}
	return genesisState
}
