package platforms

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
	"github.com/krok-o/krokctl/pkg/formatter"
)

var (
	// ListPlatformsCmd lists all supported platforms.
	ListPlatformsCmd = &cobra.Command{
		Use:   "platforms",
		Short: "List all supported platforms",
		Run:   runListPlatformsCmd,
	}
)

func init() {
	cmd.ListCmd.AddCommand(ListPlatformsCmd)
}

func runListPlatformsCmd(c *cobra.Command, args []string) {
	platforms, err := cmd.KC.PlatformClient.List()
	if err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to list supported platforms.")
	}
	fmt.Print(formatter.FormatPlatforms(platforms, cmd.KrokArgs.Formatter))
}
