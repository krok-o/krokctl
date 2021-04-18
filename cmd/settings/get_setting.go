package settings

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
	"github.com/krok-o/krokctl/pkg/formatter"
)

var (
	// GetSettingCmd get a setting.
	GetSettingCmd = &cobra.Command{
		Use:   "setting",
		Short: "Get setting",
		Run:   runGetSettingCmd,
	}
	getSettingArgs struct {
		id int
	}
)

func init() {
	cmd.GetCmd.AddCommand(GetSettingCmd)

	f := GetSettingCmd.PersistentFlags()
	f.IntVar(&getSettingArgs.id, "id", 1, "ID of the setting to get information for.")
}

func runGetSettingCmd(c *cobra.Command, args []string) {
	setting, err := cmd.KC.SettingsClient.Get(getSettingArgs.id)
	if err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to get setting.")
	}
	fmt.Print(formatter.FormatSetting(setting, cmd.KrokArgs.Formatter))
}
