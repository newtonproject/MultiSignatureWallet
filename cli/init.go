package cli

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/console"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func (cli *CLI) buildInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "init",
		Short:                 "Initialize config file",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("Initialize config file")

			prompt := fmt.Sprintf("Enter file in which to save (%s): ", defaultConfigFile)
			configPath, err := console.Stdin.PromptInput(prompt)
			if err != nil {
				fmt.Println("PromptInput err:", err)
			}
			if configPath == "" {
				configPath = defaultConfigFile
			}
			cli.config = configPath

			walletPathV := viper.GetString("walletPath")
			prompt = fmt.Sprintf("Enter the wallet storage directory (%s): ", walletPathV)
			cli.walletPath, err = console.Stdin.PromptInput(prompt)
			if err != nil {
				fmt.Println("PromptInput err:", err)
			}
			if cli.walletPath == "" {
				cli.walletPath = walletPathV
			}
			viper.Set("walletPath", cli.walletPath)

			rpcURLV := viper.GetString("rpcURL")
			prompt = fmt.Sprintf("Enter NewChain json rpc or ipc url (%s): ", rpcURLV)
			cli.rpcURL, err = console.Stdin.PromptInput(prompt)
			if err != nil {
				fmt.Println("PromptInput err:", err)
			}
			if cli.rpcURL == "" {
				cli.rpcURL = rpcURLV
			}
			viper.Set("rpcURL", cli.rpcURL)

			prompt = fmt.Sprintf("Create a default account or not: [Y/n] ")
			createNewAddress, err := console.Stdin.PromptInput(prompt)
			if err != nil {
				fmt.Println("PromptInput err:", err)
			}
			if len(createNewAddress) <= 0 {
				createNewAddress = "Y"
			}

			if strings.ToUpper(createNewAddress[:1]) == "Y" {
				wallet := keystore.NewKeyStore(cli.walletPath,
					keystore.StandardScryptN, keystore.StandardScryptP)

				cli.walletPassword, err = getPassPhrase("Your new account is locked with a password. Please give a password. Do not forget this password.", true)
				if err == nil {
					account, err := wallet.NewAccount(cli.walletPassword)
					if err == nil {
						baseAddress := account.Address.String()
						fmt.Println("New accout is ", baseAddress)
						viper.Set("from", baseAddress)

						getFaucet(cli.rpcURL, baseAddress)

					} else {
						fmt.Println("Account error:", err)
						fmt.Println("Just create your account later.")
					}
				} else {
					fmt.Println("Error: ", err)
					fmt.Println("Just create your account later.")
				}
			}

			err = viper.WriteConfigAs(configPath)
			if err != nil {
				fmt.Println("WriteConfig:", err)
			} else {
				fmt.Println("Your configuration has been saved in ", configPath)
			}
		},
	}

	return cmd
}
