package repositories

import (
	"fmt"

	"github.com/krok-o/krok/pkg/models"
	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
	"github.com/krok-o/krokctl/pkg/formatter"
)

var (
	// ListCommandsCmd lists all commands.
	ListCommandsCmd = &cobra.Command{
		Use:   "repositories",
		Short: "List repositories",
		Run:   runListRepositoryCmd,
	}
	listCommandArgs struct {
		name string
	}
)

func init() {
	cmd.ListCmd.AddCommand(ListCommandsCmd)

	f := ListCommandsCmd.PersistentFlags()
	f.StringVar(&listCommandArgs.name, "name", "", "List repositories with names that contain this name.")
}

func runListRepositoryCmd(c *cobra.Command, args []string) {
	opts := &models.ListOptions{
		Name: listCommandArgs.name,
	}
	commands, err := cmd.KC.CommandClient.List(opts)
	if err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to list commands.")
	}
	fmt.Print(formatter.FormatCommands(commands, cmd.KrokArgs.Formatter))
}
