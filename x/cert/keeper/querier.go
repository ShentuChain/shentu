package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/certikfoundation/shentu/x/cert/types"
)

// NewQuerier is the module level router for state queries.
func NewQuerier(keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case types.QueryCertifier:
			return queryCertifier(ctx, path[1:], keeper, legacyQuerierCdc)
		case types.QueryCertifiers:
			return queryCertifiers(ctx, path[1:], keeper, legacyQuerierCdc)
		case types.QueryCertifierByAlias:
			return queryCertifierByAlias(ctx, path[1:], keeper, legacyQuerierCdc)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown cert query endpoint")
		}
	}
}

func validatePathLength(path []string, length int) error {
	if len(path) != length {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "Expecting %d args. Found %d.", length, len(path))
	}
	return nil
}

// queryCertifier returns information of a certifier.
func queryCertifier(ctx sdk.Context, path []string, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) (res []byte, err error) {
	if err := validatePathLength(path, 1); err != nil {
		return nil, err
	}
	certifierAddress, err := sdk.AccAddressFromBech32(path[0])
	if err != nil {
		return nil, err
	}
	certifier, err := keeper.GetCertifier(ctx, certifierAddress)
	if err != nil {
		return nil, err
	}
	res, err = codec.MarshalJSONIndent(legacyQuerierCdc, certifier)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}

func queryCertifiers(ctx sdk.Context, path []string, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) (res []byte, err error) {
	err = validatePathLength(path, 0)
	if err != nil {
		return nil, err
	}
	res, err = codec.MarshalJSONIndent(legacyQuerierCdc, types.QueryResCertifiers{Certifiers: keeper.GetAllCertifiers(ctx)})
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}

// queryCertifierByAlias returns information of a certifier from certifier alias
func queryCertifierByAlias(ctx sdk.Context, path []string, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) (res []byte, err error) {
	if err := validatePathLength(path, 1); err != nil {
		return nil, err
	}
	certifierAlias := path[0]
	certifier, err := keeper.GetCertifierByAlias(ctx, certifierAlias)
	if err != nil {
		return nil, err
	}
	res, err = codec.MarshalJSONIndent(legacyQuerierCdc, certifier)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}
