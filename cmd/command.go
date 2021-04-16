package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// CommandCmd is root for various `relationship command...` commands
	CommandCmd = &cobra.Command{
		Use:   "command",
		Short: "Command resources",
		Run:   ShowUsage,
	}
)

func init() {
	RelationshipCmd.AddCommand(CommandCmd)
}
