package cli

import (
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// OwnerList OwnerList
func (cli *CLI) OwnerList() {
	var err error

	simpleRegistry, err := cli.GetSimpleRegistry()
	if err != nil {
		fmt.Println("GetSimpleRegistry Error: ", err)
		return
	}
	owners, err := simpleRegistry.GetOwners(nil)
	if err != nil {
		log.Fatalf("Failed to call GetOwners: %v[%s]", err, cli.contractAddress)
	}

	for _, v := range owners {
		fmt.Println(v.String())
	}

}

// OwnerCheck OwnerCheck
func (cli *CLI) OwnerCheck(owner common.Address) {
	var err error

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
	fmt.Println(isowner)
}

// GetMethodData GetMethodData
func (cli *CLI) GetMethodData(name string, args ...interface{}) ([]byte, error) {

	parsed, err := abi.JSON(strings.NewReader(MultiSigWalletWithDailyLimitABI))
	if err != nil {
		return nil, fmt.Errorf("JSON err: %v", err)
	}

	method, exist := parsed.Methods[name]
	if !exist {
		return nil, fmt.Errorf("method '%s' not found", name)
	}

	arguments, err := method.Inputs.Pack(args...)
	if err != nil {
		return nil, err
	}

	// Pack up the method ID too if not a constructor and return
	return append(method.Id(), arguments...), nil

}
