package types

const (
	EventTypeCreatePool             = "create_pool"
	EventTypeUpdatePool             = "update_pool"
	EventTypePausePool              = "pause_pool"
	EventTypeResumePool             = "resume_pool"
	EventTypeDepositCollateral      = "deposit_collateral"
	EventTypeWithdrawCollateral     = "withdraw_collateral"
	EventTypePurchaseShield         = "purchase_shield"
	EventTypeWithdrawRewards        = "withdraw_rewards"
	EventTypeWithdrawForeignRewards = "withdraw_foreign_rewards"
	EventTypeClearPayouts           = "clear_payouts"
	EventTypeCreateCompensation     = "create_compensation"
	EventTypeWithdrawReimbursement  = "withdraw_reimbursement"

	AttributeKeyShield             = "shield"
	AttributeKeyDeposit            = "deposit"
	AttributeKeySponsor            = "sponsor"
	AttributeKeyPoolID             = "pool_id"
	AttributeKeyAdditionalTime     = "additional_time"
	AttributeKeyTimeOfCoverage     = "time_of_coverage"
	AttributeKeyBlocksOfCoverage   = "blocks_of_coverage"
	AttributeKeyCollateral         = "collateral"
	AttributeKeyDenom              = "denom"
	AttributeKeyToAddr             = "to_address"
	AttributeKeyAccountAddress     = "account_address"
	AttributeKeyAmount             = "amount"
	AttributeKeyPurchaseTxHash     = "purchase_txhash"
	AttributeKeyCompensationAmount = "compensation_amount"
	AttributeKeyBeneficiary        = "beneficiary"
	AttributeValueCategory         = ModuleName
)
