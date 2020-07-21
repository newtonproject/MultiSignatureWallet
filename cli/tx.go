package cli

import (
	"bytes"
	"fmt"
	"math/big"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func (cli *CLI) buildTxSubmitCmd() *cobra.Command {
	TxSubmitCmd := &cobra.Command{
		Use:                   "submit <amount> <-t target> [-u NEW,WEI] [-f source]",
		Short:                 "Submit a transaction, pay amount in unit to target address",
		Aliases:               []string{"pay"},
		Long:                  "Allows an owner to submit and confirm a transaction",
		Args:                  cobra.MinimumNArgs(1),
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {

			amountStr := args[0]
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

			fromAddress := viper.GetString("from")
			if fromAddress == "" || !common.IsHexAddress(fromAddress) {
				fmt.Println("Error: not set from address of owner")
				fmt.Fprint(os.Stderr, cmd.UsageString())
				return
			}

			toAddressStr, err := cmd.Flags().GetString("to")
			if err != nil {
				fmt.Println("Error: required flag(s) \"to\" not set")
				fmt.Fprint(os.Stderr, cmd.UsageString())
				return
			}
			toAddress := common.HexToAddress(toAddressStr)

			data, err := cmd.Flags().GetString("data")
			if err != nil {
				fmt.Println("Error:", err)
				fmt.Fprint(os.Stderr, cmd.UsageString())
				return
			}

			cli.SubmitTransaction(fromAddress, toAddress, amountWei, []byte(data))

		},
	}

	TxSubmitCmd.Flags().StringP("to", "t", "", "target account address or name")
	TxSubmitCmd.Flags().StringP("unit", "u", "NEW", fmt.Sprintf("unit for pay amount. %s.", DenominationString))
	TxSubmitCmd.Flags().String("data", "", "custom data message (use quotes if there are spaces)")

	TxSubmitCmd.MarkFlagRequired("to")

	return TxSubmitCmd
}

func (cli *CLI) buildTxConfirmCmd() *cobra.Command {
	TxSubmitCmd := &cobra.Command{
		Use:                   "confirm <transactionId>",
		Short:                 "Confirm transactionId",
		Long:                  "Allows an owner to confirm a transaction",
		Args:                  cobra.MinimumNArgs(1),
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {

			txIDStr := args[0]
			txID := new(big.Int)
			txID, ok := txID.SetString(txIDStr, 10)
			if !ok {
				fmt.Println("transactionId Error")
				return
			}

			fromAddress := viper.GetString("from")
			if fromAddress == "" || !common.IsHexAddress(fromAddress) {
				fmt.Println("Error: not set from address of owner")
				fmt.Fprint(os.Stderr, cmd.UsageString())
				return
			}

			cli.ConfirmTransaction(fromAddress, txID)

		},
	}

	return TxSubmitCmd
}

func (cli *CLI) buildTxRevokeCmd() *cobra.Command {
	TxSubmitCmd := &cobra.Command{
		Use:                   "revoke <transactionId>",
		Short:                 "Revoke transactionId",
		Long:                  "Allows an owner to revoke a confirmation for a transaction",
		Args:                  cobra.MinimumNArgs(1),
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {

			txIDStr := args[0]
			txID := new(big.Int)
			txID, ok := txID.SetString(txIDStr, 10)
			if !ok {
				fmt.Println("transactionId Error")
				return
			}

			fromAddress := viper.GetString("from")
			if fromAddress == "" || !common.IsHexAddress(fromAddress) {
				fmt.Println("Error: not set from address of owner")
				fmt.Fprint(os.Stderr, cmd.UsageString())
				return
			}

			cli.RevokeConfirmation(fromAddress, txID)

		},
	}

	return TxSubmitCmd
}

func (cli *CLI) buildTxExecuteCmd() *cobra.Command {
	TxSubmitCmd := &cobra.Command{
		Use:                   "execute <transactionId>",
		Short:                 "Execute transactionId",
		Long:                  "Allows anyone to execute a confirmed transaction",
		Args:                  cobra.MinimumNArgs(1),
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {

			txIDStr := args[0]
			txID := new(big.Int)
			txID, ok := txID.SetString(txIDStr, 10)
			if !ok {
				fmt.Println("transactionId Error")
				return
			}

			fromAddress := viper.GetString("from")
			if fromAddress == "" || !common.IsHexAddress(fromAddress) {
				fmt.Println("Error: not set from address of owner")
				fmt.Fprint(os.Stderr, cmd.UsageString())
				return
			}

			cli.ExecuteTransaction(fromAddress, txID)

		},
	}

	return TxSubmitCmd
}

func (cli *CLI) buildTxListCmd() *cobra.Command {
	TxListCmd := &cobra.Command{
		Use:                   "list [--pending] [--executed]",
		Short:                 "List of transaction IDs in defined range",
		Long:                  "List of transaction IDs in defined range",
		Args:                  cobra.MinimumNArgs(0),
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			var err error

			fromIndex, _ := cmd.Flags().GetInt64("fromindex")
			if err != nil {
				fmt.Println("from error")
				return
			}
			if fromIndex < 0 {
				fromIndex = 0
			}
			fromIndexBig := big.NewInt(fromIndex)

			var toIndexBig *big.Int
			toIndex, _ := cmd.Flags().GetInt64("toindex")
			if err != nil {
				fmt.Println("toindex error")
				return
			}

			pending, _ := cmd.Flags().GetBool("pending")
			executed, _ := cmd.Flags().GetBool("executed")
			if !(pending || executed) {
				executed = true
				pending = true
			}

			simpleRegistry, err := cli.GetSimpleRegistry()
			if err != nil {
				fmt.Println("GetSimpleRegistry Error: ", err)
				fmt.Println(cmd.UsageString())
				return
			}

			count, err := simpleRegistry.TransactionCount(nil)
			if err != nil {
				fmt.Println("TransactionCount err: ", err)
				return
			}

			if count.Cmp(big.NewInt(0)) <= 0 {
				fmt.Println("NO transaction ID")
				return
			}

			if toIndex == 0 || count.Cmp(big.NewInt(toIndex)) < 0 {
				toIndexBig = count
			} else {
				toIndexBig = big.NewInt(toIndex)
			}

			show := false
			for i := fromIndexBig; i.Cmp(toIndexBig) < 0; i = i.Add(i, big.NewInt(1)) {

				ok, err := simpleRegistry.IsConfirmed(nil, i)
				if err != nil {
					fmt.Printf("%s IsConfirmed Error(%v)\n", i, err)
					continue
				}

				t, err := simpleRegistry.Transactions(nil, i)
				if err != nil {
					fmt.Printf("%s Transactions Error(%v)\n", i, err)
					continue
				}

				var buffer bytes.Buffer
				buffer.WriteString(i.String())

				if ok {
					buffer.WriteString(" Confirmed")
				} else {
					buffer.WriteString(" Unconfirmed")
				}

				if t.Executed {
					if executed {
						show = true
						buffer.WriteString(" Executed")
						fmt.Println(buffer.String())
					}
				} else {
					if pending {
						show = true
						buffer.WriteString(" Pending")
						fmt.Println(buffer.String())
					}
				}
			}
			if !show {
				fmt.Println("NO matching transaction ID")
			}

		},
	}

	TxListCmd.Flags().Int64("fromindex", 0, "Index start position of transaction array")
	TxListCmd.Flags().Int64("toindex", 0, "Index end position of transaction array")
	TxListCmd.Flags().Bool("pending", false, "Only show pending transactions")
	TxListCmd.Flags().Bool("executed", false, "Only show executed transactions")

	return TxListCmd
}

func (cli *CLI) showTxInfo(cmd *cobra.Command, args []string) {
	if len(args) < 0 {
		return
	}
	txIDStr := args[0]
	txID := new(big.Int)
	txID, ok := txID.SetString(txIDStr, 10)
	if !ok {
		fmt.Println("transactionId Error")
		return
	}

	unit, _ := cmd.Flags().GetString("unit")
	if unit != "" && !stringInSlice(unit, DenominationList) {
		fmt.Printf("Unit(%s) for amount error. %s.\n", unit, DenominationString)
		fmt.Fprint(os.Stderr, cmd.UsageString())
		return
	}

	simpleRegistry, err := cli.GetSimpleRegistry()
	if err != nil {
		fmt.Println("GetSimpleRegistry Error: ", err)
		fmt.Println(cmd.UsageString())
		return
	}

	count, err := simpleRegistry.GetTransactionCount(nil, true, true)
	if err != nil {
		fmt.Printf("Failed to call GetTransactionCount: %v[%s]\n", err, cli.contractAddress)
		return
	}
	if txID.Cmp(count) >= 0 {
		fmt.Printf("TxID[%s] exceeds total number[%s] of transactions\n", txIDStr, count.String())
		return
	}

	t, err := simpleRegistry.Transactions(nil, txID)
	if err != nil {
		fmt.Printf("Failed to call GetConfirmations: %v[%s]\n", err, cli.contractAddress)
		return
	}

	fmt.Printf("Transaction ID %s basic information is as follows:\n", txIDStr)
	if t.Destination.String() == cli.contractAddress {
		fmt.Printf("Destination Address: %s (contract itself)\n", t.Destination.String())
	} else {
		fmt.Println("Destination Address: ", t.Destination.String())
	}
	fmt.Println("Value: ", getWeiAmountTextUnitByUnit(t.Value, unit))

	cli.CheckTransactionStatus(txID)

	if t.Executed {
		fmt.Println("Execution status: Executed")
	} else {
		fmt.Println("Execution status: Pending")
	}

	fmt.Printf("Data: 0x%s\n", common.Bytes2Hex(t.Data))
	showDataAuto(t.Data, "\t", unit)

}

func showDataAuto(data []byte, indent, unit string) {
	if len(data) >= 4 {
		parsed, err := abi.JSON(strings.NewReader(MultiSigWalletWithDailyLimitABI))
		if err != nil {
			fmt.Printf("JSON err: %v\n", err)
			return
		}
		sigdata, argdata := data[:4], data[4:]
		method, err := parsed.MethodById(sigdata)
		if err != nil || method == nil {
			// just data
			if utf8.Valid(data) {
				fmt.Printf("%s%s\n", indent, data)
			}
			return
		}
		fmt.Printf("%s%s\n", indent, method.String())

		nonIndexed := method.Inputs.NonIndexed()
		data, err := method.Inputs.UnpackValues(argdata)

		for k, v := range data {
			if nonIndexed[k].Type.T == abi.AddressTy {
				fmt.Printf("%s\t%s: %s \n", indent, nonIndexed[k].Name, v.(common.Address).String())
			} else if nonIndexed[k].Type.T == abi.BytesTy {
				vb := v.([]byte)
				fmt.Printf("%s\t%s: 0x%s\n", indent, nonIndexed[k].Name, common.Bytes2Hex(vb))
				if len(vb) > 0 && utf8.Valid(vb) {
					fmt.Printf("%s\t\t%s\n", indent, vb)
				}
			} else if nonIndexed[k].Name == "value" {
				vbig, ok := v.(*big.Int)
				if ok {
					fmt.Printf("%s\t%s: %v\n", indent, nonIndexed[k].Name, getWeiAmountTextUnitByUnit(vbig, unit))
				} else {
					fmt.Printf("%s\t%s: %v %v\n", indent, nonIndexed[k].Name, v, nonIndexed[k].Type.T)
				}
			} else {
				fmt.Printf("%s\t%s: %v\n", indent, nonIndexed[k].Name, v)
			}

		}
	}
}
