package runs

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
	"github.com/krok-o/krokctl/pkg/formatter"
)

var (
	// GetCommandRunCmd get a command run.
	GetCommandRunCmd = &cobra.Command{
		Use:   "run",
		Short: "Get a command run",
		Run:   runGetCommandRunCmd,
	}
	getCommandRunArgs struct {
		id int
	}
)

func init() {
	cmd.GetCmd.AddCommand(GetCommandRunCmd)

	f := GetCommandRunCmd.PersistentFlags()
	f.IntVar(&getCommandRunArgs.id, "id", 1, "ID of the command run to get.")

	if err := GetCommandRunCmd.MarkPersistentFlagRequired("id"); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to mark required flag.")
	}
}

func runGetCommandRunCmd(c *cobra.Command, args []string) {
	run, err := cmd.KC.CommandRunClient.Get(getCommandRunArgs.id)
	if err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to get a command run.")
	}
	fmt.Print(formatter.FormatCommandRun(run, cmd.KrokArgs.Formatter))
}
