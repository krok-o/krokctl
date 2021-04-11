package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/krok-o/krokctl/pkg"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

const (
	// Prefix of all environment variables
	envKeyPrefix = "KROK_"
	sessionFile  = ".session"
)

var (
	krokCmd = &cobra.Command{
		Use:              "krokctl",
		Short:            "Krok Control",
		Long:             "Krok, handle multiple repositories from different providers with ease.",
		Run:              ShowUsage,
		PersistentPreRun: krokCmdPersistentPreRun,
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

	// sessionFileLocation
	sessionFileLocation string
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
	f.StringVar(&KrokArgs.Token, "token", "", "Token used to authenticate with the server")
	f.StringVar(&KrokArgs.Formatter, "format", "table", "Format to display data in: json|table")
	f.StringVar(&KrokArgs.endpoint, "endpoint", defaultEndpoint, "API endpoint of the Krok server")

	// Set up the main client.
	KC = pkg.NewKrokClient(pkg.Config{
		Address: KrokArgs.endpoint,
	}, CLILog)

	// Set up session file location
	home, err := os.UserHomeDir()
	if err != nil {
		CLILog.Fatal().Err(err).Msg("Failed to get user home directory.")
	}
	sessionFileLocation = filepath.Join(home, ".config", "krokctl")
	if _, err := os.Stat(sessionFileLocation); err != nil && os.IsNotExist(err) {
		if err := os.MkdirAll(sessionFileLocation, os.ModeDir); err != nil {
			CLILog.Fatal().Err(err).Msg("Failed to create krokctl config folder.")
		}
	}
}

// ShowUsage shows usage of the given command on stdout.
func ShowUsage(cmd *cobra.Command, args []string) {
	_ = cmd.Usage()
}

// Execute runs the main krok command.
func Execute() error {
	return krokCmd.Execute()
}

// set up a token before the actual command is executed.
func krokCmdPersistentPreRun(cmd *cobra.Command, args []string) {
	if KrokArgs.Token == "" {
		var err error
		KrokArgs.Token, err = getTokenFromSessionFile()
		if err != nil {
			CLILog.Fatal().Err(err).Msg("Failed to locate session file. Please log in.")
		}
	}
}

func getTokenFromSessionFile() (string, error) {
	if _, err := os.Stat(filepath.Join(sessionFileLocation, sessionFile)); err != nil {
		return "", err
	}

	content, err := ioutil.ReadFile(filepath.Join(sessionFileLocation, sessionFile))
	if err != nil {
		return "", err
	}
	return string(content), nil
}
