package users

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
	"github.com/krok-o/krokctl/pkg/formatter"
)

var (
	// ListUserCmd lists users.
	ListUserCmd = &cobra.Command{
		Use:   "users",
		Short: "List users",
		Run:   runListUsersCmd,
	}
)

func init() {
	cmd.ListCmd.AddCommand(ListUserCmd)
}

func runListUsersCmd(c *cobra.Command, args []string) {
	users, err := cmd.KC.UserClient.List()
	if err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to list users.")
	}
	fmt.Print(formatter.FormatUsers(users, cmd.KrokArgs.Formatter))
}
