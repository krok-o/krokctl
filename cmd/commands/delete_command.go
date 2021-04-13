package repositories

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
)

var (
	// DeleteCommandCmd deletes a Command with the given id.
	DeleteCommandCmd = &cobra.Command{
		Use:   "command",
		Short: "Delete command",
		Run:   runDeleteCommandCmd,
	}
	deleteCommandArgs struct {
		id int
	}
)

func init() {
	cmd.DeleteCmd.AddCommand(DeleteCommandCmd)

	f := DeleteCommandCmd.PersistentFlags()
	f.IntVar(&deleteCommandArgs.id, "id", -1, "ID of the Command to delete.")
}

func runDeleteCommandCmd(c *cobra.Command, args []string) {
	if deleteCommandArgs.id < 0 {
		cmd.CLILog.Error().Int("id", deleteCommandArgs.id).Msg("Please provide a valid ID.")
	}
	if err := cmd.KC.CommandClient.Delete(deleteCommandArgs.id); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to delete Command.")
	}

	fmt.Println("Success!")
}
