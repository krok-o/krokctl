package cmd

import (
	"os"

	"github.com/krok-o/krokctl/pkg"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

const (
	// Prefix of all environment variables
	envKeyPrefix = "KROK_"
)

var (
	krokCmd = &cobra.Command{
		Use:   "krokctl",
		Short: "Krok Control",
		Long:  "Krok, handle multiple repositories from different providers with ease.",
		Run:   ShowUsage,
	}

	// CLILog is the krokctl's logger.
	CLILog = zerolog.New(zerolog.ConsoleWriter{
		Out: os.Stderr,
	}).With().Timestamp().Logger()
	// KrokArgs define the root arguments for all commands.
	KrokArgs struct {
		Token     string
		Formatter string
		endpoint  string
	}
	// KC is the krok client with all the clients bundled together.
	KC *pkg.KrokClient
)

func getEnvOrDefault(key, def string) string {
	if v := os.Getenv(envKeyPrefix + key); v != "" {
		return v
	}
	return def
}

func init() {
	f := krokCmd.PersistentFlags()
	// Persistent flags
	defaultEndpoint := getEnvOrDefault("ENDPOINT", "http://localhost:9998")
	token := getEnvOrDefault("TOKEN", "")
	f.StringVar(&KrokArgs.Token, "token", token, "Token used to authenticate with the server")
	f.StringVar(&KrokArgs.Formatter, "format", "table", "Format to display data in: json|table")
	f.StringVar(&KrokArgs.endpoint, "endpoint", defaultEndpoint, "API endpoint of the Krok server")

	if token == "" {
		CLILog.Fatal().Msg("Token is empty. Please either provide one with --token or use KROK_TOKEN environment property.")
	}
	// Set up the main client.
	KC = pkg.NewKrokClient(pkg.Config{
		Address: KrokArgs.endpoint,
		Token:   KrokArgs.Token,
	}, CLILog)
}

// ShowUsage shows usage of the given command on stdout.
func ShowUsage(cmd *cobra.Command, args []string) {
	_ = cmd.Usage()
}

// Execute runs the main krok command.
func Execute() error {
	return krokCmd.Execute()
}
