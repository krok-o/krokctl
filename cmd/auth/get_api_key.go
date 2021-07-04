package auth

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
	"github.com/krok-o/krokctl/pkg/formatter"
)

var (
	// GetApiKeyCmd get an api key.
	GetApiKeyCmd = &cobra.Command{
		Use:   "apikey",
		Short: "Get api key",
		Run:   runGetApiKeyCmd,
	}
	getApiKeyCmdArgs struct {
		id int
	}
)

func init() {
	cmd.GetCmd.AddCommand(GetApiKeyCmd)

	f := GetApiKeyCmd.PersistentFlags()
	f.IntVar(&getApiKeyCmdArgs.id, "id", 1, "ID of the apikey to get information for.")

	if err := GetApiKeyCmd.MarkPersistentFlagRequired("id"); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to mark required flag.")
	}
}

func runGetApiKeyCmd(c *cobra.Command, args []string) {
	key, err := cmd.KC.ApiKeyClient.Get(getApiKeyCmdArgs.id)
	if err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to get api key.")
	}
	fmt.Print(formatter.FormatApiKey(key, cmd.KrokArgs.Formatter))
}
