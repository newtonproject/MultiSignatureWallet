package cli

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

var (
	buildCommit string
	buildDate   string
)

// CLI represents a command-line interface. This class is
// not threadsafe.
type CLI struct {
	Name       string
	rootCmd    *cobra.Command
	version    string
	walletPath string
	rpcURL     string
	config     string
	//testing    bool

	contractAddress string
	client          *ethclient.Client
	wallet          *keystore.KeyStore
	account         accounts.Account
	simpleRegistry  *SimpleRegistry
	walletPassword  string
	address         string

	tran *Transaction
}

type SimpleRegistry struct {
	*MultiSigWalletWithDailyLimit
}

// NewCLI returns an initialized CLI
func NewCLI() *CLI {
	version := "v1.1.0"
	if buildCommit != "" {
		version = fmt.Sprintf("%s-%s", version, buildCommit)
	}
	if buildDate != "" {
		version = fmt.Sprintf("%s-%s", version, buildDate)
	}

	cli := &CLI{
		Name:       "MultiSigWallet",
		rootCmd:    nil,
		version:    version,
		walletPath: "",
		rpcURL:     "",
		//	testing:         false,
		config:          "",
		contractAddress: "",
		client:          nil,
		simpleRegistry:  nil,
		walletPassword:  "",
	}

	cli.buildRootCmd()
	return cli
}

// BuildClient BuildClient
func (cli *CLI) BuildClient() error {
	var err error
	if cli.client == nil {
		cli.client, err = ethclient.Dial(cli.rpcURL)
		if err != nil {
			return fmt.Errorf("Failed to connect to the NewChain client: %v", err)
		}
	}
	return nil
}

// BuildSimpleRegistry BuildClient
func (cli *CLI) buildSimpleRegistry() (*SimpleRegistry, error) {
	if cli.client == nil {
		if err := cli.BuildClient(); err != nil {
			return nil, err
		}
	}

	if !common.IsHexAddress(cli.contractAddress) {
		return nil, fmt.Errorf("contract address is invalid")
	}

	m, _ := NewMultiSigWalletWithDailyLimit(common.HexToAddress(cli.contractAddress), cli.client)
	cli.simpleRegistry = &SimpleRegistry{m}
	return cli.simpleRegistry, nil
}

//GetSimpleRegistry GetSimpleRegistry
func (cli *CLI) GetSimpleRegistry() (*SimpleRegistry, error) {
	if cli.simpleRegistry == nil {
		return cli.buildSimpleRegistry()
	}
	return cli.simpleRegistry, nil
}

func (cli *CLI) buildWallet() error {
	if cli.wallet == nil {
		cli.wallet = keystore.NewKeyStore(cli.walletPath,
			keystore.LightScryptN, keystore.LightScryptP)
		if len(cli.wallet.Accounts()) == 0 {
			return fmt.Errorf("Empty wallet, create account first")
		}
	}

	return nil
}

func (cli *CLI) buildAccount(address string) error {

	err := cli.buildWallet()
	if err != nil {
		return err
	}

	if !common.IsHexAddress(address) {
		if common.IsHexAddress(cli.address) {
			address = cli.address
		} else {
			return fmt.Errorf("Error: address(%s) invalid", address)
		}
	}
	for _, a := range cli.wallet.Accounts() {
		if a.Address == common.HexToAddress(address) {
			cli.account = a
			break
		}
	}
	if cli.account == (accounts.Account{}) {
		return fmt.Errorf("Error: Can NOT get the keystore file of address %s", address)
	}
	cli.address = address

	return nil
}

func (cli *CLI) getTransactOpts(address string) (*bind.TransactOpts, error) {
	err := cli.buildAccount(address)
	if err != nil {
		return nil, err
	}

	var trials int
	//var walletPassword string
	var keyJSON []byte
	for trials = 0; trials <= 3; trials++ {
		keyJSON, err = cli.wallet.Export(cli.account, cli.walletPassword, cli.walletPassword)
		if err == nil {
			break
		}
		if trials >= 3 {
			return nil, fmt.Errorf("Error: Failed to unlock account %s (%v)", cli.account.Address.String(), err)

		}
		prompt := fmt.Sprintf("Unlocking account %s | Attempt %d/%d", cli.account.Address.String(), trials+1, 3)
		cli.walletPassword, _ = getPassPhrase(prompt, false)
	}

	opts, err := bind.NewTransactor(bytes.NewReader(keyJSON), cli.walletPassword)
	if err != nil {
		fmt.Println("NewTransactor Error: ", err)
		return nil, err
	}

	if err = cli.BuildClient(); err != nil {
		fmt.Println("Build client Error: ", err)
		return nil, err
	}
	networkID, err := cli.client.NetworkID(context.Background())
	if err != nil {
		fmt.Println("NetworkID Error: ", err)
		return nil, err
	}
	bind.NetworkID = networkID

	return opts, err
}

// Execute parses the command line and processes it.
func (cli *CLI) Execute() {
	cli.rootCmd.Execute()
}

// setup turns up the CLI environment, and gets called by Cobra before
// a command is executed.
func (cli *CLI) setup(cmd *cobra.Command, args []string) {
	err := setupConfig(cli)
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(os.Stderr, cmd.UsageString())
		os.Exit(1)
	}
}

func (cli *CLI) help(cmd *cobra.Command, args []string) {
	fmt.Fprint(os.Stderr, cmd.UsageString())

	os.Exit(-1)

}

// TestCommand test command
func (cli *CLI) TestCommand(command string) string {
	//cli.testing = true
	result := cli.Run(strings.Fields(command)...)
	//	cli.testing = false
	return result
}

// Run executes CLI with the given arguments. Used for testing. Not thread safe.
func (cli *CLI) Run(args ...string) string {
	oldStdout := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	cli.rootCmd.SetArgs(args)
	cli.rootCmd.Execute()
	cli.buildRootCmd()

	w.Close()

	os.Stdout = oldStdout

	var stdOut bytes.Buffer
	io.Copy(&stdOut, r)
	return stdOut.String()
}

// Embeddable returns a CLI that you can embed into your own Go programs. This
// is not thread-safe.
func (cli *CLI) Embeddable() *CLI {

	return cli
}

// SetPassword SetPassword
func (cli *CLI) SetPassword(_passPhrase string) *CLI {
	cli.walletPassword = _passPhrase
	return cli
}
