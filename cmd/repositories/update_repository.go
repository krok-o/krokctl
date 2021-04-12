package repositories

import (
	"fmt"

	"github.com/krok-o/krok/pkg/models"
	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
	"github.com/krok-o/krokctl/pkg/formatter"
)

var (
	// UpdateRepositoryCmd creates a repository with the given values.
	UpdateRepositoryCmd = &cobra.Command{
		Use:   "repository",
		Short: "Update repository",
		Run:   runUpdateRepositoryCmd,
	}
	updateRepoArgs struct {
		name string
		id   int
	}
)

func init() {
	cmd.UpdateCmd.AddCommand(UpdateRepositoryCmd)

	f := UpdateRepositoryCmd.PersistentFlags()
	f.StringVar(&updateRepoArgs.name, "name", "", "The name of the repository.")
	f.IntVar(&updateRepoArgs.id, "id", -1, "The ID of the repository to update.")
}

func runUpdateRepositoryCmd(c *cobra.Command, args []string) {
	cmd.CLILog.Debug().Msg("Creating repository...")
	repo := &models.Repository{
		Name: updateRepoArgs.name,
		ID:   updateRepoArgs.id,
	}
	repo, err := cmd.KC.RepositoryClient.Update(repo)
	if err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to update repository.")
	}

	fmt.Print(formatter.FormatRepository(repo, cmd.KrokArgs.Formatter))
}
