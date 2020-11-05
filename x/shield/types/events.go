package types

const (
	EventTypeCreatePool             = "create_pool"
	EventTypeUpdatePool             = "update_pool"
	EventTypePausePool              = "pause_pool"
	EventTypeResumePool             = "resume_pool"
	EventTypeDepositCollateral      = "deposit_collateral"
	EventTypeWithdrawCollateral     = "withdraw_collateral"
	EventTypePurchaseShield         = "purchase_shield"
	EventTypeStakeForShield         = "stake_for_shield"
	EventTypeUnstakeFromShield      = "unstake_from_shield"
	EventTypeWithdrawRewards        = "withdraw_rewards"
	EventTypeWithdrawForeignRewards = "withdraw_foreign_rewards"
	EventTypeClearPayouts           = "clear_payouts"
	EventTypeCreateCompensation     = "create_compensation"
	EventTypeWithdrawReimbursement  = "withdraw_reimbursement"
	EventTypeUpdateSponsor          = "update_sponsor"

	AttributeKeyShield              = "shield"
	AttributeKeyDeposit             = "deposit"
	AttributeKeySponsor             = "sponsor"
	AttributeKeySponsorAddress      = "sponsor_address"
	AttributeKeyPoolID              = "pool_id"
	AttributeKeyAdditionalTime      = "additional_time"
	AttributeKeyTimeOfCoverage      = "time_of_coverage"
	AttributeKeyBlocksOfCoverage    = "blocks_of_coverage"
	AttributeKeyCollateral          = "collateral"
	AttributeKeyDenom               = "denom"
	AttributeKeyToAddr              = "to_address"
	AttributeKeyAccountAddress      = "account_address"
	AttributeKeyAmount              = "amount"
	AttributeKeyPurchaseID          = "purchase_id"
	AttributeKeyCompensationAmount  = "compensation_amount"
	AttributeKeyBeneficiary         = "beneficiary"
	AttributeKeyPurchaseDescription = "purchase_description"
	AttributeKeyServiceFees         = "service_fees"
	AttributeKeyProtectionEndTime   = "protection_end_time"
	AttributeValueCategory          = ModuleName
)
