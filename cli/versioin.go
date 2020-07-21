package cli

import (
	"github.com/spf13/cobra"
)

func (cli *CLI) buildVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Get version of " + cli.Name + " CLI",
		Run: func(cmd *cobra.Command, args []string) {
			showSuccess(cli.version)
		},
	}

	return cmd
}
