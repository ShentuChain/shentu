package types

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateProgram      = "create_program"
	TypeMsgEditProgram        = "edit_program"
	TypeMsgActivateProgram    = "activate_program"
	TypeMsgCloseProgram       = "close_program"
	TypeMsgSubmitFinding      = "submit_finding"
	TypeMsgEditFinding        = "edit_finding"
	TypeMsgActivateFinding    = "activate_finding"
	TypeMsgConfirmFinding     = "confirm_finding"
	TypeMsgConfirmFindingPaid = "confirm_finding_paid"
	TypeMsgCloseFinding       = "close_finding"
	TypeMsgPublishFinding     = "publish_finding"
)

var (
	_, _, _, _       sdk.Msg = &MsgCreateProgram{}, &MsgEditProgram{}, &MsgActivateProgram{}, &MsgCloseProgram{}
	_, _, _, _, _, _ sdk.Msg = &MsgSubmitFinding{}, &MsgEditFinding{}, &MsgActivateFinding{}, &MsgConfirmFinding{}, &MsgCloseFinding{}, &MsgPublishFinding{}
)

// NewMsgCreateProgram creates a new NewMsgCreateProgram instance.
// Delegator address and validator address are the same.
func NewMsgCreateProgram(pid, name, detail string, operator sdk.AccAddress) *MsgCreateProgram {
	return &MsgCreateProgram{
		ProgramId:       pid,
		Name:            name,
		Detail:          detail,
		OperatorAddress: operator.String(),
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgCreateProgram) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgCreateProgram) Type() string { return TypeMsgCreateProgram }

// GetSigners implements the sdk.Msg interface. It returns the address(es) that
// must sign over msg.GetSignBytes().
// If the validator address is not same as delegator's, then the validator must
// sign the msg as well.
func (msg MsgCreateProgram) GetSigners() []sdk.AccAddress {
	// creator should sign the message
	cAddr, err := sdk.AccAddressFromBech32(msg.OperatorAddress)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{cAddr}
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgCreateProgram) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgCreateProgram) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OperatorAddress)
	if err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
	}
	if len(msg.ProgramId) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "empty programId")
	}
	if len(msg.Name) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "empty name")
	}
	if len(msg.Detail) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "empty detail")
	}
	return nil
}

// NewMsgEditProgram edit a program.
func NewMsgEditProgram(pid, name, detail string, operator sdk.AccAddress) *MsgEditProgram {
	return &MsgEditProgram{
		ProgramId:       pid,
		Name:            name,
		Detail:          detail,
		OperatorAddress: operator.String(),
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgEditProgram) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgEditProgram) Type() string { return TypeMsgEditProgram }

// GetSigners implements the sdk.Msg interface. It returns the address(es) that
// must sign over msg.GetSignBytes().
// If the validator address is not same as delegator's, then the validator must
// sign the msg as well.
func (msg MsgEditProgram) GetSigners() []sdk.AccAddress {
	// creator should sign the message
	cAddr, err := sdk.AccAddressFromBech32(msg.OperatorAddress)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{cAddr}
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgEditProgram) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgEditProgram) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OperatorAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, err.Error())
	}
	if len(msg.ProgramId) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "empty programId")
	}
	return nil
}

// NewMsgSubmitFinding submit a new finding.
func NewMsgSubmitFinding(pid, fid, title, detail, hash string, operator sdk.AccAddress, level SeverityLevel) *MsgSubmitFinding {
	return &MsgSubmitFinding{
		ProgramId:       pid,
		FindingId:       fid,
		Title:           title,
		FindingHash:     hash,
		OperatorAddress: operator.String(),
		SeverityLevel:   level,
		Detail:          detail,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgSubmitFinding) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgSubmitFinding) Type() string { return TypeMsgSubmitFinding }

// GetSigners implements the sdk.Msg interface. It returns the address(es) that
// must sign over msg.GetSignBytes().
// If the validator address is not same as delegator's, then the validator must
// sign the msg as well.
func (msg MsgSubmitFinding) GetSigners() []sdk.AccAddress {
	// creator should sign the message
	cAddr, err := sdk.AccAddressFromBech32(msg.OperatorAddress)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{cAddr}
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgSubmitFinding) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgSubmitFinding) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OperatorAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, err.Error())
	}
	if len(msg.ProgramId) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "empty programId")
	}
	if len(msg.FindingId) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "empty findingId")
	}
	if len(msg.Title) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "empty title")
	}
	if len(msg.FindingHash) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "empty findingHash")
	}
	if len(msg.Detail) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "empty detail")
	}
	if !ValidFindingSeverityLevel(msg.SeverityLevel) {
		return errorsmod.Wrap(ErrFindingSeverityLevelInvalid, msg.SeverityLevel.String())
	}
	return nil
}

// NewMsgEditFinding submit a new finding.
func NewMsgEditFinding(fid, title, detail, hash, paymentHash string, operator sdk.AccAddress, level SeverityLevel) *MsgEditFinding {
	return &MsgEditFinding{
		FindingId:       fid,
		Title:           title,
		FindingHash:     hash,
		OperatorAddress: operator.String(),
		SeverityLevel:   level,
		Detail:          detail,
		PaymentHash:     paymentHash,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgEditFinding) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgEditFinding) Type() string { return TypeMsgEditFinding }

// GetSigners implements the sdk.Msg interface. It returns the address(es) that
// must sign over msg.GetSignBytes().
// If the validator address is not same as delegator's, then the validator must
// sign the msg as well.
func (msg MsgEditFinding) GetSigners() []sdk.AccAddress {
	// creator should sign the message
	cAddr, err := sdk.AccAddressFromBech32(msg.OperatorAddress)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{cAddr}
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgEditFinding) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgEditFinding) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OperatorAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, err.Error())
	}
	if len(msg.FindingId) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "empty findingId")
	}
	if !ValidFindingSeverityLevel(msg.SeverityLevel) {
		return errorsmod.Wrap(ErrFindingSeverityLevelInvalid, msg.SeverityLevel.String())
	}
	return nil
}

func NewMsgActivateProgram(pid string, accAddr sdk.AccAddress) *MsgActivateProgram {
	return &MsgActivateProgram{
		ProgramId:       pid,
		OperatorAddress: accAddr.String(),
	}
}

// Route implements sdk.Msg interface.
func (msg MsgActivateProgram) Route() string { return RouterKey }

// Type implements sdk.Msg interface.
func (msg MsgActivateProgram) Type() string { return TypeMsgActivateProgram }

// GetSigners implements sdk.Msg interface. It returns the address(es) that
// must sign over msg.GetSignBytes().
func (msg MsgActivateProgram) GetSigners() []sdk.AccAddress {
	cAddr, _ := sdk.AccAddressFromBech32(msg.OperatorAddress)
	return []sdk.AccAddress{cAddr}
}

// GetSignBytes implements the sdk.Msg interface, returns the message bytes to sign over.
func (msg MsgActivateProgram) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgActivateProgram) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OperatorAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, err.Error())
	}
	if len(msg.ProgramId) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "empty programId")
	}
	return nil
}

func NewMsgCloseProgram(pid string, accAddr sdk.AccAddress) *MsgCloseProgram {
	return &MsgCloseProgram{
		ProgramId:       pid,
		OperatorAddress: accAddr.String(),
	}
}

// Route implements sdk.Msg interface.
func (msg MsgCloseProgram) Route() string { return RouterKey }

// Type implements sdk.Msg interface.
func (msg MsgCloseProgram) Type() string { return TypeMsgCloseProgram }

// GetSigners implements sdk.Msg interface. It returns the address(es) that
// must sign over msg.GetSignBytes().
func (msg MsgCloseProgram) GetSigners() []sdk.AccAddress {
	cAddr, _ := sdk.AccAddressFromBech32(msg.OperatorAddress)
	return []sdk.AccAddress{cAddr}
}

// GetSignBytes implements the sdk.Msg interface, returns the message bytes to sign over.
func (msg MsgCloseProgram) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgCloseProgram) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OperatorAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, err.Error())
	}
	if len(msg.ProgramId) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "empty programId")
	}
	return nil
}

func NewMsgActivateFinding(findingID string, hostAddr sdk.AccAddress) *MsgActivateFinding {
	return &MsgActivateFinding{
		FindingId:       findingID,
		OperatorAddress: hostAddr.String(),
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgActivateFinding) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgActivateFinding) Type() string { return TypeMsgActivateFinding }

// GetSignBytes returns the message bytes to sign over.
func (msg MsgActivateFinding) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements the sdk.Msg interface. It returns the address(es) that
// must sign over msg.GetSignBytes().
func (msg MsgActivateFinding) GetSigners() []sdk.AccAddress {
	// host should sign the message
	hostAddr, err := sdk.AccAddressFromBech32(msg.OperatorAddress)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{hostAddr}
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgActivateFinding) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OperatorAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, err.Error())
	}
	if len(msg.FindingId) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "empty findingId")
	}
	return nil
}

func NewMsgConfirmFinding(findingID, fingerprint string, hostAddr sdk.AccAddress) *MsgConfirmFinding {
	return &MsgConfirmFinding{
		FindingId:       findingID,
		OperatorAddress: hostAddr.String(),
		Fingerprint:     fingerprint,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgConfirmFinding) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgConfirmFinding) Type() string { return TypeMsgConfirmFinding }

// GetSignBytes returns the message bytes to sign over.
func (msg MsgConfirmFinding) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements the sdk.Msg interface. It returns the address(es) that
// must sign over msg.GetSignBytes().
func (msg MsgConfirmFinding) GetSigners() []sdk.AccAddress {
	// host should sign the message
	hostAddr, err := sdk.AccAddressFromBech32(msg.OperatorAddress)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{hostAddr}
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgConfirmFinding) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OperatorAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, err.Error())
	}
	if len(msg.FindingId) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "empty findingId")
	}
	if len(msg.Fingerprint) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "empty fingerprint")
	}
	return nil
}

func NewMsgConfirmFindingPaid(findingID string, hostAddr sdk.AccAddress) *MsgConfirmFindingPaid {
	return &MsgConfirmFindingPaid{
		FindingId:       findingID,
		OperatorAddress: hostAddr.String(),
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgConfirmFindingPaid) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgConfirmFindingPaid) Type() string { return TypeMsgConfirmFindingPaid }

// GetSignBytes returns the message bytes to sign over.
func (msg MsgConfirmFindingPaid) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements the sdk.Msg interface. It returns the address(es) that
// must sign over msg.GetSignBytes().
func (msg MsgConfirmFindingPaid) GetSigners() []sdk.AccAddress {
	// host should sign the message
	hostAddr, err := sdk.AccAddressFromBech32(msg.OperatorAddress)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{hostAddr}
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgConfirmFindingPaid) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OperatorAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, err.Error())
	}

	if len(msg.FindingId) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "empty findingId")
	}
	return nil
}

func NewMsgCloseFinding(findingID string, hostAddr sdk.AccAddress) *MsgCloseFinding {
	return &MsgCloseFinding{
		FindingId:       findingID,
		OperatorAddress: hostAddr.String(),
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgCloseFinding) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgCloseFinding) Type() string { return TypeMsgCloseFinding }

// GetSignBytes returns the message bytes to sign over.
func (msg MsgCloseFinding) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements the sdk.Msg interface. It returns the address(es) that
// must sign over msg.GetSignBytes().
func (msg MsgCloseFinding) GetSigners() []sdk.AccAddress {
	// host should sign the message
	hostAddr, err := sdk.AccAddressFromBech32(msg.OperatorAddress)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{hostAddr}
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgCloseFinding) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OperatorAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, err.Error())
	}

	if len(msg.FindingId) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "empty findingId")
	}
	return nil
}

// NewMsgPublishFinding publish finding.
func NewMsgPublishFinding(fid, desc, poc string, operator sdk.AccAddress) *MsgPublishFinding {
	return &MsgPublishFinding{
		FindingId:       fid,
		Description:     desc,
		ProofOfConcept:  poc,
		OperatorAddress: operator.String(),
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgPublishFinding) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgPublishFinding) Type() string { return TypeMsgPublishFinding }

// GetSigners implements the sdk.Msg interface. It returns the address(es) that
// must sign over msg.GetSignBytes().
func (msg MsgPublishFinding) GetSigners() []sdk.AccAddress {
	// releaser should sign the message
	cAddr, err := sdk.AccAddressFromBech32(msg.OperatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{cAddr}
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgPublishFinding) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgPublishFinding) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OperatorAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, err.Error())
	}

	if len(msg.FindingId) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "empty findingId")
	}
	if len(msg.Description) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "description")
	}
	if len(msg.ProofOfConcept) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "empty proofOfConcept")
	}
	return nil
}
