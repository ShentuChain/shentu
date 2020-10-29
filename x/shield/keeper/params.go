package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/certikfoundation/shentu/x/shield/types"
)

// SetPoolParams sets parameters subspace for shield pool parameters.
func (k Keeper) SetPoolParams(ctx sdk.Context, poolParams types.PoolParams) {
	k.paramSpace.Set(ctx, types.ParamStoreKeyPoolParams, &poolParams)
}

// GetPoolParams returns shield pool parameters.
func (k Keeper) GetPoolParams(ctx sdk.Context) types.PoolParams {
	var poolParams types.PoolParams
	k.paramSpace.Get(ctx, types.ParamStoreKeyPoolParams, &poolParams)
	return poolParams
}

// SetClaimProposalParams sets parameters subspace for shield claim proposal parameters.
func (k Keeper) SetClaimProposalParams(ctx sdk.Context, claimProposalParams types.ClaimProposalParams) {
	k.paramSpace.Set(ctx, types.ParamStoreKeyClaimProposalParams, &claimProposalParams)
}

// GetClaimProposalParams returns shield claim proposal parameters.
func (k Keeper) GetClaimProposalParams(ctx sdk.Context) types.ClaimProposalParams {
	var claimProposalParams types.ClaimProposalParams
	k.paramSpace.Get(ctx, types.ParamStoreKeyClaimProposalParams, &claimProposalParams)
	return claimProposalParams
}

// GetStakingPurchaseRate returns shield to staked rate.
func (k Keeper) GetStakingPurchaseRate(ctx sdk.Context) (rate sdk.Dec) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetStakingPurchaseRateKey())
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &rate)
	return
}

// SetStakingPurchaseRate sets shield to staked rate.
func (k Keeper) SetStakingPurchaseRate(ctx sdk.Context, rate sdk.Dec) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(rate)
	store.Set(types.GetStakingPurchaseRateKey(), bz)
}
