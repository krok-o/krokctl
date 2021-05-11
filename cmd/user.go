package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// UserCmd is root for various `generate user ...` commands
	UserCmd = &cobra.Command{
		Use:   "user",
		Short: "User resources",
		Run:   ShowUsage,
	}
)

func init() {
	GenerateCmd.AddCommand(UserCmd)
}
