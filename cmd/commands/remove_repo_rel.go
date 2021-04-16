package repositories

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
)

var (
	// RemoveRepoRelCmd remove a relationship to a repository.
	RemoveRepoRelCmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove command repository relationship",
		Run:   runRemoveRepoRelCmd,
	}
	removeRepoRelArgs struct {
		commandID int
		repoID    int
	}
)

func init() {
	cmd.CommandCmd.AddCommand(RemoveRepoRelCmd)

	f := RemoveRepoRelCmd.PersistentFlags()
	f.IntVar(&removeRepoRelArgs.commandID, "command-id", -1, "ID of the command to remove to repository.")
	f.IntVar(&removeRepoRelArgs.repoID, "repository-id", -1, "ID of the repository to remove the command to.")
}

func runRemoveRepoRelCmd(c *cobra.Command, args []string) {
	if removeRepoRelArgs.repoID < 0 {
		cmd.CLILog.Fatal().Msg("Please provide a valid repository ID.")
	}
	if removeRepoRelArgs.commandID < 0 {
		cmd.CLILog.Fatal().Msg("Please provide a valid command ID.")
	}
	if err := cmd.KC.CommandClient.RemoveRelationshipToRepository(removeRepoRelArgs.commandID, removeRepoRelArgs.repoID); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to remove relationship to repository.")
	}
	fmt.Println("Success!")
}
