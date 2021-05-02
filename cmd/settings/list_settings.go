package settings

import (
	"fmt"

	"github.com/krok-o/krokctl/pkg/formatter"
	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
)

var (
	// ListSettingsCmd creates a setting with the given values.
	ListSettingsCmd = &cobra.Command{
		Use:   "settings",
		Short: "List settings",
		Run:   runListSettingsCmd,
	}
	listSettingArgs struct {
		commandID int
	}
)

func init() {
	cmd.ListCmd.AddCommand(ListSettingsCmd)

	f := ListSettingsCmd.PersistentFlags()
	f.IntVar(&listSettingArgs.commandID, "command-id", -1, "The id of the command.")
	if err := ListSettingsCmd.MarkPersistentFlagRequired("command-id"); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to mark required flag.")
	}
}

func runListSettingsCmd(c *cobra.Command, args []string) {
	settings, err := cmd.KC.SettingsClient.List(listSettingArgs.commandID)
	if err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to list setting.")
	}

	fmt.Print(formatter.FormatSettings(settings, cmd.KrokArgs.Formatter))
}
