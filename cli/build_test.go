package cli

import "testing"

func TestBuild(t *testing.T) {
	cli := NewCLI()

	cli.TestCommand("build")
}
