package types

const (
	EventTypeCreateProgram     = "create_program"
	EventTypeSubmitFinding     = "submit_finding"
	EventTypeAcceptFinding     = "accept_finding"
	EventTypeRejectFinding     = "reject_finding"
	EventTypeWithdrawalFinding = "withdrawal_finding"

	AttributeKeyProgramID = "program_id"
	AttributeKeyDeposit   = "deposit"

	AttributeKeyFindingID = "finding_id"

	AttributeValueCategory = ModuleName
)
