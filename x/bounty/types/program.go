package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Programs is an array of program
type Programs []Program

// Findings is an array of finding
type Findings []Finding

func (m *Program) ValidateBasic() error {
	if m.ProgramId == 0 {
		return ErrProgramID
	}
	if m.EncryptionKey == nil {
		return ErrProgramPubKey
	}
	if _, err := sdk.AccAddressFromBech32(m.CreatorAddress); err != nil {
		return err
	}

	for _, deposit := range m.Deposit {
		if !deposit.IsValid() {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "Deposit is invalid [%s]", deposit.String())
		}
	}
	return nil
}

func (m *Finding) ValidateBasic() error {
	if m.ProgramId == 0 {
		return ErrProgramID
	} else if m.FindingId == 0 {
		return ErrFindingID
	} else if m.SeverityLevel < 0 || int(m.SeverityLevel) >= len(SeverityLevel_name) {
		return fmt.Errorf("invalid SeverityLevel:%d", m.SeverityLevel)
	} else if int(m.FindingStatus) >= len(FindingStatus_name) {
		return fmt.Errorf("invalid finding status:%d", m.FindingStatus)
	}

	if _, err := sdk.AccAddressFromBech32(m.SubmitterAddress); err != nil {
		return err
	}
	return nil
}