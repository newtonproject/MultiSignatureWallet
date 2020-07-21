package cli

import "testing"

func TestUpdate(t *testing.T) {
	cli := NewCLI()

	cli.TestCommand("update dailylimit 1024")
	cli.TestCommand("update dailylimit 10 -u Wei")
	cli.TestCommand("update required 2")
}
