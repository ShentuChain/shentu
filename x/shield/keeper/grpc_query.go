package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/certikfoundation/shentu/v2/x/shield/types"
)

var _ types.QueryServer = Keeper{}

// Pool queries a pool based on the ID or sponsor.
func (q Keeper) Pool(c context.Context, req *types.QueryPoolRequest) (*types.QueryPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	// query by ID
	pool, found := q.GetPool(ctx, req.PoolId)
	if !found {
		return nil, status.Errorf(codes.NotFound, "pool under ID %d doesn't exist", req.PoolId)
	}

	return &types.QueryPoolResponse{Pool: pool}, nil
}

// Pool queries a pool based on the ID or sponsor.
func (q Keeper) Sponsor(c context.Context, req *types.QuerySponsorRequest) (*types.QuerySponsorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	// query by ID
	pool, found := q.GetPoolsBySponsor(ctx, req.Sponsor)
	if !found {
		return nil, status.Errorf(codes.NotFound, "there is no pool with sponsor %s", req.Sponsor)
	}

	return &types.QuerySponsorResponse{Pools: pool}, nil
}

// Pools queries all pools.
func (q Keeper) Pools(c context.Context, req *types.QueryPoolsRequest) (*types.QueryPoolsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryPoolsResponse{Pools: q.GetAllPools(ctx)}, nil
}

// Provider queries a provider given the address.
func (q Keeper) Provider(c context.Context, req *types.QueryProviderRequest) (*types.QueryProviderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}

	provider, found := q.GetProvider(ctx, address)
	if !found {
		return nil, types.ErrProviderNotFound
	}

	return &types.QueryProviderResponse{Provider: provider}, nil
}

// Providers queries all providers.
func (q Keeper) Providers(c context.Context, req *types.QueryProvidersRequest) (*types.QueryProvidersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryProvidersResponse{Providers: q.GetAllProviders(ctx)}, nil
}

// PoolParams queries pool parameters.
func (q Keeper) PoolParams(c context.Context, req *types.QueryPoolParamsRequest) (*types.QueryPoolParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryPoolParamsResponse{Params: q.GetPoolParams(ctx)}, nil
}

// ClaimParams queries claim proposal parameters.
func (q Keeper) ClaimParams(c context.Context, req *types.QueryClaimParamsRequest) (*types.QueryClaimParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryClaimParamsResponse{Params: q.GetClaimProposalParams(ctx)}, nil
}

// ShieldStatus queries the global status of the shield module.
func (q Keeper) ShieldStatus(c context.Context, req *types.QueryShieldStatusRequest) (*types.QueryShieldStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryShieldStatusResponse{
		TotalCollateral:         q.GetTotalCollateral(ctx),
		TotalShield:             q.GetTotalShield(ctx),
		TotalWithdrawing:        q.GetTotalWithdrawing(ctx),
		GlobalShieldStakingPool: q.GetGlobalStakingPool(ctx),
	}, nil
}

// ShieldStaking queries staked-for-shield for pool-purchaser pair.
func (q Keeper) Purchase(c context.Context, req *types.QueryPurchaseRequest) (*types.QueryPurchaseResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	purchaser, err := sdk.AccAddressFromBech32(req.Purchaser)
	if err != nil {
		return nil, err
	}
	shieldStaking, found := q.GetPurchase(ctx, req.PoolId, purchaser)
	if !found {
		return nil, types.ErrPurchaseNotFound
	}

	return &types.QueryPurchaseResponse{Purchase: shieldStaking}, nil
}

// Reserve queries all shield reserve amount.
func (q Keeper) Reserve(c context.Context, req *types.QueryReserveRequest) (*types.QueryReserveResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	return &types.QueryReserveResponse{Reserve: q.GetReserve(ctx)}, nil
}

// PoolPurchases queries for all purchases for a specific pool.
func (k Keeper) PoolPurchases(c context.Context, req *types.QueryPoolPurchasesRequest) (*types.QueryPurchasesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	res := k.GetPoolPurchases(ctx, req.PoolId)
	return &types.QueryPurchasesResponse{Purchases: res}, nil
}

// Purchases queries for all purchases.
func (k Keeper) Purchases(c context.Context, req *types.QueryAllPurchasesRequest) (*types.QueryPurchasesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	res := k.GetAllPurchase(ctx)
	return &types.QueryPurchasesResponse{Purchases: res}, nil
}

// Purchaser queries for information on a purchaser.
func (k Keeper) Purchaser(c context.Context, req *types.QueryPurchaserRequest) (*types.QueryPurchaserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	purchaser, err := sdk.AccAddressFromBech32(req.Purchaser)
	if err != nil {
		return nil, err
	}
	res := k.GetPurchaserPurchases(ctx, purchaser)

	shield := sdk.ZeroInt()
	deposit := sdk.ZeroInt()
	for _, p := range res {
		shield = shield.Add(p.Shield)
		deposit = deposit.Add(p.Amount)
	}
	return &types.QueryPurchaserResponse{
		Purchases:    res,
		TotalShield:  shield,
		TotalDeposit: deposit,
	}, nil
}
