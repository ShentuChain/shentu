package keeper

import (
	"context"
	"crypto/rand"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/crypto/ecies"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/shentufoundation/shentu/v2/x/bounty/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) CreateProgram(goCtx context.Context, msg *types.MsgCreateProgram) (*types.MsgCreateProgramResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, err := sdk.AccAddressFromBech32(msg.CreatorAddress)
	if err != nil {
		return nil, err
	}

	err = k.bk.SendCoinsFromAccountToModule(ctx, creatorAddr, types.ModuleName, msg.Deposit)
	if err != nil {
		return nil, err
	}

	nextID := k.GetNextProgramID(ctx)

	program := types.Program{
		ProgramId:         nextID,
		CreatorAddress:    msg.CreatorAddress,
		SubmissionEndTime: msg.SubmissionEndTime,
		Description:       msg.Description,
		EncryptionKey:     msg.EncryptionKey,
		Deposit:           msg.Deposit,
		CommissionRate:    msg.CommissionRate,
	}

	k.SetProgram(ctx, program)

	k.SetNextProgramID(ctx, nextID+1)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateProgram,
			sdk.NewAttribute(types.AttributeKeyProgramID, strconv.FormatUint(program.ProgramId, 10)),
			sdk.NewAttribute(types.AttributeKeyDeposit, sdk.NewCoins(msg.Deposit...).String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.CreatorAddress),
		),
	})

	return &types.MsgCreateProgramResponse{ProgramId: nextID}, nil
}

func (k msgServer) SubmitFinding(goCtx context.Context, msg *types.MsgSubmitFinding) (*types.MsgSubmitFindingResponse, error) {
	_, err := sdk.AccAddressFromBech32(msg.SubmitterAddress)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	program, isExist := k.GetProgram(ctx, msg.Pid)
	if !isExist {
		return nil, fmt.Errorf("no program id:%d", msg.Pid)
	}

	if !program.Active {
		return nil, fmt.Errorf("program id:%d is closed", msg.Pid)
	}

	var eciesEncKey ecies.PublicKey
	err = k.cdc.UnpackAny(program.EncryptionKey, &eciesEncKey)
	if err != nil {
		return nil, fmt.Errorf("EncryptionKey error")
	}

	encryptedDesc, err := ecies.Encrypt(rand.Reader, &eciesEncKey, []byte(msg.Desc), nil, nil)
	if err != nil {
		return nil, err
	}

	encryptedPoc, err := ecies.Encrypt(rand.Reader, &eciesEncKey, []byte(msg.Poc), nil, nil)
	if err != nil {
		return nil, err
	}

	nextID := k.GetNextFindingID(ctx)

	var descAny *codectypes.Any
	var pocAny *codectypes.Any

	encDesc := types.EciesEncryptedDesc{
		Desc: encryptedDesc,
	}
	if descAny, err = codectypes.NewAnyWithValue(&encDesc); err != nil {
		return nil, err
	}

	encPoc := types.EciesEncryptedPoc{
		Poc: encryptedPoc,
	}
	if pocAny, err = codectypes.NewAnyWithValue(&encPoc); err != nil {
		return nil, err
	}

	finding := types.Finding{
		FindingId:        nextID,
		Title:            msg.Title,
		EncryptedDesc:    descAny,
		Pid:              msg.Pid,
		SeverityLevel:    msg.SeverityLevel,
		EncryptedPoc:     pocAny,
		SubmitterAddress: msg.SubmitterAddress,
	}

	err = k.AppendFidToFidList(ctx, msg.Pid, nextID)
	if err != nil {
		return nil, err
	}

	k.SetFinding(ctx, finding)
	k.SetNextFindingID(ctx, nextID+1)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSubmitFinding,
			sdk.NewAttribute(types.AttributeKeyFindingID, strconv.FormatUint(finding.FindingId, 10)),
			sdk.NewAttribute(types.AttributeKeyProgramID, strconv.FormatUint(finding.Pid, 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.SubmitterAddress),
		),
	})

	return &types.MsgSubmitFindingResponse{
		Fid: finding.FindingId,
	}, nil
}

func (k msgServer) WithdrawalFinding(goCtx context.Context, msg *types.MsgWithdrawalFinding) (*types.MsgWithdrawalFindingResponse, error) {
	fromAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	_, err = k.Keeper.WithdrawalFinding(ctx, fromAddr, msg.Fid)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeWithdrawalFinding,
			sdk.NewAttribute(types.AttributeKeyFindingID, strconv.FormatUint(msg.Fid, 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.From),
		),
	})

	return &types.MsgWithdrawalFindingResponse{}, nil
}

func (k msgServer) ReactivateFinding(goCtx context.Context, msg *types.MsgReactivateFinding) (*types.MsgReactivateFindingResponse, error) {
	fromAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	_, err = k.Keeper.ReactivateFinding(ctx, fromAddr, msg.Fid)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeReactivateFinding,
			sdk.NewAttribute(types.AttributeKeyFindingID, strconv.FormatUint(msg.Fid, 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.From),
		),
	})

	return &types.MsgReactivateFindingResponse{}, nil
}
