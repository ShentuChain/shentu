package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	ModuleName = "auth"
	RouterKey  = ModuleName
)

// MsgTriggerVesting triggers vesting of a triggered vesting account.
type MsgTriggerVesting struct {
	Certifier sdk.AccAddress `json:"certifier" yaml:"certifier"`
	Account   sdk.AccAddress `json:"account_address" yaml:"account_address"`
}

var _ sdk.Msg = MsgTriggerVesting{}

func NewMsgTriggerVesting(certifier, account sdk.AccAddress) MsgTriggerVesting {
	return MsgTriggerVesting{
		Certifier: certifier,
		Account:   account,
	}
}

// Route returns the name of the module.
func (m MsgTriggerVesting) Route() string { return ModuleName }

// Type returns a human-readable string for the message.
func (m MsgTriggerVesting) Type() string { return "trigger_vesting" }

// ValidateBasic runs stateless checks on the message.
func (m MsgTriggerVesting) ValidateBasic() error {
	if m.Certifier.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing from address")
	}
	if m.Account.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing account address")
	}
	return nil
}

// GetSignBytes encodes the message for signing.
func (m MsgTriggerVesting) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

// GetSigners defines whose signature is required.
func (m MsgTriggerVesting) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Certifier}
}

// MsgManualVesting unlocks the specified amount in a manual vesting account.
type MsgManualVesting struct {
	Certifier    sdk.AccAddress `json:"certifier" yaml:"certifier"`
	Account      sdk.AccAddress `json:"account_address" yaml:"account_address"`
	UnlockAmount sdk.Coin       `json:"unlock_amount" yaml:"unlock_amount"`
}

var _ sdk.Msg = MsgManualVesting{}

func NewMsgManualVesting(certifier, account sdk.AccAddress, unlockAmount sdk.Coin) MsgManualVesting {
	return MsgManualVesting{
		Certifier:    certifier,
		Account:      account,
		UnlockAmount: unlockAmount,
	}
}

// Route returns the name of the module.
func (m MsgManualVesting) Route() string { return ModuleName }

// Type returns a human-readable string for the message.
func (m MsgManualVesting) Type() string { return "manual_vesting" }

// ValidateBasic runs stateless checks on the message.
func (m MsgManualVesting) ValidateBasic() error {
	if m.Certifier.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing from address")
	}
	if m.Account.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing account address")
	}
	return nil
}

// GetSignBytes encodes the message for signing.
func (m MsgManualVesting) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

// GetSigners defines whose signature is required.
func (m MsgManualVesting) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Certifier}
}
