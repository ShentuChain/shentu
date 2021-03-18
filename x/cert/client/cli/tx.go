package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/certikfoundation/shentu/x/cert/types"
)

const (
	FlagAlias        = "alias"
	FlagContentType  = "content-type"
	FlagContent      = "content"
	FlagCompiler     = "compiler"
	FlagBytecodeHash = "bytecode-hash"
	FlagDescription  = "description"
	FlagCertifier    = "certifier"
	FlagPage         = "page"
	FlagLimit        = "limit"
)

// NewTxCmd returns the transaction commands for the certification module.
func NewTxCmd() *cobra.Command {
	certTxCmds := &cobra.Command{
		Use:   "cert",
		Short: "Certification transactions subcommands",
	}

	certTxCmds.AddCommand(
		GetCmdCertifyValidator(),
		GetCmdDecertifyValidator(),
		GetCmdCertifyPlatform(),
		GetCmdIssueCertificate(),
		GetCmdRevokeCertificate(),
	)

	return certTxCmds
}

// GetCmdCertifyValidator returns the validator certification transaction command.
func GetCmdCertifyValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "certify-validator <validator pubkey>",
		Short: "Certify a validator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			txf := tx.NewFactoryCLI(cliCtx, cmd.Flags()).WithTxConfig(cliCtx.TxConfig).WithAccountRetriever(cliCtx.AccountRetriever)

			from := cliCtx.GetFromAddress()
			if err := txf.AccountRetriever().EnsureExists(cliCtx, from); err != nil {
				return err
			}

			validator, err := sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeConsPub, args[0])
			if err != nil {
				return err
			}
			msg, err := types.NewMsgCertifyValidator(from, validator)
			if err != nil {
				return err
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxWithFactory(cliCtx, txf, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdDecertifyValidator returns the validator de-certification tx command.
func GetCmdDecertifyValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "decertify-validator <validator pubkey>",
		Short: "De-certify a validator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			txf := tx.NewFactoryCLI(cliCtx, cmd.Flags()).WithTxConfig(cliCtx.TxConfig).WithAccountRetriever(cliCtx.AccountRetriever)

			from := cliCtx.GetFromAddress()
			if err := txf.AccountRetriever().EnsureExists(cliCtx, from); err != nil {
				return err
			}

			validator, err := sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeConsPub, args[0])
			if err != nil {
				return err
			}
			msg, err := types.NewMsgDecertifyValidator(from, validator)
			if err != nil {
				return err
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxWithFactory(cliCtx, txf, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdIssueCertificate returns the certificate transaction command.
func GetCmdIssueCertificate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "issue-certificate <certificate type> <request content type> <request content> [<flags>]",
		Short: "Issue a certificate",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			txf := tx.NewFactoryCLI(cliCtx, cmd.Flags()).WithTxConfig(cliCtx.TxConfig).WithAccountRetriever(cliCtx.AccountRetriever)

			from := cliCtx.GetFromAddress()
			if err := txf.AccountRetriever().EnsureExists(cliCtx, from); err != nil {
				return err
			}

			certificateTypeString := strings.ToLower(args[0])
			switch certificateTypeString {
			case "compilation":
				contentType := types.RequestContentTypeFromString(args[1])
				if contentType != types.RequestContentTypeSourceCodeHash {
					return types.ErrInvalidRequestContentType
				}
				compiler, bytecodeHash, description, err := parseCertifyCompilationFlags()
				if err != nil {
					return err
				}
				msg := types.NewMsgCertifyCompilation(args[2], compiler, bytecodeHash, description, from)
				if err := msg.ValidateBasic(); err != nil {
					return err
				}
				return tx.GenerateOrBroadcastTxWithFactory(cliCtx, txf, msg)

			default:
				description := viper.GetString(FlagDescription)
				msg := types.NewMsgCertifyGeneral(certificateTypeString, args[1], args[2], description, from)
				if err := msg.ValidateBasic(); err != nil {
					return err
				}
				return tx.GenerateOrBroadcastTxWithFactory(cliCtx, txf, msg)
			}
		},
	}

	cmd.Flags().String(FlagCompiler, "", "compiler version")
	cmd.Flags().String(FlagBytecodeHash, "", "bytecode hash")
	cmd.Flags().String(FlagDescription, "", "description")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// parseCertifyCompilation parses flags for compilation certificate.
func parseCertifyCompilationFlags() (string, string, string, error) {
	compiler := viper.GetString(FlagCompiler)
	if compiler == "" {
		return "", "", "", fmt.Errorf("compiler version is required to issue a compilation certificate")
	}
	bytecodeHash := viper.GetString(FlagBytecodeHash)
	if bytecodeHash == "" {
		return "", "", "", fmt.Errorf("bytecode hash is required to issue a compilation certificate")
	}
	description := viper.GetString(FlagDescription)
	return compiler, bytecodeHash, description, nil
}

// GetCmdCertifyPlatform returns the validator host platform certification transaction command.
func GetCmdCertifyPlatform() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "certify-platform <validator pubkey> <platform>",
		Short: "Certify a validator's host platform",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			txf := tx.NewFactoryCLI(cliCtx, cmd.Flags()).WithTxConfig(cliCtx.TxConfig).WithAccountRetriever(cliCtx.AccountRetriever)

			from := cliCtx.GetFromAddress()
			if err := txf.AccountRetriever().EnsureExists(cliCtx, from); err != nil {
				return err
			}

			validator, err := sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeConsPub, args[0])
			if err != nil {
				return err
			}

			msg, err := types.NewMsgCertifyPlatform(from, validator, args[1])
			if err != nil {
				return err
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxWithFactory(cliCtx, txf, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdRevokeCertificate returns the certificate revoke command
func GetCmdRevokeCertificate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke-certificate <certificateID> [<description>]",
		Short: "revoke a certificate",
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			txf := tx.NewFactoryCLI(cliCtx, cmd.Flags()).WithTxConfig(cliCtx.TxConfig).WithAccountRetriever(cliCtx.AccountRetriever)

			from := cliCtx.GetFromAddress()
			if err := txf.AccountRetriever().EnsureExists(cliCtx, from); err != nil {
				return err
			}

			certificateID, err := strconv.ParseUint(args[0],10, 64)
			if err != nil {
				return err
			}

			var description string
			if len(args) > 1 {
				description = args[1]
			}

			msg := types.NewMsgRevokeCertificate(from, certificateID, description)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxWithFactory(cliCtx, txf, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdSubmitProposal implements the command to submit a certifier-update proposal
func GetCmdSubmitProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "certifier-update [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a certifier update proposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit a certifier update proposal along with an initial deposit.
The proposal details must be supplied via a JSON file.
Example:
$ %s tx gov submit-proposal certifier-update <path/to/proposal.json> --from=<key_or_address>
Where proposal.json contains:
{
  "title": "New Certifier, Joe Shmoe",
  "description": "Why we should make Joe Shmoe a certifier",
  "certifier": "certik1s5afhd6gxevu37mkqcvvsj8qeylhn0rz46zdlq",
  "add_or_remove": "add",
  "alias": "joe",
  "deposit": [
    {
      "denom": "ctk",
      "amount": "100"
    }
  ]
}
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			txf := tx.NewFactoryCLI(cliCtx, cmd.Flags()).WithTxConfig(cliCtx.TxConfig).WithAccountRetriever(cliCtx.AccountRetriever)

			from := cliCtx.GetFromAddress()
			if err := txf.AccountRetriever().EnsureExists(cliCtx, from); err != nil {
				return err
			}

			proposal, err := ParseCertifierUpdateProposalJSON(cliCtx.LegacyAmino, args[0])
			if err != nil {
				return err
			}

			content := types.NewCertifierUpdateProposal(
				proposal.Title,
				proposal.Description,
				proposal.Certifier,
				proposal.Alias,
				from,
				proposal.AddOrRemove,
			)

			msg, err := govtypes.NewMsgSubmitProposal(content, proposal.Deposit, from)
			if err != nil {
				return err
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxWithFactory(cliCtx, txf, msg)
		},
	}

	return cmd
}
