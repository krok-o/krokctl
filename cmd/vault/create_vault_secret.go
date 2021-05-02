package vault

import (
	"fmt"

	"github.com/krok-o/krok/pkg/models"
	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
)

var (
	// CreateSecretCmd get a vault secret.
	CreateSecretCmd = &cobra.Command{
		Use:   "secret",
		Short: "Create secret",
		Run:   runCreateSecretCmd,
	}
	createSecretArgs struct {
		key   string
		value string
	}
)

func init() {
	cmd.CreateCmd.AddCommand(CreateSecretCmd)

	f := CreateSecretCmd.PersistentFlags()
	f.StringVar(&createSecretArgs.key, "key", "", "Key of the secret.")
	f.StringVar(&createSecretArgs.value, "value", "", "Value of the secret.")

	if err := CreateSecretCmd.MarkPersistentFlagRequired("key"); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to mark required flag.")
	}
	if err := CreateSecretCmd.MarkPersistentFlagRequired("value"); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to mark required flag.")
	}
}

func runCreateSecretCmd(c *cobra.Command, args []string) {
	if err := cmd.KC.VaultClient.Create(&models.VaultSetting{Key: createSecretArgs.key, Value: createSecretArgs.value}); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to get secret.")
	}
	fmt.Println("Success!")
}
