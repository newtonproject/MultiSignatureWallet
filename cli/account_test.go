package cli

import "testing"

func TestAccount(t *testing.T) {
	cli := NewCLI()

	cli.TestCommand("account new")
	cli.TestCommand("account new -n 10")
	cli.TestCommand("account new -n 10 --faucet")

	cli.TestCommand("account list")

}
