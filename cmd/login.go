package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// LoginCmd authenticates the current user with the Krok server.
	LoginCmd = &cobra.Command{
		Use:   "login",
		Short: "Login",
		Run:   runLoginCmd,
	}
	loginArgs struct {
		token    string
		username string
	}
)

func init() {
	f := LoginCmd.PersistentFlags()
	f.StringVar(&loginArgs.token, "token", "", "The user token for remote managing of Krok server.")
	f.StringVar(&loginArgs.username, "username", "", "The username of the owner for the token.")
}

func runLoginCmd(cmd *cobra.Command, args []string) {

}
