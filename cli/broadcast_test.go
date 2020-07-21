package cli

import "testing"

func TestBroadcast(t *testing.T) {
	cli := NewCLI()

	cli.TestCommand("broadcast")
}
