package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// DeleteCmd is root for various `delete ...` commands
	DeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete resources",
		Run:   ShowUsage,
	}
)

func init() {
	krokCmd.AddCommand(DeleteCmd)
}
