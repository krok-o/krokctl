package repositories

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
	"github.com/krok-o/krokctl/pkg/formatter"
)

var (
	// UploadCommandCmd upload a command.
	UploadCommandCmd = &cobra.Command{
		Use:   "command",
		Short: "Upload command",
		Run:   runUploadCommandCmd,
	}
	uploadCommandArgs struct {
		file string
	}
)

func init() {
	cmd.UploadCmd.AddCommand(UploadCommandCmd)

	f := UploadCommandCmd.PersistentFlags()
	f.StringVar(&uploadCommandArgs.file, "file", "", "Location of the file to upload")

	if err := UploadCommandCmd.MarkPersistentFlagRequired("file"); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to mark required flag.")
	}
}

func runUploadCommandCmd(c *cobra.Command, args []string) {
	command, err := cmd.KC.CommandClient.Upload(uploadCommandArgs.file)
	if err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to upload a command.")
	}
	fmt.Print(formatter.FormatCommand(command, cmd.KrokArgs.Formatter))
}
