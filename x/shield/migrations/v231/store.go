package v231

import (
	"github.com/gogo/protobuf/grpc"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/certikfoundation/shentu/v2/x/shield/types"
	"github.com/certikfoundation/shentu/v2/x/shield/types/v1alpha1"
	"github.com/certikfoundation/shentu/v2/x/shield/types/v1beta1"
)

const (
	stakingParamsPath = "/cosmos.staking.v1beta1.Query/Params"
)

func migratePools(store sdk.KVStore, cdc codec.BinaryCodec) error {
	oldStore := prefix.NewStore(store, types.PoolKey)

	oldStoreIter := oldStore.Iterator(nil, nil)
	defer oldStoreIter.Close()

	for ; oldStoreIter.Valid(); oldStoreIter.Next() {
		var oldPool v1alpha1.Pool
		cdc.MustUnmarshal(oldStoreIter.Value(), &oldPool)

		newPool := v1beta1.Pool{
			Id:          oldPool.Id,
			Description: oldPool.Description,
			SponsorAddr: oldPool.SponsorAddr,
			Active:      oldPool.Active,
			Shield:      oldPool.Shield,
			ShieldRate:  v1beta1.DefaultShieldRate,
		}

		newPoolBz := cdc.MustMarshal(&newPool)
		store.Set(oldStoreIter.Key(), newPoolBz)
	}
	return nil
}

func resolvePurchases(store sdk.KVStore, cdc codec.BinaryCodec, bondDenom string) error {
	oldStore := prefix.NewStore(store, types.PurchaseKey)

	oldStoreIter := oldStore.Iterator(nil, nil)
	defer oldStoreIter.Close()

	for ; oldStoreIter.Valid(); oldStoreIter.Next() {
		oldStore.Delete(oldStoreIter.Key())
	}

	queueStore := prefix.NewStore(store, types.PurchaseQueueKey)

	queueStoreIter := queueStore.Iterator(nil, nil)
	defer queueStoreIter.Close()

	for ; queueStoreIter.Valid(); queueStoreIter.Next() {
		oldStore.Delete(queueStoreIter.Key())
	}

	listStore := prefix.NewStore(store, types.PurchaseListKey)

	listStoreIter := listStore.Iterator(nil, nil)
	defer listStoreIter.Close()

	total := sdk.ZeroDec()
	for ; listStoreIter.Valid(); listStoreIter.Next() {
		var pl v1alpha1.PurchaseList
		cdc.MustUnmarshal(oldStoreIter.Value(), &pl)

		for _, e := range pl.Entries {
			total = total.Add(e.ServiceFees.Native.AmountOf(bondDenom))
		}
		oldStore.Delete(listStoreIter.Key())
	}
	var reserve v1beta1.Reserve
	reserveBz := store.Get(types.ReserveKey)
	cdc.MustUnmarshal(reserveBz, &reserve)
	reserve.Amount = reserve.Amount.Add(total.TruncateInt())
	reserveBz = cdc.MustMarshal(&reserve)
	store.Set(types.ReserveKey, reserveBz)

	return nil
}

func deleteUnusedStores(store sdk.KVStore, cdc codec.BinaryCodec) error {
	store.Delete(types.GetNextPurchaseIDKey())
	store.Delete(types.GetLastUpdateTimeKey())
	return nil
}

func resolveReimbursements(store sdk.KVStore, cdc codec.BinaryCodec, bondDenom string) error {
	var reserve v1beta1.Reserve
	reserveBz := store.Get(types.ReserveKey)
	cdc.MustUnmarshal(reserveBz, &reserve)

	oldStore := prefix.NewStore(store, types.ReimbursementKey)

	oldStoreIter := oldStore.Iterator(nil, nil)
	defer oldStoreIter.Close()
	for ; oldStoreIter.Valid(); oldStoreIter.Next() {
		var reimbursement v1alpha1.Reimbursement
		cdc.MustUnmarshal(oldStoreIter.Value(), &reimbursement)

		reserve.Amount = reserve.Amount.Add(reimbursement.Amount.AmountOf(bondDenom))
		oldStore.Delete(oldStoreIter.Key())
	}
	reserveBz = cdc.MustMarshal(&reserve)
	store.Set(types.ReserveKey, reserveBz)
	return nil
}

func migrateparams(store sdk.KVStore, cdc codec.BinaryCodec, ps types.ParamSubspace) error {

	return nil
}

func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, cdc codec.BinaryCodec, paramSpace types.ParamSubspace, queryServer grpc.Server) error {
	bondDenom, err := getBondDenom(ctx, queryServer)
	if err != nil {
		return err
	}
	store := ctx.KVStore(storeKey)
	err = migratePools(store, cdc)
	if err != nil {
		return err
	}

	err = resolvePurchases(store, cdc, bondDenom)
	if err != nil {
		return err
	}

	err = resolveReimbursements(store, cdc, bondDenom)
	if err != nil {
		return err
	}

	err = migrateparams(store, cdc, paramSpace)
	if err != nil {
		return err
	}

	return deleteUnusedStores(store, cdc)
}
