package cli

import (
	"github.com/spf13/cobra"
)

func (cli *CLI) buildRootCmd() {

	if cli.rootCmd != nil {
		cli.rootCmd.ResetFlags()
		cli.rootCmd.ResetCommands()
	}

	rootCmd := &cobra.Command{
		Use:              cli.Name,
		Short:            cli.Name + " is commandline client for users to interact with the MultiSigWallet contract.",
		Run:              cli.help,
		PersistentPreRun: cli.setup,
	}
	cli.rootCmd = rootCmd

	// Global flags
	rootCmd.PersistentFlags().StringVarP(&cli.config, "config", "c", defaultConfigFile, "The `path` to config file")
	rootCmd.PersistentFlags().StringP("walletPath", "w", defaultWalletPath, "Wallet storage `directory`")
	rootCmd.PersistentFlags().StringP("rpcURL", "i", defaultRPCURL, "NewChain json rpc or ipc `url`")
	rootCmd.PersistentFlags().StringP("contractAddress", "a", defaultContractAddress, "Contract `address`")
	rootCmd.PersistentFlags().StringP("from", "f", "", "the from `address` who pay gas")

	// Basic commands
	rootCmd.AddCommand(cli.buildInitCmd())    // init
	rootCmd.AddCommand(cli.buildVersionCmd()) // version

	// deploy
	rootCmd.AddCommand(cli.buildDeployCmd())

	// account
	rootCmd.AddCommand(cli.buildAccountCmd())

	// info
	rootCmd.AddCommand(cli.buildInfoCmd())

	// tx
	rootCmd.AddCommand(cli.buildTxSubmitCmd())
	rootCmd.AddCommand(cli.buildTxConfirmCmd())
	rootCmd.AddCommand(cli.buildTxRevokeCmd())
	rootCmd.AddCommand(cli.buildTxExecuteCmd())
	rootCmd.AddCommand(cli.buildTxListCmd())

	// owner
	rootCmd.AddCommand(cli.buildOwnerCmd())

	// update
	rootCmd.AddCommand(cli.buildUpdateCmd())

	// tx for offline
	rootCmd.AddCommand(cli.buildBuildCmd())     // build
	rootCmd.AddCommand(cli.buildSignCmd())      // sign
	rootCmd.AddCommand(cli.buildBroadcastCmd()) // broadcast

}
