package simulation

import (
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/certikfoundation/shentu/x/oracle/internal/types"
)

// RandomizedGenState creates a random genesis state for module simulation.
func RandomizedGenState(simState *module.SimulationState) {
	var poolParams types.LockedPoolParams
	simState.AppParams.GetOrGenerate(
		simState.Cdc, string(types.ParamsStoreKeyPoolParams), &poolParams, simState.Rand,
		func(r *rand.Rand) {
			poolParams = GenPoolParams(r)
		})

	var taskParams types.TaskParams
	simState.AppParams.GetOrGenerate(
		simState.Cdc, string(types.ParamsStoreKeyTaskParams), &taskParams, simState.Rand,
		func(r *rand.Rand) {
			taskParams = GenTaskParams(r)
		})

	gs := types.NewGenesisState(
		nil,
		nil,
		poolParams,
		taskParams,
		nil,
		nil,
	)

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(gs)
}

// GenPoolParams returns a randomized LockedPoolParams object.
func GenPoolParams(r *rand.Rand) types.LockedPoolParams {
	return types.LockedPoolParams{
		LockedInBlocks:    r.Int63n(60),
		MinimumCollateral: r.Int63n(100000),
	}
}

// GenTaskParams returns a randomized TaskParams object.
func GenTaskParams(r *rand.Rand) types.TaskParams {
	return types.TaskParams{
		ExpirationDuration: time.Duration(r.Int63n(1000 * 1000 * 1000 * 60 * 60 * 48)),
		AggregationWindow:  r.Int63n(40),
		AggregationResult:  sdk.NewInt(r.Int63n(3)),
		ThresholdScore:     sdk.NewInt(r.Int63n(257)),
		Epsilon1:           sdk.NewInt(r.Int63n(3)),
		Epsilon2:           sdk.NewInt(r.Int63n(201)),
	}
}
