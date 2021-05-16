package users

import (
	"fmt"

	"github.com/krok-o/krok/pkg/models"
	"github.com/spf13/cobra"

	"github.com/krok-o/krokctl/cmd"
	"github.com/krok-o/krokctl/pkg/formatter"
)

var (
	// CreateUserCmd create a user.
	CreateUserCmd = &cobra.Command{
		Use:   "user",
		Short: "Create user",
		Run:   runCreateUserCmd,
	}
	createUserArgs struct {
		displayName string
		email       string
	}
)

func init() {
	cmd.CreateCmd.AddCommand(CreateUserCmd)

	f := CreateUserCmd.PersistentFlags()
	f.StringVar(&createUserArgs.displayName, "display-name", "", "Display name of the user.")
	f.StringVar(&createUserArgs.email, "email", "", "Email of the user.")

	if err := CreateUserCmd.MarkPersistentFlagRequired("email"); err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to mark required flag.")
	}
}

func runCreateUserCmd(c *cobra.Command, args []string) {
	user, err := cmd.KC.UserClient.Create(&models.User{
		DisplayName: createUserArgs.displayName,
		Email:       createUserArgs.email,
	})
	if err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to create a user.")
	}
	fmt.Print(formatter.FormatNewUser(user, cmd.KrokArgs.Formatter))
}
