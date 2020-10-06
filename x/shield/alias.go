package shield

import (
	"github.com/certikfoundation/shentu/x/shield/client"
	"github.com/certikfoundation/shentu/x/shield/keeper"
	"github.com/certikfoundation/shentu/x/shield/types"
)

const (
	ModuleName   = types.ModuleName
	StoreKey     = types.StoreKey
	RouterKey    = types.RouterKey
	QuerierRoute = types.QuerierRoute
)

type (
	Keeper              = keeper.Keeper
	GenesisState        = types.GenesisState
	Provider            = types.Provider
	ClaimProposal       = types.ShieldClaimProposal
	ClaimProposalParams = types.ClaimProposalParams
	Purchase            = types.Purchase
)

var (
	// functions aliases
	RegisterInvariants          = keeper.RegisterInvariants
	NewKeeper                   = keeper.NewKeeper
	NewQuerier                  = keeper.NewQuerier
	ModuleCdc                   = types.ModuleCdc
	ProposalHandler             = client.ProposalHandler
	GetGenesisStateFromAppState = types.GetGenesisStateFromAppState
	ValidateGenesis             = types.ValidateGenesis

	DefaultParamSpace       = types.DefaultParamspace
	ProposalTypeShieldClaim = types.ProposalTypeShieldClaim
)
