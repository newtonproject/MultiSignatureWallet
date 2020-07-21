package cli

import "testing"

func TestSign(t *testing.T) {
	cli := NewCLI()

	cli.TestCommand("sign")
}
