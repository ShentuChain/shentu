package types_test

import (
	"strings"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/certikfoundation/shentu/x/shield/types"
)

type GenesisTestSuite struct {
	suite.Suite
}

func (suite *GenesisTestSuite) TestGenesisValidation() {
	type args struct {
		ShieldAdmin         sdk.AccAddress
		NextPoolID          uint64
		NextPurchaseID      uint64
		PoolParams          types.PoolParams
		ClaimProposalParams types.ClaimProposalParams
		Pools               []types.Pool
		Collaterals         []types.Collateral
		Providers           []types.Provider
		PurchaseLists       []types.PurchaseList
		Withdraws           types.Withdraws
	}
	testCases := []struct {
		name        string
		args        args
		expectPass  bool
		expectedErr string
	}{
		{
			name: "default",
			args: args{
				NextPoolID:          uint64(1),
				NextPurchaseID:      uint64(1),
				PoolParams:          types.DefaultPoolParams(),
				ClaimProposalParams: types.DefaultClaimProposalParams(),
			},
			expectPass:  true,
			expectedErr: "",
		},
		{
			name: "MinPoolLife <= ProtectionPeriod",
			args: args{
				NextPoolID:          uint64(1),
				NextPurchaseID:      uint64(1),
				PoolParams:          types.NewPoolParams(time.Hour*24*14, time.Hour*24*7, time.Hour*24*210, sdk.NewDecWithPrec(1, 2)),
				ClaimProposalParams: types.DefaultClaimProposalParams(),
			},
			expectPass:  false,
			expectedErr: "",
		},
	}
	for _, tc := range testCases {
		tc := tc // scopelint doesn't complain
		suite.Run(tc.name, func() {
			gs := types.NewGenesisState(tc.args.ShieldAdmin, tc.args.NextPoolID, tc.args.NextPurchaseID, tc.args.PoolParams,
				tc.args.ClaimProposalParams, tc.args.Pools, tc.args.Collaterals,
				tc.args.Providers, tc.args.PurchaseLists, tc.args.Withdraws)
			err := gs.Validate()
			if tc.expectPass {
				suite.NoError(err)
			} else {
				suite.Error(err)
				suite.Require().True(strings.Contains(err.Error(), tc.expectedErr))
			}
		})
	}
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}
