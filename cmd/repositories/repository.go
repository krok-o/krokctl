package repositories

import (
	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
)

var (
	// RepositoryCmd creates a repository with the given values.
	RepositoryCmd = &cobra.Command{
		Use:   "repository",
		Short: "Create repository",
		Run:   runRepositoryCmd,
	}
)

func init() {
	cmd.CreateCmd.AddCommand(RepositoryCmd)
}

func runRepositoryCmd(c *cobra.Command, args []string) {
	cmd.CLILog.Info().Msg("Creating repository...")

}
