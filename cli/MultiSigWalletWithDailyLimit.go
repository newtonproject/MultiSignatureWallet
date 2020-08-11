// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cli

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MultiSigWalletABI is the input ABI used to generate the binding from.
const MultiSigWalletABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"removeOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"revokeConfirmation\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"confirmations\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"pending\",\"type\":\"bool\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"addOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"isConfirmed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmationCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"from\",\"type\":\"uint256\"},{\"name\":\"to\",\"type\":\"uint256\"},{\"name\":\"pending\",\"type\":\"bool\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionIds\",\"outputs\":[{\"name\":\"_transactionIds\",\"type\":\"uint256[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmations\",\"outputs\":[{\"name\":\"_confirmations\",\"type\":\"address[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transactionCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_required\",\"type\":\"uint256\"}],\"name\":\"changeRequirement\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"confirmTransaction\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"submitTransaction\",\"outputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_OWNER_COUNT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"required\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"replaceOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[]\"},{\"name\":\"_required\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"constructor\"},{\"payable\":true,\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Confirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Revocation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Submission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Execution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"ExecutionFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"RequirementChange\",\"type\":\"event\"}]"

// MultiSigWalletBin is the compiled bytecode used for deploying new contracts.
const MultiSigWalletBin = `0x606060405234156200001057600080fd5b6040516200192538038062001925833981016040528080518201919060200180519150505b6000825182603282111580156200004c5750818111155b80156200005857508015155b80156200006457508115155b15156200007057600080fd5b600092505b84518310156200014157600260008685815181106200009057fe5b90602001906020020151600160a060020a0316815260208101919091526040016000205460ff16158015620000e35750848381518110620000cd57fe5b90602001906020020151600160a060020a031615155b1515620000ef57600080fd5b6001600260008786815181106200010257fe5b90602001906020020151600160a060020a031681526020810191909152604001600020805460ff19169115159190911790555b60019092019162000075565b60038580516200015692916020019062000169565b5060048490555b5b505050505062000204565b828054828255906000526020600020908101928215620001c3579160200282015b82811115620001c35782518254600160a060020a031916600160a060020a0391909116178255602092909201916001909101906200018a565b5b50620001d2929150620001d6565b5090565b6200020191905b80821115620001d2578054600160a060020a0319168155600101620001dd565b5090565b90565b61171180620002146000396000f3006060604052361561011a5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663025e7c278114610165578063173825d91461019757806320ea8d86146101b85780632f54bf6e146101d05780633411c81c1461020357806354741525146102395780637065cb4814610268578063784547a7146102895780638b51d13f146102b35780639ace38c2146102db578063a0e67e2b1461039a578063a8abe69a14610401578063b5dc40c314610478578063b77bf600146104e2578063ba51a6df14610507578063c01a8c841461051f578063c642747414610537578063d74f8edd146105ae578063dc8452cd146105d3578063e20056e6146105f8578063ee22610b1461061f575b5b60003411156101625733600160a060020a03167fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c3460405190815260200160405180910390a25b5b005b341561017057600080fd5b61017b600435610637565b604051600160a060020a03909116815260200160405180910390f35b34156101a257600080fd5b610162600160a060020a0360043516610669565b005b34156101c357600080fd5b61016260043561081a565b005b34156101db57600080fd5b6101ef600160a060020a03600435166108fc565b604051901515815260200160405180910390f35b341561020e57600080fd5b6101ef600435600160a060020a0360243516610911565b604051901515815260200160405180910390f35b341561024457600080fd5b61025660043515156024351515610931565b60405190815260200160405180910390f35b341561027357600080fd5b610162600160a060020a03600435166109a0565b005b341561029457600080fd5b6101ef600435610add565b604051901515815260200160405180910390f35b34156102be57600080fd5b610256600435610b71565b60405190815260200160405180910390f35b34156102e657600080fd5b6102f1600435610bf0565b604051600160a060020a03851681526020810184905281151560608201526080604082018181528454600260001961010060018416150201909116049183018290529060a0830190859080156103885780601f1061035d57610100808354040283529160200191610388565b820191906000526020600020905b81548152906001019060200180831161036b57829003601f168201915b50509550505050505060405180910390f35b34156103a557600080fd5b6103ad610c24565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156103ed5780820151818401525b6020016103d4565b505050509050019250505060405180910390f35b341561040c57600080fd5b6103ad60043560243560443515156064351515610c8d565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156103ed5780820151818401525b6020016103d4565b505050509050019250505060405180910390f35b341561048357600080fd5b6103ad600435610dbb565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156103ed5780820151818401525b6020016103d4565b505050509050019250505060405180910390f35b34156104ed57600080fd5b610256610f3d565b60405190815260200160405180910390f35b341561051257600080fd5b610162600435610f43565b005b341561052a57600080fd5b610162600435610fd9565b005b341561054257600080fd5b61025660048035600160a060020a03169060248035919060649060443590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506110cb95505050505050565b60405190815260200160405180910390f35b34156105b957600080fd5b6102566110eb565b60405190815260200160405180910390f35b34156105de57600080fd5b6102566110f0565b60405190815260200160405180910390f35b341561060357600080fd5b610162600160a060020a03600435811690602435166110f6565b005b341561062a57600080fd5b6101626004356112b7565b005b600380548290811061064557fe5b906000526020600020900160005b915054906101000a9004600160a060020a031681565b600030600160a060020a031633600160a060020a031614151561068b57600080fd5b600160a060020a038216600090815260026020526040902054829060ff1615156106b457600080fd5b600160a060020a0383166000908152600260205260408120805460ff1916905591505b600354600019018210156107af5782600160a060020a03166003838154811015156106fe57fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316600160a060020a031614156107a35760038054600019810190811061073f57fe5b906000526020600020900160005b9054906101000a9004600160a060020a031660038381548110151561076e57fe5b906000526020600020900160005b6101000a815481600160a060020a030219169083600160a060020a031602179055506107af565b5b6001909101906106d7565b6003805460001901906107c290826115cd565b5060035460045411156107db576003546107db90610f43565b5b82600160a060020a03167f8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b9060405160405180910390a25b5b505b5050565b33600160a060020a03811660009081526002602052604090205460ff16151561084257600080fd5b600082815260016020908152604080832033600160a060020a038116855292529091205483919060ff16151561087757600080fd5b600084815260208190526040902060030154849060ff161561089857600080fd5b6000858152600160209081526040808320600160a060020a033316808552925291829020805460ff1916905586917ff6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9905160405180910390a35b5b505b50505b5050565b60026020526000908152604090205460ff1681565b600160209081526000928352604080842090915290825290205460ff1681565b6000805b6005548110156109985783801561095e575060008181526020819052604090206003015460ff16155b806109825750828015610982575060008181526020819052604090206003015460ff165b5b1561098f576001820191505b5b600101610935565b5b5092915050565b30600160a060020a031633600160a060020a03161415156109c057600080fd5b600160a060020a038116600090815260026020526040902054819060ff16156109e857600080fd5b81600160a060020a03811615156109fe57600080fd5b60038054905060010160045460328211158015610a1b5750818111155b8015610a2657508015155b8015610a3157508115155b1515610a3c57600080fd5b600160a060020a0385166000908152600260205260409020805460ff191660019081179091556003805490918101610a7483826115cd565b916000526020600020900160005b8154600160a060020a03808a166101009390930a8381029102199091161790915590507ff39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d60405160405180910390a25b5b50505b505b505b50565b600080805b600354811015610b695760008481526001602052604081206003805491929184908110610b0b57fe5b906000526020600020900160005b9054600160a060020a036101009290920a900416815260208101919091526040016000205460ff1615610b4d576001820191505b600454821415610b605760019250610b69565b5b600101610ae2565b5b5050919050565b6000805b600354811015610be95760008381526001602052604081206003805491929184908110610b9e57fe5b906000526020600020900160005b9054600160a060020a036101009290920a900416815260208101919091526040016000205460ff1615610be0576001820191505b5b600101610b75565b5b50919050565b6000602081905290815260409020805460018201546003830154600160a060020a0390921692909160029091019060ff1684565b610c2c611621565b6003805480602002602001604051908101604052809291908181526020018280548015610c8257602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610c64575b505050505090505b90565b610c95611621565b610c9d611621565b600080600554604051805910610cb05750595b908082528060200260200182016040525b50925060009150600090505b600554811015610d4857858015610cf6575060008181526020819052604090206003015460ff16155b80610d1a5750848015610d1a575060008181526020819052604090206003015460ff165b5b15610d3f5780838381518110610d2d57fe5b60209081029091010152600191909101905b5b600101610ccd565b878703604051805910610d585750595b908082528060200260200182016040525b5093508790505b86811015610daf57828181518110610d8457fe5b906020019060200201518489830381518110610d9c57fe5b602090810290910101525b600101610d70565b5b505050949350505050565b610dc3611621565b610dcb611621565b6003546000908190604051805910610de05750595b908082528060200260200182016040525b50925060009150600090505b600354811015610ec35760008581526001602052604081206003805491929184908110610e2657fe5b906000526020600020900160005b9054600160a060020a036101009290920a900416815260208101919091526040016000205460ff1615610eba576003805482908110610e6f57fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316838381518110610e9b57fe5b600160a060020a03909216602092830290910190910152600191909101905b5b600101610dfd565b81604051805910610ed15750595b908082528060200260200182016040525b509350600090505b81811015610f3457828181518110610efe57fe5b90602001906020020151848281518110610f1457fe5b600160a060020a039092166020928302909101909101525b600101610eea565b5b505050919050565b60055481565b30600160a060020a031633600160a060020a0316141515610f6357600080fd5b6003548160328211801590610f785750818111155b8015610f8357508015155b8015610f8e57508115155b1515610f9957600080fd5b60048390557fa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a8360405190815260200160405180910390a15b5b50505b50565b33600160a060020a03811660009081526002602052604090205460ff16151561100157600080fd5b6000828152602081905260409020548290600160a060020a0316151561102657600080fd5b600083815260016020908152604080832033600160a060020a038116855292529091205484919060ff161561105a57600080fd5b6000858152600160208181526040808420600160a060020a033316808652925292839020805460ff191690921790915586917f4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef905160405180910390a36108f2856112b7565b5b5b50505b505b5050565b60006110d88484846114a6565b90506110e381610fd9565b5b9392505050565b603281565b60045481565b600030600160a060020a031633600160a060020a031614151561111857600080fd5b600160a060020a038316600090815260026020526040902054839060ff16151561114157600080fd5b600160a060020a038316600090815260026020526040902054839060ff161561116957600080fd5b600092505b6003548310156112115784600160a060020a031660038481548110151561119157fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316600160a060020a0316141561120557836003848154811015156111d057fe5b906000526020600020900160005b6101000a815481600160a060020a030219169083600160a060020a03160217905550611211565b5b60019092019161116e565b600160a060020a03808616600081815260026020526040808220805460ff199081169091559388168252908190208054909316600117909255907f8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90905160405180910390a283600160a060020a03167ff39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d60405160405180910390a25b5b505b505b505050565b33600160a060020a03811660009081526002602052604081205490919060ff1615156112e257600080fd5b600083815260016020908152604080832033600160a060020a038116855292529091205484919060ff16151561131757600080fd5b600085815260208190526040902060030154859060ff161561133857600080fd5b61134186610add565b15611499576000868152602081815260409182902060038101805460ff1916600190811790915581548183015460028085018054959c5061142897600160a060020a039094169692956000199581161561010002959095019094160492918391601f83018190048102019051908101604052809291908181526020018280546001816001161561010002031660029004801561141e5780601f106113f35761010080835404028352916020019161141e565b820191906000526020600020905b81548152906001019060200180831161140157829003601f168201915b50505050506115a5565b1561145f57857f33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed7560405160405180910390a2611499565b857f526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b7923660405160405180910390a260038501805460ff191690555b5b5b5b505b50505b505050565b600083600160a060020a03811615156114be57600080fd5b600554915060806040519081016040908152600160a060020a0387168252602080830187905281830186905260006060840181905285815290819052208151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039190911617815560208201518160010155604082015181600201908051611549929160200190611645565b506060820151600391909101805460ff191691151591909117905550600580546001019055817fc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e5160405160405180910390a25b5b509392505050565b6000806040516020840160008287838a8c6187965a03f1925050508091505b50949350505050565b815481835581811511610813576000838152602090206108139181019083016116c4565b5b505050565b815481835581811511610813576000838152602090206108139181019083016116c4565b5b505050565b60206040519081016040526000815290565b60206040519081016040526000815290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061168657805160ff19168380011785556116b3565b828001600101855582156116b3579182015b828111156116b3578251825591602001919060010190611698565b5b506116c09291506116c4565b5090565b610c8a91905b808211156116c057600081556001016116ca565b5090565b905600a165627a7a7230582094a46bb8184ec26cc0b14c67a423f94269a9f13e192c03df65fc963282f0fa200029`

// DeployMultiSigWallet deploys a new Ethereum contract, binding an instance of MultiSigWallet to it.
func DeployMultiSigWallet(auth *bind.TransactOpts, backend bind.ContractBackend, _owners []common.Address, _required *big.Int) (common.Address, *types.Transaction, *MultiSigWallet, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiSigWalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MultiSigWalletBin), backend, _owners, _required)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MultiSigWallet{MultiSigWalletCaller: MultiSigWalletCaller{contract: contract}, MultiSigWalletTransactor: MultiSigWalletTransactor{contract: contract}, MultiSigWalletFilterer: MultiSigWalletFilterer{contract: contract}}, nil
}

// MultiSigWallet is an auto generated Go binding around an Ethereum contract.
type MultiSigWallet struct {
	MultiSigWalletCaller     // Read-only binding to the contract
	MultiSigWalletTransactor // Write-only binding to the contract
	MultiSigWalletFilterer   // Log filterer for contract events
}

// MultiSigWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiSigWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiSigWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiSigWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiSigWalletSession struct {
	Contract     *MultiSigWallet   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultiSigWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiSigWalletCallerSession struct {
	Contract *MultiSigWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MultiSigWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiSigWalletTransactorSession struct {
	Contract     *MultiSigWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MultiSigWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiSigWalletRaw struct {
	Contract *MultiSigWallet // Generic contract binding to access the raw methods on
}

// MultiSigWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiSigWalletCallerRaw struct {
	Contract *MultiSigWalletCaller // Generic read-only contract binding to access the raw methods on
}

// MultiSigWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiSigWalletTransactorRaw struct {
	Contract *MultiSigWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiSigWallet creates a new instance of MultiSigWallet, bound to a specific deployed contract.
func NewMultiSigWallet(address common.Address, backend bind.ContractBackend) (*MultiSigWallet, error) {
	contract, err := bindMultiSigWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiSigWallet{MultiSigWalletCaller: MultiSigWalletCaller{contract: contract}, MultiSigWalletTransactor: MultiSigWalletTransactor{contract: contract}, MultiSigWalletFilterer: MultiSigWalletFilterer{contract: contract}}, nil
}

// NewMultiSigWalletCaller creates a new read-only instance of MultiSigWallet, bound to a specific deployed contract.
func NewMultiSigWalletCaller(address common.Address, caller bind.ContractCaller) (*MultiSigWalletCaller, error) {
	contract, err := bindMultiSigWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletCaller{contract: contract}, nil
}

// NewMultiSigWalletTransactor creates a new write-only instance of MultiSigWallet, bound to a specific deployed contract.
func NewMultiSigWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiSigWalletTransactor, error) {
	contract, err := bindMultiSigWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletTransactor{contract: contract}, nil
}

// NewMultiSigWalletFilterer creates a new log filterer instance of MultiSigWallet, bound to a specific deployed contract.
func NewMultiSigWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiSigWalletFilterer, error) {
	contract, err := bindMultiSigWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletFilterer{contract: contract}, nil
}

// bindMultiSigWallet binds a generic wrapper to an already deployed contract.
func bindMultiSigWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiSigWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiSigWallet *MultiSigWalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MultiSigWallet.Contract.MultiSigWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiSigWallet *MultiSigWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.MultiSigWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiSigWallet *MultiSigWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.MultiSigWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiSigWallet *MultiSigWalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MultiSigWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiSigWallet *MultiSigWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiSigWallet *MultiSigWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.contract.Transact(opts, method, params...)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() constant returns(uint256)
func (_MultiSigWallet *MultiSigWalletCaller) MAXOWNERCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWallet.contract.Call(opts, out, "MAX_OWNER_COUNT")
	return *ret0, err
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() constant returns(uint256)
func (_MultiSigWallet *MultiSigWalletSession) MAXOWNERCOUNT() (*big.Int, error) {
	return _MultiSigWallet.Contract.MAXOWNERCOUNT(&_MultiSigWallet.CallOpts)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() constant returns(uint256)
func (_MultiSigWallet *MultiSigWalletCallerSession) MAXOWNERCOUNT() (*big.Int, error) {
	return _MultiSigWallet.Contract.MAXOWNERCOUNT(&_MultiSigWallet.CallOpts)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) constant returns(bool)
func (_MultiSigWallet *MultiSigWalletCaller) Confirmations(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MultiSigWallet.contract.Call(opts, out, "confirmations", arg0, arg1)
	return *ret0, err
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) constant returns(bool)
func (_MultiSigWallet *MultiSigWalletSession) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _MultiSigWallet.Contract.Confirmations(&_MultiSigWallet.CallOpts, arg0, arg1)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) constant returns(bool)
func (_MultiSigWallet *MultiSigWalletCallerSession) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _MultiSigWallet.Contract.Confirmations(&_MultiSigWallet.CallOpts, arg0, arg1)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) constant returns(uint256 count)
func (_MultiSigWallet *MultiSigWalletCaller) GetConfirmationCount(opts *bind.CallOpts, transactionId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWallet.contract.Call(opts, out, "getConfirmationCount", transactionId)
	return *ret0, err
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) constant returns(uint256 count)
func (_MultiSigWallet *MultiSigWalletSession) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _MultiSigWallet.Contract.GetConfirmationCount(&_MultiSigWallet.CallOpts, transactionId)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) constant returns(uint256 count)
func (_MultiSigWallet *MultiSigWalletCallerSession) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _MultiSigWallet.Contract.GetConfirmationCount(&_MultiSigWallet.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) constant returns(address[] _confirmations)
func (_MultiSigWallet *MultiSigWalletCaller) GetConfirmations(opts *bind.CallOpts, transactionId *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _MultiSigWallet.contract.Call(opts, out, "getConfirmations", transactionId)
	return *ret0, err
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) constant returns(address[] _confirmations)
func (_MultiSigWallet *MultiSigWalletSession) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _MultiSigWallet.Contract.GetConfirmations(&_MultiSigWallet.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) constant returns(address[] _confirmations)
func (_MultiSigWallet *MultiSigWalletCallerSession) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _MultiSigWallet.Contract.GetConfirmations(&_MultiSigWallet.CallOpts, transactionId)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_MultiSigWallet *MultiSigWalletCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _MultiSigWallet.contract.Call(opts, out, "getOwners")
	return *ret0, err
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_MultiSigWallet *MultiSigWalletSession) GetOwners() ([]common.Address, error) {
	return _MultiSigWallet.Contract.GetOwners(&_MultiSigWallet.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_MultiSigWallet *MultiSigWalletCallerSession) GetOwners() ([]common.Address, error) {
	return _MultiSigWallet.Contract.GetOwners(&_MultiSigWallet.CallOpts)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) constant returns(uint256 count)
func (_MultiSigWallet *MultiSigWalletCaller) GetTransactionCount(opts *bind.CallOpts, pending bool, executed bool) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWallet.contract.Call(opts, out, "getTransactionCount", pending, executed)
	return *ret0, err
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) constant returns(uint256 count)
func (_MultiSigWallet *MultiSigWalletSession) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _MultiSigWallet.Contract.GetTransactionCount(&_MultiSigWallet.CallOpts, pending, executed)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) constant returns(uint256 count)
func (_MultiSigWallet *MultiSigWalletCallerSession) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _MultiSigWallet.Contract.GetTransactionCount(&_MultiSigWallet.CallOpts, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) constant returns(uint256[] _transactionIds)
func (_MultiSigWallet *MultiSigWalletCaller) GetTransactionIds(opts *bind.CallOpts, from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _MultiSigWallet.contract.Call(opts, out, "getTransactionIds", from, to, pending, executed)
	return *ret0, err
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) constant returns(uint256[] _transactionIds)
func (_MultiSigWallet *MultiSigWalletSession) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _MultiSigWallet.Contract.GetTransactionIds(&_MultiSigWallet.CallOpts, from, to, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) constant returns(uint256[] _transactionIds)
func (_MultiSigWallet *MultiSigWalletCallerSession) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _MultiSigWallet.Contract.GetTransactionIds(&_MultiSigWallet.CallOpts, from, to, pending, executed)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) constant returns(bool)
func (_MultiSigWallet *MultiSigWalletCaller) IsConfirmed(opts *bind.CallOpts, transactionId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MultiSigWallet.contract.Call(opts, out, "isConfirmed", transactionId)
	return *ret0, err
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) constant returns(bool)
func (_MultiSigWallet *MultiSigWalletSession) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _MultiSigWallet.Contract.IsConfirmed(&_MultiSigWallet.CallOpts, transactionId)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) constant returns(bool)
func (_MultiSigWallet *MultiSigWalletCallerSession) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _MultiSigWallet.Contract.IsConfirmed(&_MultiSigWallet.CallOpts, transactionId)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) constant returns(bool)
func (_MultiSigWallet *MultiSigWalletCaller) IsOwner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MultiSigWallet.contract.Call(opts, out, "isOwner", arg0)
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) constant returns(bool)
func (_MultiSigWallet *MultiSigWalletSession) IsOwner(arg0 common.Address) (bool, error) {
	return _MultiSigWallet.Contract.IsOwner(&_MultiSigWallet.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) constant returns(bool)
func (_MultiSigWallet *MultiSigWalletCallerSession) IsOwner(arg0 common.Address) (bool, error) {
	return _MultiSigWallet.Contract.IsOwner(&_MultiSigWallet.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) constant returns(address)
func (_MultiSigWallet *MultiSigWalletCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MultiSigWallet.contract.Call(opts, out, "owners", arg0)
	return *ret0, err
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) constant returns(address)
func (_MultiSigWallet *MultiSigWalletSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _MultiSigWallet.Contract.Owners(&_MultiSigWallet.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) constant returns(address)
func (_MultiSigWallet *MultiSigWalletCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _MultiSigWallet.Contract.Owners(&_MultiSigWallet.CallOpts, arg0)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() constant returns(uint256)
func (_MultiSigWallet *MultiSigWalletCaller) Required(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWallet.contract.Call(opts, out, "required")
	return *ret0, err
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() constant returns(uint256)
func (_MultiSigWallet *MultiSigWalletSession) Required() (*big.Int, error) {
	return _MultiSigWallet.Contract.Required(&_MultiSigWallet.CallOpts)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() constant returns(uint256)
func (_MultiSigWallet *MultiSigWalletCallerSession) Required() (*big.Int, error) {
	return _MultiSigWallet.Contract.Required(&_MultiSigWallet.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() constant returns(uint256)
func (_MultiSigWallet *MultiSigWalletCaller) TransactionCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWallet.contract.Call(opts, out, "transactionCount")
	return *ret0, err
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() constant returns(uint256)
func (_MultiSigWallet *MultiSigWalletSession) TransactionCount() (*big.Int, error) {
	return _MultiSigWallet.Contract.TransactionCount(&_MultiSigWallet.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() constant returns(uint256)
func (_MultiSigWallet *MultiSigWalletCallerSession) TransactionCount() (*big.Int, error) {
	return _MultiSigWallet.Contract.TransactionCount(&_MultiSigWallet.CallOpts)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) constant returns(address destination, uint256 value, bytes data, bool executed)
func (_MultiSigWallet *MultiSigWalletCaller) Transactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	ret := new(struct {
		Destination common.Address
		Value       *big.Int
		Data        []byte
		Executed    bool
	})
	out := ret
	err := _MultiSigWallet.contract.Call(opts, out, "transactions", arg0)
	return *ret, err
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) constant returns(address destination, uint256 value, bytes data, bool executed)
func (_MultiSigWallet *MultiSigWalletSession) Transactions(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _MultiSigWallet.Contract.Transactions(&_MultiSigWallet.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) constant returns(address destination, uint256 value, bytes data, bool executed)
func (_MultiSigWallet *MultiSigWalletCallerSession) Transactions(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _MultiSigWallet.Contract.Transactions(&_MultiSigWallet.CallOpts, arg0)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_MultiSigWallet *MultiSigWalletTransactor) AddOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _MultiSigWallet.contract.Transact(opts, "addOwner", owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_MultiSigWallet *MultiSigWalletSession) AddOwner(owner common.Address) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.AddOwner(&_MultiSigWallet.TransactOpts, owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_MultiSigWallet *MultiSigWalletTransactorSession) AddOwner(owner common.Address) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.AddOwner(&_MultiSigWallet.TransactOpts, owner)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_MultiSigWallet *MultiSigWalletTransactor) ChangeRequirement(opts *bind.TransactOpts, _required *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.contract.Transact(opts, "changeRequirement", _required)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_MultiSigWallet *MultiSigWalletSession) ChangeRequirement(_required *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.ChangeRequirement(&_MultiSigWallet.TransactOpts, _required)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_MultiSigWallet *MultiSigWalletTransactorSession) ChangeRequirement(_required *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.ChangeRequirement(&_MultiSigWallet.TransactOpts, _required)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_MultiSigWallet *MultiSigWalletTransactor) ConfirmTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.contract.Transact(opts, "confirmTransaction", transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_MultiSigWallet *MultiSigWalletSession) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.ConfirmTransaction(&_MultiSigWallet.TransactOpts, transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_MultiSigWallet *MultiSigWalletTransactorSession) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.ConfirmTransaction(&_MultiSigWallet.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_MultiSigWallet *MultiSigWalletTransactor) ExecuteTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.contract.Transact(opts, "executeTransaction", transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_MultiSigWallet *MultiSigWalletSession) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.ExecuteTransaction(&_MultiSigWallet.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_MultiSigWallet *MultiSigWalletTransactorSession) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.ExecuteTransaction(&_MultiSigWallet.TransactOpts, transactionId)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_MultiSigWallet *MultiSigWalletTransactor) RemoveOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _MultiSigWallet.contract.Transact(opts, "removeOwner", owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_MultiSigWallet *MultiSigWalletSession) RemoveOwner(owner common.Address) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.RemoveOwner(&_MultiSigWallet.TransactOpts, owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_MultiSigWallet *MultiSigWalletTransactorSession) RemoveOwner(owner common.Address) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.RemoveOwner(&_MultiSigWallet.TransactOpts, owner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_MultiSigWallet *MultiSigWalletTransactor) ReplaceOwner(opts *bind.TransactOpts, owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _MultiSigWallet.contract.Transact(opts, "replaceOwner", owner, newOwner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_MultiSigWallet *MultiSigWalletSession) ReplaceOwner(owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.ReplaceOwner(&_MultiSigWallet.TransactOpts, owner, newOwner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_MultiSigWallet *MultiSigWalletTransactorSession) ReplaceOwner(owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.ReplaceOwner(&_MultiSigWallet.TransactOpts, owner, newOwner)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_MultiSigWallet *MultiSigWalletTransactor) RevokeConfirmation(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.contract.Transact(opts, "revokeConfirmation", transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_MultiSigWallet *MultiSigWalletSession) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.RevokeConfirmation(&_MultiSigWallet.TransactOpts, transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_MultiSigWallet *MultiSigWalletTransactorSession) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.RevokeConfirmation(&_MultiSigWallet.TransactOpts, transactionId)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_MultiSigWallet *MultiSigWalletTransactor) SubmitTransaction(opts *bind.TransactOpts, destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MultiSigWallet.contract.Transact(opts, "submitTransaction", destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_MultiSigWallet *MultiSigWalletSession) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.SubmitTransaction(&_MultiSigWallet.TransactOpts, destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_MultiSigWallet *MultiSigWalletTransactorSession) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.SubmitTransaction(&_MultiSigWallet.TransactOpts, destination, value, data)
}

// MultiSigWalletConfirmationIterator is returned from FilterConfirmation and is used to iterate over the raw logs and unpacked data for Confirmation events raised by the MultiSigWallet contract.
type MultiSigWalletConfirmationIterator struct {
	Event *MultiSigWalletConfirmation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletConfirmation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletConfirmation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletConfirmation represents a Confirmation event raised by the MultiSigWallet contract.
type MultiSigWalletConfirmation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterConfirmation is a free log retrieval operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWallet *MultiSigWalletFilterer) FilterConfirmation(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*MultiSigWalletConfirmationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.FilterLogs(opts, "Confirmation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletConfirmationIterator{contract: _MultiSigWallet.contract, event: "Confirmation", logs: logs, sub: sub}, nil
}

// WatchConfirmation is a free log subscription operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWallet *MultiSigWalletFilterer) WatchConfirmation(opts *bind.WatchOpts, sink chan<- *MultiSigWalletConfirmation, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.WatchLogs(opts, "Confirmation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletConfirmation)
				if err := _MultiSigWallet.contract.UnpackLog(event, "Confirmation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MultiSigWalletDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the MultiSigWallet contract.
type MultiSigWalletDepositIterator struct {
	Event *MultiSigWalletDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletDeposit represents a Deposit event raised by the MultiSigWallet contract.
type MultiSigWalletDeposit struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 value)
func (_MultiSigWallet *MultiSigWalletFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*MultiSigWalletDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MultiSigWallet.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletDepositIterator{contract: _MultiSigWallet.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 value)
func (_MultiSigWallet *MultiSigWalletFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MultiSigWalletDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MultiSigWallet.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletDeposit)
				if err := _MultiSigWallet.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MultiSigWalletExecutionIterator is returned from FilterExecution and is used to iterate over the raw logs and unpacked data for Execution events raised by the MultiSigWallet contract.
type MultiSigWalletExecutionIterator struct {
	Event *MultiSigWalletExecution // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletExecution)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletExecution)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletExecution represents a Execution event raised by the MultiSigWallet contract.
type MultiSigWalletExecution struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecution is a free log retrieval operation binding the contract event 0x33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed75.
//
// Solidity: event Execution(uint256 indexed transactionId)
func (_MultiSigWallet *MultiSigWalletFilterer) FilterExecution(opts *bind.FilterOpts, transactionId []*big.Int) (*MultiSigWalletExecutionIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.FilterLogs(opts, "Execution", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletExecutionIterator{contract: _MultiSigWallet.contract, event: "Execution", logs: logs, sub: sub}, nil
}

// WatchExecution is a free log subscription operation binding the contract event 0x33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed75.
//
// Solidity: event Execution(uint256 indexed transactionId)
func (_MultiSigWallet *MultiSigWalletFilterer) WatchExecution(opts *bind.WatchOpts, sink chan<- *MultiSigWalletExecution, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.WatchLogs(opts, "Execution", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletExecution)
				if err := _MultiSigWallet.contract.UnpackLog(event, "Execution", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MultiSigWalletExecutionFailureIterator is returned from FilterExecutionFailure and is used to iterate over the raw logs and unpacked data for ExecutionFailure events raised by the MultiSigWallet contract.
type MultiSigWalletExecutionFailureIterator struct {
	Event *MultiSigWalletExecutionFailure // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletExecutionFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletExecutionFailure)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletExecutionFailure)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletExecutionFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletExecutionFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletExecutionFailure represents a ExecutionFailure event raised by the MultiSigWallet contract.
type MultiSigWalletExecutionFailure struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionFailure is a free log retrieval operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(uint256 indexed transactionId)
func (_MultiSigWallet *MultiSigWalletFilterer) FilterExecutionFailure(opts *bind.FilterOpts, transactionId []*big.Int) (*MultiSigWalletExecutionFailureIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.FilterLogs(opts, "ExecutionFailure", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletExecutionFailureIterator{contract: _MultiSigWallet.contract, event: "ExecutionFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFailure is a free log subscription operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(uint256 indexed transactionId)
func (_MultiSigWallet *MultiSigWalletFilterer) WatchExecutionFailure(opts *bind.WatchOpts, sink chan<- *MultiSigWalletExecutionFailure, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.WatchLogs(opts, "ExecutionFailure", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletExecutionFailure)
				if err := _MultiSigWallet.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MultiSigWalletOwnerAdditionIterator is returned from FilterOwnerAddition and is used to iterate over the raw logs and unpacked data for OwnerAddition events raised by the MultiSigWallet contract.
type MultiSigWalletOwnerAdditionIterator struct {
	Event *MultiSigWalletOwnerAddition // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletOwnerAdditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletOwnerAddition)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletOwnerAddition)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletOwnerAdditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletOwnerAdditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletOwnerAddition represents a OwnerAddition event raised by the MultiSigWallet contract.
type MultiSigWalletOwnerAddition struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerAddition is a free log retrieval operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address indexed owner)
func (_MultiSigWallet *MultiSigWalletFilterer) FilterOwnerAddition(opts *bind.FilterOpts, owner []common.Address) (*MultiSigWalletOwnerAdditionIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MultiSigWallet.contract.FilterLogs(opts, "OwnerAddition", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletOwnerAdditionIterator{contract: _MultiSigWallet.contract, event: "OwnerAddition", logs: logs, sub: sub}, nil
}

// WatchOwnerAddition is a free log subscription operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address indexed owner)
func (_MultiSigWallet *MultiSigWalletFilterer) WatchOwnerAddition(opts *bind.WatchOpts, sink chan<- *MultiSigWalletOwnerAddition, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MultiSigWallet.contract.WatchLogs(opts, "OwnerAddition", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletOwnerAddition)
				if err := _MultiSigWallet.contract.UnpackLog(event, "OwnerAddition", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MultiSigWalletOwnerRemovalIterator is returned from FilterOwnerRemoval and is used to iterate over the raw logs and unpacked data for OwnerRemoval events raised by the MultiSigWallet contract.
type MultiSigWalletOwnerRemovalIterator struct {
	Event *MultiSigWalletOwnerRemoval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletOwnerRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletOwnerRemoval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletOwnerRemoval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletOwnerRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletOwnerRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletOwnerRemoval represents a OwnerRemoval event raised by the MultiSigWallet contract.
type MultiSigWalletOwnerRemoval struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerRemoval is a free log retrieval operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address indexed owner)
func (_MultiSigWallet *MultiSigWalletFilterer) FilterOwnerRemoval(opts *bind.FilterOpts, owner []common.Address) (*MultiSigWalletOwnerRemovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MultiSigWallet.contract.FilterLogs(opts, "OwnerRemoval", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletOwnerRemovalIterator{contract: _MultiSigWallet.contract, event: "OwnerRemoval", logs: logs, sub: sub}, nil
}

// WatchOwnerRemoval is a free log subscription operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address indexed owner)
func (_MultiSigWallet *MultiSigWalletFilterer) WatchOwnerRemoval(opts *bind.WatchOpts, sink chan<- *MultiSigWalletOwnerRemoval, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MultiSigWallet.contract.WatchLogs(opts, "OwnerRemoval", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletOwnerRemoval)
				if err := _MultiSigWallet.contract.UnpackLog(event, "OwnerRemoval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MultiSigWalletRequirementChangeIterator is returned from FilterRequirementChange and is used to iterate over the raw logs and unpacked data for RequirementChange events raised by the MultiSigWallet contract.
type MultiSigWalletRequirementChangeIterator struct {
	Event *MultiSigWalletRequirementChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletRequirementChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletRequirementChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletRequirementChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletRequirementChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletRequirementChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletRequirementChange represents a RequirementChange event raised by the MultiSigWallet contract.
type MultiSigWalletRequirementChange struct {
	Required *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRequirementChange is a free log retrieval operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_MultiSigWallet *MultiSigWalletFilterer) FilterRequirementChange(opts *bind.FilterOpts) (*MultiSigWalletRequirementChangeIterator, error) {

	logs, sub, err := _MultiSigWallet.contract.FilterLogs(opts, "RequirementChange")
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletRequirementChangeIterator{contract: _MultiSigWallet.contract, event: "RequirementChange", logs: logs, sub: sub}, nil
}

// WatchRequirementChange is a free log subscription operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_MultiSigWallet *MultiSigWalletFilterer) WatchRequirementChange(opts *bind.WatchOpts, sink chan<- *MultiSigWalletRequirementChange) (event.Subscription, error) {

	logs, sub, err := _MultiSigWallet.contract.WatchLogs(opts, "RequirementChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletRequirementChange)
				if err := _MultiSigWallet.contract.UnpackLog(event, "RequirementChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MultiSigWalletRevocationIterator is returned from FilterRevocation and is used to iterate over the raw logs and unpacked data for Revocation events raised by the MultiSigWallet contract.
type MultiSigWalletRevocationIterator struct {
	Event *MultiSigWalletRevocation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletRevocationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletRevocation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletRevocation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletRevocationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletRevocationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletRevocation represents a Revocation event raised by the MultiSigWallet contract.
type MultiSigWalletRevocation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRevocation is a free log retrieval operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWallet *MultiSigWalletFilterer) FilterRevocation(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*MultiSigWalletRevocationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.FilterLogs(opts, "Revocation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletRevocationIterator{contract: _MultiSigWallet.contract, event: "Revocation", logs: logs, sub: sub}, nil
}

// WatchRevocation is a free log subscription operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWallet *MultiSigWalletFilterer) WatchRevocation(opts *bind.WatchOpts, sink chan<- *MultiSigWalletRevocation, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.WatchLogs(opts, "Revocation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletRevocation)
				if err := _MultiSigWallet.contract.UnpackLog(event, "Revocation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MultiSigWalletSubmissionIterator is returned from FilterSubmission and is used to iterate over the raw logs and unpacked data for Submission events raised by the MultiSigWallet contract.
type MultiSigWalletSubmissionIterator struct {
	Event *MultiSigWalletSubmission // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletSubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletSubmission)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletSubmission)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletSubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletSubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletSubmission represents a Submission event raised by the MultiSigWallet contract.
type MultiSigWalletSubmission struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSubmission is a free log retrieval operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 indexed transactionId)
func (_MultiSigWallet *MultiSigWalletFilterer) FilterSubmission(opts *bind.FilterOpts, transactionId []*big.Int) (*MultiSigWalletSubmissionIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.FilterLogs(opts, "Submission", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletSubmissionIterator{contract: _MultiSigWallet.contract, event: "Submission", logs: logs, sub: sub}, nil
}

// WatchSubmission is a free log subscription operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 indexed transactionId)
func (_MultiSigWallet *MultiSigWalletFilterer) WatchSubmission(opts *bind.WatchOpts, sink chan<- *MultiSigWalletSubmission, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.WatchLogs(opts, "Submission", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletSubmission)
				if err := _MultiSigWallet.contract.UnpackLog(event, "Submission", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MultiSigWalletWithDailyLimitABI is the input ABI used to generate the binding from.
const MultiSigWalletWithDailyLimitABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"removeOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"revokeConfirmation\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"confirmations\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"calcMaxWithdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"pending\",\"type\":\"bool\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dailyLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastDay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"addOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"isConfirmed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmationCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"from\",\"type\":\"uint256\"},{\"name\":\"to\",\"type\":\"uint256\"},{\"name\":\"pending\",\"type\":\"bool\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionIds\",\"outputs\":[{\"name\":\"_transactionIds\",\"type\":\"uint256[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmations\",\"outputs\":[{\"name\":\"_confirmations\",\"type\":\"address[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transactionCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_required\",\"type\":\"uint256\"}],\"name\":\"changeRequirement\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"confirmTransaction\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"submitTransaction\",\"outputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_dailyLimit\",\"type\":\"uint256\"}],\"name\":\"changeDailyLimit\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_OWNER_COUNT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"required\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"replaceOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spentToday\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[]\"},{\"name\":\"_required\",\"type\":\"uint256\"},{\"name\":\"_dailyLimit\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"constructor\"},{\"payable\":true,\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"dailyLimit\",\"type\":\"uint256\"}],\"name\":\"DailyLimitChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Confirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Revocation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Submission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Execution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"ExecutionFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"RequirementChange\",\"type\":\"event\"}]"

// MultiSigWalletWithDailyLimitBin is the compiled bytecode used for deploying new contracts.
const MultiSigWalletWithDailyLimitBin = `0x606060405234156200001057600080fd5b60405162001b1b38038062001b1b83398101604052808051820191906020018051919060200180519150505b82825b600082518260328211158015620000565750818111155b80156200006257508015155b80156200006e57508115155b15156200007a57600080fd5b600092505b84518310156200014b57600260008685815181106200009a57fe5b90602001906020020151600160a060020a0316815260208101919091526040016000205460ff16158015620000ed5750848381518110620000d757fe5b90602001906020020151600160a060020a031615155b1515620000f957600080fd5b6001600260008786815181106200010c57fe5b90602001906020020151600160a060020a031681526020810191909152604001600020805460ff19169115159190911790555b6001909201916200007f565b6003858051620001609291602001906200017c565b5060048490555b5b505050600683905550505b50505062000217565b828054828255906000526020600020908101928215620001d6579160200282015b82811115620001d65782518254600160a060020a031916600160a060020a0391909116178255602092909201916001909101906200019d565b5b50620001e5929150620001e9565b5090565b6200021491905b80821115620001e5578054600160a060020a0319168155600101620001f0565b5090565b90565b6118f480620002276000396000f300606060405236156101515763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663025e7c27811461019c578063173825d9146101ce57806320ea8d86146101ef5780632f54bf6e146102075780633411c81c1461023a5780634bc9fdc214610270578063547415251461029557806367eeba0c146102c45780636b0c932d146102e95780637065cb481461030e578063784547a71461032f5780638b51d13f146103595780639ace38c214610381578063a0e67e2b14610440578063a8abe69a146104a7578063b5dc40c31461051e578063b77bf60014610588578063ba51a6df146105ad578063c01a8c84146105c5578063c6427474146105dd578063cea0862114610654578063d74f8edd1461066c578063dc8452cd14610691578063e20056e6146106b6578063ee22610b146106dd578063f059cf2b146106f5575b5b60003411156101995733600160a060020a03167fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c3460405190815260200160405180910390a25b5b005b34156101a757600080fd5b6101b260043561071a565b604051600160a060020a03909116815260200160405180910390f35b34156101d957600080fd5b610199600160a060020a036004351661074c565b005b34156101fa57600080fd5b6101996004356108fd565b005b341561021257600080fd5b610226600160a060020a03600435166109df565b604051901515815260200160405180910390f35b341561024557600080fd5b610226600435600160a060020a03602435166109f4565b604051901515815260200160405180910390f35b341561027b57600080fd5b610283610a14565b60405190815260200160405180910390f35b34156102a057600080fd5b61028360043515156024351515610a4e565b60405190815260200160405180910390f35b34156102cf57600080fd5b610283610abd565b60405190815260200160405180910390f35b34156102f457600080fd5b610283610ac3565b60405190815260200160405180910390f35b341561031957600080fd5b610199600160a060020a0360043516610ac9565b005b341561033a57600080fd5b610226600435610c06565b604051901515815260200160405180910390f35b341561036457600080fd5b610283600435610c9a565b60405190815260200160405180910390f35b341561038c57600080fd5b610397600435610d19565b604051600160a060020a03851681526020810184905281151560608201526080604082018181528454600260001961010060018416150201909116049183018290529060a08301908590801561042e5780601f106104035761010080835404028352916020019161042e565b820191906000526020600020905b81548152906001019060200180831161041157829003601f168201915b50509550505050505060405180910390f35b341561044b57600080fd5b610453610d4d565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156104935780820151818401525b60200161047a565b505050509050019250505060405180910390f35b34156104b257600080fd5b61045360043560243560443515156064351515610db6565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156104935780820151818401525b60200161047a565b505050509050019250505060405180910390f35b341561052957600080fd5b610453600435610ee4565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156104935780820151818401525b60200161047a565b505050509050019250505060405180910390f35b341561059357600080fd5b610283611066565b60405190815260200160405180910390f35b34156105b857600080fd5b61019960043561106c565b005b34156105d057600080fd5b610199600435611102565b005b34156105e857600080fd5b61028360048035600160a060020a03169060248035919060649060443590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506111f495505050505050565b60405190815260200160405180910390f35b341561065f57600080fd5b610199600435611214565b005b341561067757600080fd5b610283611271565b60405190815260200160405180910390f35b341561069c57600080fd5b610283611276565b60405190815260200160405180910390f35b34156106c157600080fd5b610199600160a060020a036004358116906024351661127c565b005b34156106e857600080fd5b61019960043561143d565b005b341561070057600080fd5b610283611663565b60405190815260200160405180910390f35b600380548290811061072857fe5b906000526020600020900160005b915054906101000a9004600160a060020a031681565b600030600160a060020a031633600160a060020a031614151561076e57600080fd5b600160a060020a038216600090815260026020526040902054829060ff16151561079757600080fd5b600160a060020a0383166000908152600260205260408120805460ff1916905591505b600354600019018210156108925782600160a060020a03166003838154811015156107e157fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316600160a060020a031614156108865760038054600019810190811061082257fe5b906000526020600020900160005b9054906101000a9004600160a060020a031660038381548110151561085157fe5b906000526020600020900160005b6101000a815481600160a060020a030219169083600160a060020a03160217905550610892565b5b6001909101906107ba565b6003805460001901906108a590826117b0565b5060035460045411156108be576003546108be9061106c565b5b82600160a060020a03167f8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b9060405160405180910390a25b5b505b5050565b33600160a060020a03811660009081526002602052604090205460ff16151561092557600080fd5b600082815260016020908152604080832033600160a060020a038116855292529091205483919060ff16151561095a57600080fd5b600084815260208190526040902060030154849060ff161561097b57600080fd5b6000858152600160209081526040808320600160a060020a033316808552925291829020805460ff1916905586917ff6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9905160405180910390a35b5b505b50505b5050565b60026020526000908152604090205460ff1681565b600160209081526000928352604080842090915290825290205460ff1681565b60006007546201518001421115610a2e5750600654610a4b565b6008546006541015610a4257506000610a4b565b50600854600654035b90565b6000805b600554811015610ab557838015610a7b575060008181526020819052604090206003015460ff16155b80610a9f5750828015610a9f575060008181526020819052604090206003015460ff165b5b15610aac576001820191505b5b600101610a52565b5b5092915050565b60065481565b60075481565b30600160a060020a031633600160a060020a0316141515610ae957600080fd5b600160a060020a038116600090815260026020526040902054819060ff1615610b1157600080fd5b81600160a060020a0381161515610b2757600080fd5b60038054905060010160045460328211158015610b445750818111155b8015610b4f57508015155b8015610b5a57508115155b1515610b6557600080fd5b600160a060020a0385166000908152600260205260409020805460ff191660019081179091556003805490918101610b9d83826117b0565b916000526020600020900160005b8154600160a060020a03808a166101009390930a8381029102199091161790915590507ff39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d60405160405180910390a25b5b50505b505b505b50565b600080805b600354811015610c925760008481526001602052604081206003805491929184908110610c3457fe5b906000526020600020900160005b9054600160a060020a036101009290920a900416815260208101919091526040016000205460ff1615610c76576001820191505b600454821415610c895760019250610c92565b5b600101610c0b565b5b5050919050565b6000805b600354811015610d125760008381526001602052604081206003805491929184908110610cc757fe5b906000526020600020900160005b9054600160a060020a036101009290920a900416815260208101919091526040016000205460ff1615610d09576001820191505b5b600101610c9e565b5b50919050565b6000602081905290815260409020805460018201546003830154600160a060020a0390921692909160029091019060ff1684565b610d55611804565b6003805480602002602001604051908101604052809291908181526020018280548015610dab57602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610d8d575b505050505090505b90565b610dbe611804565b610dc6611804565b600080600554604051805910610dd95750595b908082528060200260200182016040525b50925060009150600090505b600554811015610e7157858015610e1f575060008181526020819052604090206003015460ff16155b80610e435750848015610e43575060008181526020819052604090206003015460ff165b5b15610e685780838381518110610e5657fe5b60209081029091010152600191909101905b5b600101610df6565b878703604051805910610e815750595b908082528060200260200182016040525b5093508790505b86811015610ed857828181518110610ead57fe5b906020019060200201518489830381518110610ec557fe5b602090810290910101525b600101610e99565b5b505050949350505050565b610eec611804565b610ef4611804565b6003546000908190604051805910610f095750595b908082528060200260200182016040525b50925060009150600090505b600354811015610fec5760008581526001602052604081206003805491929184908110610f4f57fe5b906000526020600020900160005b9054600160a060020a036101009290920a900416815260208101919091526040016000205460ff1615610fe3576003805482908110610f9857fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316838381518110610fc457fe5b600160a060020a03909216602092830290910190910152600191909101905b5b600101610f26565b81604051805910610ffa5750595b908082528060200260200182016040525b509350600090505b8181101561105d5782818151811061102757fe5b9060200190602002015184828151811061103d57fe5b600160a060020a039092166020928302909101909101525b600101611013565b5b505050919050565b60055481565b30600160a060020a031633600160a060020a031614151561108c57600080fd5b60035481603282118015906110a15750818111155b80156110ac57508015155b80156110b757508115155b15156110c257600080fd5b60048390557fa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a8360405190815260200160405180910390a15b5b50505b50565b33600160a060020a03811660009081526002602052604090205460ff16151561112a57600080fd5b6000828152602081905260409020548290600160a060020a0316151561114f57600080fd5b600083815260016020908152604080832033600160a060020a038116855292529091205484919060ff161561118357600080fd5b6000858152600160208181526040808420600160a060020a033316808652925292839020805460ff191690921790915586917f4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef905160405180910390a36109d58561143d565b5b5b50505b505b5050565b6000611201848484611669565b905061120c81611102565b5b9392505050565b30600160a060020a031633600160a060020a031614151561123457600080fd5b60068190557fc71bdc6afaf9b1aa90a7078191d4fc1adf3bf680fca3183697df6b0dc226bca28160405190815260200160405180910390a15b5b50565b603281565b60045481565b600030600160a060020a031633600160a060020a031614151561129e57600080fd5b600160a060020a038316600090815260026020526040902054839060ff1615156112c757600080fd5b600160a060020a038316600090815260026020526040902054839060ff16156112ef57600080fd5b600092505b6003548310156113975784600160a060020a031660038481548110151561131757fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316600160a060020a0316141561138b578360038481548110151561135657fe5b906000526020600020900160005b6101000a815481600160a060020a030219169083600160a060020a03160217905550611397565b5b6001909201916112f4565b600160a060020a03808616600081815260026020526040808220805460ff199081169091559388168252908190208054909316600117909255907f8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90905160405180910390a283600160a060020a03167ff39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d60405160405180910390a25b5b505b505b505050565b33600160a060020a0381166000908152600260205260408120549091829160ff16151561146957600080fd5b600084815260016020908152604080832033600160a060020a038116855292529091205485919060ff16151561149e57600080fd5b600086815260208190526040902060030154869060ff16156114bf57600080fd5b600087815260208190526040902095506114d887610c06565b9450848061150b575060028087015460001961010060018316150201160415801561150b575061150b8660010154611768565b5b5b156116545760038601805460ff191660011790558415156115375760018601546008805490910190555b85546001870154600160a060020a03909116906002880160405180828054600181600116156101000203166002900480156115b35780601f10611588576101008083540402835291602001916115b3565b820191906000526020600020905b81548152906001019060200180831161159657829003601f168201915b505091505060006040518083038185876187965a03f1925050501561160457867f33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed7560405160405180910390a2611654565b867f526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b7923660405160405180910390a260038601805460ff19169055841515611654576001860154600880549190910390555b5b5b5b5b505b50505b50505050565b60085481565b600083600160a060020a038116151561168157600080fd5b600554915060806040519081016040908152600160a060020a0387168252602080830187905281830186905260006060840181905285815290819052208151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03919091161781556020820151816001015560408201518160020190805161170c929160200190611828565b506060820151600391909101805460ff191691151591909117905550600580546001019055817fc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e5160405160405180910390a25b5b509392505050565b60006007546201518001421115611783574260075560006008555b6006548260085401118061179a5750600854828101105b156117a7575060006117ab565b5060015b919050565b8154818355818115116108f6576000838152602090206108f69181019083016118a7565b5b505050565b8154818355818115116108f6576000838152602090206108f69181019083016118a7565b5b505050565b60206040519081016040526000815290565b60206040519081016040526000815290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061186957805160ff1916838001178555611896565b82800160010185558215611896579182015b8281111561189657825182559160200191906001019061187b565b5b506118a39291506118a7565b5090565b610a4b91905b808211156118a357600081556001016118ad565b5090565b905600a165627a7a72305820e75e9113bdeed0d8495f54139848753af6dda4daa0bac3977bea26696ae8ee2a0029`

// DeployMultiSigWalletWithDailyLimit deploys a new Ethereum contract, binding an instance of MultiSigWalletWithDailyLimit to it.
func DeployMultiSigWalletWithDailyLimit(auth *bind.TransactOpts, backend bind.ContractBackend, _owners []common.Address, _required *big.Int, _dailyLimit *big.Int) (common.Address, *types.Transaction, *MultiSigWalletWithDailyLimit, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiSigWalletWithDailyLimitABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MultiSigWalletWithDailyLimitBin), backend, _owners, _required, _dailyLimit)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MultiSigWalletWithDailyLimit{MultiSigWalletWithDailyLimitCaller: MultiSigWalletWithDailyLimitCaller{contract: contract}, MultiSigWalletWithDailyLimitTransactor: MultiSigWalletWithDailyLimitTransactor{contract: contract}, MultiSigWalletWithDailyLimitFilterer: MultiSigWalletWithDailyLimitFilterer{contract: contract}}, nil
}

// MultiSigWalletWithDailyLimit is an auto generated Go binding around an Ethereum contract.
type MultiSigWalletWithDailyLimit struct {
	MultiSigWalletWithDailyLimitCaller     // Read-only binding to the contract
	MultiSigWalletWithDailyLimitTransactor // Write-only binding to the contract
	MultiSigWalletWithDailyLimitFilterer   // Log filterer for contract events
}

// MultiSigWalletWithDailyLimitCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiSigWalletWithDailyLimitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigWalletWithDailyLimitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiSigWalletWithDailyLimitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigWalletWithDailyLimitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiSigWalletWithDailyLimitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigWalletWithDailyLimitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiSigWalletWithDailyLimitSession struct {
	Contract     *MultiSigWalletWithDailyLimit // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MultiSigWalletWithDailyLimitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiSigWalletWithDailyLimitCallerSession struct {
	Contract *MultiSigWalletWithDailyLimitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// MultiSigWalletWithDailyLimitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiSigWalletWithDailyLimitTransactorSession struct {
	Contract     *MultiSigWalletWithDailyLimitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// MultiSigWalletWithDailyLimitRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiSigWalletWithDailyLimitRaw struct {
	Contract *MultiSigWalletWithDailyLimit // Generic contract binding to access the raw methods on
}

// MultiSigWalletWithDailyLimitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiSigWalletWithDailyLimitCallerRaw struct {
	Contract *MultiSigWalletWithDailyLimitCaller // Generic read-only contract binding to access the raw methods on
}

// MultiSigWalletWithDailyLimitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiSigWalletWithDailyLimitTransactorRaw struct {
	Contract *MultiSigWalletWithDailyLimitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiSigWalletWithDailyLimit creates a new instance of MultiSigWalletWithDailyLimit, bound to a specific deployed contract.
func NewMultiSigWalletWithDailyLimit(address common.Address, backend bind.ContractBackend) (*MultiSigWalletWithDailyLimit, error) {
	contract, err := bindMultiSigWalletWithDailyLimit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimit{MultiSigWalletWithDailyLimitCaller: MultiSigWalletWithDailyLimitCaller{contract: contract}, MultiSigWalletWithDailyLimitTransactor: MultiSigWalletWithDailyLimitTransactor{contract: contract}, MultiSigWalletWithDailyLimitFilterer: MultiSigWalletWithDailyLimitFilterer{contract: contract}}, nil
}

// NewMultiSigWalletWithDailyLimitCaller creates a new read-only instance of MultiSigWalletWithDailyLimit, bound to a specific deployed contract.
func NewMultiSigWalletWithDailyLimitCaller(address common.Address, caller bind.ContractCaller) (*MultiSigWalletWithDailyLimitCaller, error) {
	contract, err := bindMultiSigWalletWithDailyLimit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitCaller{contract: contract}, nil
}

// NewMultiSigWalletWithDailyLimitTransactor creates a new write-only instance of MultiSigWalletWithDailyLimit, bound to a specific deployed contract.
func NewMultiSigWalletWithDailyLimitTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiSigWalletWithDailyLimitTransactor, error) {
	contract, err := bindMultiSigWalletWithDailyLimit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitTransactor{contract: contract}, nil
}

// NewMultiSigWalletWithDailyLimitFilterer creates a new log filterer instance of MultiSigWalletWithDailyLimit, bound to a specific deployed contract.
func NewMultiSigWalletWithDailyLimitFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiSigWalletWithDailyLimitFilterer, error) {
	contract, err := bindMultiSigWalletWithDailyLimit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitFilterer{contract: contract}, nil
}

// bindMultiSigWalletWithDailyLimit binds a generic wrapper to an already deployed contract.
func bindMultiSigWalletWithDailyLimit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiSigWalletWithDailyLimitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MultiSigWalletWithDailyLimit.Contract.MultiSigWalletWithDailyLimitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.MultiSigWalletWithDailyLimitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.MultiSigWalletWithDailyLimitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MultiSigWalletWithDailyLimit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.contract.Transact(opts, method, params...)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() constant returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) MAXOWNERCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "MAX_OWNER_COUNT")
	return *ret0, err
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() constant returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) MAXOWNERCOUNT() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.MAXOWNERCOUNT(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() constant returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) MAXOWNERCOUNT() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.MAXOWNERCOUNT(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// CalcMaxWithdraw is a free data retrieval call binding the contract method 0x4bc9fdc2.
//
// Solidity: function calcMaxWithdraw() constant returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) CalcMaxWithdraw(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "calcMaxWithdraw")
	return *ret0, err
}

// CalcMaxWithdraw is a free data retrieval call binding the contract method 0x4bc9fdc2.
//
// Solidity: function calcMaxWithdraw() constant returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) CalcMaxWithdraw() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.CalcMaxWithdraw(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// CalcMaxWithdraw is a free data retrieval call binding the contract method 0x4bc9fdc2.
//
// Solidity: function calcMaxWithdraw() constant returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) CalcMaxWithdraw() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.CalcMaxWithdraw(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) constant returns(bool)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) Confirmations(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "confirmations", arg0, arg1)
	return *ret0, err
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) constant returns(bool)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _MultiSigWalletWithDailyLimit.Contract.Confirmations(&_MultiSigWalletWithDailyLimit.CallOpts, arg0, arg1)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) constant returns(bool)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _MultiSigWalletWithDailyLimit.Contract.Confirmations(&_MultiSigWalletWithDailyLimit.CallOpts, arg0, arg1)
}

// DailyLimit is a free data retrieval call binding the contract method 0x67eeba0c.
//
// Solidity: function dailyLimit() constant returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) DailyLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "dailyLimit")
	return *ret0, err
}

// DailyLimit is a free data retrieval call binding the contract method 0x67eeba0c.
//
// Solidity: function dailyLimit() constant returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) DailyLimit() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.DailyLimit(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// DailyLimit is a free data retrieval call binding the contract method 0x67eeba0c.
//
// Solidity: function dailyLimit() constant returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) DailyLimit() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.DailyLimit(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) constant returns(uint256 count)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) GetConfirmationCount(opts *bind.CallOpts, transactionId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "getConfirmationCount", transactionId)
	return *ret0, err
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) constant returns(uint256 count)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetConfirmationCount(&_MultiSigWalletWithDailyLimit.CallOpts, transactionId)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) constant returns(uint256 count)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetConfirmationCount(&_MultiSigWalletWithDailyLimit.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) constant returns(address[] _confirmations)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) GetConfirmations(opts *bind.CallOpts, transactionId *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "getConfirmations", transactionId)
	return *ret0, err
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) constant returns(address[] _confirmations)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetConfirmations(&_MultiSigWalletWithDailyLimit.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) constant returns(address[] _confirmations)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetConfirmations(&_MultiSigWalletWithDailyLimit.CallOpts, transactionId)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "getOwners")
	return *ret0, err
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) GetOwners() ([]common.Address, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetOwners(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) GetOwners() ([]common.Address, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetOwners(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) constant returns(uint256 count)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) GetTransactionCount(opts *bind.CallOpts, pending bool, executed bool) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "getTransactionCount", pending, executed)
	return *ret0, err
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) constant returns(uint256 count)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetTransactionCount(&_MultiSigWalletWithDailyLimit.CallOpts, pending, executed)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) constant returns(uint256 count)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetTransactionCount(&_MultiSigWalletWithDailyLimit.CallOpts, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) constant returns(uint256[] _transactionIds)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) GetTransactionIds(opts *bind.CallOpts, from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "getTransactionIds", from, to, pending, executed)
	return *ret0, err
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) constant returns(uint256[] _transactionIds)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetTransactionIds(&_MultiSigWalletWithDailyLimit.CallOpts, from, to, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) constant returns(uint256[] _transactionIds)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetTransactionIds(&_MultiSigWalletWithDailyLimit.CallOpts, from, to, pending, executed)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) constant returns(bool)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) IsConfirmed(opts *bind.CallOpts, transactionId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "isConfirmed", transactionId)
	return *ret0, err
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) constant returns(bool)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _MultiSigWalletWithDailyLimit.Contract.IsConfirmed(&_MultiSigWalletWithDailyLimit.CallOpts, transactionId)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) constant returns(bool)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _MultiSigWalletWithDailyLimit.Contract.IsConfirmed(&_MultiSigWalletWithDailyLimit.CallOpts, transactionId)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) constant returns(bool)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) IsOwner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "isOwner", arg0)
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) constant returns(bool)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) IsOwner(arg0 common.Address) (bool, error) {
	return _MultiSigWalletWithDailyLimit.Contract.IsOwner(&_MultiSigWalletWithDailyLimit.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) constant returns(bool)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) IsOwner(arg0 common.Address) (bool, error) {
	return _MultiSigWalletWithDailyLimit.Contract.IsOwner(&_MultiSigWalletWithDailyLimit.CallOpts, arg0)
}

// LastDay is a free data retrieval call binding the contract method 0x6b0c932d.
//
// Solidity: function lastDay() constant returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) LastDay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "lastDay")
	return *ret0, err
}

// LastDay is a free data retrieval call binding the contract method 0x6b0c932d.
//
// Solidity: function lastDay() constant returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) LastDay() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.LastDay(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// LastDay is a free data retrieval call binding the contract method 0x6b0c932d.
//
// Solidity: function lastDay() constant returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) LastDay() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.LastDay(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) constant returns(address)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "owners", arg0)
	return *ret0, err
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) constant returns(address)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _MultiSigWalletWithDailyLimit.Contract.Owners(&_MultiSigWalletWithDailyLimit.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) constant returns(address)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _MultiSigWalletWithDailyLimit.Contract.Owners(&_MultiSigWalletWithDailyLimit.CallOpts, arg0)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() constant returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) Required(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "required")
	return *ret0, err
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() constant returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) Required() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.Required(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() constant returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) Required() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.Required(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// SpentToday is a free data retrieval call binding the contract method 0xf059cf2b.
//
// Solidity: function spentToday() constant returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) SpentToday(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "spentToday")
	return *ret0, err
}

// SpentToday is a free data retrieval call binding the contract method 0xf059cf2b.
//
// Solidity: function spentToday() constant returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) SpentToday() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.SpentToday(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// SpentToday is a free data retrieval call binding the contract method 0xf059cf2b.
//
// Solidity: function spentToday() constant returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) SpentToday() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.SpentToday(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() constant returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) TransactionCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "transactionCount")
	return *ret0, err
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() constant returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) TransactionCount() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.TransactionCount(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() constant returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) TransactionCount() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.TransactionCount(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) constant returns(address destination, uint256 value, bytes data, bool executed)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) Transactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	ret := new(struct {
		Destination common.Address
		Value       *big.Int
		Data        []byte
		Executed    bool
	})
	out := ret
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "transactions", arg0)
	return *ret, err
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) constant returns(address destination, uint256 value, bytes data, bool executed)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) Transactions(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _MultiSigWalletWithDailyLimit.Contract.Transactions(&_MultiSigWalletWithDailyLimit.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) constant returns(address destination, uint256 value, bytes data, bool executed)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) Transactions(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _MultiSigWalletWithDailyLimit.Contract.Transactions(&_MultiSigWalletWithDailyLimit.CallOpts, arg0)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactor) AddOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.contract.Transact(opts, "addOwner", owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) AddOwner(owner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.AddOwner(&_MultiSigWalletWithDailyLimit.TransactOpts, owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorSession) AddOwner(owner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.AddOwner(&_MultiSigWalletWithDailyLimit.TransactOpts, owner)
}

// ChangeDailyLimit is a paid mutator transaction binding the contract method 0xcea08621.
//
// Solidity: function changeDailyLimit(uint256 _dailyLimit) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactor) ChangeDailyLimit(opts *bind.TransactOpts, _dailyLimit *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.contract.Transact(opts, "changeDailyLimit", _dailyLimit)
}

// ChangeDailyLimit is a paid mutator transaction binding the contract method 0xcea08621.
//
// Solidity: function changeDailyLimit(uint256 _dailyLimit) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) ChangeDailyLimit(_dailyLimit *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ChangeDailyLimit(&_MultiSigWalletWithDailyLimit.TransactOpts, _dailyLimit)
}

// ChangeDailyLimit is a paid mutator transaction binding the contract method 0xcea08621.
//
// Solidity: function changeDailyLimit(uint256 _dailyLimit) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorSession) ChangeDailyLimit(_dailyLimit *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ChangeDailyLimit(&_MultiSigWalletWithDailyLimit.TransactOpts, _dailyLimit)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactor) ChangeRequirement(opts *bind.TransactOpts, _required *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.contract.Transact(opts, "changeRequirement", _required)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) ChangeRequirement(_required *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ChangeRequirement(&_MultiSigWalletWithDailyLimit.TransactOpts, _required)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorSession) ChangeRequirement(_required *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ChangeRequirement(&_MultiSigWalletWithDailyLimit.TransactOpts, _required)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactor) ConfirmTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.contract.Transact(opts, "confirmTransaction", transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ConfirmTransaction(&_MultiSigWalletWithDailyLimit.TransactOpts, transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorSession) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ConfirmTransaction(&_MultiSigWalletWithDailyLimit.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactor) ExecuteTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.contract.Transact(opts, "executeTransaction", transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ExecuteTransaction(&_MultiSigWalletWithDailyLimit.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorSession) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ExecuteTransaction(&_MultiSigWalletWithDailyLimit.TransactOpts, transactionId)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactor) RemoveOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.contract.Transact(opts, "removeOwner", owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) RemoveOwner(owner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.RemoveOwner(&_MultiSigWalletWithDailyLimit.TransactOpts, owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorSession) RemoveOwner(owner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.RemoveOwner(&_MultiSigWalletWithDailyLimit.TransactOpts, owner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactor) ReplaceOwner(opts *bind.TransactOpts, owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.contract.Transact(opts, "replaceOwner", owner, newOwner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) ReplaceOwner(owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ReplaceOwner(&_MultiSigWalletWithDailyLimit.TransactOpts, owner, newOwner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorSession) ReplaceOwner(owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ReplaceOwner(&_MultiSigWalletWithDailyLimit.TransactOpts, owner, newOwner)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactor) RevokeConfirmation(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.contract.Transact(opts, "revokeConfirmation", transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.RevokeConfirmation(&_MultiSigWalletWithDailyLimit.TransactOpts, transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorSession) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.RevokeConfirmation(&_MultiSigWalletWithDailyLimit.TransactOpts, transactionId)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactor) SubmitTransaction(opts *bind.TransactOpts, destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.contract.Transact(opts, "submitTransaction", destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.SubmitTransaction(&_MultiSigWalletWithDailyLimit.TransactOpts, destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorSession) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.SubmitTransaction(&_MultiSigWalletWithDailyLimit.TransactOpts, destination, value, data)
}

// MultiSigWalletWithDailyLimitConfirmationIterator is returned from FilterConfirmation and is used to iterate over the raw logs and unpacked data for Confirmation events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitConfirmationIterator struct {
	Event *MultiSigWalletWithDailyLimitConfirmation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitConfirmation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitConfirmation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitConfirmation represents a Confirmation event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitConfirmation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterConfirmation is a free log retrieval operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterConfirmation(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*MultiSigWalletWithDailyLimitConfirmationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "Confirmation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitConfirmationIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "Confirmation", logs: logs, sub: sub}, nil
}

// WatchConfirmation is a free log subscription operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchConfirmation(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitConfirmation, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "Confirmation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitConfirmation)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "Confirmation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MultiSigWalletWithDailyLimitDailyLimitChangeIterator is returned from FilterDailyLimitChange and is used to iterate over the raw logs and unpacked data for DailyLimitChange events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitDailyLimitChangeIterator struct {
	Event *MultiSigWalletWithDailyLimitDailyLimitChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitDailyLimitChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitDailyLimitChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitDailyLimitChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitDailyLimitChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitDailyLimitChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitDailyLimitChange represents a DailyLimitChange event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitDailyLimitChange struct {
	DailyLimit *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDailyLimitChange is a free log retrieval operation binding the contract event 0xc71bdc6afaf9b1aa90a7078191d4fc1adf3bf680fca3183697df6b0dc226bca2.
//
// Solidity: event DailyLimitChange(uint256 dailyLimit)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterDailyLimitChange(opts *bind.FilterOpts) (*MultiSigWalletWithDailyLimitDailyLimitChangeIterator, error) {

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "DailyLimitChange")
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitDailyLimitChangeIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "DailyLimitChange", logs: logs, sub: sub}, nil
}

// WatchDailyLimitChange is a free log subscription operation binding the contract event 0xc71bdc6afaf9b1aa90a7078191d4fc1adf3bf680fca3183697df6b0dc226bca2.
//
// Solidity: event DailyLimitChange(uint256 dailyLimit)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchDailyLimitChange(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitDailyLimitChange) (event.Subscription, error) {

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "DailyLimitChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitDailyLimitChange)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "DailyLimitChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MultiSigWalletWithDailyLimitDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitDepositIterator struct {
	Event *MultiSigWalletWithDailyLimitDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitDeposit represents a Deposit event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitDeposit struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 value)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*MultiSigWalletWithDailyLimitDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitDepositIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 value)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitDeposit)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MultiSigWalletWithDailyLimitExecutionIterator is returned from FilterExecution and is used to iterate over the raw logs and unpacked data for Execution events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitExecutionIterator struct {
	Event *MultiSigWalletWithDailyLimitExecution // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitExecution)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitExecution)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitExecution represents a Execution event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitExecution struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecution is a free log retrieval operation binding the contract event 0x33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed75.
//
// Solidity: event Execution(uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterExecution(opts *bind.FilterOpts, transactionId []*big.Int) (*MultiSigWalletWithDailyLimitExecutionIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "Execution", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitExecutionIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "Execution", logs: logs, sub: sub}, nil
}

// WatchExecution is a free log subscription operation binding the contract event 0x33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed75.
//
// Solidity: event Execution(uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchExecution(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitExecution, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "Execution", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitExecution)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "Execution", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MultiSigWalletWithDailyLimitExecutionFailureIterator is returned from FilterExecutionFailure and is used to iterate over the raw logs and unpacked data for ExecutionFailure events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitExecutionFailureIterator struct {
	Event *MultiSigWalletWithDailyLimitExecutionFailure // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitExecutionFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitExecutionFailure)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitExecutionFailure)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitExecutionFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitExecutionFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitExecutionFailure represents a ExecutionFailure event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitExecutionFailure struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionFailure is a free log retrieval operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterExecutionFailure(opts *bind.FilterOpts, transactionId []*big.Int) (*MultiSigWalletWithDailyLimitExecutionFailureIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "ExecutionFailure", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitExecutionFailureIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "ExecutionFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFailure is a free log subscription operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchExecutionFailure(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitExecutionFailure, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "ExecutionFailure", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitExecutionFailure)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MultiSigWalletWithDailyLimitOwnerAdditionIterator is returned from FilterOwnerAddition and is used to iterate over the raw logs and unpacked data for OwnerAddition events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitOwnerAdditionIterator struct {
	Event *MultiSigWalletWithDailyLimitOwnerAddition // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitOwnerAdditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitOwnerAddition)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitOwnerAddition)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitOwnerAdditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitOwnerAdditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitOwnerAddition represents a OwnerAddition event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitOwnerAddition struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerAddition is a free log retrieval operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address indexed owner)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterOwnerAddition(opts *bind.FilterOpts, owner []common.Address) (*MultiSigWalletWithDailyLimitOwnerAdditionIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "OwnerAddition", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitOwnerAdditionIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "OwnerAddition", logs: logs, sub: sub}, nil
}

// WatchOwnerAddition is a free log subscription operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address indexed owner)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchOwnerAddition(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitOwnerAddition, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "OwnerAddition", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitOwnerAddition)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "OwnerAddition", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MultiSigWalletWithDailyLimitOwnerRemovalIterator is returned from FilterOwnerRemoval and is used to iterate over the raw logs and unpacked data for OwnerRemoval events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitOwnerRemovalIterator struct {
	Event *MultiSigWalletWithDailyLimitOwnerRemoval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitOwnerRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitOwnerRemoval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitOwnerRemoval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitOwnerRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitOwnerRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitOwnerRemoval represents a OwnerRemoval event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitOwnerRemoval struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerRemoval is a free log retrieval operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address indexed owner)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterOwnerRemoval(opts *bind.FilterOpts, owner []common.Address) (*MultiSigWalletWithDailyLimitOwnerRemovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "OwnerRemoval", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitOwnerRemovalIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "OwnerRemoval", logs: logs, sub: sub}, nil
}

// WatchOwnerRemoval is a free log subscription operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address indexed owner)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchOwnerRemoval(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitOwnerRemoval, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "OwnerRemoval", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitOwnerRemoval)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "OwnerRemoval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MultiSigWalletWithDailyLimitRequirementChangeIterator is returned from FilterRequirementChange and is used to iterate over the raw logs and unpacked data for RequirementChange events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitRequirementChangeIterator struct {
	Event *MultiSigWalletWithDailyLimitRequirementChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitRequirementChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitRequirementChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitRequirementChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitRequirementChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitRequirementChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitRequirementChange represents a RequirementChange event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitRequirementChange struct {
	Required *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRequirementChange is a free log retrieval operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterRequirementChange(opts *bind.FilterOpts) (*MultiSigWalletWithDailyLimitRequirementChangeIterator, error) {

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "RequirementChange")
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitRequirementChangeIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "RequirementChange", logs: logs, sub: sub}, nil
}

// WatchRequirementChange is a free log subscription operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchRequirementChange(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitRequirementChange) (event.Subscription, error) {

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "RequirementChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitRequirementChange)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "RequirementChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MultiSigWalletWithDailyLimitRevocationIterator is returned from FilterRevocation and is used to iterate over the raw logs and unpacked data for Revocation events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitRevocationIterator struct {
	Event *MultiSigWalletWithDailyLimitRevocation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitRevocationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitRevocation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitRevocation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitRevocationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitRevocationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitRevocation represents a Revocation event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitRevocation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRevocation is a free log retrieval operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterRevocation(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*MultiSigWalletWithDailyLimitRevocationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "Revocation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitRevocationIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "Revocation", logs: logs, sub: sub}, nil
}

// WatchRevocation is a free log subscription operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchRevocation(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitRevocation, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "Revocation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitRevocation)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "Revocation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MultiSigWalletWithDailyLimitSubmissionIterator is returned from FilterSubmission and is used to iterate over the raw logs and unpacked data for Submission events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitSubmissionIterator struct {
	Event *MultiSigWalletWithDailyLimitSubmission // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitSubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitSubmission)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitSubmission)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitSubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitSubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitSubmission represents a Submission event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitSubmission struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSubmission is a free log retrieval operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterSubmission(opts *bind.FilterOpts, transactionId []*big.Int) (*MultiSigWalletWithDailyLimitSubmissionIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "Submission", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitSubmissionIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "Submission", logs: logs, sub: sub}, nil
}

// WatchSubmission is a free log subscription operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchSubmission(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitSubmission, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "Submission", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitSubmission)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "Submission", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
