package cli

import (
	"bufio"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/certikfoundation/shentu/x/auth/internal/types"
)

// GetTxCmd returns the transaction commands for this module.
// NOTE: Auth tx commands from Cosmos are mounted directly under the root.
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Auth transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	txCmd.AddCommand(
		GetCmdManualVesting(cdc),
	)
	return txCmd
}

// GetCmdManualVesting implements the command for unlocking
// the specified amount in a manual vesting account.
func GetCmdManualVesting(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "manual-vesting [address] [amount]",
		Short: "Unlock the amount from a manual vesting account's vesting coins.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := authtxb.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			accGetter := authtxb.NewAccountRetriever(cliCtx)

			if _, err := accGetter.GetAccount(cliCtx.GetFromAddress()); err != nil {
				return err
			}

			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoin(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgManualVesting(cliCtx.GetFromAddress(), addr, amount)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	cmd = flags.PostCommands(cmd)[0]
	return cmd
}
