package types

import (
	"encoding/json"
	"time"
	"github.com/gogo/protobuf/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewTask returns a new task.
func NewTask(
	contract string,
	function string,
	beginBlock int64,
	bounty sdk.Coins,
	description string,
	expiration time.Time,
	creator sdk.AccAddress,
	closingBlock int64,
	waitingBlocks int64,
) Task {
	return Task{
		Contract:      contract,
		Function:      function,
		BeginBlock:    beginBlock,
		Bounty:        bounty,
		Description:   description,
		Expiration:    expiration,
		Creator:       creator.String(),
		ClosingBlock:  closingBlock,
		WaitingBlocks: waitingBlocks,
		Status:        TaskStatusPending,
	}
}

// NewResponse returns a new response.
func NewResponse(score sdk.Int, operator sdk.AccAddress) Response {
	return Response{
		Operator: operator.String(),
		Score:    score,
	}
}

type Responses []Response

// String implements the Stringer interface.
func (r Responses) String() string {
	jsonBytes, err := json.Marshal(r)
	if err != nil {
		return "[]"
	}
	return string(jsonBytes)
}

type TaskI interface {
	proto.Message

	GetID() []byte
	GetCreator() string
	GetResponses() []Response
	IsExpired(ctx sdk.Context) bool
	GetValidTime(ctx sdk.Context) (int64, time.Time)
	GetStatus() TaskStatus
	GetScore() int64
}

func (t Task) GetID() []byte {
	return append([]byte(t.Contract), []byte(t.Function)...)
}

func (t Task) GetCreator() string {
	return t.Creator
}

func (t Task) GetResponses() []Response {
	return t.Responses
}

func (t Task) IsExpired(ctx sdk.Context) bool {
	return t.Expiration.Before(ctx.BlockTime())
}

func (t Task) GetValidTime(ctx sdk.Context) (int64, time.Time) {
	return t.ClosingBlock, time.Time{}
}

func (t Task) GetStatus() TaskStatus {
	return t.Status
}

func (t Task) GetScore() int64 {
	return t.Result.Int64()
}

func (t TxTask) GetID() []byte {
	return t.TxHash
}

func (t TxTask) GetCreator() string {
	return t.Creator
}

func (t TxTask) GetResponses() []Response {
	return t.Responses
}

func (t TxTask) IsExpired(ctx sdk.Context) bool {
	return t.Expiration.Before(ctx.BlockTime())
}

func (t TxTask) GetValidTime(ctx sdk.Context) (int64, time.Time) {
	return -1, t.ValidTime
}

func (t TxTask) GetStatus() TaskStatus {
	return t.Status
}

func (t TxTask) GetScore() int64 {
	return -100
}