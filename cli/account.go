package cli

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func (cli *CLI) buildAccountCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account [new|list|balance]",
		Short: fmt.Sprintf("Manage %s accounts", cli.bc.String()),
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			return
		},
	}

	cmd.AddCommand(cli.buildAccountNewCmd())
	cmd.AddCommand(cli.buildAccountListCmd())
	cmd.AddCommand(cli.buildBalanceCmd())

	return cmd
}

func (cli *CLI) buildAccountNewCmd() *cobra.Command {
	accountNewCmd := &cobra.Command{
		Use:                   "new [-n number] [--faucet] [-s] [-l]",
		Short:                 "create a new account",
		Args:                  cobra.MinimumNArgs(0),
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			walletPath := cli.walletPath
			if cmd.Flags().Changed("light") {
				light, _ := cmd.Flags().GetBool("light")
				standard, _ := cmd.Flags().GetBool("standard")
				if light && !standard {
					cli.wallet = keystore.NewKeyStore(walletPath,
						keystore.LightScryptN, keystore.LightScryptP)
				}
			}
			if cli.wallet == nil {
				cli.wallet = keystore.NewKeyStore(walletPath,
					keystore.StandardScryptN, keystore.StandardScryptP)
			}

			if cli.walletPassword == "" {
				cli.walletPassword, err = getPassPhrase("Your new account is locked with a password. Please give a password. Do not forget this password.", true)
				if err != nil {
					fmt.Println("Error: ", err)
					return
				}
			}

			numOfNew, err := cmd.Flags().GetInt("numOfNew")
			if err != nil {
				numOfNew = viper.GetInt("account.numOfNew")
			}
			if numOfNew <= 0 {
				fmt.Printf("number[%d] of new account less then 1\n", numOfNew)
				numOfNew = 1
			}

			faucet, _ := cmd.Flags().GetBool("faucet")

			for i := 0; i < numOfNew; i++ {
				account, err := cli.wallet.NewAccount(cli.walletPassword)
				if err != nil {
					fmt.Println("Account error:", err)
					return
				}
				fmt.Println(account.Address.Hex())
				if faucet {
					getFaucet(cli.rpcURL, account.Address.String())
				}
				if cli.address == "" {
					cli.address = account.Address.String()
				}
			}
		},
	}

	accountNewCmd.Flags().IntP("numOfNew", "n", 1, "number of the new account")
	accountNewCmd.Flags().Bool("faucet", false, "get faucet for new account")
	accountNewCmd.Flags().BoolP("standard", "s", false, "use the standard scrypt for keystore")
	accountNewCmd.Flags().BoolP("light", "l", false, "use the light scrypt for keystore")
	return accountNewCmd
}

func (cli *CLI) buildAccountListCmd() *cobra.Command {
	accountListCmd := &cobra.Command{
		Use:   "list",
		Short: "list all accounts in the wallet path",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			walletPath := cli.walletPath
			wallet := keystore.NewKeyStore(walletPath,
				keystore.LightScryptN, keystore.LightScryptP)
			if len(wallet.Accounts()) == 0 {
				fmt.Println("Empty wallet, create account first.")
				return
			}

			for _, account := range wallet.Accounts() {
				fmt.Println(account.Address.Hex())
			}
		},
	}

	return accountListCmd
}

func (cli *CLI) buildBalanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   fmt.Sprintf("balance [-u %s] [address1] [address2]...", strings.Join(UnitList, "|")),
		Short:                 "Get balance of address",
		Args:                  cobra.MinimumNArgs(0),
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {

			unit, _ := cmd.Flags().GetString("unit")
			if unit != "" && !stringInSlice(unit, UnitList) {
				fmt.Printf("Unit(%s) for invalid. %s.\n", unit, fmt.Sprintf("Available unit: %s", strings.Join(UnitList, ",")))
				fmt.Fprint(os.Stderr, cmd.UsageString())
				return
			}

			var addressList []common.Address

			if len(args) <= 0 {
				if err := cli.openWallet(true); err != nil {
					fmt.Println(err)
					return
				}

				for _, account := range cli.wallet.Accounts() {
					addressList = append(addressList, account.Address)
				}

			} else {
				for _, addressStr := range args {
					addressList = append(addressList, common.HexToAddress(addressStr))
				}
			}

			for _, address := range addressList {
				balance, err := cli.getBalance(address)
				if err != nil {
					fmt.Println("Balance error:", err)
					return
				}
				fmt.Printf("Address[%s] Balance[%s]\n", address.Hex(), getWeiAmountTextUnitByUnit(balance, unit))
			}

			return
		},
	}

	cmd.Flags().StringP("unit", "u", "", fmt.Sprintf("unit for balance. %s.", fmt.Sprintf("Available unit: %s", strings.Join(UnitList, ","))))

	return cmd
}

func (cli *CLI) getBalance(address common.Address) (*big.Int, error) {
	if err := cli.BuildClient(); err != nil {
		return nil, err
	}
	return cli.client.BalanceAt(context.Background(), address, nil)
}
