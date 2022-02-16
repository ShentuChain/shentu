package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/certikfoundation/shentu/v2/x/shield/types"
)

// RegisterInvariants registers all shield invariants.
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, "module-account", ModuleAccountInvariant(k))
	ir.RegisterRoute(types.ModuleName, "provider", ProviderInvariant(k))
	ir.RegisterRoute(types.ModuleName, "shield", ShieldInvariant(k))
	ir.RegisterRoute(types.ModuleName, "global-staking-pool", GlobalStakingPoolInvariant(k))
}

// ModuleAccountInvariant checks that the module account coins reflects the sum of
// remaining services and rewards held on store
func ModuleAccountInvariant(keeper Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		bondDenom := keeper.BondDenom(ctx)

		moduleCoins := keeper.bk.GetAllBalances(ctx, keeper.ak.GetModuleAccount(ctx, types.ModuleName).GetAddress())

		// remaining service fees
		serviceFees := keeper.GetServiceFees(ctx)

		// rewards
		var rewards sdk.DecCoins
		for _, provider := range keeper.GetAllProviders(ctx) {
			rewards = rewards.Add(provider.Rewards...)
		}

		totalInt, change := serviceFees.Add(rewards...).TruncateDecimal()

		// shield stake
		shieldStake := sdk.ZeroInt()
		for _, stake := range keeper.GetAllPurchase(ctx) {
			shieldStake = shieldStake.Add(stake.Amount)
		}

		// pending payouts
		pending_payouts := sdk.ZeroInt()
		for _, pp := range keeper.GetAllPendingPayouts(ctx) {
			pending_payouts = pending_payouts.Add(pp.Amount)
		}

		// reserve payouts
		reserve := keeper.GetReserve(ctx).Amount

		totalInt = totalInt.Add(sdk.NewCoin(bondDenom, shieldStake)).Add(sdk.NewCoin(bondDenom, pending_payouts)).Add(sdk.NewCoin(bondDenom, reserve))

		broken := !totalInt.IsEqual(moduleCoins) || !change.Empty()

		return sdk.FormatInvariant(types.ModuleName, "module-account",
			fmt.Sprintf("\n\tshield ModuleAccount coins: %s"+
				"\n\tsum of remaining service fees & rewards & staked & reimbursement & pending payouts amount:  %s"+
				"\n\tremaining change amount: %s\n",
				moduleCoins, totalInt, change)), broken
	}
}

// ProviderInvariant checks that the providers' coin amounts equal to the tracked value.
func ProviderInvariant(keeper Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		providers := keeper.GetAllProviders(ctx)
		withdrawSum := sdk.NewInt(0)
		collateralSum := sdk.NewInt(0)
		for _, prov := range providers {
			withdrawSum = withdrawSum.Add(prov.Withdrawing)
			collateralSum = collateralSum.Add(prov.Collateral)
		}

		totalWithdraw := keeper.GetTotalWithdrawing(ctx)
		totalCollateral := keeper.GetTotalCollateral(ctx)
		broken := !totalWithdraw.Equal(withdrawSum) || !totalCollateral.Equal(collateralSum)

		return sdk.FormatInvariant(types.ModuleName, "provider",
			fmt.Sprintf("\n\ttotal withdraw amount: %s"+
				"\n\tsum of providers' withdrawing amount:  %s"+
				"\n\ttotal collateral amount: %s"+
				"\n\tsum of providers' collateral amount: %s\n",
				totalWithdraw, withdrawSum, totalCollateral, collateralSum)), broken
	}
}

// ShieldInvariant checks that the sum of individual pools' shield is
// equal to the total shield.
func ShieldInvariant(keeper Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		pools := keeper.GetAllPools(ctx)
		shieldSum := sdk.NewInt(0)
		for _, pool := range pools {
			shieldSum = shieldSum.Add(pool.Shield)
		}

		purchases := keeper.GetAllPurchase(ctx)
		pShieldSum := sdk.NewInt(0)
		for _, purchase := range purchases {
			pShieldSum = pShieldSum.Add(purchase.Shield)
		}

		totalShield := keeper.GetTotalShield(ctx)
		broken := !totalShield.Equal(shieldSum) || !totalShield.Equal(pShieldSum) || totalShield.IsNegative()

		return sdk.FormatInvariant(types.ModuleName, "shield",
			fmt.Sprintf("\n\ttotal shield amount: %s"+
				"\n\tsum of pools' shield amount:  %s\n",
				totalShield, shieldSum)), broken
	}
}

// GlobalStakingPoolInvariant checks the total staked sum equals to the global staking pool amount.
func GlobalStakingPoolInvariant(keeper Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		stakedCoin := sdk.NewCoin(keeper.BondDenom(ctx), sdk.ZeroInt())
		for _, staked := range keeper.GetAllPurchase(ctx) {
			stakedCoin = stakedCoin.Add(sdk.NewCoin(keeper.BondDenom(ctx), staked.Amount))
		}
		stakedInt := stakedCoin.Amount
		globalStakingPool := keeper.GetGlobalStakingPool(ctx)
		broken := !stakedInt.Equal(globalStakingPool)

		return sdk.FormatInvariant(types.ModuleName, "global-staking-pool",
			fmt.Sprintf("\n\tsum of staked amount:  %s"+
				"\n\tglobal staking pool amount: %s\n",
				stakedInt, globalStakingPool.String())), broken
	}
}
