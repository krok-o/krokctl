package users

import (
	"fmt"

	"github.com/krok-o/krok/pkg/models"
	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
	"github.com/krok-o/krokctl/pkg/formatter"
)

var (
	// UpdateUserCmd update a user.
	UpdateUserCmd = &cobra.Command{
		Use:   "user",
		Short: "Update user",
		Run:   runUpdateUserCmd,
	}
	updateUserArgs struct {
		id          int
		displayName string
	}
)

func init() {
	cmd.UpdateCmd.AddCommand(UpdateUserCmd)

	f := UpdateUserCmd.PersistentFlags()
	f.StringVar(&updateUserArgs.displayName, "display-name", "", "Display name of the user.")
	f.IntVar(&updateUserArgs.id, "id", -1, "ID of the user you would like to update.")
}

func runUpdateUserCmd(c *cobra.Command, args []string) {
	user, err := cmd.KC.UserClient.Update(&models.User{
		ID:          updateUserArgs.id,
		DisplayName: updateUserArgs.displayName,
	})
	if err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to update a user.")
	}
	fmt.Print(formatter.FormatUser(user, cmd.KrokArgs.Formatter))
}
