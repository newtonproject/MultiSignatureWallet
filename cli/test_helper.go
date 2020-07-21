package cli

import (
	"strings"
	"testing"
)

func expectOutput(t *testing.T, cli *CLI, want string, command string) {
	got := cli.TestCommand(command)

	if strings.TrimSpace(got) != want {
		t.Errorf("(%s) wrong output: want %v, got %v", command, want, got)
	}
}

func newTestCLI() *CLI {
	cli := NewCLI()

	return cli
}

// TestCLIVersion test cli version
func TestCLIVersion(t *testing.T) {
	cli := newTestCLI()
	expectOutput(t, cli, cli.version, "version")
}
