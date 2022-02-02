package keeper

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/certikfoundation/shentu/v2/x/shield/types"
)

// PurchaseShield purchases shield of a pool with standard fee rate.
func (k Keeper) PurchaseShield(ctx sdk.Context, poolID uint64, amount sdk.Coins, description string, purchaser sdk.AccAddress) (types.Purchase, error) {
	poolParams := k.GetPoolParams(ctx)
	if poolParams.MinShieldPurchase.IsAnyGT(amount) {
		return types.Purchase{}, types.ErrPurchaseTooSmall
	}
	bondDenom := k.BondDenom(ctx)
	if amount.AmountOf(bondDenom).Equal(sdk.ZeroInt()) {
		return types.Purchase{}, types.ErrInsufficientStaking
	}
	pool, found := k.GetPool(ctx, poolID)
	if !found {
		return types.Purchase{}, types.ErrNoPoolFound
	}
	if !pool.Active {
		return types.Purchase{}, types.ErrPoolInactive
	}
	if amount.Empty() {
		return types.Purchase{}, types.ErrNoShield
	}

	sp, err := k.AddStaking(ctx, poolID, purchaser, description, amount)

	return sp, err
}

// SetNextPurchaseID sets the latest pool ID to store.
func (k Keeper) SetNextPurchaseID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.LittleEndian.PutUint64(bz, id)
	store.Set(types.GetNextPurchaseIDKey(), bz)
}

// GetNextPurchaseID gets the latest pool ID from store.
func (k Keeper) GetNextPurchaseID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	opBz := store.Get(types.GetNextPurchaseIDKey())
	return binary.LittleEndian.Uint64(opBz)
}

func (k Keeper) GetPurchaserPurchases(ctx sdk.Context, address sdk.AccAddress) (res []types.Purchase) {
	purchases := k.GetAllPurchase(ctx)

	for _, p := range purchases {
		if p.Purchaser == address.String() {
			res = append(res, p)
		}
	}
	return
}

func (k Keeper) GetPoolPurchases(ctx sdk.Context, poolID uint64) (res []types.Purchase) {
	purchases := k.GetAllPurchase(ctx)

	for _, p := range purchases {
		if p.PoolId == poolID {
			res = append(res, p)
		}
	}
	return
}

// DistributeFees distributes rewards for current block plus leftover rewards for last block.
func (k Keeper) DistributeFees(ctx sdk.Context) {
	serviceFees := sdk.NewDecCoins()

	// TODO: Add support for any denoms.

	// Add leftover block service fees from last block
	remainingServiceFees := k.GetRemainingServiceFees(ctx)
	serviceFees = serviceFees.Add(remainingServiceFees...)

	// Add block service fees that need to be distributed for this block.
	blockServiceFees := k.GetBlockServiceFees(ctx)
	serviceFees = serviceFees.Add(blockServiceFees...)
	k.DeleteBlockServiceFees(ctx)

	// Distribute service fees.
	totalCollateral := k.GetTotalCollateral(ctx)
	providers := k.GetAllProviders(ctx)
	for _, provider := range providers {
		providerAddr, err := sdk.AccAddressFromBech32(provider.Address)
		if err != nil {
			panic(err)
		}

		// fees * providerCollateral / totalCollateral
		nativeFees := serviceFees.MulDec(sdk.NewDecFromInt(provider.Collateral).QuoInt(totalCollateral))
		provider.Rewards = provider.Rewards.Add(nativeFees...)
		k.SetProvider(ctx, providerAddr, provider)
		serviceFees = serviceFees.Sub(nativeFees)
	}

	// Store remaining block reward as new leftover
	k.SetRemainingServiceFees(ctx, serviceFees)
}

func (k Keeper) RecoverPurchases(ctx sdk.Context) {
	bondDenom := k.BondDenom(ctx)
	k.IteratePurchases(ctx, func(purchase types.Purchase) bool {
		var updated []types.RecoveringEntry
		for _, e := range purchase.RecoveringEntries {
			if e.RecoverTime.Before(ctx.BlockTime()) {
				purchase.Shield = purchase.Shield.Add(e.Amount.AmountOf(bondDenom))
			} else {
				updated = append(updated, e)
			}
		}
		purchase.RecoveringEntries = updated
		k.SetPurchase(ctx, purchase)
		return false
	})
}
