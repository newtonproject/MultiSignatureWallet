package cli

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/console"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"
)

//  PayTransaction for send Transaction
type PayTransaction struct {
	To    common.Address `json:"to"`
	Value *big.Int       `json:"value"`
	Unit  string         `json:"unit"`
	Data  []byte         `json:"data"`
}

func (cli *CLI) buildBuildCmd() *cobra.Command {
	signTxCmd := &cobra.Command{
		Use:                   "build [--out outfile]",
		Short:                 "Build transaction",
		Long:                  "Build transaction in guide",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			var inStr string

			if cli.tran == nil {
				cli.tran = new(Transaction)
				cli.applyTranDefault()
			}

			if cmd.Flags().Changed("in") {
				inStr, err = cmd.Flags().GetString("in")
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				if err := cli.applyTxFile(inStr); err != nil {
					fmt.Printf("Error: apply infile(%s): %v\n", inStr, err)
					return
				}
			}

			offline, _ := cmd.Flags().GetBool("offline")

			if cmd.Flags().Changed("noguide") {
				if ok, _ := cmd.Flags().GetBool("noguide"); !ok {
					fmt.Println("Error: flag noguide changed but is false")
					return
				}
			} else {
				if err := cli.applyTxGuide(offline); err != nil {
					fmt.Println("Error:", err)
					return
				}
			}

			// update nonce, gasPrice, gasLimit, networkID from node
			if !offline {
				opts, err := cli.getNoSignTransactOpts()
				if err != nil {
					fmt.Println("Error:", err)
					return
				}

				simpleRegistry, err := cli.GetSimpleRegistry()
				if err != nil {
					fmt.Println("Error: GetSimpleRegistry Error: ", err)
					return
				}
				switch cli.tran.action {
				case Submit:
					//payTran := cli.tran.params.(*PayTransaction)
					if len(cli.tran.params) < 3 {
						err = errors.New("submitTransaction params length error")
					} else {
						_, err = simpleRegistry.SubmitTransaction(opts, cli.tran.params[0].(common.Address), cli.tran.params[1].(*big.Int), cli.tran.params[2].([]byte))
					}
				case Confirm:
					if len(cli.tran.params) < 1 {
						err = errors.New("confirmTransaction params length error")
					} else {
						_, err = simpleRegistry.ConfirmTransaction(opts, cli.tran.params[0].(*big.Int))
					}
				case Revoke:
					if len(cli.tran.params) < 1 {
						err = errors.New("revokeConfirmation params length error")
					} else {
						_, err = simpleRegistry.RevokeConfirmation(opts, cli.tran.params[0].(*big.Int))
					}
				case Execute:
					if len(cli.tran.params) < 1 {
						err = errors.New("executeTransaction params length error")
					} else {
						_, err = simpleRegistry.ExecuteTransaction(opts, cli.tran.params[0].(*big.Int))
					}
				// case OwnerAdd:
				// 	if len(cli.tran.params) < 1 {
				// 		err = errors.New("add owner params length error")
				// 	} else {
				// 		_, err = simpleRegistry.AddOwner(opts, cli.tran.params[0].(common.Address))
				// 	}
				// case OwnerRemove:
				// 	if len(cli.tran.params) < 1 {
				// 		err = errors.New("remove owner params length error")
				// 	} else {
				// 		_, err = simpleRegistry.RemoveOwner(opts, cli.tran.params[0].(common.Address))
				// 	}
				// case OwnerReplace:
				// 	if len(cli.tran.params) < 2 {
				// 		err = errors.New("replace owner params length error")
				// 	} else {
				// 		_, err = simpleRegistry.ReplaceOwner(opts, cli.tran.params[0].(common.Address), cli.tran.params[1].(common.Address))
				// 	}
				// case DailyLimit:
				// 	if len(cli.tran.params) < 1 {
				// 		err = errors.New("change daily limit params length error")
				// 	} else {
				// 		_, err = simpleRegistry.ChangeDailyLimit(opts, cli.tran.params[0].(*big.Int))
				// 	}
				// case Required:
				// 	if len(cli.tran.params) < 1 {
				// 		err = errors.New("change required params length error")
				// 	} else {
				// 		fmt.Println(cli.tran.params[0].(*big.Int).String())
				// 		_, err = simpleRegistry.ChangeRequirement(opts, cli.tran.params[0].(*big.Int))
				// 	}
				default:
					err = errors.New("unsupported function")
				}

				if err != nil {
					if err == errNoSignTransactor {
						// ok
					} else {
						fmt.Println("Error: ", err)
						return
					}
				} else {
					fmt.Println("ERROR")
					return
				}
			}

			tByte, err := cli.tran.MarshalJSON(true)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Transaction details are as follows:")
			fmt.Println(string(tByte))

			var outStr string
			defaultOutStr := time.Now().Format("20060102150405") + ".tx" // bitcoin 2009-01-03 18:15:05
			if cmd.Flags().Changed("out") {
				outStr, err = cmd.Flags().GetString("out")
			} else {
				prompt := fmt.Sprintf("Enter file to save transaction (default: %s): ", defaultOutStr)
				outStr, err = console.Stdin.PromptInput(prompt)
			}
			if err != nil {
				fmt.Println("Error:", err)
			}
			if outStr == "" {
				outStr = defaultOutStr
			}
			if err := cli.saveTranToFile(outStr); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Successfully save transaction to file", outStr)
			}

			// if sign, _ := cmd.Flags().GetBool("sign"); !sign && !offline {
			// 	return
			// }

			// if inStr == "" {
			// 	inStr = "tx"
			// }
			// cli.signTxAndSave(inStr + ".sign")
		},
	}

	signTxCmd.Flags().String("out", "", "file `path` to save built transaction")
	signTxCmd.Flags().String("in", "", "file `path` to load transaction to be built")
	signTxCmd.Flags().Bool("noguide", false, "disable guide to build transaction")
	// signTxCmd.Flags().Bool("sign", false, "sign transaction after build")
	signTxCmd.Flags().Bool("offline", false, "build offline transaction")

	return signTxCmd
}

var (
	errNoSignTransactor = errors.New("no sign transactor")
)

func (cli *CLI) getNoSignTransactOpts() (*bind.TransactOpts, error) {

	if cli.tran == nil {
		return nil, errCliTranNil
	}

	if err := cli.BuildClient(); err != nil {
		fmt.Println("NetworkID Error: ", err)
		return nil, err
	}
	networkID, err := cli.client.NetworkID(context.Background())
	if err != nil {
		fmt.Println("NetworkID Error: ", err)
		return nil, err
	}

	keyAddr := cli.tran.From
	opts := &bind.TransactOpts{
		From: keyAddr,
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != keyAddr {
				return nil, errors.New("not authorized to sign this account")
			}

			cli.copyTx(tx)
			cli.tran.From = address
			cli.tran.NetworkID = networkID

			return nil, errNoSignTransactor
		},
	}

	return opts, err
}

func (cli *CLI) copyTx(tx *types.Transaction) {
	if cli.tran == nil {
		cli.tran = new(Transaction)
	}
	if tx.To() != nil {
		cli.tran.To = *tx.To()
	}
	cli.tran.Value = tx.Value()
	cli.tran.Unit = UnitWEI
	cli.tran.Data = tx.Data()
	cli.tran.Nonce = tx.Nonce()
	cli.tran.GasPrice = tx.GasPrice()
	cli.tran.GasLimit = tx.Gas()
	cli.tran.NetworkID = tx.ChainId()
}
