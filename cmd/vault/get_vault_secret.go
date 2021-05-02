package vault

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
	"github.com/krok-o/krokctl/pkg/formatter"
)

var (
	// GetSecretCmd get a command.
	GetSecretCmd = &cobra.Command{
		Use:   "secret",
		Short: "Get secret",
		Run:   runGetSecretCmd,
	}
	getSecretArgs struct {
		name string
	}
)

func init() {
	cmd.GetCmd.AddCommand(GetSecretCmd)

	f := GetSecretCmd.PersistentFlags()
	f.StringVar(&getSecretArgs.name, "name", "", "Name of the secret to retrieve.")

	if err := GetSecretCmd.MarkPersistentFlagRequired("name"); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to mark required flag.")
	}
}

func runGetSecretCmd(c *cobra.Command, args []string) {
	command, err := cmd.KC.VaultClient.Get(getSecretArgs.name)
	if err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to get secret.")
	}
	fmt.Print(formatter.FormatVaultSecret(command, cmd.KrokArgs.Formatter))
}
