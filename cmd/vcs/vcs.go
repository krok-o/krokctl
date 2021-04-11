package vcs

import (
	"fmt"

	"github.com/krok-o/krok/pkg/models"
	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
)

var (
	// VcsCmd creates a token for a platform
	VcsCmd = &cobra.Command{
		Use:   "vcs",
		Short: "Create token for a platform like Github, Gitlab...",
		Run:   runVcsCmd,
	}
	vcsArgs struct {
		token string
		vcs   int
	}
)

func init() {
	cmd.CreateCmd.AddCommand(VcsCmd)

	f := VcsCmd.PersistentFlags()
	f.StringVar(&vcsArgs.token, "token", "", "Token for a version control system.")
	f.IntVar(&vcsArgs.vcs, "vcs", 1, "Version control system. Please refer to krok documentation to find out what is supported. 1 = Github.")
}

func runVcsCmd(c *cobra.Command, args []string) {
	err := cmd.KC.VcsClient.Create(&models.VCSToken{
		Token: vcsArgs.token,
		VCS:   vcsArgs.vcs,
	})
	if err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to create vcs token.")
	}

	fmt.Println("Success!")
}
