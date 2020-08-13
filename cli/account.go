package cli

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func (cli *CLI) buildAccountCmd() *cobra.Command {
	use := "account [new|list|balance]"
	if cli.bc == NewChain {
		use = "account [new|list|balance|convert]"
	}

	cmd := &cobra.Command{
		Use:   use,
		Short: fmt.Sprintf("Manage %s accounts", cli.bc.String()),
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			return
		},
	}

	cmd.AddCommand(cli.buildAccountNewCmd())
	cmd.AddCommand(cli.buildAccountListCmd())
	cmd.AddCommand(cli.buildBalanceCmd())
	if cli.bc == NewChain {
		cmd.AddCommand(cli.buildAccountConvertCmd())
	}

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

func (cli *CLI) buildAccountConvertCmd() *cobra.Command {
	accountListCmd := &cobra.Command{
		Use:                   "convert",
		Short:                 "convert address to NewChainAddress",
		Args:                  cobra.MinimumNArgs(1),
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {

			err := cli.BuildClient()
			if err != nil {
				fmt.Printf("Error: build client error(%v)\n", err)
				return
			}
			chainID, err := cli.client.NetworkID(context.Background())
			if err != nil {
				fmt.Printf("Error: get chainID error(%v), use chainID as %s\n", err, chainID.String())
				return
			}

			for _, addressStr := range args {
				if common.IsHexAddress(addressStr) {
					address := common.HexToAddress(addressStr)
					fmt.Println(address.String(), addressToNew(chainID.Bytes(), address))
					continue
				}

				address, err := newToAddress(chainID.Bytes(), addressStr)
				if err != nil {
					fmt.Println(err, addressStr)
					continue
				}
				fmt.Println(address.String(), addressStr)
			}

		},
	}

	return accountListCmd
}

func addressToNew(chainID []byte, address common.Address) string {
	input := append(chainID, address.Bytes()...)
	return "NEW" + base58.CheckEncode(input, 0)
}

func newToAddress(chainID []byte, newAddress string) (common.Address, error) {
	if newAddress[:3] != "NEW" {
		return common.Address{}, errors.New("not NEW address")
	}

	decoded, version, err := base58.CheckDecode(newAddress[3:])
	if err != nil {
		return common.Address{}, err
	}
	if version != 0 {
		return common.Address{}, errors.New("illegal version")
	}
	if len(decoded) < 20 {
		return common.Address{}, errors.New("illegal decoded length")
	}
	if !bytes.Equal(decoded[:len(decoded)-20], chainID) {
		return common.Address{}, errors.New("illegal ChainID")
	}

	address := common.BytesToAddress(decoded[len(decoded)-20:])

	return address, nil
}
