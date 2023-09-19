package cmd

import (
	"peter/resolver"

	"github.com/spf13/cobra"
)

// signupCmd represents the signup command
var signupCmd = &cobra.Command{
	Use:   "signup emailAddress password",
	Args:  cobra.ExactArgs(2),
	Short: "signup for peter",
	Run: func(cmd *cobra.Command, args []string) {
		err := resolver.Signup(args[0], args[1])
		if err != nil {
			cmd.Println("Error signing up. " + err.Error())
		} else {
			cmd.Println("Successfully signed up!")
		}
	},
}

func init() {
	rootCmd.AddCommand(signupCmd)
}
