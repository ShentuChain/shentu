package keeper

import (
	"encoding/binary"
	"time"

	gogotypes "github.com/gogo/protobuf/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/certikfoundation/shentu/x/shield/types"
)

// Keeper implements the shield keeper.
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        codec.BinaryMarshaler
	ak         types.AccountKeeper
	bk         types.BankKeeper
	sk         types.StakingKeeper
	gk         types.GovKeeper
	paramSpace types.ParamSubspace
}

// NewKeeper creates a shield keeper.
func NewKeeper(cdc codec.BinaryMarshaler, shieldStoreKey sdk.StoreKey, ak types.AccountKeeper, bk types.BankKeeper,
	sk types.StakingKeeper, gk types.GovKeeper, paramSpace types.ParamSubspace) Keeper {
	return Keeper{
		storeKey:   shieldStoreKey,
		cdc:        cdc,
		ak:         ak,
		bk:         bk,
		sk:         sk,
		gk:         gk,
		paramSpace: paramSpace,
	}
}

// GetValidator returns info of a validator given its operator address.
func (k Keeper) GetValidator(ctx sdk.Context, addr sdk.ValAddress) (stakingtypes.ValidatorI, bool) {
	return k.sk.GetValidator(ctx, addr)
}

// SetLatestPoolID sets the latest pool ID to store.
func (k Keeper) SetNextPoolID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.LittleEndian.PutUint64(bz, id)
	store.Set(types.GetNextPoolIDKey(), bz)
}

// GetNextPoolID gets the latest pool ID from store.
func (k Keeper) GetNextPoolID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	opBz := store.Get(types.GetNextPoolIDKey())
	return binary.LittleEndian.Uint64(opBz)
}

// GetPoolBySponsor search store for a pool object with given pool ID.
func (k Keeper) GetPoolBySponsor(ctx sdk.Context, sponsor string) (types.Pool, bool) {
	ret := types.Pool{}
	k.IterateAllPools(ctx, func(pool types.Pool) bool {
		if pool.Sponsor == sponsor {
			ret = pool
			return true
		} else {
			return false
		}
	})
	if ret.Id == 0 {
		return ret, false
	}
	return ret, true
}

// BondDenom returns staking bond denomination.
func (k Keeper) BondDenom(ctx sdk.Context) string {
	return k.sk.BondDenom(ctx)
}

// GetVotingParams returns gov keeper's voting params.
func (k Keeper) GetVotingParams(ctx sdk.Context) govtypes.VotingParams {
	return k.gk.GetVotingParams(ctx)
}

// SetLastUpdateTime sets the last update time.
// Last update time will be set when the first purchase is made or distributing service fees.
func (k Keeper) SetLastUpdateTime(ctx sdk.Context, prevUpdateTime time.Time) {
	store := ctx.KVStore(k.storeKey)

	timeProto, err := gogotypes.TimestampProto(prevUpdateTime)
	if err != nil {
		panic(err)
	}
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(timeProto)
	store.Set(types.GetLastUpdateTimeKey(), bz)
}

// GetLastUpdateTime returns the last update time.
func (k Keeper) GetLastUpdateTime(ctx sdk.Context) (time.Time, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetLastUpdateTimeKey())
	if bz == nil {
		return time.Time{}, false
	}

	var protoTimestamp *gogotypes.Timestamp
	timestamp, err := gogotypes.TimestampFromProto(protoTimestamp)
	if err != nil {
		panic(err)
	}
	return timestamp, true
}
