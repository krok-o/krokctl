package cmd

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

const (
	// Prefix of all environment variables
	envKeyPrefix  = "KROK_"
	envKeyProgram = envKeyPrefix + "EXECUTABLE"
	apiPortSuffix = ":9998"
)

var (
	krokCmd = &cobra.Command{
		Use:              "krokctl",
		Short:            "Krok Control",
		Long:             "Krok, handle multiple repositories from different providers with ease.",
		Run:              ShowUsage,
		PersistentPreRun: krokCmdPersistentPreRun,
	}

	CLILog = zerolog.New(zerolog.ConsoleWriter{
		Out: os.Stderr,
	}).With().Timestamp().Logger()
	krokArgs struct {
		Token    string
		endpoint string
	}
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
	defaultEndpoint := getEnvOrDefault("ENDPOINT", "localhost")
	f.StringVar(&krokArgs.Token, "token", "", "Token used to authenticate with the server")
	f.StringVar(&krokArgs.endpoint, "endpoint", defaultEndpoint, "API endpoint of the Krok server")
}

// ShowUsage shows usage of the given command on stdout.
func ShowUsage(cmd *cobra.Command, args []string) {
	cmd.Usage()
}

// Execute runs the main krok command.
func Execute() error {
	return krokCmd.Execute()
}

// set up a token before the actual command is executed.
func krokCmdPersistentPreRun(cmd *cobra.Command, args []string) {
	if krokArgs.Token == "" {
		krokArgs.Token = getEnvOrDefault("TOKEN", "")
	}
}
