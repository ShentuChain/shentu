package shield

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/certikfoundation/shentu/x/shield/types"
)

// NewHandler creates an sdk.Handler for all the slashing type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case types.MsgCreatePool:
			return handleMsgCreatePool(ctx, msg, k)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", ModuleName, msg)
		}
	}
}

func handleMsgCreatePool(ctx sdk.Context, msg types.MsgCreatePool, k Keeper) (*sdk.Result, error) {
	_, err := k.CreatePool(ctx, msg.From, msg.Coverage, msg.Deposit, msg.Sponsor)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreatePool,
			sdk.NewAttribute(types.AttributeKeyCoverage, msg.Coverage.String()),
			sdk.NewAttribute(types.AttributeKeyDeposit, msg.Deposit.String()),
			sdk.NewAttribute(types.AttributeKeySponsor, msg.Sponsor),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.From.String()),
		),
	})
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func NewShieldClaimProposalHandler(k Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case types.ShieldClaimProposal:
			return handleShieldClaimProposal(ctx, k, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized shield proposal content type: %T", c)
		}
	}
}

func handleShieldClaimProposal(ctx sdk.Context, k Keeper, p types.ShieldClaimProposal) error {
	return nil
}
