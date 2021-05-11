package users

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
	"github.com/krok-o/krokctl/pkg/formatter"
)

var (
	// GetUserCmd get a user.
	GetUserCmd = &cobra.Command{
		Use:   "user",
		Short: "Get user",
		Run:   runGetUserCmd,
	}
	getUserArgs struct {
		id int
	}
)

func init() {
	cmd.GetCmd.AddCommand(GetUserCmd)

	f := GetUserCmd.PersistentFlags()
	f.IntVar(&getUserArgs.id, "id", 1, "ID of the user to get information for.")

	if err := GetUserCmd.MarkPersistentFlagRequired("id"); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to mark required flag.")
	}
}

func runGetUserCmd(c *cobra.Command, args []string) {
	user, err := cmd.KC.UserClient.Get(getUserArgs.id)
	if err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to get a user.")
	}
	fmt.Print(formatter.FormatUser(user, cmd.KrokArgs.Formatter))
}
