package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/newtonproject/MultiSignatureWallet/cli"
)

// test address
const taddr = "0xDB2C9C06E186D58EFe19f213b3d5FaF8B8c99481"

func getTempFile() (string, func()) {
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		log.Fatal(err)
	}

	file := dir + string(os.PathSeparator) + "lumen-integration-test.json"

	return file, func() {
		logrus.Debugf("cleaning up temp file: %s", file)
		os.RemoveAll(dir)
	}
}

func run(cli *cli.CLI, command string) string {
	fmt.Printf("$ ./%s %s\n", cli.Name, command)
	got := cli.TestCommand(command)
	fmt.Printf("%s\n", got)
	return strings.TrimSpace(got)
}

func runArgs(cli *cli.CLI, args ...string) string {
	fmt.Printf("$ ./%s %s\n", cli.Name, strings.Join(args, " "))
	got := cli.Embeddable().Run(args...)
	fmt.Printf("%s\n", got)
	return strings.TrimSpace(got)
}

func expectOutput(t *testing.T, cli *cli.CLI, want string, command string) {
	got := run(cli, command)

	if got != want {
		t.Errorf("(%s) wrong output: want %v, got %v", command, want, got)
	}
}

func newCLI() (*cli.CLI, func()) {
	_, cleanupFunc := getTempFile()

	dpos := cli.NewCLI()
	dpos.TestCommand("version")
	run(dpos, fmt.Sprintf("version"))

	return dpos, cleanupFunc
}

func getBalance(cli *cli.CLI, account string) float64 {
	balanceString := run(cli, "balance "+account)

	balance, err := strconv.ParseFloat(balanceString, 64)

	if err != nil {
		return 0
	}

	return balance
}

// Create new funded test account
func TestTx(t *testing.T) {
	cli, _ := newCLI()
	cli.SetPassword("test")
	address := run(cli, "account new") // no password

	run(cli, "deploy -o 0xdDeB86Dd09F16316B67322199E288d7AF35E0806,0x9B3deA9C636BA262f870f98a1c64d444BF0f6544,0xc8B5c4cB6DB7254d082b24A96627F143E8A80c31 -r 2 -l 1024 --from "+address)
	run(cli, "info")

	run(cli, "tx submit 1 --to 0xdDeB86Dd09F16316B67322199E288d7AF35E0806")

}

func TestAll(t *testing.T) {
	cli, _ := newCLI()

	cli.SetPassword("test")

	address := run(cli, "account new")

	fmt.Println(address)

}
