//go:generate abigen --sol contract/MultiSigWalletWithDailyLimit.sol --pkg cli --out cli/MultiSigWalletWithDailyLimit.go
package cli

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

// Deploy deploy contract
func (cli *CLI) Deploy(address string, owners []common.Address, required, dailyLimit *big.Int) error {
	var err error

	opts, err := cli.getTransactOpts(address)
	if err != nil {
		return fmt.Errorf("getTransactOpts error(%s)", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	opts.Context = ctx

	if err := cli.BuildClient(); err != nil {
		return fmt.Errorf("Build client error(%s)", err)
	}
	client := cli.client
	contractAddress, tx, _, err := DeployMultiSigWalletWithDailyLimit(opts, client, owners, required, dailyLimit)
	if err != nil {
		return fmt.Errorf("DeployContract error(%s)", err)
	}

	fmt.Printf("Contract deploy: %s\n", contractAddress.String())
	fmt.Printf("Transaction waiting to be mined: 0x%x\n", tx.Hash())
	cli.contractAddress = contractAddress.String()
	viper.Set("contractaddress", cli.contractAddress)
	bind.WaitDeployed(opts.Context, client, tx)

	return nil
}
