package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/certikfoundation/shentu/v2/x/shield/types"
)

func (k Keeper) GetGlobalStakingPool(ctx sdk.Context) (pool sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetGlobalStakeForShieldPoolKey())
	if bz == nil {
		return sdk.NewInt(0)
	}

	ip := sdk.IntProto{}
	k.cdc.MustUnmarshalLengthPrefixed(bz, &ip)
	return ip.Int
}

func (k Keeper) SetGlobalStakingPool(ctx sdk.Context, value sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalLengthPrefixed(&sdk.IntProto{Int: value})
	store.Set(types.GetGlobalStakeForShieldPoolKey(), bz)
}

func (k Keeper) DeletePurchase(ctx sdk.Context, poolID uint64, purchaser sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetPurchaseKey(poolID, purchaser))
}

func (k Keeper) GetPurchase(ctx sdk.Context, poolID uint64, purchaser sdk.AccAddress) (purchase types.Purchase, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetPurchaseKey(poolID, purchaser))
	if bz != nil {
		k.cdc.MustUnmarshalLengthPrefixed(bz, &purchase)
		found = true
	}
	return
}

func (k Keeper) SetPurchase(ctx sdk.Context, purchase types.Purchase) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalLengthPrefixed(&purchase)
	purchaser, err := sdk.AccAddressFromBech32(purchase.Purchaser)
	if err != nil {
		panic(err)
	}
	store.Set(types.GetPurchaseKey(purchase.PoolId, purchaser), bz)
}

func (k Keeper) AddStaking(ctx sdk.Context, poolID uint64, purchaser sdk.AccAddress, description string, amount sdk.Coins) (types.Purchase, error) {
	if err := k.bk.SendCoinsFromAccountToModule(ctx, purchaser, types.ModuleName, amount); err != nil {
		return types.Purchase{}, err
	}
	pool, found := k.GetPool(ctx, poolID)
	if !found {
		return types.Purchase{}, types.ErrNoPoolFound
	}
	pool.Shield = pool.Shield.Add(amount.AmountOf(k.BondDenom(ctx)).ToDec().Mul(pool.ShieldRate).TruncateInt())
	k.SetPool(ctx, pool)

	bondDenomAmt := amount.AmountOf(k.BondDenom(ctx))
	shieldAmt := bondDenomAmt.ToDec().Mul(pool.ShieldRate).TruncateInt()
	gSPool := k.GetGlobalStakingPool(ctx)
	gSPool = gSPool.Add(bondDenomAmt)
	k.SetGlobalStakingPool(ctx, gSPool)

	sp, found := k.GetPurchase(ctx, poolID, purchaser)
	if !found {
		sp = types.NewPurchase(poolID, purchaser, description, bondDenomAmt, shieldAmt)
	} else {
		sp.Amount = sp.Amount.Add(bondDenomAmt)
		sp.Shield = sp.Shield.Add(shieldAmt)
	}
	sp.StartTime = ctx.BlockTime()
	k.SetPurchase(ctx, sp)

	totalShield := k.GetTotalShield(ctx)
	totalShield = totalShield.Add(shieldAmt)
	k.SetTotalShield(ctx, totalShield)
	return sp, nil
}

func (k Keeper) Unstake(ctx sdk.Context, poolID uint64, purchaser sdk.AccAddress, amount sdk.Coins) error {
	bdAmount := amount.AmountOf(k.BondDenom(ctx))

	sp, found := k.GetPurchase(ctx, poolID, purchaser)
	if !found {
		return types.ErrPurchaseNotFound
	}
	if sp.Amount.LT(bdAmount) {
		return types.ErrInsufficientStaking
	}
	poolParams := k.GetPoolParams(ctx)
	cd := poolParams.CooldownPeriod
	if sp.StartTime.Add(cd).After(ctx.BlockTime()) {
		return types.ErrBeforeCooldownEnd
	}
	sp.Amount = sp.Amount.Sub(bdAmount)
	if sp.Amount.Equal(sdk.ZeroInt()) {
		k.DeletePurchase(ctx, poolID, purchaser)
	} else {
		sp.StartTime = ctx.BlockTime()
		k.SetPurchase(ctx, sp)
	}

	// update pool
	pool, found := k.GetPool(ctx, poolID)
	if !found {
		return types.ErrNoPoolFound
	}
	pool.Shield = pool.Shield.Sub(bdAmount.ToDec().Mul(pool.ShieldRate).TruncateInt())

	// update global pool
	bondDenomAmt := bdAmount
	gSPool := k.GetGlobalStakingPool(ctx)
	gSPool = gSPool.Sub(bondDenomAmt)
	k.SetGlobalStakingPool(ctx, gSPool)

	withdrawCoins := amount

	return k.bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, purchaser, withdrawCoins)
}

func (k Keeper) FundShieldBlockRewards(ctx sdk.Context, amount sdk.Coins, sender sdk.AccAddress) error {
	if err := k.bk.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, amount); err != nil {
		return err
	}
	blockServiceFee := k.GetBlockServiceFees(ctx)
	blockServiceFee = blockServiceFee.Add(types.NewMixedDecCoins(sdk.NewDecCoinsFromCoins(amount...), sdk.NewDecCoins()))
	k.SetBlockServiceFees(ctx, blockServiceFee)
	return nil
}

func (k Keeper) GetAllPurchase(ctx sdk.Context) (purchases []types.Purchase) {
	k.IteratePurchases(ctx, func(purchase types.Purchase) bool {
		purchases = append(purchases, purchase)
		return false
	})
	return
}

// IteratePurchases iterates through purchase lists in a pool
func (k Keeper) IteratePurchases(ctx sdk.Context, callback func(purchase types.Purchase) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PurchaseKey)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var purchase types.Purchase
		k.cdc.MustUnmarshalLengthPrefixed(iterator.Value(), &purchase)

		if callback(purchase) {
			break
		}
	}
}
