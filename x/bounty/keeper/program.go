package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/shentufoundation/shentu/v2/x/bounty/types"
)

func (k Keeper) GetProgram(ctx sdk.Context, id string) (types.Program, bool) {
	store := ctx.KVStore(k.storeKey)

	pBz := store.Get(types.GetProgramKey(id))
	if pBz == nil {
		return types.Program{}, false
	}

	var program types.Program
	k.cdc.MustUnmarshal(pBz, &program)
	return program, true
}

func (k Keeper) GetAllPrograms(ctx sdk.Context) []types.Program {
	store := ctx.KVStore(k.storeKey)

	var programs []types.Program
	var program types.Program

	iterator := sdk.KVStorePrefixIterator(store, types.ProgramsKey)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		k.cdc.MustUnmarshal(iterator.Value(), &program)
		programs = append(programs, program)
	}
	return programs
}

func (k Keeper) SetProgram(ctx sdk.Context, program types.Program) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&program)
	store.Set(types.GetProgramKey(program.ProgramId), bz)
}

func (k Keeper) OpenProgram(ctx sdk.Context, caller sdk.AccAddress, id string) error {
	program, found := k.GetProgram(ctx, id)
	if !found {
		return types.ErrProgramNotExists
	}

	// Check the permissions. Only the cert address can operate.
	if !k.certKeeper.IsCertifier(ctx, caller) {
		return sdkerrors.Wrapf(govtypes.ErrInvalidVote, "%s is not a certified identity", caller.String())
	}

	if program.Status != types.ProgramStatusInactive {
		return types.ErrProgramNotInactive
	}

	program.Status = types.ProgramStatusActive
	k.SetProgram(ctx, program)
	return nil
}

func (k Keeper) CloseProgram(ctx sdk.Context, caller sdk.AccAddress, id string) error {
	program, found := k.GetProgram(ctx, id)
	if !found {
		return types.ErrProgramNotExists
	}

	// Check the permissions. Only the admin of the program or cert address can operate.
	if program.AdminAddress != caller.String() || !k.certKeeper.IsCertifier(ctx, caller) {
		return types.ErrFindingOperatorNotAllowed
	}

	if program.Status != types.ProgramStatusActive {
		return types.ErrProgramNotActive
	}

	program.Status = types.ProgramStatusClosed
	k.SetProgram(ctx, program)
	return nil
}
