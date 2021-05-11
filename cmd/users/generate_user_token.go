package users

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
)

var (
	// GenerateTokenUserCmd create a new token for a user.
	GenerateTokenUserCmd = &cobra.Command{
		Use:   "token",
		Short: "Generate a new api token for a user.",
		Run:   runGenerateTokenUserCmd,
	}
)

func init() {
	cmd.UserCmd.AddCommand(GenerateTokenUserCmd)
}

func runGenerateTokenUserCmd(c *cobra.Command, args []string) {
	token, err := cmd.KC.UserClient.Generate()
	if err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to reset token.")
	}
	fmt.Print(token["token"])
}
