package repositories

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
	"github.com/krok-o/krokctl/pkg/formatter"
)

var (
	// GetRepositoriesCmd get a repository.
	GetRepositoriesCmd = &cobra.Command{
		Use:   "repository",
		Short: "Get repository",
		Run:   runGetRepositoryCmd,
	}
	getRepoArgs struct {
		id int
	}
)

func init() {
	cmd.GetCmd.AddCommand(GetRepositoriesCmd)

	f := GetRepositoriesCmd.PersistentFlags()
	f.IntVar(&getRepoArgs.id, "id", 1, "ID of the repository to get information for.")
	if err := GetRepositoriesCmd.MarkPersistentFlagRequired("id"); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to mark required flag.")
	}
}

func runGetRepositoryCmd(c *cobra.Command, args []string) {
	repo, err := cmd.KC.RepositoryClient.Get(getRepoArgs.id)
	if err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to get repository.")
	}
	fmt.Print(formatter.FormatRepository(repo, cmd.KrokArgs.Formatter))
}
