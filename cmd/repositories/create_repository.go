package repositories

import (
	"fmt"

	"github.com/krok-o/krok/pkg/models"
	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
	"github.com/krok-o/krokctl/pkg/formatter"
)

var (
	// RepositoryCmd creates a repository with the given values.
	RepositoryCmd = &cobra.Command{
		Use:   "repository",
		Short: "Create repository",
		Run:   runRepositoryCmd,
	}
	repoArgs struct {
		name   string
		url    string
		events []string
		auth   struct {
			ssh      string
			username string
			password string
			secret   string
		}
		vcs int
	}
)

func init() {
	cmd.CreateCmd.AddCommand(RepositoryCmd)

	f := RepositoryCmd.PersistentFlags()
	f.StringVar(&repoArgs.name, "name", "", "The name of the repository.")
	f.StringVar(&repoArgs.url, "url", "", "The URL of the repository.")
	f.StringSliceVar(&repoArgs.events, "events", []string{"push"}, "The events to subscribe to for this repository. Exp: push")
	f.StringVar(&repoArgs.auth.secret, "secret", "", "The hook secret.")
	f.StringVar(&repoArgs.auth.ssh, "ssh", "", "An SSH key to access the repository.")
	f.StringVar(&repoArgs.auth.username, "username", "", "A username to access the repository")
	f.StringVar(&repoArgs.auth.password, "password", "", "A password to access the repository.")
	f.IntVar(&repoArgs.vcs, "vcs", 1, "Version control system. Please refer to krok documentation to find out what is supported. 1 = Github.")
}

func runRepositoryCmd(c *cobra.Command, args []string) {
	cmd.CLILog.Debug().Msg("Creating repository...")
	repo := &models.Repository{
		Name: repoArgs.name,
		URL:  repoArgs.url,
		VCS:  repoArgs.vcs,
		Auth: &models.Auth{
			Secret:   repoArgs.auth.secret,
			SSH:      repoArgs.auth.ssh,
			Username: repoArgs.auth.username,
			Password: repoArgs.auth.password,
		},
		Events: repoArgs.events,
	}
	repo, err := cmd.KC.RepositoryClient.Create(repo)
	if err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to create repository.")
	}

	fmt.Println(formatter.FormatRepository(repo, cmd.KrokArgs.Formatter))
}
