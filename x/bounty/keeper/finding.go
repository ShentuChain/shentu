package keeper

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/shentufoundation/shentu/v2/x/bounty/types"
)

func (k Keeper) GetFinding(ctx sdk.Context, id string) (types.Finding, bool) {
	store := ctx.KVStore(k.storeKey)

	findingData := store.Get(types.GetFindingKey(id))
	if findingData == nil {
		return types.Finding{}, false
	}

	var finding types.Finding
	k.cdc.MustUnmarshal(findingData, &finding)
	return finding, true
}

func (k Keeper) SetFinding(ctx sdk.Context, finding types.Finding) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&finding)
	store.Set(types.GetFindingKey(finding.FindingId), bz)
}

func (k Keeper) DeleteFinding(ctx sdk.Context, id string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetFindingKey(id))
}

func (k Keeper) SetPidFindingIDList(ctx sdk.Context, pid string, findingIds []string) error {
	store := ctx.KVStore(k.storeKey)
	bytes, err := StringsToBytes(findingIds)
	if err != nil {
		return err
	}
	store.Set(types.GetProgramIDFindingListKey(pid), bytes)
	return nil
}

func (k Keeper) GetPidFindingIDList(ctx sdk.Context, pid string) ([]string, error) {
	store := ctx.KVStore(k.storeKey)
	findingIDs := store.Get(types.GetProgramIDFindingListKey(pid))

	if findingIDs == nil {
		return []string{}, nil
	}

	findingIDList, err := BytesToStrings(findingIDs)
	if err != nil {
		return nil, err
	}
	return findingIDList, nil
}

func (k Keeper) AppendFidToFidList(ctx sdk.Context, pid, fid string) error {
	fids, err := k.GetPidFindingIDList(ctx, pid)
	if err != nil {
		return err
	}

	fids = append(fids, fid)
	err = k.SetPidFindingIDList(ctx, pid, fids)
	return err
}

func (k Keeper) DeleteFidFromFidList(ctx sdk.Context, pid, fid string) error {
	fids, err := k.GetPidFindingIDList(ctx, pid)
	if err != nil {
		return err
	}
	for idx, id := range fids {
		if id == fid {
			if len(fids) == 1 {
				// Delete fid list if empty
				store := ctx.KVStore(k.storeKey)
				store.Delete(types.GetProgramIDFindingListKey(pid))
				return nil
			}
			fids = append(fids[:idx], fids[idx+1:]...)
			return k.SetPidFindingIDList(ctx, pid, fids)
		}
	}
	return types.ErrFindingNotExists
}

func (k Keeper) GetAllFindings(ctx sdk.Context) []types.Finding {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.FindingKey)

	var findings []types.Finding
	var finding types.Finding

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		k.cdc.MustUnmarshal(iterator.Value(), &finding)
		findings = append(findings, finding)
	}
	return findings
}

func StringsToBytes(list []string) ([]byte, error) {
	marshal, err := json.Marshal(list)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}

func BytesToStrings(list []byte) ([]string, error) {
	var fids []string
	err := json.Unmarshal(list, &fids)
	if err != nil {
		return nil, err
	}

	return fids, nil
}

func (k Keeper) ConfirmFinding(ctx sdk.Context, msg *types.MsgConfirmFinding) (types.Finding, error) {
	var finding types.Finding
	// get finding
	finding, found := k.GetFinding(ctx, msg.FindingId)
	if !found {
		return finding, types.ErrFindingNotExists
	}
	// get program
	program, isExist := k.GetProgram(ctx, finding.ProgramId)
	if !isExist {
		return finding, types.ErrProgramNotExists
	}
	if program.Status != types.ProgramStatusActive {
		return finding, types.ErrProgramNotActive
	}

	// only host can update finding comment
	if program.AdminAddress != msg.OperatorAddress {
		return finding, types.ErrProgramCreatorInvalid
	}

	// fingerprint comparison
	fingerprintHash := k.GetFindingFingerPrintHash(&finding)
	if msg.Fingerprint != fingerprintHash {
		return finding, types.ErrFindingHashInvalid
	}
	return finding, nil
}

func (k Keeper) GetFindingFingerPrintHash(finding *types.Finding) string {
	findingFingerPrint := &types.FindingFingerPrint{
		ProgramId:     finding.ProgramId,
		FindingId:     finding.FindingId,
		Title:         finding.Title,
		FindingHash:   finding.FindingHash,
		SeverityLevel: finding.SeverityLevel,
		Status:        finding.Status,
		Detail:        finding.Detail,
	}

	bz := k.cdc.MustMarshal(findingFingerPrint)
	hash := sha256.Sum256(bz)
	return hex.EncodeToString(hash[:])
}
