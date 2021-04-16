package repositories

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
)

var (
	// AddRepoRelCmd add a relationship to a repository.
	AddRepoRelCmd = &cobra.Command{
		Use:   "add",
		Short: "Add command repository relationship",
		Run:   runAddRepoRelCmd,
	}
	addRepoRelArgs struct {
		commandID int
		repoID    int
	}
)

func init() {
	cmd.CommandCmd.AddCommand(AddRepoRelCmd)

	f := AddRepoRelCmd.PersistentFlags()
	f.IntVar(&addRepoRelArgs.commandID, "command-id", -1, "ID of the command to add to repository.")
	f.IntVar(&addRepoRelArgs.repoID, "repository-id", -1, "ID of the repository to add the command to.")
}

func runAddRepoRelCmd(c *cobra.Command, args []string) {
	if addRepoRelArgs.repoID < 0 {
		cmd.CLILog.Fatal().Msg("Please provide a valid repository ID.")
	}
	if addRepoRelArgs.commandID < 0 {
		cmd.CLILog.Fatal().Msg("Please provide a valid command ID.")
	}
	if err := cmd.KC.CommandClient.AddRelationshipToRepository(addRepoRelArgs.commandID, addRepoRelArgs.repoID); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to add relationship to repository.")
	}
	fmt.Println("Success!")
}
