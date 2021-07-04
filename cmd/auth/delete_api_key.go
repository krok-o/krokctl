package auth

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
)

var (
	// DeleteApiKeyCmd deletes a Command with the given id.
	DeleteApiKeyCmd = &cobra.Command{
		Use:   "apikey",
		Short: "Delete api key",
		Run:   runDeleteCommandCmd,
	}
	deleteApiKeyCmdArgs struct {
		id int
	}
)

func init() {
	cmd.DeleteCmd.AddCommand(DeleteApiKeyCmd)

	f := DeleteApiKeyCmd.PersistentFlags()
	f.IntVar(&deleteApiKeyCmdArgs.id, "id", -1, "ID of the api key to delete.")

	if err := DeleteApiKeyCmd.MarkPersistentFlagRequired("id"); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to mark required flag.")
	}
}

func runDeleteCommandCmd(c *cobra.Command, args []string) {
	if deleteApiKeyCmdArgs.id < 0 {
		cmd.CLILog.Error().Int("id", deleteApiKeyCmdArgs.id).Msg("Please provide a valid ID.")
	}
	if err := cmd.KC.ApiKeyClient.Delete(deleteApiKeyCmdArgs.id); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to delete api key.")
	}

	fmt.Println("Success!")
}
