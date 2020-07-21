package cli

import "testing"

func TestInfo(t *testing.T) {
	cli := NewCLI()

	cli.TestCommand("info")
}
