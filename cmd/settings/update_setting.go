package settings

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/krok-o/krok/pkg/models"
	"github.com/krok-o/krokctl/cmd"
)

var (
	// UpdateSettingCmd updates a setting with the given values.
	UpdateSettingCmd = &cobra.Command{
		Use:   "setting",
		Short: "Update setting",
		Run:   runUpdateSettingCmd,
	}
	updateSettingArgs struct {
		id    int
		value string
	}
)

func init() {
	cmd.UpdateCmd.AddCommand(UpdateSettingCmd)

	f := UpdateSettingCmd.PersistentFlags()
	f.StringVar(&updateSettingArgs.value, "value", "", "The value of the setting.")
	f.IntVar(&updateSettingArgs.id, "id", -1, "ID of the setting to update the value for.")
	if err := UpdateSettingCmd.MarkPersistentFlagRequired("id"); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to mark required flag.")
	}
}

func runUpdateSettingCmd(c *cobra.Command, args []string) {
	setting := &models.CommandSetting{
		ID:    updateSettingArgs.id,
		Value: settingArgs.value,
	}
	if err := cmd.KC.SettingsClient.Update(setting); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to update setting.")
	}

	fmt.Println("Success!")
}
