package cli

import (
	"crypto/elliptic"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

const (
	newchainPublicKey = "c829d38b9fc274c8cb13b239a2b473ec04363167a84f2b4d6666b286f78c92515228bb895ac3802285cde0bac18592efbaffeb1bc14e1da00139b7dbf5248375"
	ethereumPublicKey = "979b7fa28feeb35a4741660a16076f1943202cb72b6af70d327f053e248bab9ba81760f39d0701ef1d8f89cc1fbd2cacba0710a12cd5314d5e0c9021aa3637f9"
)

// BlockChain type
type BlockChain int

const (
	UnknownChain BlockChain = iota
	NewChain
	Ethereum
)

func (bc BlockChain) String() string {
	switch bc {
	case NewChain:
		return "NewChain"
	case Ethereum:
		return "Ethereum"
	}

	return "UnknownChain"
}

func getBlockChain() (BlockChain, error) {
	// Check NewChain
	b, err := hex.DecodeString(newchainPublicKey)
	if err != nil {
		return UnknownChain, err
	} else if len(b) != 64 {
		return UnknownChain, fmt.Errorf("wrong length, want %d hex chars\n", 128)
	}
	b = append([]byte{0x4}, b...)

	x, _ := elliptic.Unmarshal(crypto.S256(), b)
	if x != nil {
		// OK
		return NewChain, nil
	}

	// Check Ethereum
	be, err := hex.DecodeString(ethereumPublicKey)
	if err != nil {
		return UnknownChain, err
	} else if len(be) != 64 {
		return UnknownChain, fmt.Errorf("wrong length, want %d hex chars\n", 128)
	}
	be = append([]byte{0x4}, be...)

	xb, _ := elliptic.Unmarshal(crypto.S256(), be)
	if xb != nil {
		// OK
		return Ethereum, nil
	}

	return UnknownChain, nil
}

func (bc BlockChain) Init() {
	InitRPCUrl(bc)
	InitUnit(bc)
	InitERC(bc)
}

// Unit
var (
	UnitETH = "ETH"
	UnitWEI = "WEI"

	// UnitList is array for Unit string
	// UnitList = []string{"Wei", "Ada", "Babbage", "Shannon", "Szabo", "Finney", "Ether", "Einstein", "Douglas", "Gwei"}
	// UnitLis = []string{"ISAAC", "NEW"}
	UnitList []string
)

func InitUnit(bc BlockChain) {
	if bc == NewChain {
		UnitETH = "NEW"
		UnitWEI = "ISAAC"

	}

	UnitList = []string{UnitETH, UnitWEI}
}

var (
	// ModeERC20 https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20.md
	ModeERC20 = "ERC20"
	// ModeERC721 https://github.com/ethereum/EIPs/blob/master/EIPS/eip-721.md
	ModeERC721 = "ERC721"

	// UnitList is array for Unit string
	// UnitList = []string{"Wei", "Ada", "Babbage", "Shannon", "Szabo", "Finney", "Ether", "Einstein", "Douglas", "Gwei"}
	ModeERCList []string
)

var (
	errOnlyERC20  error
	errOnlyERC721 error
)

func InitERC(bc BlockChain) {
	if bc == NewChain {
		ModeERC20 = "NEP6"
		ModeERC721 = "NEP7"
	}

	ModeERCList = []string{ModeERC20, ModeERC721}

	// error
	errOnlyERC20 = fmt.Errorf("only %s support", ModeERC20)
	errOnlyERC721 = fmt.Errorf("only %s support", ModeERC721)
}

// rpc
var defaultRPCURL string

const defaultNEWRPCURL = "https://rpc1.newchain.newtonproject.org"
const defaultETHRPCUrl = "https://ethrpc.service.newtonproject.org"

func InitRPCUrl(bc BlockChain) {
	// default RPC Url
	defaultRPCURL = defaultETHRPCUrl
	if bc == NewChain {
		defaultRPCURL = defaultNEWRPCURL
	}
}
