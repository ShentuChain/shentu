package keeper

import (
	"sort"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking"

	"github.com/certikfoundation/shentu/x/shield/types"
)

// ClaimLock locks collaterals after a claim proposal is submitted.
func (k Keeper) ClaimLock(ctx sdk.Context, proposalID uint64, poolID uint64,
	loss sdk.Coins, purchaseTxHash []byte, lockPeriod time.Duration) error {
	pool, err := k.GetPool(ctx, poolID)
	if err != nil {
		return err
	}

	// update shield of purchase
	purchase, err := k.GetPurchase(ctx, purchaseTxHash)
	if err != nil {
		return err
	}
	if !purchase.Shield.IsAllGTE(loss) {
		return types.ErrNotEnoughShield
	}
	purchase.Shield = purchase.Shield.Sub(loss)
	k.SetPurchase(ctx, purchaseTxHash, purchase)

	// update the shield of pool
	if !pool.Shield.IsAllGTE(loss) {
		// TODO this should never happen?
		return types.ErrNotEnoughShield
	}
	pool.Shield = pool.Shield.Sub(loss)

	// update locked collaterals for community
	collaterals := k.GetAllPoolCollaterals(ctx, pool)
	for _, collateral := range collaterals {
		lockedCoins := GetLockedCoins(loss, pool.TotalCollateral, collateral.Amount)
		lockedCollateral := types.NewLockedCollateral(proposalID, lockedCoins)
		collateral.LockedCollaterals = append(collateral.LockedCollaterals, lockedCollateral)
		collateral.Amount = collateral.Amount.Sub(lockedCoins)
		k.LockProvider(ctx, collateral.Provider, lockedCoins, lockPeriod)
		k.SetCollateral(ctx, pool, collateral.Provider, collateral)
	}

	k.SetPool(ctx, pool)
	return nil
}

// LockProvider checks if delegations of an account can cover the loss.
// It modifies unbonding time if the totals delegations cannot cover the loss.
func (k Keeper) LockProvider(ctx sdk.Context, delAddr sdk.AccAddress, locked sdk.Coins, lockPeriod time.Duration) {
	provider, found := k.GetProvider(ctx, delAddr)
	if !found {
		panic(types.ErrProviderNotFound)
	}
	if !provider.Collateral.IsAllGTE(locked) {
		panic(types.ErrNotEnoughCollateral)
	}

	// update provider
	provider.TotalLocked = provider.TotalLocked.Add(locked...)
	k.SetProvider(ctx, delAddr, provider)

	// if there are enough delegations
	// TODO logics for providers are changing, check if it outdated
	if provider.DelegationBonded.IsAllGTE(provider.TotalLocked) {
		return
	}

	// if there are not enough delegations, check unbondings
	unbondingDelegations := k.GetSortedUnbondingDelegations(ctx, delAddr)
	short := provider.TotalLocked.Sub(provider.DelegationBonded).AmountOf(k.sk.BondDenom(ctx))
	endTime := ctx.BlockTime().Add(lockPeriod)
	for _, ubd := range unbondingDelegations {
		if !short.IsPositive() {
			return
		}
		if ubd.Entry.CompletionTime.Before(endTime) {
			// change unbonding completion time
			k.UpdateUnbonding(ctx, ubd, ubd.DelegatorAddress, endTime)
		}
		short = short.Sub(ubd.Entry.Balance)
	}
	if short.IsPositive() {
		panic("not enough bonded and unbonding delegations")
	}
}

// GetSortedUnbondingDelegations gets unbonding delegations sorted by completion time.
func (k Keeper) GetSortedUnbondingDelegations(ctx sdk.Context, delAddr sdk.AccAddress) []types.UnbondingDelegation {
	unbondingDelegation := k.sk.GetAllUnbondingDelegations(ctx, delAddr)
	var unbondingDelegations []types.UnbondingDelegation
	for _, ubd := range unbondingDelegation {
		for _, entry := range ubd.Entries {
			unbondingDelegations = append(
				unbondingDelegations,
				types.NewUnbondingDelegation(ubd.DelegatorAddress, ubd.ValidatorAddress, entry),
			)
		}
	}
	sort.SliceStable(unbondingDelegations, func(i, j int) bool {
		return unbondingDelegations[i].Entry.CompletionTime.After(unbondingDelegations[j].Entry.CompletionTime)
	})
	return unbondingDelegations
}

// UpdateUnbondingCompletionTime updates the completion time of a unbonding delegation.
func (k Keeper) UpdateUnbonding(
	ctx sdk.Context, unbondingDelegation types.UnbondingDelegation, recipient sdk.AccAddress, endTime time.Time) {
	delAddr := unbondingDelegation.DelegatorAddress
	valAddr := unbondingDelegation.ValidatorAddress
	ubd, found := k.sk.GetUnbondingDelegation(ctx, delAddr, valAddr)
	if !found {
		panic("unbonding delegation was not found")
	}
	for i, entry := range ubd.Entries {
		// TODO: this check is not enough. Need to fix it.
		if entry.CreationHeight == unbondingDelegation.Entry.CreationHeight {
			// remove unbonding delegation with old completion time from UBDQueue
			timeSlice := k.sk.GetUBDQueueTimeSlice(ctx, entry.CompletionTime)
			for i := 0; i < len(timeSlice); i++ {
				if timeSlice[i].DelegatorAddress.Equals(delAddr) && timeSlice[i].ValidatorAddress.Equals(valAddr) {
					timeSlice = append(timeSlice[:i], timeSlice[i+1:]...)
					break
				}
			}

			// update and add unbonding delegation
			entry.CompletionTime = endTime
			ubd.Entries[i] = entry
			ubd.DelegatorAddress = recipient
			k.sk.SetUnbondingDelegation(ctx, ubd)
			k.sk.InsertUBDQueue(ctx, ubd, endTime)
			break
		}
	}
}

// ClaimUnlock unlocks locked collaterals.
func (k Keeper) ClaimUnlock(ctx sdk.Context, proposalID uint64, poolID uint64, loss sdk.Coins) error {
	pool, err := k.GetPool(ctx, poolID)
	if err != nil {
		return err
	}

	// unlock collaterals for community
	collaterals := k.GetAllPoolCollaterals(ctx, pool)
	for _, collateral := range collaterals {
		for j, locked := range collateral.LockedCollaterals {
			if locked.ProposalID == proposalID {
				collateral.Amount = collateral.Amount.Add(locked.LockedCoins...)
				collateral.LockedCollaterals = append(collateral.LockedCollaterals[:j], collateral.LockedCollaterals[j+1:]...)
				break
			}
		}
		k.SetCollateral(ctx, pool, collateral.Provider, collateral)
	}

	// unlock collaterals for CertiK
	collateral := k.GetPoolCertiKCollateral(ctx, pool)
	for i, locked := range collateral.LockedCollaterals {
		if locked.ProposalID == proposalID {
			collateral.Amount = collateral.Amount.Add(locked.LockedCoins...)
			collateral.LockedCollaterals[i] = collateral.LockedCollaterals[len(collateral.LockedCollaterals)-1]
			collateral.LockedCollaterals = collateral.LockedCollaterals[:len(collateral.LockedCollaterals)-1]
			break
		}
	}
	k.SetCollateral(ctx, pool, k.GetAdmin(ctx), collateral)

	k.SetPool(ctx, pool)
	return nil
}

// RestoreShield restores shield for proposer.
func (k Keeper) RestoreShield(ctx sdk.Context, poolID uint64, loss sdk.Coins, purchaseTxHash []byte) error {
	pool, err := k.GetPool(ctx, poolID)
	if err != nil {
		return err
	}

	// update shield of purchase
	purchase, err := k.GetPurchase(ctx, purchaseTxHash)
	if err != nil {
		return err
	}
	purchase.Shield = purchase.Shield.Sub(loss)
	k.SetPurchase(ctx, purchaseTxHash, purchase)

	// update shield of pool
	pool.Shield = pool.Shield.Add(loss...)
	k.SetPool(ctx, pool)

	return nil
}

// UndelegateCoinsToShieldModule undelegates delegations and send coins the the shield module.
func (k Keeper) UndelegateCoinsToShieldModule(ctx sdk.Context, delAddr sdk.AccAddress, loss sdk.Coins) error {
	delegations := k.sk.GetAllDelegatorDelegations(ctx, delAddr)
	var totalDelAmountDec sdk.Dec
	for _, del := range delegations {
		val, found := k.sk.GetValidator(ctx, del.GetValidatorAddr())
		if !found {
			panic("validator was not found")
		}
		totalDelAmountDec = totalDelAmountDec.Add(val.TokensFromShares(del.GetShares()))
	}

	// start with bonded delegations
	lossAmountDec := loss.AmountOf(k.sk.BondDenom(ctx)).ToDec()
	for _, del := range delegations {
		val, found := k.sk.GetValidator(ctx, del.GetValidatorAddr())
		if !found {
			panic("validator is not found")
		}
		delAmountDec := val.TokensFromShares(del.GetShares())
		var ubdAmountDec sdk.Dec
		if totalDelAmountDec.GT(lossAmountDec) {
			ubdAmountDec = lossAmountDec.Mul(delAmountDec).Quo(totalDelAmountDec)
		} else {
			ubdAmountDec = delAmountDec
		}
		ubdShares, err := val.SharesFromTokens(ubdAmountDec.TruncateInt())
		if err != nil {
			panic(err)
		}
		k.UndelegateShares(ctx, del.DelegatorAddress, del.ValidatorAddress, ubdShares)
	}
	if totalDelAmountDec.GTE(lossAmountDec) {
		return nil
	}

	// if bonded delegations are not enough, track unbonding delegations
	moduleAddr := k.supplyKeeper.GetModuleAddress(types.ModuleName)
	unbondingDelegations := k.GetSortedUnbondingDelegations(ctx, delAddr)
	shortDec := lossAmountDec.Sub(totalDelAmountDec)
	for _, ubd := range unbondingDelegations {
		if !shortDec.IsPositive() {
			return nil
		}
		k.UpdateUnbonding(ctx, ubd, moduleAddr, ubd.Entry.CompletionTime)
		ubdAmountDec := ubd.Entry.InitialBalance.ToDec()
		if ubdAmountDec.GT(shortDec) {
			// FIXME not a good way to go maybe?
			overflowCoins := sdk.NewDecCoins(sdk.NewDecCoin(k.sk.BondDenom(ctx), ubdAmountDec.Sub(shortDec).TruncateInt()))
			overflowMixedCoins := types.MixedDecCoins{Native: overflowCoins}
			k.AddRewards(ctx, delAddr, overflowMixedCoins)
			break
		}
		shortDec = shortDec.Sub(ubdAmountDec)
	}
	if shortDec.IsPositive() {
		panic("not enough bonded stake")
	}
	return nil
}

// UndelegateShares undelegates delegations of a delegator to a validator by shares.
func (k Keeper) UndelegateShares(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress, shares sdk.Dec) {
	delegation, found := k.sk.GetDelegation(ctx, delAddr, valAddr)
	if !found {
		panic("delegation is not found")
	}
	k.sk.BeforeDelegationSharesModified(ctx, delAddr, valAddr)

	// undelegate coins from the staking module account to the shield module account
	val, found := k.sk.GetValidator(ctx, valAddr)
	if !found {
		panic("validator was not found")
	}
	ubdCoins := sdk.NewCoins(sdk.NewCoin(k.sk.BondDenom(ctx), val.TokensFromShares(shares).TruncateInt()))
	if err := k.supplyKeeper.SendCoinsFromModuleToModule(ctx, staking.BondedPoolName, types.ModuleName, ubdCoins); err != nil {
		panic(err)
	}

	// update delegation records
	delegation.Shares = delegation.Shares.Sub(shares)
	k.sk.SetDelegation(ctx, delegation)

	k.sk.AfterDelegationModified(ctx, delegation.DelegatorAddress, delegation.ValidatorAddress)
}

// SetReimbursement sets a reimbursement in store.
func (k Keeper) SetReimbursement(ctx sdk.Context, proposalID uint64, payout types.Reimbursement) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(payout)
	store.Set(types.GetReimbursementKey(proposalID), bz)
}

// GetReimbursement get a reimbursement in store.
func (k Keeper) GetReimbursement(ctx sdk.Context, proposalID uint64) (types.Reimbursement, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetReimbursementKey(proposalID))
	if bz != nil {
		var reimbursement types.Reimbursement
		k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &reimbursement)
		return reimbursement, nil
	}
	return types.Reimbursement{}, types.ErrCompensationNotFound
}

// DeleteReimbursement deletes a reimbursement.
func (k Keeper) DeleteReimbursement(ctx sdk.Context, proposalID uint64) error {
	store := ctx.KVStore(k.storeKey)
	if _, err := k.GetReimbursement(ctx, proposalID); err != nil {
		return err
	}
	store.Delete(types.GetReimbursementKey(proposalID))
	return nil
}

// CreateReimbursement creates a reimbursement.
func (k Keeper) CreateReimbursement(
	ctx sdk.Context, proposalID uint64, poolID uint64, amount sdk.Coins, beneficiary sdk.AccAddress,
) error {
	pool, err := k.GetPool(ctx, poolID)
	if err != nil {
		return err
	}

	// for each community member, get coins from delegations
	var loss sdk.Coins
	collaterals := k.GetAllPoolCollaterals(ctx, pool)
	for _, collateral := range collaterals {
		for j := range collateral.LockedCollaterals {
			if collateral.LockedCollaterals[j].ProposalID == proposalID {
				loss = collateral.LockedCollaterals[j].LockedCoins
				collateral.LockedCollaterals = append(
					collateral.LockedCollaterals[:j],
					collateral.LockedCollaterals[j+1],
				)
				k.SetCollateral(ctx, pool, collateral.Provider, collateral)
				break
			}
		}
		if err := k.UndelegateCoinsToShieldModule(ctx, collateral.Provider, loss); err != nil {
			return err
		}
	}

	proposalParams := k.GetClaimProposalParams(ctx)
	reimbursement := types.NewReimbursement(amount, beneficiary, ctx.BlockTime().Add(proposalParams.PayoutPeriod))
	k.SetReimbursement(ctx, proposalID, reimbursement)
	return nil
}

// WithdrawReimbursement checks a reimbursement and pays the beneficiary.
func (k Keeper) WithdrawReimbursement(ctx sdk.Context, proposalID uint64, beneficiary sdk.AccAddress) (sdk.Coins, error) {
	reimbursement, err := k.GetReimbursement(ctx, proposalID)
	if err != nil {
		return sdk.Coins{}, err
	}

	// check beneficiary and time
	if !reimbursement.Beneficiary.Equals(beneficiary) {
		return sdk.Coins{}, types.ErrInvalidBeneficiary
	}
	if reimbursement.PayoutTime.After(ctx.BlockTime()) {
		return sdk.Coins{}, types.ErrNotPayoutTime
	}

	if err := k.supplyKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, beneficiary, reimbursement.Amount); err != nil {
		return sdk.Coins{}, types.ErrNotPayoutTime
	}
	if err := k.DeleteReimbursement(ctx, proposalID); err != nil {
		return sdk.Coins{}, err
	}
	return reimbursement.Amount, nil
}
