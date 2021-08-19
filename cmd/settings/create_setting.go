package settings

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/krok-o/krok/pkg/models"

	"github.com/krok-o/krokctl/cmd"
	"github.com/krok-o/krokctl/pkg/formatter"
)

var (
	// CreateSettingCmd creates a setting with the given values.
	CreateSettingCmd = &cobra.Command{
		Use:   "setting",
		Short: "Create setting",
		Run:   runSettingCmd,
	}
	settingArgs struct {
		commandID int
		key       string
		value     string
		inVault   bool
	}
)

func init() {
	cmd.CreateCmd.AddCommand(CreateSettingCmd)

	f := CreateSettingCmd.PersistentFlags()
	f.StringVar(&settingArgs.key, "key", "", "The key of the setting.")
	f.StringVar(&settingArgs.value, "value", "", "The value of the setting.")
	f.IntVar(&settingArgs.commandID, "command-id", -1, "The id of the command.")
	f.BoolVar(&settingArgs.inVault, "in-vault", false, "Whether value is secret or not and should be placed in secure storage.")
}

func runSettingCmd(c *cobra.Command, args []string) {
	setting := &models.CommandSetting{
		InVault:   settingArgs.inVault,
		Key:       settingArgs.key,
		Value:     settingArgs.value,
		CommandID: settingArgs.commandID,
	}
	setting, err := cmd.KC.SettingsClient.Create(setting)
	if err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to create setting.")
	}

	fmt.Print(formatter.FormatSetting(setting, cmd.KrokArgs.Formatter))
}
