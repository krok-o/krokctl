package repositories

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
)

var (
	// RemovePlatformRelCmd removes a relationship to a repository.
	RemovePlatformRelCmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove command platform relationship",
		Run:   runRemovePlatformRelCmd,
	}
	removePlatformRelArgs struct {
		commandID  int
		platformID int
	}
)

func init() {
	cmd.PlatformCmd.AddCommand(RemovePlatformRelCmd)

	f := RemovePlatformRelCmd.PersistentFlags()
	f.IntVar(&removePlatformRelArgs.commandID, "command-id", -1, "ID of the command to add to platform.")
	f.IntVar(&removePlatformRelArgs.platformID, "platform-id", -1, "ID of the platform to add the command to.")

	if err := RemovePlatformRelCmd.MarkPersistentFlagRequired("command-id"); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to mark required flag.")
	}
	if err := RemovePlatformRelCmd.MarkPersistentFlagRequired("platform-id"); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to mark required flag.")
	}
}

func runRemovePlatformRelCmd(c *cobra.Command, args []string) {
	if removePlatformRelArgs.platformID < 0 {
		cmd.CLILog.Fatal().Msg("Please provide a valid platform ID.")
	}
	if removePlatformRelArgs.commandID < 0 {
		cmd.CLILog.Fatal().Msg("Please provide a valid command ID.")
	}
	if err := cmd.KC.CommandClient.RemoveRelationshipToPlatform(removePlatformRelArgs.commandID, removePlatformRelArgs.platformID); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to remove relationship to platform.")
	}
	fmt.Println("Success!")
}
