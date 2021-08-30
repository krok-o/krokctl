package repositories

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/krok-o/krok/pkg/models"

	"github.com/krok-o/krokctl/cmd"
	"github.com/krok-o/krokctl/pkg/formatter"
)

var (
	// CreateCommandCmd create a command.
	CreateCommandCmd = &cobra.Command{
		Use:   "command",
		Short: "Create command",
		Run:   runCreateCommandCmd,
	}
	createCommandArgs struct {
		name     string
		image    string
		schedule string
	}
)

func init() {
	cmd.CreateCmd.AddCommand(CreateCommandCmd)

	f := CreateCommandCmd.PersistentFlags()
	f.StringVar(&createCommandArgs.name, "name", "", "Name of the new command. Must be unique.")
	f.StringVar(&createCommandArgs.schedule, "schedule", "", "Schedule when to run this command. Must follow cronjob syntax.")
	f.StringVar(&createCommandArgs.image, "image", "", "Image of the command.")

	if err := CreateCommandCmd.MarkPersistentFlagRequired("name"); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to mark required flag.")
	}
	if err := CreateCommandCmd.MarkPersistentFlagRequired("image"); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to mark required flag.")
	}
}

func runCreateCommandCmd(c *cobra.Command, args []string) {
	newCommand := &models.Command{
		Name:     createCommandArgs.name,
		Image:    createCommandArgs.image,
		Schedule: createCommandArgs.schedule,
		Enabled:  true,
	}
	command, err := cmd.KC.CommandClient.Create(newCommand)
	if err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to create a command.")
	}
	fmt.Print(formatter.FormatCommand(command, cmd.KrokArgs.Formatter))
}
