package app

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

func (app CertiKApp) setUpgradeHandler() {
	app.upgradeKeeper.SetUpgradeHandler("shentu-2.1", func(ctx sdk.Context, plan upgradetypes.Plan) {
	})

	upgradeInfo, err := app.upgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
		// handle error
	}

	if upgradeInfo.Name == "shentu-2.1" && !app.upgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := storetypes.StoreUpgrades{
			Renamed: []storetypes.StoreRename{{}},
			Deleted: []string{},
		}

		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}
