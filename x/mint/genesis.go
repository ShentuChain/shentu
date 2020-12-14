package mint

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/mint/types"

	"github.com/certikfoundation/shentu/common"
)

// DefaultGenesisState creates a default GenesisState object.
func DefaultGenesisState() *types.GenesisState {
	return &types.GenesisState{
		Minter: types.InitialMinter(sdk.NewDecWithPrec(4, 2)),
		Params: types.NewParams(
			common.MicroCTKDenom,
			sdk.NewDecWithPrec(10, 2), // max inflation rate change
			sdk.NewDecWithPrec(14, 2), // max inflation rate
			sdk.NewDecWithPrec(4, 2),  // min inflation rate
			sdk.NewDecWithPrec(67, 2), // target staked coin percentage
			common.BlocksPerYear,      // blocks per year, 5 second block time
		),
	}
}
