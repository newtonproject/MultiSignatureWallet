package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/spf13/cobra"
)

func (cli *CLI) buildSignCmd() *cobra.Command {
	signTxCmd := &cobra.Command{
		Use:                   "sign <filepath> [-u NEW|WEI]",
		Short:                 "Sign the transaction in the file",
		Args:                  cobra.MinimumNArgs(1),
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			var unit string
			if cmd.Flags().Changed("unit") {
				unit, _ = cmd.Flags().GetString("unit")
				if unit != "" && !stringInSlice(unit, UnitList) {
					fmt.Printf("Unit(%s) for amount error. %s.\n", unit, fmt.Sprintf("Available unit: %s", strings.Join(UnitList, ",")))
					fmt.Fprint(os.Stderr, cmd.UsageString())
					return
				}
			}

			if cli.tran == nil {
				cli.tran = new(Transaction)
				cli.applyTranDefault()
			}

			infileStr := args[0]

			if err := cli.applyTxFile(infileStr); err != nil {
				fmt.Printf("Error apply infile(%s): %v\n", infileStr, err)
				return
			}

			fmt.Println("Transaction details are as follows:")
			cli.printTxIndent()
			fmt.Println("The data is as follows:")
			showDataAuto(cli.tran.Data, "", unit)

			var outStr string
			var err error
			if cmd.Flags().Changed("out") {
				outStr, err = cmd.Flags().GetString("out")
				if err != nil {
					fmt.Println(err)
				}
			}

			if outStr == "" {
				if infileStr == "" {
					outStr = "tx.sign"
				} else {
					outStr = infileStr + ".sign"
				}
			}
			cli.signTxAndSave(outStr)
		},
	}

	signTxCmd.Flags().String("out", "", "file `path` to save signed transaction")
	signTxCmd.Flags().StringP("unit", "u", UnitETH, fmt.Sprintf("unit for pay amount. %s.", fmt.Sprintf("Available unit: %s", strings.Join(UnitList, ","))))

	return signTxCmd
}

func (cli *CLI) printTxIndent() {
	if cli.tran != nil {
		tByte, err := cli.tran.MarshalJSON(true)
		if err == nil {
			fmt.Println(string(tByte))
		}
	}
}

func (cli *CLI) signTxAndSave(filepath string) {
	signTx, err := cli.unlockAndSignTx()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Signed Transaction Hash: ", signTx.Hash().String())

	data, err := rlp.EncodeToBytes(signTx)
	if err != nil {
		fmt.Println(err)
		return
	}
	dataHex := common.ToHex(data)
	fmt.Printf("Signed Transaction: %s\n", dataHex)

	if err := saveStringToFile(dataHex, filepath); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully save signed transacion hex to file", filepath)
}

func (cli *CLI) unlockAndSignTx() (*types.Transaction, error) {
	account := accounts.Account{Address: cli.tran.From}
	if err := cli.unlockWallet(account); err != nil {
		return nil, err
	}

	return cli.signTx()
}

func (cli *CLI) unlockWallet(account accounts.Account) error {
	if cli.wallet == nil {
		if err := cli.openWallet(true); err != nil {
			return err
		}
	}
	if account.Address == (common.Address{}) {
		return errRequiredFromAddress
	}
	if _, err := cli.wallet.Find(account); err != nil {
		return fmt.Errorf("%v (%s)", err, account.Address.String())
	}

	var trials int
	var err error
	walletPassword := cli.tran.Password
	for trials = 0; trials < 3; trials++ {
		prompt := fmt.Sprintf("Unlocking account %s | Attempt %d/%d", account.Address.String(), trials+1, 3)
		if walletPassword == "" {
			walletPassword, _ = getPassPhrase(prompt, false)
		} else {
			fmt.Println(prompt, "\nUse the the password has set")
		}
		err = cli.wallet.Unlock(account, walletPassword)
		if err == nil {
			break
		}
		walletPassword = ""
	}

	if trials >= 3 {
		if err != nil {
			return err
		}
		return fmt.Errorf("Error: Failed to unlock account %s (%v)", account.Address.String(), err)
	}

	return nil
}

func (cli *CLI) openWallet(check bool) error {
	if cli.wallet == nil {
		cli.wallet = keystore.NewKeyStore(cli.walletPath,
			keystore.LightScryptN, keystore.LightScryptP)
	}

	if check && len(cli.wallet.Accounts()) == 0 {
		return errWalletPathEmppty
	}
	return nil
}

func (cli *CLI) signTx() (*types.Transaction, error) {
	if cli.tran == nil {
		return nil, errCliTranNil
	}
	tx := types.NewTransaction(cli.tran.Nonce, cli.tran.To, cli.tran.Value, cli.tran.GasLimit, cli.tran.GasPrice, cli.tran.Data)
	signTx, err := cli.wallet.SignTx(accounts.Account{Address: cli.tran.From}, tx, cli.tran.NetworkID)
	if err != nil {
		return nil, err
	}
	return signTx, nil
}
