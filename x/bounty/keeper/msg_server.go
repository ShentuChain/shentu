package keeper

import (
	"context"
	"fmt"
	"strconv"

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
		Active:            true,
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

	program, isExist := k.GetProgram(ctx, msg.ProgramId)
	if !isExist {
		return nil, fmt.Errorf("no program id:%d", msg.ProgramId)
	}

	if !program.Active {
		return nil, fmt.Errorf("program id:%d is closed", msg.ProgramId)
	}

	findingID := k.GetNextFindingID(ctx)

	finding := types.Finding{
		FindingId:        findingID,
		Title:            msg.Title,
		EncryptedDesc:    msg.EncryptedDesc,
		ProgramId:        msg.ProgramId,
		SeverityLevel:    msg.SeverityLevel,
		EncryptedPoc:     msg.EncryptedPoc,
		SubmitterAddress: msg.SubmitterAddress,
		FindingStatus:    types.FindingStatusUnConfirmed,
	}

	err = k.AppendFidToFidList(ctx, msg.ProgramId, findingID)
	if err != nil {
		return nil, err
	}

	k.SetFinding(ctx, finding)
	k.SetNextFindingID(ctx, findingID+1)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSubmitFinding,
			sdk.NewAttribute(types.AttributeKeyFindingID, strconv.FormatUint(finding.FindingId, 10)),
			sdk.NewAttribute(types.AttributeKeyProgramID, strconv.FormatUint(finding.ProgramId, 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.SubmitterAddress),
		),
	})

	return &types.MsgSubmitFindingResponse{
		FindingId: finding.FindingId,
	}, nil
}

func (k msgServer) HostAcceptFinding(goCtx context.Context, msg *types.MsgHostAcceptFinding) (*types.MsgHostAcceptFindingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	finding, err := k.hostProcess(ctx, msg.FindingId, msg.HostAddress, msg.EncryptedComment)
	if err != nil {
		return nil, err
	}

	finding.FindingStatus = types.FindingStatusValid
	k.SetFinding(ctx, *finding)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeAcceptFinding,
			sdk.NewAttribute(types.AttributeKeyFindingID, strconv.FormatUint(finding.FindingId, 10)),
			sdk.NewAttribute(types.AttributeKeyProgramID, strconv.FormatUint(finding.ProgramId, 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.HostAddress),
		),
	})

	return &types.MsgHostAcceptFindingResponse{}, nil
}

func (k msgServer) HostRejectFinding(goCtx context.Context, msg *types.MsgHostRejectFinding) (*types.MsgHostRejectFindingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	finding, err := k.hostProcess(ctx, msg.FindingId, msg.HostAddress, msg.EncryptedComment)
	if err != nil {
		return nil, err
	}

	finding.FindingStatus = types.FindingStatusInvalid
	k.SetFinding(ctx, *finding)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeRejectFinding,
			sdk.NewAttribute(types.AttributeKeyFindingID, strconv.FormatUint(finding.FindingId, 10)),
			sdk.NewAttribute(types.AttributeKeyProgramID, strconv.FormatUint(finding.ProgramId, 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.HostAddress),
		),
	})

	return &types.MsgHostRejectFindingResponse{}, nil
}

func (k msgServer) hostProcess(ctx sdk.Context, fid uint64, hostAddr string, encryptedCommentAny *codectypes.Any) (*types.Finding, error) {

	// get finding
	finding, isExist := k.GetFinding(ctx, fid)
	if !isExist {
		return nil, fmt.Errorf("no finding id:%d", fid)
	}
	// get program
	program, isExist := k.GetProgram(ctx, finding.ProgramId)
	if !isExist {
		return nil, fmt.Errorf("no program id:%d", finding.ProgramId)
	}
	if !program.Active {
		return nil, fmt.Errorf("program id:%d is closed", finding.ProgramId)
	}

	// only creator can update finding comment
	if program.CreatorAddress != hostAddr {
		return nil, fmt.Errorf("%s not the program creator, expect %s", hostAddr, program.CreatorAddress)
	}

	finding.EncryptedComment = encryptedCommentAny
	return &finding, nil
}

func (k msgServer) ReleaseFinding(goCtx context.Context, msg *types.MsgReleaseFinding) (*types.MsgReleaseFindingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get finding
	finding, isExist := k.GetFinding(ctx, msg.FindingId)
	if !isExist {
		return nil, fmt.Errorf("no finding id:%d", msg.FindingId)
	}
	// get program
	program, isExist := k.GetProgram(ctx, finding.ProgramId)
	if !isExist {
		return nil, fmt.Errorf("no program id:%d", finding.ProgramId)
	}
	if !program.Active {
		return nil, fmt.Errorf("program id:%d is closed", finding.ProgramId)
	}

	// only creator can update finding comment
	if program.CreatorAddress != msg.HostAddress {
		return nil, fmt.Errorf("%s not the program creator, expect %s", msg.HostAddress, program.CreatorAddress)
	}

	finding.Desc = msg.Desc
	finding.Poc = msg.Poc
	finding.Comment = msg.Comment

	k.SetFinding(ctx, finding)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeReleaseFinding,
			sdk.NewAttribute(types.AttributeKeyFindingID, strconv.FormatUint(finding.FindingId, 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.HostAddress),
		),
	})

	return &types.MsgReleaseFindingResponse{}, nil
}
