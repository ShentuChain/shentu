package shield

import (
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/certikfoundation/shentu/common"
	"github.com/certikfoundation/shentu/x/shield/types"
)

// BeginBlock executes logics to begin a block
func BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock, k Keeper) {
}

// EndBlocker processes premium payment at every block.
func EndBlocker(ctx sdk.Context, k Keeper) {
	pools := k.GetAllPools(ctx)
	for _, pool := range pools {
		if k.PoolEnded(ctx, pool) && (pool.Premium.Native.Empty() && pool.Premium.Foreign.Empty()) {
			k.ClosePool(ctx, pool)
			continue
		}
		// compute premiums for current block
		var currentBlockPremium types.MixedDecCoins
		timeUntilEnd := pool.EndTime.Sub(ctx.BlockTime())
		blocksUntilEnd := sdk.MaxDec(sdk.NewDec(timeUntilEnd.Milliseconds()/1000).QuoInt64(int64(common.SecondsPerBlock)), sdk.OneDec())
		if ctx.BlockTime().After(pool.EndTime) {
			// must spend all premium
			currentBlockPremium = pool.Premium
		} else {
			currentBlockPremium = pool.Premium.QuoDec(blocksUntilEnd)
		}

		// distribute to A and C in proportion
		totalCollateralAmount := pool.TotalCollateral
		recipients := k.GetAllPoolCollaterals(ctx, pool)
		for _, recipient := range recipients {
			stakeProportion := sdk.NewDecFromInt(recipient.Amount).QuoInt(totalCollateralAmount)
			nativePremium := currentBlockPremium.Native.MulDecTruncate(stakeProportion)
			foreignPremium := currentBlockPremium.Foreign.MulDecTruncate(stakeProportion)

			pool.Premium.Native = pool.Premium.Native.Sub(nativePremium)

			pool.Premium.Foreign = pool.Premium.Foreign.Sub(foreignPremium)

			rewards := types.NewMixedDecCoins(nativePremium, foreignPremium)
			k.AddRewards(ctx, recipient.Provider, rewards)
		}

		k.SetPool(ctx, pool)
	} // for each pool

	// remove expired purchases
	k.RemoveExpiredPurchases(ctx)

	// process completed withdraws
	k.DequeueCompletedWithdrawQueue(ctx)
}
