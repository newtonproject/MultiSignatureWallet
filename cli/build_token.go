package cli

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/console"
)

func (cli *CLI) applyTxGuideTokenTransfer() error {
	if cli.tran == nil {
		return errCliTranNil
	}

	if cli.tran == nil {
		return errCliTranNil
	}

	token, err := promptAddress("Enter the address of token: ")
	if err != nil {
		return err
	}

	recipient, err := promptAddress("Enter the address of recipient: ")
	if err != nil {
		return err
	}

	ERC20TransferABI := `[{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transfer","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"}]`
	parsed, err := abi.JSON(strings.NewReader(ERC20TransferABI))
	if err != nil {
		return fmt.Errorf("JSON err: %v", err)
	}

	// get number
	amountStr, err := console.Stdin.PromptInput("Enter the amount to transfer: ")
	if err != nil {
		fmt.Println("PromptInput err:", err)
		return err
	}

	var decimals uint8
	{
		// get deicmals
		erc20 := bind.NewBoundContract(token, parsed, cli.client, cli.client, cli.client)
		var (
			ret0 = new(uint8)
		)
		out := ret0
		err = erc20.Call(nil, out, "decimals")
		if err != nil {
			return err
		}

		decimals = *out
	}
	fmt.Println("The decimals of the token is : ", decimals)

	amount, err := GetAmountISAACFromTextWithDecimals(amountStr, decimals)
	if err != nil {
		return err
	}

	var data []byte
	{
		name := "transfer"
		method, exist := parsed.Methods[name]
		if !exist {
			return fmt.Errorf("method '%s' not found", name)
		}

		arguments, err := method.Inputs.Pack(recipient, amount)
		if err != nil {
			return err
		}

		// Pack up the method ID too if not a constructor and return
		data = append(method.Id(), arguments...)
	}
	fmt.Println("The data to token is: ", hex.EncodeToString(data))

	cli.tran.action = Submit
	cli.tran.params = append(cli.tran.params, token)
	cli.tran.params = append(cli.tran.params, big.NewInt(0))
	cli.tran.params = append(cli.tran.params, data)

	return nil
}

// GetAmountISAACFromText convert 1 NEW to 10000000000 ISAAC
func GetAmountISAACFromTextWithDecimals(amountStr string, decimals uint8) (*big.Int, error) {
	return getAmountISAACFromTextWithDecimals(amountStr, int(decimals))
}

func getAmountISAACFromTextWithDecimals(amountStr string, decimals int) (*big.Int, error) {
	index := strings.IndexByte(amountStr, '.')
	if index <= 0 {
		amountISAAC, ok := new(big.Int).SetString(amountStr, 10)
		if !ok {
			return nil, errors.New("convert string to big error")
		}
		return new(big.Int).Mul(amountISAAC, new(big.Int).Exp(big10, big.NewInt(int64(decimals)), nil)), nil
	}
	amountStrInt := amountStr[:index]
	amountStrDec := amountStr[index+1:]
	amountStrDecLen := len(amountStrDec)
	if amountStrDecLen > decimals {
		return nil, errors.New("convert string to big error")
	}
	amountStrDec = amountStrDec + strings.Repeat("0", decimals-amountStrDecLen)
	amountStrInt = amountStrInt + amountStrDec

	amountStrIntBig, ok := new(big.Int).SetString(amountStrInt, 10)
	if !ok {
		return nil, errors.New("convert string to big error")
	}

	return amountStrIntBig, nil
}
