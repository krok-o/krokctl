package vault

import (
	"fmt"

	"github.com/krok-o/krok/pkg/models"
	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
)

var (
	// UpdateSecretCmd get a vault secret.
	UpdateSecretCmd = &cobra.Command{
		Use:   "secret",
		Short: "Create secret",
		Run:   runUpdateSecretCmd,
	}
	updateSecretArgs struct {
		key   string
		value string
	}
)

func init() {
	cmd.UpdateCmd.AddCommand(UpdateSecretCmd)

	f := UpdateSecretCmd.PersistentFlags()
	f.StringVar(&updateSecretArgs.key, "key", "", "Key of the secret.")
	f.StringVar(&updateSecretArgs.value, "value", "", "Value of the secret.")

	if err := UpdateSecretCmd.MarkPersistentFlagRequired("key"); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to mark required flag.")
	}
	if err := UpdateSecretCmd.MarkPersistentFlagRequired("value"); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to mark required flag.")
	}
}

func runUpdateSecretCmd(c *cobra.Command, args []string) {
	if err := cmd.KC.VaultClient.Update(&models.VaultSetting{Key: updateSecretArgs.key, Value: updateSecretArgs.value}); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to get secret.")
	}
	fmt.Println("Success!")
}
