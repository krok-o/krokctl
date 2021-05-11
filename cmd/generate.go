package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// GenerateCmd is root for various `generate ...` commands
	GenerateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generate resources",
		Run:   ShowUsage,
	}
)

func init() {
	krokCmd.AddCommand(GenerateCmd)
}
