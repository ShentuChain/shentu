package v1alpha1

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/certikfoundation/shentu/v2/x/shield/types"
)

const (
	TypeMsgCreatePool             = "create_pool"
	TypeMsgUpdatePool             = "update_pool"
	TypeMsgPausePool              = "pause_pool"
	TypeMsgResumePool             = "resume_pool"
	TypeMsgDepositCollateral      = "deposit_collateral"
	TypeMsgWithdrawCollateral     = "withdraw_collateral"
	TypeMsgWithdrawRewards        = "withdraw_rewards"
	TypeMsgWithdrawForeignRewards = "withdraw_foreign_rewards"
	TypeMsgPurchaseShield         = "purchase_shield"
	TypeMsgWithdrawReimbursement  = "withdraw_reimbursement"
	TypeMsgStakeForShield         = "stake_for_shield"
	TypeMsgUnstakeFromShield      = "unstake_from_shield"
	TypeMsgUpdateSponsor          = "update_sponsor"
)

// NewMsgCreatePool creates a new NewMsgCreatePool instance.
func NewMsgCreatePool(accAddr sdk.AccAddress, shield sdk.Coins, deposit MixedCoins, sponsor string, sponsorAddr sdk.AccAddress, description string, shieldLimit sdk.Int) *MsgCreatePool {
	return &MsgCreatePool{
		From:        accAddr.String(),
		Shield:      shield,
		Deposit:     deposit,
		Sponsor:     sponsor,
		SponsorAddr: sponsorAddr.String(),
		Description: description,
		ShieldLimit: shieldLimit,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgCreatePool) Route() string { return types.RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgCreatePool) Type() string { return TypeMsgCreatePool }

// GetSigners implements the sdk.Msg interface.
func (msg MsgCreatePool) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgCreatePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgCreatePool) ValidateBasic() error {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return err
	}
	if from.Empty() {
		return types.ErrEmptySender
	}

	if strings.TrimSpace(msg.Sponsor) == "" {
		return types.ErrEmptySponsor
	}
	if !msg.Shield.IsValid() || msg.Shield.IsZero() {
		return types.ErrNoShield
	}
	return nil
}

// NewMsgUpdatePool creates a new MsgUpdatePool instance.
func NewMsgUpdatePool(accAddr sdk.AccAddress, shield sdk.Coins, serviceFees MixedCoins, id uint64, description string, shieldLimit sdk.Int) *MsgUpdatePool {
	return &MsgUpdatePool{
		From:        accAddr.String(),
		Shield:      shield,
		ServiceFees: serviceFees,
		PoolId:      id,
		Description: description,
		ShieldLimit: shieldLimit,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgUpdatePool) Route() string { return types.RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgUpdatePool) Type() string { return TypeMsgUpdatePool }

// GetSigners implements the sdk.Msg interface.
func (msg MsgUpdatePool) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgUpdatePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgUpdatePool) ValidateBasic() error {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	if from.Empty() {
		return types.ErrEmptySender
	}

	if msg.PoolId == 0 {
		return types.ErrInvalidPoolID
	}
	if !msg.Shield.IsValid() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid shield")
	}
	return nil
}

// NewMsgPausePool creates a new NewMsgPausePool instance.
func NewMsgPausePool(accAddr sdk.AccAddress, id uint64) *MsgPausePool {
	return &MsgPausePool{
		From:   accAddr.String(),
		PoolId: id,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgPausePool) Route() string { return types.RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgPausePool) Type() string { return TypeMsgPausePool }

// GetSigners implements the sdk.Msg interface.
func (msg MsgPausePool) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgPausePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgPausePool) ValidateBasic() error {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	if from.Empty() {
		return types.ErrEmptySender
	}

	if msg.PoolId == 0 {
		return types.ErrInvalidPoolID
	}
	return nil
}

func NewMsgResumePool(accAddr sdk.AccAddress, id uint64) *MsgResumePool {
	return &MsgResumePool{
		From:   accAddr.String(),
		PoolId: id,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgResumePool) Route() string { return types.RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgResumePool) Type() string { return TypeMsgResumePool }

// GetSigners implements the sdk.Msg interface.
func (msg MsgResumePool) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgResumePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgResumePool) ValidateBasic() error {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	if from.Empty() {
		return types.ErrEmptySender
	}

	if msg.PoolId == 0 {
		return types.ErrInvalidPoolID
	}
	return nil
}

// NewMsgDepositCollateral creates a new MsgDepositCollateral instance.
func NewMsgDepositCollateral(sender sdk.AccAddress, collateral sdk.Coins) *MsgDepositCollateral {
	return &MsgDepositCollateral{
		From:       sender.String(),
		Collateral: collateral,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgDepositCollateral) Route() string { return types.RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgDepositCollateral) Type() string { return "deposit_collateral" }

// GetSigners implements the sdk.Msg interface.
func (msg MsgDepositCollateral) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgDepositCollateral) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgDepositCollateral) ValidateBasic() error {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	if from.Empty() {
		return types.ErrEmptySender
	}

	if !msg.Collateral.IsValid() || msg.Collateral.IsZero() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "Collateral amount: %s", msg.Collateral)
	}
	return nil
}

// NewMsgDepositCollateral creates a new MsgDepositCollateral instance.
func NewMsgWithdrawCollateral(sender sdk.AccAddress, collateral sdk.Coins) *MsgWithdrawCollateral {
	return &MsgWithdrawCollateral{
		From:       sender.String(),
		Collateral: collateral,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgWithdrawCollateral) Route() string { return types.RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgWithdrawCollateral) Type() string { return "withdraw_collateral" }

// GetSigners implements the sdk.Msg interface.
func (msg MsgWithdrawCollateral) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgWithdrawCollateral) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgWithdrawCollateral) ValidateBasic() error {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	if from.Empty() {
		return types.ErrEmptySender
	}

	if !msg.Collateral.IsValid() || msg.Collateral.IsZero() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "Collateral amount: %s", msg.Collateral)
	}
	return nil
}

// NewMsgWithdrawRewards creates a new MsgWithdrawRewards instance.
func NewMsgWithdrawRewards(sender sdk.AccAddress) *MsgWithdrawRewards {
	return &MsgWithdrawRewards{
		From: sender.String(),
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgWithdrawRewards) Route() string { return types.RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgWithdrawRewards) Type() string { return TypeMsgWithdrawRewards }

// GetSigners implements the sdk.Msg interface.
func (msg MsgWithdrawRewards) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgWithdrawRewards) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgWithdrawRewards) ValidateBasic() error {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	if from.Empty() {
		return types.ErrEmptySender
	}

	return nil
}

// NewMsgWithdrawForeignRewards creates a new MsgWithdrawForeignRewards instance.
func NewMsgWithdrawForeignRewards(sender sdk.AccAddress, denom, toAddr string) *MsgWithdrawForeignRewards {
	return &MsgWithdrawForeignRewards{
		From:   sender.String(),
		Denom:  denom,
		ToAddr: toAddr,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgWithdrawForeignRewards) Route() string { return types.RouterKey }

// Type implements the sdk.Msg interface
func (msg MsgWithdrawForeignRewards) Type() string { return TypeMsgWithdrawForeignRewards }

// GetSigners implements the sdk.Msg interface
func (msg MsgWithdrawForeignRewards) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgWithdrawForeignRewards) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgWithdrawForeignRewards) ValidateBasic() error {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	if from.Empty() {
		return types.ErrEmptySender
	}
	if strings.TrimSpace(msg.ToAddr) == "" {
		return types.ErrInvalidToAddr
	}
	return nil
}

// NewMsgPurchaseShield creates a new MsgPurchaseShield instance.
func NewMsgPurchaseShield(poolID uint64, shield sdk.Coins, description string, from sdk.AccAddress) *MsgPurchaseShield {
	return &MsgPurchaseShield{
		PoolId:      poolID,
		Shield:      shield,
		Description: description,
		From:        from.String(),
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgPurchaseShield) Route() string { return types.RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgPurchaseShield) Type() string { return TypeMsgPurchaseShield }

// GetSigners implements the sdk.Msg interface.
func (msg MsgPurchaseShield) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgPurchaseShield) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgPurchaseShield) ValidateBasic() error {
	if msg.PoolId == 0 {
		return types.ErrInvalidPoolID
	}
	if !msg.Shield.IsValid() || msg.Shield.IsZero() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "shield amount: %s", msg.Shield)
	}
	if strings.TrimSpace(msg.Description) == "" {
		return types.ErrPurchaseMissingDescription
	}

	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	if from.Empty() {
		return types.ErrEmptySender
	}
	return nil
}

// NewMsgWithdrawReimbursement creates a new MsgWithdrawReimbursement instance.
func NewMsgWithdrawReimbursement(proposalID uint64, from sdk.AccAddress) *MsgWithdrawReimbursement {
	return &MsgWithdrawReimbursement{
		ProposalId: proposalID,
		From:       from.String(),
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgWithdrawReimbursement) Route() string { return types.RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgWithdrawReimbursement) Type() string { return TypeMsgWithdrawReimbursement }

// GetSigners implements the sdk.Msg interface.
func (msg MsgWithdrawReimbursement) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgWithdrawReimbursement) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgWithdrawReimbursement) ValidateBasic() error {
	return nil
}

// NewMsgStakeForShield creates a new MsgPurchaseShield instance.
func NewMsgStakeForShield(poolID uint64, shield sdk.Coins, description string, from sdk.AccAddress) *MsgStakeForShield {
	return &MsgStakeForShield{
		PoolId:      poolID,
		Shield:      shield,
		Description: description,
		From:        from.String(),
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgStakeForShield) Route() string { return types.RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgStakeForShield) Type() string { return TypeMsgStakeForShield }

// GetSigners implements the sdk.Msg interface.
func (msg MsgStakeForShield) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgStakeForShield) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgStakeForShield) ValidateBasic() error {
	return nil
}

// NewMsgUnstakeFromShield creates a new MsgPurchaseShield instance.
func NewMsgUnstakeFromShield(poolID uint64, shield sdk.Coins, from sdk.AccAddress) *MsgUnstakeFromShield {
	return &MsgUnstakeFromShield{
		PoolId: poolID,
		Shield: shield,
		From:   from.String(),
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgUnstakeFromShield) Route() string { return types.RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgUnstakeFromShield) Type() string { return TypeMsgUnstakeFromShield }

// GetSigners implements the sdk.Msg interface.
func (msg MsgUnstakeFromShield) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgUnstakeFromShield) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgUnstakeFromShield) ValidateBasic() error {
	return nil
}

// NewMsgUpdateSponsor creates a new NewMsgUpdateSponsor instance.
func NewMsgUpdateSponsor(poolID uint64, sponsor string, sponsorAddr, fromAddr sdk.AccAddress) *MsgUpdateSponsor {
	return &MsgUpdateSponsor{
		PoolId:      poolID,
		Sponsor:     sponsor,
		SponsorAddr: sponsorAddr.String(),
		From:        fromAddr.String(),
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgUpdateSponsor) Route() string { return types.RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgUpdateSponsor) Type() string { return TypeMsgUpdateSponsor }

// GetSigners implements the sdk.Msg interface.
func (msg MsgUpdateSponsor) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgUpdateSponsor) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgUpdateSponsor) ValidateBasic() error {
	fromAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	if fromAddr.Empty() {
		return types.ErrEmptySender
	}

	sponsorAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	if sponsorAddr.Empty() || strings.TrimSpace(msg.Sponsor) == "" {
		return types.ErrEmptySponsor
	}
	return nil
}
