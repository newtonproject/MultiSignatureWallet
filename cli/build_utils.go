package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/console"
)

// method or action
const (
	Submit = iota + 1
	Confirm
	Revoke
	Execute
	OwnerAdd
	OwnerRemove
	OwnerReplace
	DailyLimit
	Required
	TokenTransfer
)

// Transaction for send Transaction
type Transaction struct {
	// Contract  common.Address `json:"to"`
	From      common.Address `json:"from"`
	To        common.Address `json:"to"`
	Value     *big.Int       `json:"value"`
	Unit      string         `json:"unit"`
	Data      []byte         `json:"data"`
	Nonce     uint64         `json:"nonce"`
	GasPrice  *big.Int       `json:"gasPrice"`
	GasLimit  uint64         `json:"gas"`
	NetworkID *big.Int       `json:"networkID"`
	Password  string         `json:"password,omitempty"`

	action int
	params []interface{}
}

func (t Transaction) MarshalJSON(indent bool) ([]byte, error) {
	type transaction struct {
		From      common.Address  `json:"from"`
		To        *common.Address `json:"to"`
		Value     string          `json:"value"`
		Unit      string          `json:"unit"`
		Data      *hexutil.Bytes  `json:"data"`
		Nonce     uint64          `json:"nonce"`
		GasPrice  *big.Int        `json:"gasPrice"`
		GasLimit  uint64          `json:"gas"`
		NetworkID *big.Int        `json:"networkID"`
		// Password  string          `json:"password,omitempty"`
	}
	var tran transaction
	tran.From = t.From
	tran.To = &t.To
	tran.Value = getWeiAmountTextByUnit(t.Value, t.Unit)
	tran.Unit = t.Unit
	tran.Data = (*hexutil.Bytes)(&t.Data)
	tran.Nonce = t.Nonce
	tran.GasPrice = t.GasPrice
	tran.GasLimit = t.GasLimit
	tran.NetworkID = t.NetworkID
	// tran.Password = t.Password

	if indent {
		return json.MarshalIndent(tran, "", " ")
	}

	return json.Marshal(&tran)
}

func (t *Transaction) UnmarshalJSON(input []byte) error {
	type transaction struct {
		From      common.Address  `json:"from"`
		To        *common.Address `json:"to"`
		Value     string          `json:"value"`
		Unit      string          `json:"unit"`
		Data      *hexutil.Bytes  `json:"data"`
		Nonce     uint64          `json:"nonce"`
		GasPrice  *big.Int        `json:"gasPrice"`
		GasLimit  uint64          `json:"gas"`
		NetworkID *big.Int        `json:"networkID"`
		Password  string          `json:"password,omitempty"`
	}
	var tran transaction
	if err := json.Unmarshal(input, &tran); err != nil {
		return err
	}

	t.From = tran.From
	if tran.To != nil {
		t.To = *tran.To
	}
	t.Unit = tran.Unit
	amountWei, err := getAmountWei(tran.Value, tran.Unit)
	if err != nil {
		return err
	}
	t.Value = amountWei
	if tran.Data != nil {
		t.Data = *tran.Data
	}
	t.Nonce = tran.Nonce
	if tran.GasPrice != nil {
		t.GasPrice = tran.GasPrice
	}
	if tran.GasLimit >= 21000 {
		t.GasLimit = tran.GasLimit
	}
	if tran.NetworkID != nil {
		t.NetworkID = tran.NetworkID
	}
	if tran.Password != "" {
		t.Password = tran.Password
	}

	return nil
}

func (cli *CLI) applyTranDefault() error {
	if common.IsHexAddress(cli.contractAddress) {
		cli.tran.To = common.HexToAddress(cli.contractAddress)
	}

	if common.IsHexAddress(cli.address) {
		cli.tran.From = common.HexToAddress(cli.address)
	}

	cli.tran.Value = new(big.Int)
	cli.tran.Unit = UnitETH
	cli.tran.GasPrice = big.NewInt(1)
	cli.tran.NetworkID = big.NewInt(16888)

	return nil
}

func (cli *CLI) applyTxFile(path string) error {
	if cli.tran == nil {
		return errCliTranNil
	}
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return cli.tran.UnmarshalJSON(f)
}

func (cli *CLI) applyTxGuide(offline bool) error {
	var prompt string

	if cli.tran == nil {
		return errCliTranNil
	}

	// get contract address
	for i := 0; ; i++ {
		if err := func() error {
			// enter contract address
			if cli.tran.To == (common.Address{}) {
				prompt = fmt.Sprintf("Enter contract address: ")
			} else {
				prompt = fmt.Sprintf("Enter contract address (default: %s): ", cli.tran.To.String())
			}
			contractAddressStr, err := console.Stdin.PromptInput(prompt)
			if err != nil {
				return fmt.Errorf("get contract address error")
			}
			if contractAddressStr == "" {
				if cli.tran.To == (common.Address{}) {
					return errRequiredContractAddress
				}
			} else {
				if !common.IsHexAddress(contractAddressStr) {
					return errContractAddressIllegal
				}
				cli.tran.To = common.HexToAddress(contractAddressStr)
				cli.contractAddress = contractAddressStr
			}
			return nil
		}(); err == nil {
			break
		} else if i < 2 {
			fmt.Println(err)
			continue
		} else {
			return err
		}
	}

	// get from address
	for i := 0; ; i++ {
		if err := func() error {
			if cli.tran.From == (common.Address{}) {
				prompt = fmt.Sprintf("Enter from address who sign tx: ")
			} else {
				prompt = fmt.Sprintf("Enter from address who sign tx (default: %s): ", cli.tran.From.String())
			}
			fromAddressStr, err := console.Stdin.PromptInput(prompt)
			if err != nil {
				return fmt.Errorf("get \"from\" error")
			}
			if fromAddressStr == "" {
				if cli.tran.From == (common.Address{}) {
					return errRequiredFromAddress
				}
			} else {
				if !common.IsHexAddress(fromAddressStr) {
					return errFromAddressIllegal
				}
				cli.tran.From = common.HexToAddress(fromAddressStr)
			}

			simpleRegistry, err := cli.GetSimpleRegistry()
			if err != nil {
				return fmt.Errorf("getSimpleRegistry Error: %v", err)
			}

			if isowner, err := simpleRegistry.IsOwner(nil, cli.tran.From); err != nil || !isowner {
				return fmt.Errorf("fromAddress is not the owner: %v", cli.tran.From.String())
			}

			return nil
		}(); err == nil {
			break
		} else if i < 2 {
			fmt.Println(err)
			continue
		} else {
			return err
		}
	}

	// get to address
	fmt.Println("Which action to use? (default = submit)")
	fmt.Printf(" %d. Submit - Submit transaction, pay to address\n", Submit)
	fmt.Printf(" %d. Confirm - Confirm transaction ID\n", Confirm)
	fmt.Printf(" %d. Revoke - Revoke a confirmation for a transaction\n", Revoke)
	fmt.Printf(" %d. Execute - Execute a confirmed transaction\n", Execute)
	fmt.Printf(" %d. OwnerAdd - Add an owner\n", OwnerAdd)
	fmt.Printf(" %d. OwnerRemove - Remove an owner\n", OwnerRemove)
	fmt.Printf(" %d. OwnerReplace - Replace an owner\n", OwnerReplace)
	fmt.Printf(" %d. DailyLimit - Update the daily limit\n", DailyLimit)
	fmt.Printf(" %d. Required - Update the number of required\n", Required)
	fmt.Printf(" %d. TokenTransfer - Transfer token from this MSW\n", TokenTransfer)

	action := Submit
	prompt = fmt.Sprintf("Enter the number of action (default: %d): ", action)
	for i := 0; ; i++ {
		if err := func() error {
			action = Submit
			actionStr, err := console.Stdin.PromptInput(prompt)
			if err != nil {
				return fmt.Errorf("get \"to\" error")
			}
			if err != nil {
				fmt.Println("PromptInput err:", err)
				return err
			}
			if actionStr != "" {
				if !IsUintString(actionStr) {
					return errIllegalAmount
				}
				action, err = strconv.Atoi(actionStr)
				if err != nil {
					return err
				}
			}

			switch action {
			case Submit:
			case Confirm:
			case Revoke:
			case Execute:
			case OwnerAdd:
			case OwnerRemove:
			case OwnerReplace:
			case DailyLimit:
			case Required:
			case TokenTransfer:
			default:
				return errIllegalAmount
			}
			return nil
		}(); err == nil {
			break
		} else if i < 2 {
			fmt.Println(err)
			continue
		} else {
			return err
		}
	}

	switch action {
	case Submit:
		// Submit - Submit transaction, pay to address
		if err := cli.applyTxGuideSubmit(); err != nil {
			return err
		}
	case Confirm:
		// Confirm - Confirm transaction ID
		if err := cli.applyTxGuideID("Enter the ID to confirm", Confirm); err != nil {
			return err
		}
	case Revoke:
		// Revoke - Revoke a confirmation for a transaction
		if err := cli.applyTxGuideID("Enter the ID to revoke", Revoke); err != nil {
			return err
		}
	case Execute:
		// Execute - Execute a confirmed transaction
		if err := cli.applyTxGuideID("Enter the ID to execute", Execute); err != nil {
			return err
		}
	case OwnerAdd:
		// OwnerAdd - Add an owner
		if err := cli.applyTxGuideOwnerAdd(); err != nil {
			return err
		}
	case OwnerRemove:
		// OwnerRemove - Remove an owner\n", OwnerRemove)
		if err := cli.applyTxGuideOwnerRemove(); err != nil {
			return err
		}
	case OwnerReplace:
		// OwnerReplace - Replace an owner
		if err := cli.applyTxGuideOwnerReplace(); err != nil {
			return err
		}
	case DailyLimit:
		// DailyLimit - Update the daily limit
		if err := cli.applyTxGuideDailyLimit(); err != nil {
			return err
		}
	case Required:
		// Required - Update the number of required
		if err := cli.applyTxGuideRequired(); err != nil {
			return err
		}
	case TokenTransfer:
		// Token - Call token transfer
		if err := cli.applyTxGuideTokenTransfer(); err != nil {
			return err
		}
	default:
		return errIllegalAmount
	}

	if !offline {
		return nil
	}

	// get nonce, gasLimit, gasPrice, chainID
	if err := cli.applyTxGuideNode(); err != nil {
		return err
	}

	return nil
}

func (cli *CLI) applyTxGuideNode() error {
	var prompt string

	if cli.tran == nil {
		return errCliTranNil
	}

	// get nonce
	prompt = fmt.Sprintf("Enter nonce of from address (default: %d): ", cli.tran.Nonce)
	nonceStr, err := console.Stdin.PromptInput(prompt)
	if err != nil {
		return err
	}
	if nonceStr != "" {
		nonce, err := strconv.ParseUint(nonceStr, 10, 64)
		if err != nil {
			return err
		}
		cli.tran.Nonce = nonce
	}

	// get gasPrice
	if cli.tran.GasPrice == nil {
		cli.tran.GasPrice = big.NewInt(1)
	}
	prompt = fmt.Sprintf("Enter gasPrice (default: %s WEI): ", cli.tran.GasPrice.String())
	gasPriceStr, err := console.Stdin.PromptInput(prompt)
	if err != nil {
		fmt.Println("get gasPrice error")
		return err
	}
	if gasPriceStr != "" {
		if !IsDecimalString(gasPriceStr) {
			return errors.New("gasPrice invalid")
		}
		gasPrice, ok := new(big.Int).SetString(gasPriceStr, 10)
		if !ok {
			return errors.New("convert gasPrice to bigInt error")
		}
		cli.tran.GasPrice = gasPrice
	}

	// get GasLimit
	if cli.tran.GasLimit < 21000 {
		cli.tran.GasLimit = 21000
	}
	prompt = fmt.Sprintf("Enter gasLimit (default: %d): ", cli.tran.GasLimit)
	gasLimitStr, err := console.Stdin.PromptInput(prompt)
	if err != nil {
		return fmt.Errorf("get gasLimit error")
	}
	if gasLimitStr != "" {
		gasLimit, err := strconv.ParseUint(nonceStr, 10, 64)
		if err != nil {
			return fmt.Errorf("conver gasLimit error")
		}
		cli.tran.GasLimit = gasLimit
	}

	// get ChainID
	if cli.tran.NetworkID == nil || cli.tran.NetworkID.Cmp(big.NewInt(0)) == 0 {
		cli.tran.NetworkID = big.NewInt(16888)
	}
	prompt = fmt.Sprintf("Enter ChainID (default: %s): ", cli.tran.NetworkID.String())
	networkIDStr, err := console.Stdin.PromptInput(prompt)
	if err != nil {
		return err
	}
	if networkIDStr != "" {
		networkID, ok := new(big.Int).SetString(networkIDStr, 10)
		if !ok {
			return errors.New("chainID conver error")
		}
		cli.tran.NetworkID = networkID
	}

	return nil
}

func (cli *CLI) saveTranToFile(filepath string) error {
	tByte, err := cli.tran.MarshalJSON(true)
	if err != nil {
		return err
	}

	return saveByteToFile(tByte, filepath)
}

func saveByteToFile(b []byte, filepath string) error {
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	bN := append(b, '\n')
	if _, err := f.Write(bN); err != nil {
		return err
	}

	return nil
}

func saveStringToFile(str, filepath string) error {
	return saveByteToFile([]byte(str), filepath)
}

func (cli *CLI) applyTxGuideSubmit() error {
	var prompt string

	if cli.tran == nil {
		return errCliTranNil
	}

	// get to address
	var to common.Address
	for i := 0; ; i++ {
		if err := func() error {
			if to == (common.Address{}) {
				prompt = fmt.Sprintf("Enter to address: ")
			} else {
				prompt = fmt.Sprintf("Enter to address (default: %s): ", to.String())
			}
			toAddressStr, err := console.Stdin.PromptInput(prompt)
			if err != nil {
				return fmt.Errorf("get \"to\" error")
			}
			if toAddressStr == "" {
				if to == (common.Address{}) {
					return errRequiredToAddress
				}
			} else {
				if !common.IsHexAddress(toAddressStr) {
					return errToAddressIllegal
				}
				to = common.HexToAddress(toAddressStr)
			}
			return nil
		}(); err == nil {
			break
		} else if i < 2 {
			fmt.Println(err)
			continue
		} else {
			return err
		}
	}

	// get pay amount unit
	var unit string
	for i := 0; ; i++ {
		if err := func() error {
			prompt = fmt.Sprintf("Enter unit for amount (NEW or WEI, default NEW): ")
			var err error
			unit, err = console.Stdin.PromptInput(prompt)
			if err != nil {
				fmt.Println("Error: get \"unit\" error")
				return err
			}
			if unit == "" {
				unit = UnitETH
			} else {
				if !stringInSlice(unit, UnitList) {
					return errIllegalUnit
				}
			}
			return nil
		}(); err == nil {
			break
		} else if i < 2 {
			fmt.Println(err)
			continue
		} else {
			return err
		}
	}

	// get pay amount
	value := new(big.Int)
	for i := 0; ; i++ {
		if err := func() error {
			prompt = fmt.Sprintf("Enter amount to pay in %s (default: %s): ", unit,
				getWeiAmountTextByUnit(value, unit))
			amountStr, err := console.Stdin.PromptInput(prompt)
			if err != nil {
				fmt.Println("PromptInput err:", err)
				return err
			}
			if amountStr == "" {
				if value == nil {
					return errAmount0
				}
			} else {
				if !IsDecimalString(amountStr) {
					return errIllegalAmount
				}
				value, err = getAmountWei(amountStr, unit)
				if err != nil {
					return err
				}
				if value.Cmp(maxValue) > 0 {
					return errValueExceeds
				}
			}
			return nil
		}(); err == nil {
			break
		} else if i < 2 {
			fmt.Println(err)
			continue
		} else {
			return err
		}
	}

	// get pay message
	var data []byte
	for i := 0; ; i++ {
		if err := func() error {
			prompt = fmt.Sprintf("Enter text message (default is empty): ")
			dataStr, err := console.Stdin.PromptInput(prompt)
			if err != nil {
				return err
			}
			if dataStr != "" {
				data = []byte(dataStr)
			}
			return nil
		}(); err == nil {
			break
		} else if i < 2 {
			fmt.Println(err)
			continue
		} else {
			return err
		}
	}

	cli.tran.action = Submit
	cli.tran.params = append(cli.tran.params, to)
	cli.tran.params = append(cli.tran.params, value)
	cli.tran.params = append(cli.tran.params, data)

	return nil
}

func (cli *CLI) applyTxGuideID(prompt string, action int) error {
	if cli.tran == nil {
		return errCliTranNil
	}

	// get number
	idStr, err := console.Stdin.PromptInput(prompt + ": ")
	if err != nil {
		fmt.Println("PromptInput err:", err)
		return err
	}
	if idStr == "" {
		return errors.New("nothing is entered")
	}
	id, ok := big.NewInt(0).SetString(idStr, 10)
	if !ok {
		return errors.New("convert string to number error")
	}

	cli.tran.action = action
	cli.tran.params = append(cli.tran.params, id)

	return nil
}

func (cli *CLI) applyTxGuideOwnerAdd() error {
	return cli.applyTxGuideAddress("Enter the address to add: ", "addOwner")
}

func (cli *CLI) applyTxGuideOwnerRemove() error {
	return cli.applyTxGuideAddress("Enter the address to remove: ", "removeOwner")
}

func (cli *CLI) applyTxGuideAddress(prompt, method string) error {

	if cli.tran == nil {
		return errCliTranNil
	}

	address, err := promptAddress(prompt)
	if err != nil {
		return err
	}

	data, err := cli.GetMethodData(method, address)
	if err != nil {
		return err
	}

	cli.tran.action = Submit
	cli.tran.params = append(cli.tran.params, cli.tran.To)
	cli.tran.params = append(cli.tran.params, big.NewInt(0))
	cli.tran.params = append(cli.tran.params, data)

	return nil
}

func promptAddress(prompt string) (common.Address, error) {
	addressStr, err := console.Stdin.PromptInput(prompt)
	if err != nil {
		fmt.Println("PromptInput err:", err)
		return common.Address{}, err
	}
	if addressStr == "" {
		return common.Address{}, errors.New("nothing is entered")
	}
	if !common.IsHexAddress(addressStr) {
		return common.Address{}, errors.New("illegal address")
	}

	return common.HexToAddress(addressStr), nil
}

func (cli *CLI) applyTxGuideOwnerReplace() error {

	if cli.tran == nil {
		return errCliTranNil
	}

	address, err := promptAddress("Enter the old address of owner: ")
	if err != nil {
		return err
	}
	newAddress, err := promptAddress("Enter the new address of owner: ")
	if err != nil {
		return err
	}

	data, err := cli.GetMethodData("replaceOwner", address, newAddress)
	if err != nil {
		return err
	}

	cli.tran.action = Submit
	cli.tran.params = append(cli.tran.params, cli.tran.To)
	cli.tran.params = append(cli.tran.params, big.NewInt(0))
	cli.tran.params = append(cli.tran.params, data)

	return nil
}

func (cli *CLI) applyTxGuideRequired() error {
	if cli.tran == nil {
		return errCliTranNil
	}

	// get number
	idStr, err := console.Stdin.PromptInput("Enter the required number to change: ")
	if err != nil {
		fmt.Println("PromptInput err:", err)
		return err
	}
	if idStr == "" {
		return errors.New("nothing is entered")
	}
	id, ok := big.NewInt(0).SetString(idStr, 10)
	if !ok {
		return errors.New("convert string to number error")
	}

	data, err := cli.GetMethodData("changeRequirement", id)
	if err != nil {
		return err
	}

	cli.tran.action = Submit
	cli.tran.params = append(cli.tran.params, cli.tran.To)
	cli.tran.params = append(cli.tran.params, big.NewInt(0))
	cli.tran.params = append(cli.tran.params, data)

	return nil
}

func (cli *CLI) applyTxGuideDailyLimit() error {
	var err error

	if cli.tran == nil {
		return errCliTranNil
	}

	// get pay amount unit
	var unit string
	prompt := fmt.Sprintf("Enter unit for daily limit (NEW or WEI, default NEW): ")
	unit, err = console.Stdin.PromptInput(prompt)
	if err != nil {
		fmt.Println("Error: get \"unit\" error")
		return err
	}
	if unit == "" {
		unit = UnitETH
	} else {
		if !stringInSlice(unit, UnitList) {
			return errIllegalUnit
		}
	}

	// get pay amount
	value := new(big.Int)
	prompt = fmt.Sprintf("Enter daily limit to change in %s (default: %s): ", unit,
		getWeiAmountTextByUnit(value, unit))
	amountStr, err := console.Stdin.PromptInput(prompt)
	if err != nil {
		fmt.Println("PromptInput err:", err)
		return err
	}
	if amountStr == "" {
		if value == nil {
			return errAmount0
		}
	} else {
		if !IsDecimalString(amountStr) {
			return errIllegalAmount
		}
		value, err = getAmountWei(amountStr, unit)
		if err != nil {
			return err
		}
	}

	data, err := cli.GetMethodData("changeDailyLimit", value)
	if err != nil {
		return err
	}

	cli.tran.action = Submit
	cli.tran.params = append(cli.tran.params, cli.tran.To)
	cli.tran.params = append(cli.tran.params, big.NewInt(0))
	cli.tran.params = append(cli.tran.params, data)

	return nil
}
