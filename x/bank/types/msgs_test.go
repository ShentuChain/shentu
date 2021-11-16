package types

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
)

type IntegrationTestSuite struct {
	suite.Suite

	app         *simapp.SimApp
	ctx         sdk.Context
	queryClient types.QueryClient
}

func (suite *IntegrationTestSuite) SetupTest() {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())
	app.BankKeeper.SetParams(ctx, types.DefaultParams())

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.BankKeeper)
	queryClient := types.NewQueryClient(queryHelper)

	suite.app = app
	suite.ctx = ctx
	suite.queryClient = queryClient
}

func TestTypesTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (suite *IntegrationTestSuite) TestMsgSendRoute() {
	addr1 := sdk.AccAddress([]byte("from"))
	addr2 := sdk.AccAddress([]byte("to"))
	unlockerAddress := sdk.AccAddress([]byte("unlocker"))
	coins := sdk.NewCoins(sdk.NewInt64Coin("uctk", 10))
	var msg = NewMsgLockedSend(addr1, addr2, unlockerAddress.String(), coins)
	suite.Require().Equal(msg.Route(), bankTypes.RouterKey)
	suite.Require().Equal(msg.Type(), "locked_send")
}

func (suite *IntegrationTestSuite) TestMsgSendValidation() {
	addr1 := sdk.AccAddress([]byte("from"))
	addr2 := sdk.AccAddress([]byte("to"))
	unlockerAddress := sdk.AccAddress([]byte("unlocker"))
	CTK123 := sdk.NewCoins(sdk.NewInt64Coin("ctk", 123))
	CTK0 := sdk.NewCoins(sdk.NewInt64Coin("ctk", 0))
	CTK123eth123 := sdk.NewCoins(sdk.NewInt64Coin("ctk", 123), sdk.NewInt64Coin("eth", 123))
	CTK123eth0 := sdk.Coins{sdk.NewInt64Coin("ctk", 123), sdk.NewInt64Coin("eth", 0)}

	var emptyAddr sdk.AccAddress

	cases := []struct {
		tx    *MsgLockedSend
		valid bool
	}{
		{NewMsgLockedSend(addr1, addr2, unlockerAddress.String(), CTK123), true},       // valid send
		{NewMsgLockedSend(addr1, addr2, unlockerAddress.String(), CTK123eth123), true}, // valid send with multiple coins
		{NewMsgLockedSend(addr1, addr2, unlockerAddress.String(), CTK0), false},        // non positive coin
		{NewMsgLockedSend(addr1, addr2, unlockerAddress.String(), CTK123eth0), false},  // non positive coin in multicoins
		{NewMsgLockedSend(emptyAddr, addr2, unlockerAddress.String(), CTK123), false},  // empty from addr
		{NewMsgLockedSend(addr1, emptyAddr, unlockerAddress.String(), CTK123), false},  // empty to addr
	}

	for _, tc := range cases {
		err := tc.tx.ValidateBasic()
		if tc.valid {
			suite.Require().Nil(err)
		}
	}
}

func (suite *IntegrationTestSuite) TestMsgSendGetSignBytes() {
	addr1 := sdk.AccAddress([]byte("input"))
	addr2 := sdk.AccAddress([]byte("output"))
	unlockerAddress := sdk.AccAddress([]byte("unlocker"))
	coins := sdk.NewCoins(sdk.NewInt64Coin("ctk", 10))
	var msg = NewMsgLockedSend(addr1, addr2, unlockerAddress.String(), coins)
	res := msg.GetSignBytes()

	expected := `{"type":"bank/MsgLockedSend","value":{"amount":[{"amount":"10","denom":"ctk"}],"from_address":"cosmos1d9h8qat57ljhcm","to_address":"cosmos1da6hgur4wsmpnjyg","unlocker_address":"cosmos1w4hxcmmrddjhy0qf5ju"}}`

	suite.Require().Equal(expected, string(res))
}

func (suite *IntegrationTestSuite)  TestMsgSendGetSigners() {
	unlockerAddress := sdk.AccAddress([]byte("unlocker"))
	var msg = NewMsgLockedSend(sdk.AccAddress([]byte("input1")), sdk.AccAddress{}, unlockerAddress.String(), sdk.NewCoins())
	res := msg.GetSigners()
	// TODO: fix this !
	suite.Require().Equal(fmt.Sprintf("%v", res), "[696E70757431]")
}

