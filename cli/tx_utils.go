package cli

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// SubmitTransaction SubmitTransaction
func (cli *CLI) SubmitTransaction(fromAddress string, toAddress common.Address, value *big.Int, data []byte) {
	var err error

	if !common.IsHexAddress(fromAddress) {
		fmt.Println("Error: fromAddress is invalid hex-encoded: ", fromAddress)
		return
	}

	simpleRegistry, err := cli.GetSimpleRegistry()
	if err != nil {
		fmt.Println("GetSimpleRegistry Error: ", err)
		return
	}

	if isowner, err := simpleRegistry.IsOwner(nil, common.HexToAddress(fromAddress)); err != nil || !isowner {
		fmt.Println("Error: fromAddress is not the owner: ", fromAddress)
		return
	}

	opts, err := cli.getTransactOpts(fromAddress)
	if err != nil {
		fmt.Println("GetTransactOpts: ", err)
		return
	}

	tx, err := simpleRegistry.SubmitTransaction(opts, toAddress, value, data)
	if err != nil {
		fmt.Println("SubmitTransaction error: ", err)
		return
	}

	fmt.Println("Transaction hash is: ", tx.Hash().String())

	//var transaction *big.Int
	ctx := context.Background()

	fmt.Println("Waiting for transaction receipt to get your transferID...")
	txp, err := bind.WaitMined(ctx, cli.client, tx)
	if err != nil {
		fmt.Printf("Error: wait tx mined error(%v)\n", err)
		return
	}
	if !cli.showSubmitID(txp) {
		fmt.Println("No transferID get, please use transaction hash to get it later")
	}
	return
}

var GasFail = "failed to estimate gas needed: gas required exceeds allowance or always failing transaction"

// ConfirmTransaction ConfirmTransaction
func (cli *CLI) ConfirmTransaction(fromAddress string, transactionId *big.Int) {
	var err error

	if !common.IsHexAddress(fromAddress) {
		fmt.Println("Error: fromAddress is invalid hex-encoded: ", fromAddress)
		return
	}

	simpleRegistry, err := cli.GetSimpleRegistry()
	if err != nil {
		fmt.Println("ConfirmTransaction Error: ", err)
		return
	}

	if isowner, err := simpleRegistry.IsOwner(nil, common.HexToAddress(fromAddress)); err != nil || !isowner {
		fmt.Println("Error: fromAddress is not the owner: ", fromAddress)
		return
	}

	count, err := simpleRegistry.GetTransactionCount(nil, true, true)
	if err != nil {
		fmt.Printf("Failed to call GetTransactionCount: %v[%s]\n", err, cli.contractAddress)
		return
	}
	if transactionId.Cmp(count) >= 0 {
		fmt.Printf("TxID[%s] exceeds total number[%s] of transactions\n", transactionId.String(), count.String())
		return
	}

	transaction, err := simpleRegistry.Transactions(nil, transactionId)
	if err != nil {
		fmt.Println("Transactions Error: ", err)
		return
	}
	if transaction.Executed {
		fmt.Printf("Transaction ID(%s) has been Executed\n", transactionId.String())
		return
	}

	isConfirmed, err := simpleRegistry.IsConfirmed(nil, transactionId)
	if err != nil {
		fmt.Println("IsConfirmed Error: ", err)
		return
	}
	if isConfirmed {
		fmt.Printf("Transaction ID(%s) has been Confirmed\n", transactionId.String())
		return
	}

	confirmation, err := simpleRegistry.Confirmations(nil, transactionId, common.HexToAddress(fromAddress))
	if err != nil {
		fmt.Println("Confirmations Error: ", err)
		return
	}
	if confirmation {
		fmt.Printf("Address[%s] has confirmated transaction ID(%s)\n", fromAddress, transactionId.String())
		return
	}

	opts, err := cli.getTransactOpts(fromAddress)
	if err != nil {
		fmt.Println("GetTransactOpts: ", err)
		return
	}

	tx, err := simpleRegistry.ConfirmTransaction(opts, transactionId)
	if err != nil {
		if err.Error() == GasFail {
			fmt.Printf("ID(%s) has been confirmed\n", transactionId.String())
		} else {
			fmt.Println("SubmitTransaction error: ", err)
		}

		return
	}

	fmt.Println("Transaction hash is: ", tx.Hash().String())

}

// RevokeConfirmation RevokeConfirmation
func (cli *CLI) RevokeConfirmation(fromAddress string, transactionId *big.Int) {
	var err error

	if !common.IsHexAddress(fromAddress) {
		fmt.Println("Error: fromAddress is invalid hex-encoded: ", fromAddress)
		return
	}

	simpleRegistry, err := cli.GetSimpleRegistry()
	if err != nil {
		fmt.Println("ConfirmTransaction Error: ", err)
		return
	}

	if isowner, err := simpleRegistry.IsOwner(nil, common.HexToAddress(fromAddress)); err != nil || !isowner {
		fmt.Println("Error: fromAddress is not the owner: ", fromAddress)
		return
	}

	count, err := simpleRegistry.GetTransactionCount(nil, true, true)
	if err != nil {
		fmt.Printf("Failed to call GetTransactionCount: %v[%s]\n", err, cli.contractAddress)
		return
	}
	if transactionId.Cmp(count) >= 0 {
		fmt.Printf("TxID[%s] exceeds total number[%s] of transactions\n", transactionId.String(), count.String())
		return
	}

	transaction, err := simpleRegistry.Transactions(nil, transactionId)
	if err != nil {
		fmt.Println("Transactions Error: ", err)
		return
	}
	if transaction.Executed {
		fmt.Printf("Transaction ID(%s) has been Executed\n", transactionId.String())
		return
	}

	confirmation, err := simpleRegistry.Confirmations(nil, transactionId, common.HexToAddress(fromAddress))
	if err != nil {
		fmt.Println("Confirmations Error: ", err)
		return
	}
	if !confirmation {
		fmt.Printf("Address[%s] NOT confirmate transaction ID(%s)\n", fromAddress, transactionId.String())
		return
	}

	opts, err := cli.getTransactOpts(fromAddress)
	if err != nil {
		fmt.Println("GetTransactOpts: ", err)
		return
	}

	tx, err := simpleRegistry.RevokeConfirmation(opts, transactionId)
	if err != nil {
		fmt.Println("SubmitTransaction error: ", err)
		return
	}

	fmt.Println("Transaction hash is: ", tx.Hash().String())
}

// ExecuteTransaction ExecuteTransaction
func (cli *CLI) ExecuteTransaction(fromAddress string, transactionId *big.Int) {
	var err error

	if !common.IsHexAddress(fromAddress) {
		fmt.Println("Error: fromAddress is invalid hex-encoded: ", fromAddress)
		return
	}

	simpleRegistry, err := cli.GetSimpleRegistry()
	if err != nil {
		fmt.Println("ConfirmTransaction Error: ", err)
		return
	}

	if isowner, err := simpleRegistry.IsOwner(nil, common.HexToAddress(fromAddress)); err != nil || !isowner {
		fmt.Println("Error: fromAddress is not the owner: ", fromAddress)
		return
	}

	count, err := simpleRegistry.GetTransactionCount(nil, true, true)
	if err != nil {
		fmt.Printf("Failed to call GetTransactionCount: %v[%s]\n", err, cli.contractAddress)
		return
	}
	if transactionId.Cmp(count) >= 0 {
		fmt.Printf("TxID[%s] exceeds total number[%s] of transactions\n", transactionId.String(), count.String())
		return
	}

	transaction, err := simpleRegistry.Transactions(nil, transactionId)
	if err != nil {
		fmt.Println("Transactions Error: ", err)
		return
	}
	if transaction.Executed {
		fmt.Printf("Transaction ID(%s) has been Executed\n", transactionId.String())
		return
	}

	opts, err := cli.getTransactOpts(fromAddress)
	if err != nil {
		fmt.Println("GetTransactOpts: ", err)
		return
	}

	tx, err := simpleRegistry.ExecuteTransaction(opts, transactionId)
	if err != nil {
		if err.Error() == GasFail {
			fmt.Printf("ID(%s) has been executed\n", transactionId.String())
		} else {
			fmt.Println("SubmitTransaction error: ", err)
		}
		return
	}

	fmt.Println("Transaction hash is: ", tx.Hash().String())

}

// CheckTransactionStatus CheckTransaction
func (cli *CLI) CheckTransactionStatus(transactionId *big.Int) {
	var err error

	simpleRegistry, err := cli.GetSimpleRegistry()
	if err != nil {
		fmt.Printf("Confirmation status: GetSimpleRegistry Error(%v) ", err)
		return
	}
	ok, err := simpleRegistry.IsConfirmed(nil, transactionId)
	if err != nil {
		fmt.Printf("Confirmation status: IsConfirmed Error(%v)", err)
		return
	}

	count, err := simpleRegistry.GetConfirmationCount(nil, transactionId)
	if err != nil {
		fmt.Printf("Confirmation status: GetTransactionCount Error(%v)\n", err)
		return
	}

	required, err := simpleRegistry.Required(nil)
	if err != nil {
		fmt.Printf("Confirmation status: Required Error(%v)", err)
		return
	}

	if ok {
		fmt.Printf("Confirmation status: Confirmed(%s/%s)\n", count.String(), required.String())
	} else {
		fmt.Printf("Confirmation status: NOT confirmed(%s/%s)\n", count.String(), required.String())
	}

	if count.Cmp(big.NewInt(0)) > 0 {
		ownerslist, err := simpleRegistry.GetConfirmations(nil, transactionId)
		if err != nil {
			fmt.Printf("Confirmed owner list: GetConfirmations Error: (%v)\n", err)
			return
		}
		fmt.Printf("Confirmed owner list (%d):\n", len(ownerslist))
		for _, v := range ownerslist {
			fmt.Printf("\t%s\n", v.String())
		}
	}
}
