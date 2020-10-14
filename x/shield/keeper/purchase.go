package keeper

import (
	"encoding/binary"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/certikfoundation/shentu/x/shield/types"
)

// SetPurchaseList sets a purchase list.
func (k Keeper) SetPurchaseList(ctx sdk.Context, purchaseList types.PurchaseList) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(purchaseList)
	store.Set(types.GetPurchaseListKey(purchaseList.PoolID, purchaseList.Purchaser), bz)
}

// AddPurchase sets a purchase of shield.
func (k Keeper) AddPurchase(ctx sdk.Context, poolID uint64, purchaser sdk.AccAddress, purchase types.Purchase) {
	purchaseList, found := k.GetPurchaseList(ctx, poolID, purchaser)
	if !found {
		purchaseList = types.NewPurchaseList(poolID, purchaser, []types.Purchase{purchase})
	} else {
		purchaseList.Entries = append(purchaseList.Entries, purchase)
	}
	k.SetPurchaseList(ctx, purchaseList)
}

// GetPurchaseList gets a purchase from store by txhash.
func (k Keeper) GetPurchaseList(ctx sdk.Context, poolID uint64, purchaser sdk.AccAddress) (types.PurchaseList, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetPurchaseListKey(poolID, purchaser))
	if bz != nil {
		var purchase types.PurchaseList
		k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &purchase)
		return purchase, true
	}
	return types.PurchaseList{}, false
}

// GetPurchase gets a purchase out of a purchase list
func GetPurchase(purchaseList types.PurchaseList, purchaseID uint64) (types.Purchase, bool) {
	for _, entry := range purchaseList.Entries {
		if entry.PurchaseID == purchaseID {
			return entry, true
		}
	}
	return types.Purchase{}, false
}

// DeletePurchaseList deletes a purchase of shield.
func (k Keeper) DeletePurchaseList(ctx sdk.Context, poolID uint64, purchaser sdk.AccAddress) error {
	store := ctx.KVStore(k.storeKey)
	_, found := k.GetPurchaseList(ctx, poolID, purchaser)
	if !found {
		return types.ErrPurchaseNotFound
	}
	store.Delete(types.GetPurchaseListKey(poolID, purchaser))
	return nil
}

// DequeuePurchase dequeues a purchase from the purchase queue
func (k Keeper) DequeuePurchase(ctx sdk.Context, purchaseList types.PurchaseList, endTime time.Time) {
	timeslice := k.GetPurchaseQueueTimeSlice(ctx, endTime)
	for i, ppPair := range timeslice {
		if (purchaseList.PoolID == ppPair.PoolID) && purchaseList.Purchaser.Equals(ppPair.Purchaser) {
			if len(timeslice) > 1 {
				timeslice = append(timeslice[:i], timeslice[i+1:]...)
				k.SetPurchaseQueueTimeSlice(ctx, endTime, timeslice)
				return
			}
			ctx.KVStore(k.storeKey).Delete(types.GetPurchaseCompletionTimeKey(endTime))
			return
		}
	}
}

// PurchaseShield purchases shield of a pool.
func (k Keeper) PurchaseShield(
	ctx sdk.Context, poolID uint64, shield sdk.Coins, description string, purchaser sdk.AccAddress,
) (types.Purchase, error) {
	pool, err := k.GetPool(ctx, poolID)
	if err != nil {
		return types.Purchase{}, err
	}
	poolParams := k.GetPoolParams(ctx)
	claimParams := k.GetClaimProposalParams(ctx)

	// check preconditions
	if err = k.PoolShieldAvailable(ctx, pool); err != nil {
		return types.Purchase{}, err
	}

	shieldAmt := shield.AmountOf(k.sk.BondDenom(ctx))
	if shieldAmt.GT(pool.Available) {
		return types.Purchase{}, types.ErrNotEnoughShield
	}
	if pool.SponsorAddr.Equals(purchaser) {
		return types.Purchase{}, types.ErrSponsorPurchase
	}

	// send tokens to shield module account
	shieldDec := sdk.NewDecCoinsFromCoins(shield...)
	premium, _ := shieldDec.MulDec(poolParams.ShieldFeesRate).TruncateDecimal()
	if err := k.DepositNativePremium(ctx, premium, purchaser); err != nil {
		return types.Purchase{}, err
	}

	// update pool premium, shield and available
	premiumMixedDec := types.NewMixedDecCoins(sdk.NewDecCoinsFromCoins(premium...), sdk.DecCoins{})
	pool.Premium = pool.Premium.Add(premiumMixedDec)
	pool.Shield = pool.Shield.Add(shield...)
	pool.Available = pool.Available.Sub(shieldAmt)
	k.SetPool(ctx, pool)

	// set purchase
	protectionEndTime := ctx.BlockTime().Add(poolParams.ProtectionPeriod)
	claimPeriodEndTime := ctx.BlockTime().Add(claimParams.ClaimPeriod)
	purchaseID := k.GetNextPurchaseID(ctx)
	purchase := types.NewPurchase(purchaseID, shield, ctx.BlockHeight(), protectionEndTime,
		claimPeriodEndTime, claimPeriodEndTime, description)
	purchaseList := types.NewPurchaseList(poolID, purchaser, []types.Purchase{purchase})
	k.AddPurchase(ctx, poolID, purchaser, purchase)
	k.InsertPurchaseQueue(ctx, purchaseList, claimPeriodEndTime)
	k.SetNextPurchaseID(ctx, purchaseID+1)

	return purchase, nil
}

// IterateAllPurchases iterates over the all the stored purchases and performs a callback function.
func (k Keeper) IterateAllPurchases(ctx sdk.Context, callback func(purchase types.Purchase) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PurchaseListKey)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var purchase types.Purchase
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iterator.Value(), &purchase)

		if callback(purchase) {
			break
		}
	}
}

// TODO improve the performance
// RemoveExpiredPurchases removes purchases whose claim period end time is before current block time.
func (k Keeper) RemoveExpiredPurchases(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	iterator := k.PurchaseQueueIterator(ctx, ctx.BlockTime())
	bondDenom := k.sk.BondDenom(ctx)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var timeslice []types.PPPair
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iterator.Value(), &timeslice)
		for _, ppPair := range timeslice {
			purchaseList, _ := k.GetPurchaseList(ctx, ppPair.PoolID, ppPair.Purchaser)

			for i := 0; i < len(purchaseList.Entries); {
				entry := purchaseList.Entries[i]
				if entry.ExpirationTime.Before(ctx.BlockTime()) {
					purchaseList.Entries = append(purchaseList.Entries[:i], purchaseList.Entries[i+1:]...)
					pool, err := k.GetPool(ctx, purchaseList.PoolID)
					if err != nil {
						// skip purchases of closed pools
						continue
					}
					pool.Available = pool.Available.Add(entry.Shield.AmountOf(bondDenom))
					pool.Shield = pool.Shield.Sub(entry.Shield)
					k.SetPool(ctx, pool)
					continue
				}
				i++
			}
			if len(purchaseList.Entries) == 0 {
				k.DeletePurchaseList(ctx, purchaseList.PoolID, purchaseList.Purchaser)
			} else {
				k.SetPurchaseList(ctx, purchaseList)
			}
		}
		store.Delete(iterator.Key())
	}
}

// GetOnesPurchases returns a purchaser's all purchases.
func (k Keeper) GetOnesPurchases(ctx sdk.Context, address sdk.AccAddress) (res []types.PurchaseList) {
	pools := k.GetAllPools(ctx)
	for _, pool := range pools {
		pList, found := k.GetPurchaseList(ctx, pool.PoolID, address)
		if !found {
			continue
		}
		res = append(res, pList)
	}
	return
}

// GetPoolPurchaseLists returns a all purchases in a given pool.
func (k Keeper) GetPoolPurchaseLists(ctx sdk.Context, poolID uint64) (purchases []types.PurchaseList) {
	k.IteratePoolPurchaseLists(ctx, poolID, func(purchaseList types.PurchaseList) bool {
		if purchaseList.PoolID == poolID {
			purchases = append(purchases, purchaseList)
		}
		return false
	})
	return purchases
}

// IteratePurchaseLists iterates through purchase lists in a pool
func (k Keeper) IteratePurchaseLists(ctx sdk.Context, callback func(purchase types.PurchaseList) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PurchaseListKey)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var purchaseList types.PurchaseList
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iterator.Value(), &purchaseList)

		if callback(purchaseList) {
			break
		}
	}
}

// IteratePoolPurchaseLists iterates through purchases in a pool
func (k Keeper) IteratePoolPurchaseLists(ctx sdk.Context, poolID uint64, callback func(purchaseList types.PurchaseList) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.LittleEndian.PutUint64(bz, poolID)
	iterator := sdk.KVStorePrefixIterator(store, append(types.PurchaseListKey, bz...))

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var purchaseList types.PurchaseList
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iterator.Value(), &purchaseList)

		if callback(purchaseList) {
			break
		}
	}
}

// GetAllPurchaseLists retrieves all purchases.
func (k Keeper) GetAllPurchaseLists(ctx sdk.Context) (purchases []types.PurchaseList) {
	k.IteratePurchaseLists(ctx, func(purchase types.PurchaseList) bool {
		purchases = append(purchases, purchase)
		return false
	})
	return
}

func (k Keeper) InsertPurchaseQueue(ctx sdk.Context, purchaseList types.PurchaseList, endTime time.Time) {
	timeSlice := k.GetPurchaseQueueTimeSlice(ctx, endTime)

	ppPair := types.PPPair{PoolID: purchaseList.PoolID, Purchaser: purchaseList.Purchaser}
	if len(timeSlice) == 0 {
		k.SetPurchaseQueueTimeSlice(ctx, endTime, []types.PPPair{ppPair})
		return
	}
	timeSlice = append(timeSlice, ppPair)
	k.SetPurchaseQueueTimeSlice(ctx, endTime, timeSlice)
}

// GetPurchaseQueueTimeSlice gets a specific purchase queue timeslice,
// which is a slice of purchases corresponding to a given time.
func (k Keeper) GetPurchaseQueueTimeSlice(ctx sdk.Context, timestamp time.Time) []types.PPPair {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetPurchaseCompletionTimeKey(timestamp))
	if bz == nil {
		return []types.PPPair{}
	}
	var ppPairs []types.PPPair
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &ppPairs)
	return ppPairs
}

func (k Keeper) SetPurchaseQueueTimeSlice(ctx sdk.Context, timestamp time.Time, ppPairs []types.PPPair) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(ppPairs)
	store.Set(types.GetPurchaseCompletionTimeKey(timestamp), bz)
}

// PurchaseQueueIterator returns all the purchase queue timeslices from time 0 until endTime
func (k Keeper) PurchaseQueueIterator(ctx sdk.Context, endTime time.Time) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return store.Iterator(types.PurchaseQueueKey,
		sdk.InclusiveEndBytes(types.GetPurchaseCompletionTimeKey(endTime)))
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

// PoolShieldAvailable checks if there is enough time for a purchase.
func (k Keeper) PoolShieldAvailable(ctx sdk.Context, pool types.Pool) error {
	if !pool.Active {
		return types.ErrPoolInactive
	}
	claimParams := k.GetClaimProposalParams(ctx)
	if pool.EndTime.Before(ctx.BlockTime().Add(claimParams.ClaimPeriod).Add(k.GetVotingParams(ctx).VotingPeriod * 2)) {
		return types.ErrPoolLifeTooShort
	}
	return nil
}
