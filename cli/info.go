package cli

import (
	"context"
	"fmt"
	"os"
	"strings"

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

			return
		},
	}

	cmd.Flags().StringP("unit", "u", "", fmt.Sprintf("unit for value. %s.", fmt.Sprintf("Available unit: %s", strings.Join(UnitList, ","))))

	return cmd
}
