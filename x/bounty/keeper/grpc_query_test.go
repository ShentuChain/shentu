package keeper_test

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"

	"github.com/shentufoundation/shentu/v2/x/bounty/types"
)

func (suite *KeeperTestSuite) TestGRPCQueryProgram() {
	queryClient := suite.queryClient

	var (
		req *types.QueryProgramRequest
	)

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"empty request",
			func() {
				req = &types.QueryProgramRequest{}
			},
			false,
		},
		{
			"non existing program request",
			func() {
				req = &types.QueryProgramRequest{ProgramId: "3"}
			},
			false,
		},
		{
			"zero program id request",
			func() {
				req = &types.QueryProgramRequest{ProgramId: "0"}
			},
			false,
		},
		{
			"valid request",
			func() {
				req = &types.QueryProgramRequest{ProgramId: "1"}
				// create programs
				suite.InitCreateProgram("1")
			},
			true,
		},
	}

	for _, testCase := range testCases {
		suite.Run(fmt.Sprintf("Case %s", testCase.msg), func() {
			testCase.malleate()

			programRes, err := queryClient.Program(context.Background(), req)

			if testCase.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				suite.Require().Nil(programRes)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestGRPCQueryPrograms() {
	queryClient := suite.queryClient

	var (
		req *types.QueryProgramsRequest
	)

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"valid request",
			func() {
				req = &types.QueryProgramsRequest{
					Pagination: nil,
				}

				// create two programs
				suite.InitCreateProgram("1")
				suite.InitCreateProgram("2")
			},
			true,
		},
	}

	for _, testCase := range testCases {
		suite.Run(fmt.Sprintf("Case %s", testCase.msg), func() {
			testCase.malleate()

			programRes, err := queryClient.Programs(context.Background(), req)

			if testCase.expPass {
				suite.Require().NoError(err)
				suite.Require().Equal(len(programRes.Programs), 2)
			} else {
				suite.Require().Error(err)
				suite.Require().Nil(programRes)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestGRPCQueryFinding() {
	queryClient := suite.queryClient

	// create programs
	pid := uuid.NewString()
	suite.InitCreateProgram(pid)
	suite.InitActivateProgram(pid)
	var (
		req *types.QueryFindingRequest
	)

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"empty request",
			func() {
				req = &types.QueryFindingRequest{}
			},
			false,
		},
		{
			"non existing finding id request",
			func() {
				req = &types.QueryFindingRequest{FindingId: "100"}
			},
			false,
		},
		{
			"zero finding id request",
			func() {
				req = &types.QueryFindingRequest{FindingId: "1"}
			},
			false,
		},
		{
			"valid request",
			func() {
				req = &types.QueryFindingRequest{FindingId: "1"}
				suite.InitSubmitFinding("1", "1")
			},
			true,
		},
		{
			"valid request",
			func() {
				req = &types.QueryFindingRequest{FindingId: "2"}
				suite.InitSubmitFinding("1", "2")

				ctx := sdk.WrapSDKContext(suite.ctx)
				suite.msgServer.ReleaseFinding(ctx, types.NewMsgReleaseFinding("2", "desc", "poc", suite.address[0]))
			},
			true,
		},
	}

	for _, testCase := range testCases {
		suite.Run(fmt.Sprintf("Case %s", testCase.msg), func() {
			testCase.malleate()

			findingRes, err := queryClient.Finding(context.Background(), req)
			if testCase.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				suite.Require().Nil(findingRes)
			}
		})
	}

}

func (suite *KeeperTestSuite) TestGRPCQueryFindings() {
	queryClient := suite.queryClient

	// create programs
	pid := uuid.NewString()
	suite.InitCreateProgram(pid)
	suite.InitActivateProgram(pid)

	var (
		req *types.QueryFindingsRequest
	)

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"valid request",
			func() {
				req = &types.QueryFindingsRequest{ProgramId: pid}
				suite.InitSubmitFinding(pid, "1")
			},
			true,
		},
		{
			"valid request with submitter address",
			func() {
				req = &types.QueryFindingsRequest{ProgramId: pid, SubmitterAddress: suite.address[0].String()}
			},
			true,
		},
	}

	for _, testCase := range testCases {
		suite.Run(fmt.Sprintf("Case %s", testCase.msg), func() {
			testCase.malleate()

			findingRes, err := queryClient.Findings(context.Background(), req)

			if testCase.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				suite.Require().Nil(findingRes)
			}
		})
	}

}
