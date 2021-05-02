package repositories

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
	"github.com/krok-o/krokctl/pkg/formatter"
)

var (
	// GetCommandCmd get a command.
	GetCommandCmd = &cobra.Command{
		Use:   "command",
		Short: "Get command",
		Run:   runGetCommandCmd,
	}
	getCommandArgs struct {
		id int
	}
)

func init() {
	cmd.GetCmd.AddCommand(GetCommandCmd)

	f := GetCommandCmd.PersistentFlags()
	f.IntVar(&getCommandArgs.id, "id", 1, "ID of the command to get information for.")

	if err := GetCommandCmd.MarkPersistentFlagRequired("id"); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to mark required flag.")
	}
}

func runGetCommandCmd(c *cobra.Command, args []string) {
	command, err := cmd.KC.CommandClient.Get(getCommandArgs.id)
	if err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to get a command.")
	}
	fmt.Print(formatter.FormatCommand(command, cmd.KrokArgs.Formatter))
}
