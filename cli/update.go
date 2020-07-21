package cli

import (
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func (cli *CLI) buildUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [dailylimit|required]",
		Short: "Manage daily limit and requirement",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			return
		},
	}

	cmd.AddCommand(cli.buildUpdateDailyLimitCmd())
	cmd.AddCommand(cli.buildRequirementCmd())
	return cmd
}

func (cli *CLI) buildUpdateDailyLimitCmd() *cobra.Command {
	DailyLimitCmd := &cobra.Command{
		Use:   "dailylimit <number> [-u NEW|WEI]",
		Short: "change the daily limit",
		Long:  "Allows to change the daily limit. Transaction has to be sent by wallet",
		Args:  cobra.MinimumNArgs(1),
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			var err error

			fromAddress := viper.GetString("from")
			if fromAddress == "" || !common.IsHexAddress(fromAddress) {
				fmt.Println("Error: not set from address of owner")
				fmt.Fprint(os.Stderr, cmd.UsageString())
				return
			}

			amountStr := args[0]
			if amountStr == "" {
				fmt.Println("Error: not set daily limit ")
				fmt.Println(cmd.UsageString())
				return
			}
			unit, err := cmd.Flags().GetString("unit")
			if err != nil {
				fmt.Println("Error: required flag(s) \"unit\" not set")
				fmt.Fprint(os.Stderr, cmd.UsageString())
				return
			}
			d := stringInSlice(unit, DenominationList)
			if !d {
				fmt.Printf("Unit(%s) for amount error. %s.\n", unit, DenominationString)
				fmt.Fprint(os.Stderr, cmd.UsageString())
				return
			}

			amountWei, err := getAmountWei(amountStr, unit)
			if err != nil {
				fmt.Println("Get amount error:", err)
				fmt.Fprint(os.Stderr, cmd.UsageString())
				return
			}

			data, err := cli.GetMethodData("changeDailyLimit", amountWei)
			if err != nil {
				fmt.Println("GetMethodData error: ", err)
				return
			}

			cli.SubmitTransaction(fromAddress, common.HexToAddress(cli.contractAddress), big.NewInt(0), data)
		},
	}

	DailyLimitCmd.Flags().StringP("unit", "u", "NEW", fmt.Sprintf("unit for pay amount. %s.", DenominationString))

	return DailyLimitCmd
}

func (cli *CLI) buildRequirementCmd() *cobra.Command {
	RequirementCmd := &cobra.Command{
		Use:   "required <number>",
		Short: "change the number of required",
		Long:  "Allows to change the number of required confirmations. Transaction has to be sent by wallet",
		Args:  cobra.MinimumNArgs(1),
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			var err error

			fromAddress := viper.GetString("from")
			if fromAddress == "" || !common.IsHexAddress(fromAddress) {
				fmt.Println("Error: not set from address of owner")
				fmt.Fprint(os.Stderr, cmd.UsageString())
				return
			}

			numberStr := args[0]
			if numberStr == "" {
				fmt.Println("Error: not set daily limit ")
				fmt.Println(cmd.UsageString())
				return
			}
			number := new(big.Int)
			number, ok := number.SetString(numberStr, 10)
			if !ok {
				fmt.Println("Error: daily limit invalid")
				return
			}

			// check number
			if number.Cmp(big.NewInt(0)) <= 0 {
				number = big.NewInt(1)
			} else {
				simpleRegistry, err := cli.GetSimpleRegistry()
				if err != nil {
					fmt.Println("GetSimpleRegistry Error: ", err)
					return
				}
				owners, err := simpleRegistry.GetOwners(nil)
				if err != nil {
					log.Fatalf("Failed to call GetOwners: %v[%s]", err, cli.contractAddress)
				}
				olen := len(owners)
				olenBig := big.NewInt(int64(olen))
				if number.Cmp(olenBig) > 0 {
					number = olenBig
				}
			}

			data, err := cli.GetMethodData("changeRequirement", number)
			if err != nil {
				fmt.Println("GetMethodData error: ", err)
				return
			}

			cli.SubmitTransaction(fromAddress, common.HexToAddress(cli.contractAddress), big.NewInt(0), data)
		},
	}

	return RequirementCmd
}
