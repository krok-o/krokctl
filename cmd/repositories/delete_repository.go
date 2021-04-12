package repositories

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
)

var (
	// DeleteRepositoryCmd deletes a repository with the given id.
	DeleteRepositoryCmd = &cobra.Command{
		Use:   "repository",
		Short: "Delete repository",
		Run:   runDeleteRepositoryCmd,
	}
	deleteRepoArgs struct {
		id int
	}
)

func init() {
	cmd.DeleteCmd.AddCommand(DeleteRepositoryCmd)

	f := DeleteRepositoryCmd.PersistentFlags()
	f.IntVar(&deleteRepoArgs.id, "id", -1, "ID of the repository to delete.")
}

func runDeleteRepositoryCmd(c *cobra.Command, args []string) {
	if deleteRepoArgs.id < 0 {
		cmd.CLILog.Error().Int("id", deleteRepoArgs.id).Msg("Please provide a valid ID.")
	}
	if err := cmd.KC.RepositoryClient.Delete(deleteRepoArgs.id); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to delete repository.")
	}

	fmt.Println("Success!")
}
