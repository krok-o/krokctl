package users

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
)

var (
	// DeleteUserCmd deletes a User with the given id.
	DeleteUserCmd = &cobra.Command{
		Use:   "user",
		Short: "Delete user",
		Run:   runDeleteUserCmd,
	}
	deleteUserArgs struct {
		id int
	}
)

func init() {
	cmd.DeleteCmd.AddCommand(DeleteUserCmd)

	f := DeleteUserCmd.PersistentFlags()
	f.IntVar(&deleteUserArgs.id, "id", -1, "ID of the User to delete.")

	if err := DeleteUserCmd.MarkPersistentFlagRequired("id"); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to mark required flag.")
	}
}

func runDeleteUserCmd(c *cobra.Command, args []string) {
	if deleteUserArgs.id < 0 {
		cmd.CLILog.Error().Int("id", deleteUserArgs.id).Msg("Please provide a valid ID.")
	}
	if err := cmd.KC.UserClient.Delete(deleteUserArgs.id); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to delete User.")
	}

	fmt.Println("Success!")
}
