package events

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
	"github.com/krok-o/krokctl/pkg/formatter"
)

var (
	// GetEventCmd get an event.
	GetEventCmd = &cobra.Command{
		Use:   "event",
		Short: "Get event",
		Run:   runGetEventCmd,
	}
	getEventArgs struct {
		id int
	}
)

func init() {
	cmd.GetCmd.AddCommand(GetEventCmd)

	f := GetEventCmd.PersistentFlags()
	f.IntVar(&getEventArgs.id, "id", 1, "ID of the event to get information for.")

	if err := GetEventCmd.MarkPersistentFlagRequired("id"); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to mark required flag.")
	}
}

func runGetEventCmd(c *cobra.Command, args []string) {
	event, err := cmd.KC.EventClient.Get(getEventArgs.id)
	if err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to get event.")
	}
	fmt.Print(formatter.FormatEvent(event, cmd.KrokArgs.Formatter))
}
