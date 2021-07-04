package auth

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
	"github.com/krok-o/krokctl/pkg/formatter"
)

var (
	// CreateApiKeyCmd create an api key.
	CreateApiKeyCmd = &cobra.Command{
		Use:   "apikey",
		Short: "Create an api key",
		Run:   runCreateApiKeyCmd,
	}
	createApiKeyArgs struct {
		name string
	}
)

func init() {
	cmd.CreateCmd.AddCommand(CreateApiKeyCmd)

	f := CreateApiKeyCmd.PersistentFlags()
	f.StringVar(&createApiKeyArgs.name, "name", "", "The name of the key.")
}

func runCreateApiKeyCmd(command *cobra.Command, args []string) {
	key, err := cmd.KC.ApiKeyClient.Create(createApiKeyArgs.name)
	if err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to create an api key.")
	}
	fmt.Print(formatter.FormatCreatedApiKey(key, cmd.KrokArgs.Formatter))
}
