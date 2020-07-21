package cli

import (
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func (cli *CLI) buildOwnerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "owner [list|add|check|remove|replace]",
		Short: "Manage contract owners",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			return
		},
	}

	cmd.AddCommand(cli.buildOwnerListCmd())
	cmd.AddCommand(cli.buildOwnerAddCmd())
	cmd.AddCommand(cli.buildOwnerRemoveCmd())
	cmd.AddCommand(cli.buildOwnerReplaceCmd())
	cmd.AddCommand(cli.buildOwnerCheckCmd())

	return cmd
}

func (cli *CLI) buildOwnerListCmd() *cobra.Command {
	OwnerListCmd := &cobra.Command{
		Use:   "list",
		Short: "List all owners",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			cli.OwnerList()
			return
		},
	}

	return OwnerListCmd
}

func (cli *CLI) buildOwnerAddCmd() *cobra.Command {
	OwnerAddCmd := &cobra.Command{
		Use:                   "add <owner>",
		Short:                 "Allows to add a new owner. Transaction has to be sent by wallet",
		Args:                  cobra.MinimumNArgs(1),
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			var err error

			fromAddress := viper.GetString("from")
			if fromAddress == "" || !common.IsHexAddress(fromAddress) {
				fmt.Println("Error: not set from address of owner")
				fmt.Fprint(os.Stderr, cmd.UsageString())
				return
			}

			ownerStr := args[0]
			if ownerStr == "" && !common.IsHexAddress(ownerStr) {
				fmt.Println("Error: not set owners or invalid")
				fmt.Println(cmd.UsageString())
				return
			}
			owner := common.HexToAddress(ownerStr)

			simpleRegistry, err := cli.GetSimpleRegistry()
			if err != nil {
				fmt.Println("GetSimpleRegistry Error: ", err)
				return
			}
			isowner, err := simpleRegistry.IsOwner(nil, owner)
			if err != nil {
				fmt.Printf("Failed to call GetOwners: %v[%s]\n", err, cli.contractAddress)
				return
			}
			if isowner {
				fmt.Printf("Address[%s] is one of the owners\n", ownerStr)
				return
			}

			data, err := cli.GetMethodData("addOwner", owner)
			if err != nil {
				fmt.Println("GetMethodData error: ", err)
				return
			}

			cli.SubmitTransaction(fromAddress, common.HexToAddress(cli.contractAddress), big.NewInt(0), data)
		},
	}

	return OwnerAddCmd
}

func (cli *CLI) buildOwnerRemoveCmd() *cobra.Command {
	OwnerRemoveCmd := &cobra.Command{
		Use:                   "remove <owner>",
		Short:                 "Allows to remove an owner. Transaction has to be sent by wallet",
		Args:                  cobra.MinimumNArgs(1),
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			fromAddress := viper.GetString("from")
			if fromAddress == "" || !common.IsHexAddress(fromAddress) {
				fmt.Println("Error: not set from address of owner")
				fmt.Fprint(os.Stderr, cmd.UsageString())
				return
			}

			ownerStr := args[0]
			if ownerStr == "" && !common.IsHexAddress(ownerStr) {
				fmt.Println("Error: not set owners or invalid")
				fmt.Println(cmd.UsageString())
				return
			}
			owner := common.HexToAddress(ownerStr)

			simpleRegistry, err := cli.GetSimpleRegistry()
			if err != nil {
				fmt.Println("GetSimpleRegistry Error: ", err)
				return
			}
			isowner, err := simpleRegistry.IsOwner(nil, owner)
			if err != nil {
				fmt.Printf("Failed to call GetOwners: %v[%s]\n", err, cli.contractAddress)
				return
			}
			if !isowner {
				fmt.Printf("Address[%s] is NOT owner\n", ownerStr)
				return
			}

			data, err := cli.GetMethodData("removeOwner", owner)
			if err != nil {
				fmt.Println("GetMethodData error: ", err)
				return
			}

			cli.SubmitTransaction(fromAddress, common.HexToAddress(cli.contractAddress), big.NewInt(0), data)
		},
	}

	return OwnerRemoveCmd
}

func (cli *CLI) buildOwnerReplaceCmd() *cobra.Command {
	OwnerReplaceCmd := &cobra.Command{
		Use:                   "replace <owner> <newOwner>",
		Short:                 "Allows to replace an owner with a new owner. Transaction has to be sent by wallet",
		Args:                  cobra.MinimumNArgs(2),
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			fromAddress := viper.GetString("from")
			if fromAddress == "" || !common.IsHexAddress(fromAddress) {
				fmt.Println("Error: not set from address of owner")
				fmt.Fprint(os.Stderr, cmd.UsageString())
				return
			}

			ownerStr := args[0]
			if ownerStr == "" || !common.IsHexAddress(ownerStr) {
				fmt.Println("Error: not set owner or owner invalid")
				fmt.Println(cmd.UsageString())
				return
			}
			owner := common.HexToAddress(ownerStr)

			newOwnerStr := args[1]
			if newOwnerStr == "" || !common.IsHexAddress(newOwnerStr) {
				fmt.Println("Error: not set new owner or new owner invalid")
				fmt.Println(cmd.UsageString())
				return
			}
			newOwner := common.HexToAddress(newOwnerStr)

			simpleRegistry, err := cli.GetSimpleRegistry()
			if err != nil {
				fmt.Println("GetSimpleRegistry Error: ", err)
				return
			}
			isowner, err := simpleRegistry.IsOwner(nil, owner)
			if err != nil {
				fmt.Printf("Failed to call GetOwners: %v[%s]\n", err, cli.contractAddress)
				return
			}
			if !isowner {
				fmt.Printf("Address[%s] is NOT owner\n", ownerStr)
				return
			}

			data, err := cli.GetMethodData("replaceOwner", owner, newOwner)
			if err != nil {
				fmt.Println("GetMethodData error: ", err)
				return
			}

			cli.SubmitTransaction(fromAddress, common.HexToAddress(cli.contractAddress), big.NewInt(0), data)
		},
	}

	return OwnerReplaceCmd
}

func (cli *CLI) buildOwnerCheckCmd() *cobra.Command {
	OwnerCheckCmd := &cobra.Command{
		Use:   "check <owner>",
		Short: "Check whether an address is owner",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			ownerStr := args[0]
			if ownerStr == "" && !common.IsHexAddress(ownerStr) {
				fmt.Println("Error: not set owners or invalid")
				fmt.Println(cmd.UsageString())
				return
			}
			owner := common.HexToAddress(ownerStr)

			cli.OwnerCheck(owner)
		},
	}

	return OwnerCheckCmd
}
