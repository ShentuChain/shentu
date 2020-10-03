package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/certikfoundation/shentu/x/shield/types"
)

// GetPoolCollateral retrieves collateral for a pool-provider pair.
func (k Keeper) GetCollateral(ctx sdk.Context, pool types.Pool, addr sdk.AccAddress) (types.Collateral, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetCollateralKey(pool.PoolID, addr))
	if bz == nil {
		return types.Collateral{}, false
	}
	var collateral types.Collateral
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &collateral)
	return collateral, true
}

// SetCollateral retrieves collateral for a pool-provider pair.
func (k Keeper) SetCollateral(ctx sdk.Context, pool types.Pool, addr sdk.AccAddress, collateral types.Collateral) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(collateral)
	store.Set(types.GetCollateralKey(pool.PoolID, addr), bz)
}

// FreeCollateral frees collaterals deposited in a pool.
func (k Keeper) FreeCollaterals(ctx sdk.Context, pool types.Pool) {
	store := ctx.KVStore(k.storeKey)
	k.IteratePoolCollaterals(ctx, pool, func(collateral types.Collateral) bool {
		provider, _ := k.GetProvider(ctx, collateral.Provider)
		provider.Collateral = provider.Collateral.Sub(collateral.Amount)
		k.SetProvider(ctx, collateral.Provider, provider)
		store.Delete(types.GetCollateralKey(pool.PoolID, collateral.Provider))
		return false
	})
}

// IteratePoolCollaterals iterates through collaterals in a pool
func (k Keeper) IteratePoolCollaterals(ctx sdk.Context, pool types.Pool, callback func(collateral types.Collateral) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, append(types.PoolKey, types.GetPoolKey(pool.PoolID)...))

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var collateral types.Collateral
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iterator.Value(), &collateral)

		if callback(collateral) {
			break
		}
	}
}

// GetOnesCollaterals returns a community member's all collaterals.
func (k Keeper) GetOnesCollaterals(ctx sdk.Context, address sdk.AccAddress) (collaterals []types.Collateral) {
	k.IterateAllPools(ctx, func(pool types.Pool) bool {
		collateral, found := k.GetCollateral(ctx, pool, address)
		if found {
			collaterals = append(collaterals, collateral)
		}
		return false
	})
	return collaterals
}

// GetPoolCertiKCollateral retrieves CertiK's provided collateral from a pool.
func (k Keeper) GetPoolCertiKCollateral(ctx sdk.Context, pool types.Pool) (collateral types.Collateral) {
	admin := k.GetAdmin(ctx)
	collateral, _ = k.GetCollateral(ctx, pool, admin)
	return
}

// GetAllPoolCollaterals retrieves all collaterals in a pool.
func (k Keeper) GetAllPoolCollaterals(ctx sdk.Context, pool types.Pool) (collaterals []types.Collateral) {
	k.IteratePoolCollaterals(ctx, pool, func(collateral types.Collateral) bool {
		collaterals = append(collaterals, collateral)
		return false
	})
	return collaterals
}

// DepositCollateral deposits a community member's collateral for a pool.
func (k Keeper) DepositCollateral(ctx sdk.Context, from sdk.AccAddress, id uint64, amount sdk.Coins) error {
	pool, err := k.GetPool(ctx, id)
	if err != nil {
		return err
	}

	// check eligibility
	provider, found := k.GetProvider(ctx, from)
	if !found {
		k.addProvider(ctx, from)
		provider, _ = k.GetProvider(ctx, from)
	}
	provider.Collateral = provider.Collateral.Add(amount...)
	if provider.Collateral.IsAnyGT(provider.DelegationBonded) {
		return types.ErrInsufficientStaking
	}

	// update the pool - update or create collateral entry
	collateral, found := k.GetCollateral(ctx, pool, from)
	if !found {
		collateral = types.NewCollateral(pool, from, amount)
	}
	collateral.Amount = collateral.Amount.Add(amount...)
	k.SetCollateral(ctx, pool, from, collateral)
	pool.TotalCollateral = pool.TotalCollateral.Add(amount...)
	k.SetPool(ctx, pool)
	k.SetProvider(ctx, from, provider)

	return nil
}

// WithdrawCollateral withdraws a community member's collateral for a pool.
func (k Keeper) WithdrawCollateral(ctx sdk.Context, from sdk.AccAddress, id uint64, amount sdk.Coins) error {
	pool, err := k.GetPool(ctx, id)
	if err != nil {
		return err
	}

	// check eligibility
	provider, found := k.GetProvider(ctx, from)
	if !found {
		return types.ErrNoDelegationAmount
	}
	if amount.IsAnyGT(provider.Collateral) {
		return types.ErrInvalidCollateralAmount
	}

	// insert into withdrawal queue
	poolParams := k.GetPoolParams(ctx)
	completionTime := ctx.BlockHeader().Time.Add(poolParams.WithdrawalPeriod)
	withdrawal := types.NewWithdrawal(id, from, amount)
	k.InsertWithdrawalQueue(ctx, withdrawal, completionTime)

	// update the pool - update or create collateral entry

	collateral, found := k.GetCollateral(ctx, pool, from)
	if !found {
		return types.ErrNoCollateralFound
	}
	collateral.Amount = collateral.Amount.Sub(amount)
	k.SetCollateral(ctx, pool, from, collateral)
	pool.TotalCollateral = pool.TotalCollateral.Sub(amount)
	k.SetPool(ctx, pool)
	return nil
}
