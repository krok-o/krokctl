package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// UploadCmd is root for various `upload ...` commands
	UploadCmd = &cobra.Command{
		Use:   "upload",
		Short: "Upload resources",
		Run:   ShowUsage,
	}
)

func init() {
	krokCmd.AddCommand(UploadCmd)
}
