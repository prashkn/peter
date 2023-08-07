package cmd

import (
	"fmt"
	"peter/resolver"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login username password",
	Args:  cobra.ExactArgs(2),
	Short: "Log into your password dashboard",
	Run: func(cmd *cobra.Command, args []string) {
		user, err := resolver.Login(args[0], args[1])
		if err != nil {
			cmd.Println("Error logging in. " + err.Error())
		} else {
			cmd.Println(fmt.Printf("Successfully logged in as %s", user.Name))
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
