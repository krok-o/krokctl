package repositories

import (
	"fmt"

	"github.com/krok-o/krok/pkg/models"
	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
	"github.com/krok-o/krokctl/pkg/formatter"
)

var (
	// UpdateCommandCmd creates a command with the given values.
	UpdateCommandCmd = &cobra.Command{
		Use:   "command",
		Short: "Update command",
		Run:   runUpdateCommandCmd,
	}
	updateCommandArgs struct {
		name     string
		schedule string
		location string
		hash     string
		filename string
		enabled  bool
		id       int
	}
)

func init() {
	cmd.UpdateCmd.AddCommand(UpdateCommandCmd)

	f := UpdateCommandCmd.PersistentFlags()
	f.StringVar(&updateCommandArgs.name, "name", "", "The name of the command.")
	f.StringVar(&updateCommandArgs.schedule, "schedule", "", "The schedule of the command.")
	f.BoolVar(&updateCommandArgs.enabled, "enabled", true, "Enable / Disable command.")
	f.IntVar(&updateCommandArgs.id, "id", -1, "The ID of the command to update.")
}

func runUpdateCommandCmd(c *cobra.Command, args []string) {
	cmd.CLILog.Debug().Msg("Creating command...")

	f := c.Flags()
	command := &models.Command{}
	hasChanged := false
	if f.Changed("name") {
		command.Name = updateCommandArgs.name
		hasChanged = true
	}
	if f.Changed("schedule") {
		command.Schedule = updateCommandArgs.schedule
		hasChanged = true
	}
	// TODO: Change this once we change enabled to a pointer.
	command.Enabled = updateCommandArgs.enabled

	if hasChanged {
		updated, err := cmd.KC.CommandClient.Update(command)
		if err != nil {
			cmd.CLILog.Fatal().Err(err).Msg("Failed to update command.")
		}
		fmt.Print(formatter.FormatCommand(updated, cmd.KrokArgs.Formatter))
	}

}
