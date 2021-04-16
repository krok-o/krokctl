package repositories

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
)

var (
	// AddPlatformRelCmd add a relationship to a repository.
	AddPlatformRelCmd = &cobra.Command{
		Use:   "add",
		Short: "Add command platform relationship",
		Run:   runAddPlatformRelCmd,
	}
	addPlatformRelArgs struct {
		commandID  int
		platformID int
	}
)

func init() {
	cmd.PlatformCmd.AddCommand(AddPlatformRelCmd)

	f := AddPlatformRelCmd.PersistentFlags()
	f.IntVar(&addPlatformRelArgs.commandID, "command-id", -1, "ID of the command to add to platform.")
	f.IntVar(&addPlatformRelArgs.platformID, "platform-id", -1, "ID of the platform to add the command to.")
}

func runAddPlatformRelCmd(c *cobra.Command, args []string) {
	if addPlatformRelArgs.platformID < 0 {
		cmd.CLILog.Fatal().Msg("Please provide a valid platform ID.")
	}
	if addPlatformRelArgs.commandID < 0 {
		cmd.CLILog.Fatal().Msg("Please provide a valid command ID.")
	}
	if err := cmd.KC.CommandClient.AddRelationshipToPlatform(addPlatformRelArgs.commandID, addPlatformRelArgs.platformID); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to add relationship to platform.")
	}
	fmt.Println("Success!")
}
