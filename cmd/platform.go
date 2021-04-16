package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// PlatformCmd is root for various `relationship platform...` commands
	PlatformCmd = &cobra.Command{
		Use:   "platform",
		Short: "Platform resources",
		Run:   ShowUsage,
	}
)

func init() {
	RelationshipCmd.AddCommand(PlatformCmd)
}
