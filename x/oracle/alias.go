package oracle

import (
	"github.com/certikfoundation/shentu/x/oracle/keeper"
	"github.com/certikfoundation/shentu/x/oracle/types"
)

const (
	ModuleName        = types.ModuleName
	QuerierRoute      = types.QuerierRoute
	StoreKey          = types.StoreKey
	DefaultParamSpace = types.ModuleName
)

var (
	NewKeeper                 = keeper.NewKeeper
	NewQuerier                = keeper.NewQuerier
	NewMsgTaskResponse        = types.NewMsgTaskResponse
	DefaultGenesisState       = types.DefaultGenesisState
	TaskStoreKeyPrefix        = types.TaskStoreKeyPrefix
	ClosingTaskStoreKeyPrefix = types.ClosingTaskStoreKeyPrefix
)

type (
	Keeper          = keeper.Keeper
	MsgTaskResponse = types.MsgTaskResponse
	MsgCreateTask   = types.MsgCreateTask
)
