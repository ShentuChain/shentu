package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/shentufoundation/shentu/v2/x/shield/types"
)

var _ types.QueryServer = Keeper{}

// Provider queries a provider given the address.
func (k Keeper) Provider(c context.Context, req *types.QueryProviderRequest) (*types.QueryProviderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}

	provider, found := k.GetProvider(ctx, address)
	if !found {
		return nil, types.ErrProviderNotFound
	}

	return &types.QueryProviderResponse{Provider: provider}, nil
}

// Providers queries all providers.
func (k Keeper) Providers(c context.Context, req *types.QueryProvidersRequest) (*types.QueryProvidersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryProvidersResponse{Providers: k.GetAllProviders(ctx)}, nil
}

// ShieldStatus queries the global status of the shield module.
func (k Keeper) ShieldStatus(c context.Context, req *types.QueryShieldStatusRequest) (*types.QueryShieldStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	grs, err := k.GetRemainingServiceFees(ctx)
	if err != nil {
		return nil, err
	}
	return &types.QueryShieldStatusResponse{
		RemainingServiceFees: grs,
	}, nil
}
