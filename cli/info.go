package cli

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

func (cli *CLI) buildInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "info [transactionID] [-a contractAddress] [-u NEW|WEI]",
		Short:                 "Show the basic info of contract wallet or a transaction ID",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) > 0 {
				cli.showTxInfo(cmd, args)
				return
			}

			unit, _ := cmd.Flags().GetString("unit")
			if unit != "" && !stringInSlice(unit, UnitList) {
				fmt.Printf("Unit(%s) for invalid. %s.\n", unit, fmt.Sprintf("Available unit: %s", strings.Join(UnitList, ",")))
				fmt.Fprint(os.Stderr, cmd.UsageString())
				return
			}

			simpleRegistry, err := cli.GetSimpleRegistry()
			if err != nil {
				fmt.Println("GetSimpleRegistry Error: ", err)
				fmt.Println(cmd.UsageString())
				return
			}

			fmt.Printf("The contract address(%s) basic information is as follows:\n", cli.contractAddress)

			ctx := context.Background()
			balance, err := cli.client.BalanceAt(ctx, common.HexToAddress(cli.contractAddress), nil)
			if err != nil {
				fmt.Printf("Balance: BalanceAt Error(%v)\n", err)
			} else {
				fmt.Println("Balance: ", getWeiAmountTextUnitByUnit(balance, unit))
			}

			owners, err := simpleRegistry.GetOwners(nil)
			if err != nil {
				fmt.Printf("Owners List: GetOwners Error(%v)\n", err)
			} else {
				fmt.Println("Owners List: ")
				for _, v := range owners {
					fmt.Println("\t", v.String())
				}
			}

			required, err := simpleRegistry.Required(nil)
			if err != nil {
				fmt.Printf("The number of required confirmations: Required Error (%v)\n", err)
			} else {
				fmt.Println("The number of required confirmations: ", required.String())
			}

			dailyLimit, err := simpleRegistry.DailyLimit(nil)
			if err != nil {
				fmt.Printf("Daily Limit: DailyLimit Error(%v)\n", err)
			} else {
				fmt.Println("Daily Limit:", getWeiAmountTextUnitByUnit(dailyLimit, unit))
			}

			max, err := simpleRegistry.CalcMaxWithdraw(nil)
			if err != nil {
				fmt.Printf("\tToday Remaining Limit: CalcMaxWithdraw Error(%v)\n", err)
			} else {
				fmt.Println("\tToday Remaining Limit:", getWeiAmountTextUnitByUnit(max, unit))
			}

			spentToday, err := simpleRegistry.SpentToday(nil)
			if err != nil {
				fmt.Printf("\tToday Spent Limit: SpentToday Error(%v)\n", err)
			} else {
				fmt.Println("\tToday Spent Limit:", getWeiAmountTextUnitByUnit(spentToday, unit))
			}

			if cmd.Flags().Changed("token") {
				tokenAddressStr, err := cmd.Flags().GetString("token")
				if err != nil {
					fmt.Printf("Balance: get token address error\n")
					return
				}
				tokenAddress := common.HexToAddress(tokenAddressStr)

				parsed, err := abi.JSON(strings.NewReader(ERC20TransferABI))
				if err != nil {
					fmt.Printf("Balance(%v): JSON err: %v\n",
						tokenAddress.String(), err)
					return
				}

				err = cli.BuildClient()
				if err != nil {
					fmt.Printf("Balance(%v): %v\n",
						tokenAddress.String(), err)
					return
				}
				erc20 := bind.NewBoundContract(tokenAddress, parsed, cli.client, cli.client, cli.client)

				var decimals uint8
				{
					// get decimals
					var (
						ret0 = new(uint8)
					)
					out := ret0
					err = erc20.Call(nil, out, "decimals")
					if err != nil {
						fmt.Printf("Balance(%v): %v\n",
							tokenAddress.String(), err)
						return
					}

					decimals = *out
				}

				var symbol string
				{
					// get name
					var (
						ret0 = new(string)
					)
					out := ret0
					err = erc20.Call(nil, out, "symbol")
					if err != nil {
						fmt.Printf("Balance(%v): %v\n",
							tokenAddress.String(), err)
						return
					}

					symbol = *out
				}

				var balance *big.Int
				{
					// get balance
					var (
						ret0 = new(*big.Int)
					)
					out := ret0

					err = erc20.Call(nil, out, "balanceOf", common.HexToAddress(cli.contractAddress))
					if err != nil {
						fmt.Printf("Balance(%v): %v\n",
							tokenAddress.String(), err)
						return
					}

					balance = big.NewInt(0).Set(*ret0)
				}

				fmt.Printf("Balance(%v): %v %s\n",
					tokenAddress.String(),
					getAmountTextByWeiWithDecimals(balance, decimals),
					symbol)
			}

			return
		},
	}

	cmd.Flags().StringP("unit", "u", "", fmt.Sprintf("unit for value. %s.", fmt.Sprintf("Available unit: %s", strings.Join(UnitList, ","))))
	cmd.Flags().String("token", "", "the address of token, if set then show the balance of this token")

	return cmd
}
