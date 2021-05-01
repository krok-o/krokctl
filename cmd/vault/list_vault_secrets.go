package vault

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
	"github.com/krok-o/krokctl/pkg/formatter"
)

var (
	// ListVaultSecretsCmd lists all vaultSecrets.
	ListVaultSecretsCmd = &cobra.Command{
		Use:   "secrets",
		Short: "List vault secrets",
		Run:   runListVaultSecretsCmd,
	}
)

func init() {
	cmd.ListCmd.AddCommand(ListVaultSecretsCmd)
}

func runListVaultSecretsCmd(c *cobra.Command, args []string) {
	vaultSecrets, err := cmd.KC.VaultClient.List()
	if err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to list vault secrets.")
	}
	fmt.Print(formatter.FormatVaultSecrets(vaultSecrets, cmd.KrokArgs.Formatter))
}
