package keeper

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto/ecies"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/shentufoundation/shentu/v2/x/bounty/client/cli"
	"github.com/shentufoundation/shentu/v2/x/bounty/types"
)

func (k Keeper) GetFinding(ctx sdk.Context, id uint64) (types.Finding, bool) {
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
	findingData := k.cdc.MustMarshal(&finding)
	store.Set(types.GetFindingKey(finding.FindingId), findingData)
}

func (k Keeper) DeleteFinding(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetFindingKey(id))
}

func (k Keeper) GetNextFindingID(ctx sdk.Context) (uint64, error) {
	store := ctx.KVStore(k.storeKey)
	findingID := store.Get(types.GetNextFindingIDKey())
	if findingID == nil {
		return 1, errorsmod.Wrap(types.ErrInvalidGenesis, "initial next finding ID hasn't been set")
	}
	return binary.LittleEndian.Uint64(findingID), nil
}

func (k Keeper) SetNextFindingID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.LittleEndian.PutUint64(bz, id)
	store.Set(types.GetNextFindingIDKey(), bz)
}

func (k Keeper) SetPidFindingIDList(ctx sdk.Context, pid uint64, findingIds []uint64) error {
	findingIDList, err := Uint64sToBytes(findingIds)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetProgramIDFindingListKey(pid), findingIDList)
	return nil
}

func (k Keeper) GetPidFindingIDList(ctx sdk.Context, pid uint64) ([]uint64, error) {
	store := ctx.KVStore(k.storeKey)
	findingIDs := store.Get(types.GetProgramIDFindingListKey(pid))

	if findingIDs == nil {
		return nil, types.ErrProgramFindingListEmpty
	}

	findingIDList, err := BytesToUint64s(findingIDs)
	if err != nil {
		return nil, err
	}
	return findingIDList, nil
}

func (k Keeper) AppendFidToFidList(ctx sdk.Context, pid, fid uint64) error {
	fids, err := k.GetPidFindingIDList(ctx, pid)
	if err != nil {
		if err == types.ErrProgramFindingListEmpty {
			fids = []uint64{}
		} else {
			return err
		}
	}

	fids = append(fids, fid)
	err = k.SetPidFindingIDList(ctx, pid, fids)
	return err
}

func (k Keeper) DeleteFidFromFidList(ctx sdk.Context, pid, fid uint64) error {
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

func Uint64sToBytes(list []uint64) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, list)
	if err != nil {
		return nil, types.ErrProgramFindingListMarshal
	}
	return buf.Bytes(), nil
}

func BytesToUint64s(list []byte) ([]uint64, error) {
	buf := bytes.NewBuffer(list)
	r64 := make([]uint64, (len(list)+7)/8)
	err := binary.Read(buf, binary.LittleEndian, &r64)
	if err != nil {
		return nil, types.ErrProgramFindingListUnmarshal
	}
	return r64, nil
}

func CheckPlainText(pubKey *ecies.PublicKey, msg *types.MsgReleaseFinding, finding types.Finding) error {
	if finding.GetFindingDesc() != nil {
		encryptedDesc, ok := finding.GetFindingDesc().(*types.EciesEncryptedDesc)
		if !ok {
			return fmt.Errorf("invalid any data")
		}
		if err := CheckEncryptedData(pubKey, msg.Desc, encryptedDesc.FindingDesc); err != nil {
			return err
		}
	} else if msg.Desc != "" {
		return types.ErrFindingPlainTextDataInvalid
	}

	if finding.GetFindingPoc() != nil {
		encryptedPoc, ok := finding.GetFindingPoc().(*types.EciesEncryptedPoc)
		if !ok {
			return fmt.Errorf("invalid any data")
		}
		if err := CheckEncryptedData(pubKey, msg.Poc, encryptedPoc.FindingPoc); err != nil {
			return err
		}
	} else if msg.Poc != "" {
		return types.ErrFindingPlainTextDataInvalid
	}

	if finding.GetFindingComment() != nil {
		encryptedComment, ok := finding.GetFindingComment().(*types.EciesEncryptedComment)
		if !ok {
			return fmt.Errorf("invalid any data")
		}
		if err := CheckEncryptedData(pubKey, msg.Comment, encryptedComment.FindingComment); err != nil {
			return err
		}
	} else if msg.Comment != "" {
		return types.ErrFindingPlainTextDataInvalid
	}

	return nil
}

func CheckEncryptedData(pubKey *ecies.PublicKey, plainText string, encryptedData []byte) error {
	if len(encryptedData) < cli.RandBytesLen {
		return types.ErrFindingEncryptedDataInvalid
	}
	randBytesStart := len(encryptedData) - cli.RandBytesLen
	encryptData := encryptedData[:randBytesStart]
	randBytes := encryptedData[randBytesStart:]

	encryptedBytes, err := ecies.Encrypt(bytes.NewReader(randBytes), pubKey, []byte(plainText), nil, nil)
	if err != nil {
		return types.ErrProgramPubKey
	}

	if !bytes.Equal(encryptedBytes, encryptData) {
		return types.ErrFindingPlainTextDataInvalid
	}
	return nil
}

// GetBase64QueryFinding This function takes a Finding and converts associated descriptions, proofs of concept, and comments into a QueryFinding with all the information encoded in base64.
func GetBase64QueryFinding(finding *types.Finding) (types.QueryFinding, error) {
	queryFinding := types.QueryFinding{
		FindingId:        finding.FindingId,
		Title:            finding.Title,
		ProgramId:        finding.ProgramId,
		SeverityLevel:    finding.SeverityLevel,
		SubmitterAddress: finding.SubmitterAddress,
		FindingStatus:    finding.FindingStatus,
	}

	processFindingData := func(data interface{}, fieldName string) (string, error) {
		if data == nil {
			return "", nil
		}
		switch v := data.(type) {
		case *types.EciesEncryptedDesc:
			randBytesStart := len(v.FindingDesc) - cli.RandBytesLen
			return base64.StdEncoding.EncodeToString(v.FindingDesc[:randBytesStart]), nil
		case *types.EciesEncryptedPoc:
			randBytesStart := len(v.FindingPoc) - cli.RandBytesLen
			return base64.StdEncoding.EncodeToString(v.FindingPoc[:randBytesStart]), nil
		case *types.EciesEncryptedComment:
			randBytesStart := len(v.FindingComment) - cli.RandBytesLen
			return base64.StdEncoding.EncodeToString(v.FindingComment[:randBytesStart]), nil
		case *types.PlainTextDesc:
			if v.FindingDesc != nil {
				return base64.StdEncoding.EncodeToString(v.FindingDesc), nil
			}
		case *types.PlainTextPoc:
			if v.FindingPoc != nil {
				return base64.StdEncoding.EncodeToString(v.FindingPoc), nil
			}
		case *types.PlainTextComment:
			if v.FindingComment != nil {
				return base64.StdEncoding.EncodeToString(v.FindingComment), nil
			}
		default:
			return "", fmt.Errorf("invalid any data for field %s", fieldName)
		}
		return "", nil
	}

	var err error
	queryFinding.FindingDesc, err = processFindingData(finding.GetFindingDesc(), "FindingDesc")
	if err != nil {
		return queryFinding, err
	}

	queryFinding.FindingPoc, err = processFindingData(finding.GetFindingPoc(), "FindingPoc")
	if err != nil {
		return queryFinding, err
	}

	queryFinding.FindingComment, err = processFindingData(finding.GetFindingComment(), "FindingComment")
	if err != nil {
		return queryFinding, err
	}

	return queryFinding, nil
}