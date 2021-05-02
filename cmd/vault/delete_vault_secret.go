package vault

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
)

var (
	// DeleteSecretCmd get a command.
	DeleteSecretCmd = &cobra.Command{
		Use:   "secret",
		Short: "Delete secret",
		Run:   runDeleteSecretCmd,
	}
	deleteSecretArgs struct {
		key string
	}
)

func init() {
	cmd.DeleteCmd.AddCommand(DeleteSecretCmd)

	f := DeleteSecretCmd.PersistentFlags()
	f.StringVar(&deleteSecretArgs.key, "key", "", "Key of the secret to delete.")

	if err := DeleteSecretCmd.MarkPersistentFlagRequired("key"); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to mark required flag.")
	}
}

func runDeleteSecretCmd(c *cobra.Command, args []string) {
	if err := cmd.KC.VaultClient.Delete(deleteSecretArgs.key); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to get secret.")
	}
	fmt.Println("Success!")
}
