package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// ListCmd is root for various `list ...` commands
	ListCmd = &cobra.Command{
		Use:   "list",
		Short: "List resources",
		Run:   ShowUsage,
	}
)

func init() {
	krokCmd.AddCommand(ListCmd)
}
