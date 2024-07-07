package keeper

import (
	"github.com/shentufoundation/shentu/v2/x/mint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
)

type Keeper struct {
	mintkeeper.Keeper
	dk            types.DistributionKeeper
	accountKeeper types.AccountKeeper
	stakingKeeper types.StakingKeeper
}

// NewKeeper implements the wrapper newkeeper on top of the original newkeeper with distribution, supply and staking keeper.
func NewKeeper(
	cdc codec.BinaryCodec, key storetypes.StoreKey,
	sk types.StakingKeeper, ak types.AccountKeeper, bk types.BankKeeper, distributionKeeper types.DistributionKeeper,
	feeCollectorName string, authority string) Keeper {
	return Keeper{
		Keeper:        mintkeeper.NewKeeper(cdc, key, sk, ak, bk, feeCollectorName, authority),
		dk:            distributionKeeper,
		accountKeeper: ak,
		stakingKeeper: sk,
	}
}

// SendToCommunityPool sends coins to the community pool using FundCommunityPool.
func (k Keeper) SendToCommunityPool(ctx sdk.Context, amount sdk.Coins) error {
	if amount.AmountOf(k.stakingKeeper.BondDenom(ctx)).Equal(sdk.ZeroInt()) {
		return nil
	}
	mintAddress := k.accountKeeper.GetModuleAddress(minttypes.ModuleName)
	return k.dk.FundCommunityPool(ctx, amount, mintAddress)
}

// GetCommunityPoolRatio returns the current ratio of the community pool compared to the total supply.
func (k Keeper) GetCommunityPoolRatio(ctx sdk.Context) sdk.Dec {
	communityPool := k.dk.GetFeePool(ctx).CommunityPool
	for _, coin := range communityPool {
		totalBondedTokensDec := sdk.NewDecFromInt(k.StakingTokenSupply(ctx))
		if coin.Denom == k.stakingKeeper.BondDenom(ctx) {
			ratio := coin.Amount.Quo(totalBondedTokensDec)
			return ratio
		}
	}
	return sdk.NewDec(0)
}

// GetPoolMint returns Coins that are about to be minted towards the community pool.
func (k Keeper) GetPoolMint(ctx sdk.Context, ratio sdk.Dec, mintedCoin sdk.Coin) sdk.Coin {
	communityPoolMintDec := ratio.MulInt(mintedCoin.Amount)
	amount := communityPoolMintDec.TruncateInt()
	return sdk.NewCoin(k.stakingKeeper.BondDenom(ctx), amount)
}
