package cmd

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/pkg"
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
		APIKeyID     string
		APIKeySecret string
		Email        string
		Formatter    string
		endpoint     string
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
	apiKeyID := getEnvOrDefault("API_KEY_ID", "")
	apiKeySecret := getEnvOrDefault("API_KEY_SECRET", "")
	email := getEnvOrDefault("EMAIL", "")
	f.StringVar(&KrokArgs.APIKeyID, "api-key-id", apiKeyID, "ID of the api key to use for authenticating with Krok.")
	f.StringVar(&KrokArgs.APIKeySecret, "api-key-secret", apiKeySecret, "Secret of the api key to use for authenticating with Krok.")
	f.StringVar(&KrokArgs.Email, "email", email, "Email of the user.")
	f.StringVar(&KrokArgs.Formatter, "format", "table", "Format to display data in: json|table")
	f.StringVar(&KrokArgs.endpoint, "endpoint", defaultEndpoint, "API endpoint of the Krok server")

	// Set up the main client.
	KC = pkg.NewKrokClient(pkg.Config{
		Address:      KrokArgs.endpoint,
		APIKeyID:     KrokArgs.APIKeyID,
		APIKeySecret: KrokArgs.APIKeySecret,
		Email:        KrokArgs.Email,
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
