package auth

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
	"github.com/krok-o/krokctl/pkg/formatter"
)

var (
	// ListApiKeyCmd create an api key.
	ListApiKeyCmd = &cobra.Command{
		Use:   "apikeys",
		Short: "List api keys",
		Run:   runListApiKeyCmd,
	}
)

func init() {
	cmd.ListCmd.AddCommand(ListApiKeyCmd)
}

func runListApiKeyCmd(command *cobra.Command, args []string) {
	key, err := cmd.KC.ApiKeyClient.List()
	if err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to create an api key.")
	}
	fmt.Print(formatter.FormatApiKeys(key, cmd.KrokArgs.Formatter))
}
