package settings

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
)

var (
	// DeleteSettingCmd deletes a setting with the given id.
	DeleteSettingCmd = &cobra.Command{
		Use:   "setting",
		Short: "Delete setting",
		Run:   runDeleteSettingCmd,
	}
	deleteSettingArgs struct {
		id int
	}
)

func init() {
	cmd.DeleteCmd.AddCommand(DeleteSettingCmd)

	f := DeleteSettingCmd.PersistentFlags()
	f.IntVar(&deleteSettingArgs.id, "id", -1, "ID of the setting to delete.")
}

func runDeleteSettingCmd(c *cobra.Command, args []string) {
	if deleteSettingArgs.id < 0 {
		cmd.CLILog.Error().Int("id", deleteSettingArgs.id).Msg("Please provide a valid ID.")
	}
	if err := cmd.KC.SettingsClient.Delete(deleteSettingArgs.id); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to delete setting.")
	}

	fmt.Println("Success!")
}
