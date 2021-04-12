package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// GetCmd is root for various `get ...` commands
	GetCmd = &cobra.Command{
		Use:   "get",
		Short: "Get resources",
		Run:   ShowUsage,
	}
)

func init() {
	krokCmd.AddCommand(GetCmd)
}
