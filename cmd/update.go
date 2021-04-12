package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// UpdateCmd is root for various `update ...` commands
	UpdateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update resources",
		Run:   ShowUsage,
	}
)

func init() {
	krokCmd.AddCommand(UpdateCmd)
}
