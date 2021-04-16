package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// RelationshipCmd is root for various `relationship ...` commands
	RelationshipCmd = &cobra.Command{
		Use:   "relationship",
		Short: "Relationship resources",
		Run:   ShowUsage,
	}
)

func init() {
	krokCmd.AddCommand(RelationshipCmd)
}
