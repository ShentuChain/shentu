package types

import (
	"time"

	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/cosmos/cosmos-sdk/x/supply/exported"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authexported "github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingexported "github.com/cosmos/cosmos-sdk/x/staking/exported"
)

// AccountKeeper expected account keeper
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authexported.Account
	IterateAccounts(ctx sdk.Context, process func(authexported.Account) (stop bool))
}

// StakingKeeper expected staking keeper
type StakingKeeper interface {
	// iterate through validators by admin address, execute func for each validator
	IterateValidators(sdk.Context, func(index int64, validator stakingexported.ValidatorI) (stop bool))

	// get a particular validator by admin address with a found flag
	GetValidator(sdk.Context, sdk.ValAddress) (staking.Validator, bool)
	// get the set of all validators with no limits, used during genesis dump
	GetAllValidators(ctx sdk.Context) []staking.Validator
	// return all delegations to a specific validator. Useful for querier
	GetValidatorDelegations(ctx sdk.Context, valAddr sdk.ValAddress) []staking.Delegation

	ValidatorByConsAddr(sdk.Context, sdk.ConsAddress) stakingexported.ValidatorI // get a particular validator by consensus address

	// slash the validator and delegators of the validator, specifying offense height, offense power, and slash fraction
	Slash(sdk.Context, sdk.ConsAddress, int64, int64, sdk.Dec)
	Jail(sdk.Context, sdk.ConsAddress)   // jail a validator
	Unjail(sdk.Context, sdk.ConsAddress) // unjail a validator

	// Delegation allows for getting a particular delegation for a given validator
	// and delegator outside the scope of the staking module.
	Delegation(sdk.Context, sdk.AccAddress, sdk.ValAddress) stakingexported.DelegationI
	GetAllDelegatorDelegations(ctx sdk.Context, delegator sdk.AccAddress) []staking.Delegation
	GetAllUnbondingDelegations(ctx sdk.Context, delegator sdk.AccAddress) []staking.UnbondingDelegation
	GetUnbondingDelegation(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) (staking.UnbondingDelegation, bool)
	SetUnbondingDelegation(ctx sdk.Context, ubd staking.UnbondingDelegation)
	GetUBDQueueTimeSlice(ctx sdk.Context, timestamp time.Time) (dvPairs []staking.DVPair)
	SetUBDQueueTimeSlice(ctx sdk.Context, timestamp time.Time, timeslice []staking.DVPair)
	InsertUBDQueue(ctx sdk.Context, ubd staking.UnbondingDelegation, completionTime time.Time)
	SetDelegation(ctx sdk.Context, delegation staking.Delegation)
	GetDelegation(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) (staking.Delegation, bool)
	BeforeDelegationSharesModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress)
	AfterDelegationModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress)
	UBDQueueIterator(ctx sdk.Context, timestamp time.Time) sdk.Iterator
	RemoveValidatorTokensAndShares(ctx sdk.Context, validator staking.Validator, sharesToRemove sdk.Dec) (valOut staking.Validator, removedTokens sdk.Int)
	RemoveUBDQueue(ctx sdk.Context, timestamp time.Time)
	GetRedelegations(ctx sdk.Context, delegator sdk.AccAddress, maxRetrieve uint16) (redelegations []staking.Redelegation)

	BondDenom(sdk.Context) string

	// MaxValidators returns the maximum amount of bonded validators
	MaxValidators(sdk.Context) uint16
}

// BankKeeper defines the expected bank keeper (noalias)
type BankKeeper interface {
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	DelegateCoins(ctx sdk.Context, fromAdd, toAddr sdk.AccAddress, amt sdk.Coins) error
	UndelegateCoins(ctx sdk.Context, fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) error

	SubtractCoins(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coins) (sdk.Coins, error)
	AddCoins(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coins) (sdk.Coins, error)
}

type SupplyKeeper interface {
	GetModuleAddress(moduleName string) sdk.AccAddress
	GetModuleAccount(ctx sdk.Context, moduleName string) exported.ModuleAccountI
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
}

type GovKeeper interface {
	GetVotingParams(ctx sdk.Context) govTypes.VotingParams
}
