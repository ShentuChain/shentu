package types

import (
	"encoding/binary"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is the name of this module.
	ModuleName = "shield"

	// RouterKey is used to route messages.
	RouterKey = ModuleName

	// StoreKey is the prefix under which we store this module's data.
	StoreKey = ModuleName

	// QuerierRoute is used to handle abci_query requests.
	QuerierRoute = ModuleName

	// DefaultParamspace is the default name for parameter store.
	DefaultParamspace = ModuleName
)

var (
	ShieldAdminKey              = []byte{0x00}
	TotalCollateralKey          = []byte{0x01}
	TotalWithdrawingKey         = []byte{0x02}
	TotalShieldKey              = []byte{0x03}
	TotalClaimedKey             = []byte{0x04}
	ServiceFeesKey              = []byte{0x05}
	RemainingServiceFeesKey     = []byte{0x06}
	PoolKey                     = []byte{0x07}
	NextPoolIDKey               = []byte{0x08}
	NextPurchaseIDKey           = []byte{0x09}
	ProviderKey                 = []byte{0x0C}
	WithdrawQueueKey            = []byte{0x0D}
	GlobalStakeForShieldPoolKey = []byte{0x0F}
	PurchaseKey                 = []byte{0x11}
	BlockServiceFeesKey         = []byte{0x12}
	ReimbursementKey            = []byte{0x14}
	ReserveKey                  = []byte{0x15}
	PendingPayoutKey            = []byte{0x16}
)

func GetTotalCollateralKey() []byte {
	return TotalCollateralKey
}

func GetTotalWithdrawingKey() []byte {
	return TotalWithdrawingKey
}

func GetTotalShieldKey() []byte {
	return TotalShieldKey
}

func GetTotalClaimedKey() []byte {
	return TotalClaimedKey
}

func GetServiceFeesKey() []byte {
	return ServiceFeesKey
}

// GetPoolKey gets the key for the pool identified by pool ID.
func GetPoolKey(id uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, id)
	return append(PoolKey, b...)
}

// GetShieldAdminKey gets the key for the shield admin.
func GetShieldAdminKey() []byte {
	return ShieldAdminKey
}

// GetNextPoolIDKey gets the key for the next pool ID.
func GetNextPoolIDKey() []byte {
	return NextPoolIDKey
}

// GetNextPurchaseIDKey gets the key for the next pool ID.
func GetNextPurchaseIDKey() []byte {
	return NextPurchaseIDKey
}

// GetProviderKey gets the key for the delegator's tracker.
func GetProviderKey(addr sdk.AccAddress) []byte {
	return append(ProviderKey, addr...)
}

// GetWithdrawCompletionTimeKey gets a withdraw queue key,
// which is obtained from the completion time.
func GetWithdrawCompletionTimeKey(timestamp time.Time) []byte {
	bz := sdk.FormatTimeBytes(timestamp)
	return append(WithdrawQueueKey, bz...)
}

func GetGlobalStakeForShieldPoolKey() []byte {
	return GlobalStakeForShieldPoolKey
}

func GetPurchaseKey(poolID uint64, purchaser sdk.AccAddress) []byte {
	bz := make([]byte, 8)
	binary.LittleEndian.PutUint64(bz, poolID)
	return append(PurchaseKey, append(bz, purchaser...)...)
}

// GetReimbursementKey gets the key for a reimbursement.
func GetReimbursementKey(proposalID uint64) []byte {
	bz := make([]byte, 8)
	binary.LittleEndian.PutUint64(bz, proposalID)
	return append(ReimbursementKey, bz...)
}

// GetReserveKey gets the key for Shield Donation Pool.
func GetReserveKey() []byte {
	return ReserveKey
}

// GetPendingPayoutKey gets the key for the pending payout
// corresponding to the given proposal ID.
func GetPendingPayoutKey(proposalID uint64) []byte {
	bz := make([]byte, 8)
	binary.LittleEndian.PutUint64(bz, proposalID)
	return append(PendingPayoutKey, bz...)
}
