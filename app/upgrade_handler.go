package app

import (
	"fmt"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/shentufoundation/shentu/v2/x/bounty"
	bountytypes "github.com/shentufoundation/shentu/v2/x/bounty/types"
)

const (
	upgradeName = "v2.7.0"
)

func (app ShentuApp) setUpgradeHandler() {
	app.UpgradeKeeper.SetUpgradeHandler(
		upgradeName,
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {

			fromVM[bountytypes.ModuleName] = app.mm.Modules[bountytypes.ModuleName].ConsensusVersion()

			ctx.Logger().Info("Start to run module migrations...")

			bountyGenesis := bountytypes.DefaultGenesisState()
			bounty.InitGenesis(ctx, app.BountyKeeper, *bountyGenesis)

			newVersionMap, err := app.mm.RunMigrations(ctx, app.configurator, fromVM)

			return newVersionMap, err
		},
	)

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	if upgradeInfo.Name == upgradeName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{bountytypes.ModuleName},
		}

		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}
