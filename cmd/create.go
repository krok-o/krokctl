package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// CreateCmd is root for various `create ...` commands
	CreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create resources",
		Run:   ShowUsage,
	}
)

func init() {
	krokCmd.AddCommand(CreateCmd)
}
