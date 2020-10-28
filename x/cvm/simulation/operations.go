package simulation

import (
	"encoding/hex"
	"math/rand"
	"strconv"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/certikfoundation/shentu/x/cvm/internal/keeper"
	"github.com/certikfoundation/shentu/x/cvm/internal/types"
)

const (
	OpWeightMsgDeploy = "op_weight_msg_deploy"
)

// WeightedOperations creates an operation with a weight for each type of message generators.
func WeightedOperations(appParams simulation.AppParams, cdc *codec.Codec, k keeper.Keeper) simulation.WeightedOperations {
	var weightMsgDeploy int
	appParams.GetOrGenerate(cdc, OpWeightMsgDeploy, &weightMsgDeploy, nil,
		func(_ *rand.Rand) {
			weightMsgDeploy = simappparams.DefaultWeightMsgSend
		})

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(weightMsgDeploy, SimulateMsgDeployHello55(k)),
		simulation.NewWeightedOperation(weightMsgDeploy, SimulateMsgDeploySimple(k)),
		simulation.NewWeightedOperation(weightMsgDeploy, SimulateMsgDeploySimpleEvent(k)),
		simulation.NewWeightedOperation(weightMsgDeploy, SimulateMsgDeployStorage(k)),
	}
}

// SimulateMsgDeployHello55 creates a massage deploying /tests/hello55.sol contract.
func SimulateMsgDeployHello55(k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		caller, _ := simulation.RandomAcc(r, accs)

		// deploy hello55.sol
		msg, contractAddr, err := DeployContract(caller, Hello55Code, Hello55Abi, k, r, ctx, chainID, app)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		// check sayHi() ret
		data, err := hex.DecodeString(Hello55SayHi)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		ret, err := k.Call(ctx, caller.Address, contractAddr, 0, data, nil, true, false, false)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		value, err := strconv.ParseInt(hex.EncodeToString(ret), 16, 32)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		if value != 55 {
			panic("return value incorrect")
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgDeploySimple creates a massage deploying /tests/simple.sol contract.
func SimulateMsgDeploySimple(k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		caller, _ := simulation.RandomAcc(r, accs)

		// deploy simple.sol
		msg, contractAddr, err := DeployContract(caller, SimpleCode, SimpleAbi, k, r, ctx, chainID, app)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		// check get() ret
		data, err := hex.DecodeString(SimpleGet)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		ret, err := k.Call(ctx, caller.Address, contractAddr, 0, data, nil, true, false, false)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		value, err := strconv.ParseInt(hex.EncodeToString(ret), 16, 64)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		if value != 0 {
			panic("return value incorrect")
		}

		futureOperations := []simulation.FutureOperation{
			{
				BlockHeight: int(ctx.BlockHeight()) + r.Intn(10),
				Op:          SimulateMsgCallSimpleSet(k, contractAddr, int(r.Uint32())),
			},
		}

		return simulation.NewOperationMsg(msg, true, ""), futureOperations, nil
	}
}

// SimulateMsgCallSimpleSet creates a message calling set() in /tests/simple.sol contract.
func SimulateMsgCallSimpleSet(k keeper.Keeper, contractAddr sdk.AccAddress, varValue int) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		caller, _ := simulation.RandomAcc(r, accs)

		hexStr := strconv.FormatInt(int64(varValue), 16)
		length := len(hexStr)
		for i := 0; i < 64-length; i++ {
			hexStr = "0" + hexStr
		}
		data, err := hex.DecodeString(SimpleSetPrefix + hexStr)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		msg := types.NewMsgCall(caller.Address, contractAddr, 0, data)

		account := k.AuthKeeper().GetAccount(ctx, caller.Address)
		fees, err := simulation.RandomFees(r, ctx, account.SpendableCoins(ctx.BlockTime()))
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			caller.PrivKey,
		)

		_, _, err = app.Deliver(tx)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		// check pure/view function ret
		data, err = hex.DecodeString(SimpleGet)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		ret, err := k.Call(ctx, caller.Address, contractAddr, 0, data, nil, true, false, false)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		value, err := strconv.ParseInt(hex.EncodeToString(ret), 16, 64)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		if value != int64(varValue) {
			panic("return value incorrect")
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

func callFunction() {

}

// SimulateMsgDeploySimpleEvent creates a massage deploying /tests/simpleevent.sol contract.
func SimulateMsgDeploySimpleEvent(k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		caller, _ := simulation.RandomAcc(r, accs)

		// deploy simpleevent.sol
		msg, contractAddr, err := DeployContract(caller, SimpleeventCode, SimpleeventAbi, k, r, ctx, chainID, app)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		// check get() ret
		data, err := hex.DecodeString(SimpleeventGet)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		ret, err := k.Call(ctx, caller.Address, contractAddr, 0, data, nil, true, false, false)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		value, err := strconv.ParseInt(hex.EncodeToString(ret), 16, 64)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		if value != 0 {
			panic("return value incorrect")
		}

		futureOperations := []simulation.FutureOperation{
			{
				BlockHeight: int(ctx.BlockHeight()) + r.Intn(10),
				Op:          SimulateMsgCallSimpleEventSet(k, contractAddr, int(r.Uint32())),
			},
		}

		return simulation.NewOperationMsg(msg, true, ""), futureOperations, nil
	}
}

// SimulateMsgCallSimpleEventSet creates a message calling set() in /tests/simpleevent.sol contract.
func SimulateMsgCallSimpleEventSet(k keeper.Keeper, contractAddr sdk.AccAddress, varValue int) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		caller, _ := simulation.RandomAcc(r, accs)

		hexStr := strconv.FormatInt(int64(varValue), 16)
		length := len(hexStr)
		for i := 0; i < 64-length; i++ {
			hexStr = "0" + hexStr
		}
		data, err := hex.DecodeString(SimpleeventSetPrefix + hexStr)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		msg := types.NewMsgCall(caller.Address, contractAddr, 0, data)

		account := k.AuthKeeper().GetAccount(ctx, caller.Address)
		fees, err := simulation.RandomFees(r, ctx, account.SpendableCoins(ctx.BlockTime()))
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			caller.PrivKey,
		)

		_, _, err = app.Deliver(tx)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		// check pure/view function ret
		data, err = hex.DecodeString(SimpleeventGet)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		ret, err := k.Call(ctx, caller.Address, contractAddr, 0, data, nil, true, false, false)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		value, err := strconv.ParseInt(hex.EncodeToString(ret), 16, 64)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		if value != int64(varValue) {
			panic("return value incorrect")
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgDeployStorage creates a massage deploying /tests/storage.sol contract.
func SimulateMsgDeployStorage(k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		caller, _ := simulation.RandomAcc(r, accs)

		// deploy storage.sol
		msg, contractAddr, err := DeployContract(caller, StorageCode, StorageAbi, k, r, ctx, chainID, app)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		// check retrieve() ret
		data, err := hex.DecodeString(StorageRetrieve)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		ret, err := k.Call(ctx, caller.Address, contractAddr, 0, data, nil, true, false, false)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		value, err := strconv.ParseInt(hex.EncodeToString(ret), 16, 64)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		if value != 0 {
			panic("return value incorrect")
		}

		// check sayMyAddres() ret
		data, err = hex.DecodeString(StorageSayMyAddres)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		ret, err = k.Call(ctx, caller.Address, contractAddr, 0, data, nil, true, false, false)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		sender := sdk.AccAddress(ret[12:])
		if !sender.Equals(caller.Address) {
			panic("return value incorrect")
		}

		futureOperations := []simulation.FutureOperation{
			{
				BlockHeight: int(ctx.BlockHeight()) + r.Intn(10),
				Op:          SimulateMsgCallStorageStore(k, contractAddr, int(r.Uint32())),
			},
		}

		return simulation.NewOperationMsg(msg, true, ""), futureOperations, nil
	}
}

// SimulateMsgCallStorageStore creates a message calling store() in /tests/storage.sol contract.
func SimulateMsgCallStorageStore(k keeper.Keeper, contractAddr sdk.AccAddress, varValue int) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		caller, _ := simulation.RandomAcc(r, accs)

		hexStr := strconv.FormatInt(int64(varValue), 16)
		length := len(hexStr)
		for i := 0; i < 64-length; i++ {
			hexStr = "0" + hexStr
		}
		data, err := hex.DecodeString(StorageStorePrefix + hexStr)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		msg := types.NewMsgCall(caller.Address, contractAddr, 0, data)

		account := k.AuthKeeper().GetAccount(ctx, caller.Address)
		fees, err := simulation.RandomFees(r, ctx, account.SpendableCoins(ctx.BlockTime()))
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			caller.PrivKey,
		)

		_, _, err = app.Deliver(tx)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		// check pure/view function ret
		data, err = hex.DecodeString(StorageRetrieve)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		ret, err := k.Call(ctx, caller.Address, contractAddr, 0, data, nil, true, false, false)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		value, err := strconv.ParseInt(hex.EncodeToString(ret), 16, 64)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		if value != int64(varValue) {
			panic("return value incorrect")
		}

		// check pure/view function ret
		data, err = hex.DecodeString(StorageSayMyAddres)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		ret, err = k.Call(ctx, caller.Address, contractAddr, 0, data, nil, true, false, false)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}
		sender := sdk.AccAddress(ret[12:])
		if !sender.Equals(caller.Address) {
			panic("return value incorrect")
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}
