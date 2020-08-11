package cli

import (
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func (cli *CLI) buildDeployCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "deploy <-o addr0,addr1,addr3> <-r number> [-l dailyLimitAmountInUnit] [-u NEW|WEI]",
		Short:                 fmt.Sprintf("Deploy %s contract", cli.bc.String()),
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			var err error

			save, _ := cmd.Flags().GetBool("save")

			unit, err := cmd.Flags().GetString("unit")
			if err != nil {
				fmt.Println("Error: required flag(s) \"unit\" not set")
				fmt.Fprint(os.Stderr, cmd.UsageString())
				return
			}
			d := stringInSlice(unit, UnitList)
			if !d {
				fmt.Printf("Unit(%s) for amount error. %s.\n", unit, fmt.Sprintf("Available unit: %s", strings.Join(UnitList, ",")))
				fmt.Fprint(os.Stderr, cmd.UsageString())
				return
			}

			fromAddress := viper.GetString("from")
			if fromAddress == "" || !common.IsHexAddress(fromAddress) {
				fmt.Println("Error: not set from address of owner")
				fmt.Println(cmd.UsageString())
				return
			}

			owners, _ := cmd.Flags().GetString("owners")
			if owners == "" {
				fmt.Println("Error: not set owners")
				fmt.Println(cmd.UsageString())
				return
			}

			required, _ := cmd.Flags().GetInt64("required")
			if required <= 0 || required > 50 {
				fmt.Println("Error: not set required")
				fmt.Println(cmd.UsageString())
				return
			}
			requiredBig := big.NewInt(required)

			dailyLimitWei := new(big.Int)
			dailyLimitStr, _ := cmd.Flags().GetString("dailylimit")
			if dailyLimitStr != "" {
				if !IsUintString(dailyLimitStr) {
					fmt.Printf("DailyLimitStr(%v) illegal\n", dailyLimitStr)
					return
				}
				dailyLimitWei, err = getAmountWei(dailyLimitStr, unit)
				if err != nil {
					fmt.Printf("Get amount error(%v): %v\n", err, dailyLimitStr)
					return
				}
			}

			ownerList := strings.Split(owners, ",")
			var ownerlist []common.Address
			ownercheck := make(map[string]bool)
			for _, owner := range ownerList {
				if !common.IsHexAddress(owner) {
					fmt.Printf("Error: address of owner(%v) illegal\n", owner)
					return
				}
				if ownercheck[owner] {
					fmt.Printf("Error: repeated owner(%s)\n", owner)
					return
				}
				ownercheck[owner] = true
				ownerlist = append(ownerlist, common.HexToAddress(owner))
			}
			if len(ownerlist) < int(required) {
				fmt.Printf("Required(%v) is greater than the number (%v) of owners\n", required, len(ownerList))
				return
			}

			if cli.contractAddress == "" {
				save = true
			}
			if err := cli.Deploy(fromAddress, ownerlist, requiredBig, dailyLimitWei); err != nil {
				fmt.Println(err)
				return
			}

			if save {
				viper.WriteConfigAs(cli.config)
			}
		},
	}

	cmd.Flags().StringP("owners", "o", "", "the list of initial owners `address`es, separated by commas(,)")
	cmd.Flags().Int64P("required", "r", 0, "the `number` of required confirmations, maximum is 50") // how to get 50, contract not deploy?
	cmd.Flags().StringP("dailylimit", "l", "0", "the `amount` in unit, which can be withdrawn without confirmations on a daily basis")
	cmd.Flags().StringP("unit", "u", UnitETH, fmt.Sprintf("unit for daily limit. %s.", fmt.Sprintf("Available unit: %s", strings.Join(UnitList, ","))))
	cmd.Flags().BoolP("save", "s", false, "save contract address to config file")

	cmd.MarkFlagRequired("owners")
	cmd.MarkFlagRequired("required")
	// cmd.MarkFlagRequired("dailylimit")

	return cmd
}
