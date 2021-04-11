package repositories

import (
	"fmt"

	"github.com/krok-o/krok/pkg/models"
	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
	"github.com/krok-o/krokctl/pkg/formatter"
)

var (
	// ListRepositoriesCmd lists all repositories.
	ListRepositoriesCmd = &cobra.Command{
		Use:   "repositories",
		Short: "List repositories",
		Run:   runListRepositoryCmd,
	}
	listRepoArgs struct {
		name string
		vcs  int
	}
)

func init() {
	cmd.ListCmd.AddCommand(ListRepositoriesCmd)

	f := ListRepositoriesCmd.PersistentFlags()
	f.StringVar(&listRepoArgs.name, "name", "", "List repositories with names that contain this name.")
	f.IntVar(&listRepoArgs.vcs, "vcs", 1, "List repositories which belong to a given vcs. Github = 1...")
}

func runListRepositoryCmd(c *cobra.Command, args []string) {
	opts := &models.ListOptions{
		Name: repoArgs.name,
		VCS:  repoArgs.vcs,
	}
	repos, err := cmd.KC.RepositoryClient.List(opts)
	if err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to list repository.")
	}
	fmt.Println(formatter.FormatRepositories(repos, cmd.KrokArgs.Formatter))
}
