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

const ERC20TransferABI = `[{"type":"constructor","stateMutability":"nonpayable","inputs":[{"type":"string","name":"name","internalType":"string"},{"type":"string","name":"symbol","internalType":"string"},{"type":"uint8","name":"decimals","internalType":"uint8"},{"type":"uint256","name":"cap","internalType":"uint256"},{"type":"uint256","name":"initialSupply","internalType":"uint256"},{"type":"bool","name":"transferEnabled","internalType":"bool"},{"type":"bool","name":"mintingFinished","internalType":"bool"}]},{"type":"event","name":"Approval","inputs":[{"type":"address","name":"owner","internalType":"address","indexed":true},{"type":"address","name":"spender","internalType":"address","indexed":true},{"type":"uint256","name":"value","internalType":"uint256","indexed":false}],"anonymous":false},{"type":"event","name":"MintFinished","inputs":[],"anonymous":false},{"type":"event","name":"OwnershipTransferred","inputs":[{"type":"address","name":"previousOwner","internalType":"address","indexed":true},{"type":"address","name":"newOwner","internalType":"address","indexed":true}],"anonymous":false},{"type":"event","name":"RoleGranted","inputs":[{"type":"bytes32","name":"role","internalType":"bytes32","indexed":true},{"type":"address","name":"account","internalType":"address","indexed":true},{"type":"address","name":"sender","internalType":"address","indexed":true}],"anonymous":false},{"type":"event","name":"RoleRevoked","inputs":[{"type":"bytes32","name":"role","internalType":"bytes32","indexed":true},{"type":"address","name":"account","internalType":"address","indexed":true},{"type":"address","name":"sender","internalType":"address","indexed":true}],"anonymous":false},{"type":"event","name":"Transfer","inputs":[{"type":"address","name":"from","internalType":"address","indexed":true},{"type":"address","name":"to","internalType":"address","indexed":true},{"type":"uint256","name":"value","internalType":"uint256","indexed":false}],"anonymous":false},{"type":"event","name":"TransferEnabled","inputs":[],"anonymous":false},{"type":"function","stateMutability":"view","outputs":[{"type":"string","name":"","internalType":"string"}],"name":"BUILT_ON","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"bytes32","name":"","internalType":"bytes32"}],"name":"DEFAULT_ADMIN_ROLE","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"bytes32","name":"","internalType":"bytes32"}],"name":"MINTER_ROLE","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"bytes32","name":"","internalType":"bytes32"}],"name":"OPERATOR_ROLE","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"allowance","inputs":[{"type":"address","name":"owner","internalType":"address"},{"type":"address","name":"spender","internalType":"address"}]},{"type":"function","stateMutability":"nonpayable","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"approve","inputs":[{"type":"address","name":"spender","internalType":"address"},{"type":"uint256","name":"amount","internalType":"uint256"}]},{"type":"function","stateMutability":"nonpayable","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"approveAndCall","inputs":[{"type":"address","name":"spender","internalType":"address"},{"type":"uint256","name":"value","internalType":"uint256"}]},{"type":"function","stateMutability":"nonpayable","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"approveAndCall","inputs":[{"type":"address","name":"spender","internalType":"address"},{"type":"uint256","name":"value","internalType":"uint256"},{"type":"bytes","name":"data","internalType":"bytes"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"balanceOf","inputs":[{"type":"address","name":"account","internalType":"address"}]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"burn","inputs":[{"type":"uint256","name":"amount","internalType":"uint256"}]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"burnFrom","inputs":[{"type":"address","name":"account","internalType":"address"},{"type":"uint256","name":"amount","internalType":"uint256"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"cap","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint8","name":"","internalType":"uint8"}],"name":"decimals","inputs":[]},{"type":"function","stateMutability":"nonpayable","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"decreaseAllowance","inputs":[{"type":"address","name":"spender","internalType":"address"},{"type":"uint256","name":"subtractedValue","internalType":"uint256"}]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"enableTransfer","inputs":[]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"finishMinting","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"bytes32","name":"","internalType":"bytes32"}],"name":"getRoleAdmin","inputs":[{"type":"bytes32","name":"role","internalType":"bytes32"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"address"}],"name":"getRoleMember","inputs":[{"type":"bytes32","name":"role","internalType":"bytes32"},{"type":"uint256","name":"index","internalType":"uint256"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"getRoleMemberCount","inputs":[{"type":"bytes32","name":"role","internalType":"bytes32"}]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"grantRole","inputs":[{"type":"bytes32","name":"role","internalType":"bytes32"},{"type":"address","name":"account","internalType":"address"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"hasRole","inputs":[{"type":"bytes32","name":"role","internalType":"bytes32"},{"type":"address","name":"account","internalType":"address"}]},{"type":"function","stateMutability":"nonpayable","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"increaseAllowance","inputs":[{"type":"address","name":"spender","internalType":"address"},{"type":"uint256","name":"addedValue","internalType":"uint256"}]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"mint","inputs":[{"type":"address","name":"to","internalType":"address"},{"type":"uint256","name":"value","internalType":"uint256"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"mintingFinished","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"string","name":"","internalType":"string"}],"name":"name","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"address"}],"name":"owner","inputs":[]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"recoverERC20","inputs":[{"type":"address","name":"tokenAddress","internalType":"address"},{"type":"uint256","name":"tokenAmount","internalType":"uint256"}]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"renounceOwnership","inputs":[]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"renounceRole","inputs":[{"type":"bytes32","name":"role","internalType":"bytes32"},{"type":"address","name":"account","internalType":"address"}]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"revokeRole","inputs":[{"type":"bytes32","name":"role","internalType":"bytes32"},{"type":"address","name":"account","internalType":"address"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"supportsInterface","inputs":[{"type":"bytes4","name":"interfaceId","internalType":"bytes4"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"string","name":"","internalType":"string"}],"name":"symbol","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"totalSupply","inputs":[]},{"type":"function","stateMutability":"nonpayable","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"transfer","inputs":[{"type":"address","name":"to","internalType":"address"},{"type":"uint256","name":"value","internalType":"uint256"}]},{"type":"function","stateMutability":"nonpayable","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"transferAndCall","inputs":[{"type":"address","name":"to","internalType":"address"},{"type":"uint256","name":"value","internalType":"uint256"}]},{"type":"function","stateMutability":"nonpayable","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"transferAndCall","inputs":[{"type":"address","name":"to","internalType":"address"},{"type":"uint256","name":"value","internalType":"uint256"},{"type":"bytes","name":"data","internalType":"bytes"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"transferEnabled","inputs":[]},{"type":"function","stateMutability":"nonpayable","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"transferFrom","inputs":[{"type":"address","name":"from","internalType":"address"},{"type":"address","name":"to","internalType":"address"},{"type":"uint256","name":"value","internalType":"uint256"}]},{"type":"function","stateMutability":"nonpayable","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"transferFromAndCall","inputs":[{"type":"address","name":"from","internalType":"address"},{"type":"address","name":"to","internalType":"address"},{"type":"uint256","name":"value","internalType":"uint256"},{"type":"bytes","name":"data","internalType":"bytes"}]},{"type":"function","stateMutability":"nonpayable","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"transferFromAndCall","inputs":[{"type":"address","name":"from","internalType":"address"},{"type":"address","name":"to","internalType":"address"},{"type":"uint256","name":"value","internalType":"uint256"}]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"transferOwnership","inputs":[{"type":"address","name":"newOwner","internalType":"address"}]}]`

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

func getAmountTextByWeiWithDecimals(amount *big.Int, decimals uint8) string {
	amountStr := amount.String()
	len := len(amountStr)

	if len <= int(decimals) {
		aStr := strings.TrimRight(amountStr, "0")
		if aStr == "" {
			return "0"
		}
		return "0." + strings.Repeat("0", int(decimals)-len) + aStr
	}

	if decimals == 0 {
		return amountStr[:len-int(decimals)]
	}
	aStr := strings.TrimRight(amountStr[len-int(decimals):], "0")
	if aStr == "" {
		return amountStr[:len-int(decimals)]
	}
	return amountStr[:len-int(decimals)] + "." + aStr
}
