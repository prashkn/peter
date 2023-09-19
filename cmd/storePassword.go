package cmd

import (
	"peter/resolver"

	"github.com/spf13/cobra"
)

// storePasswordCmd represents the storePassword command
var storePasswordCmd = &cobra.Command{
	Use:   "storePassword name/website username password",
	Args:  cobra.ExactArgs(3),
	Short: "stores your password to anything",
	Run: func(cmd *cobra.Command, args []string) {
		err := resolver.InsertPassword(args[0], args[1], args[2])
		if err != nil {
			cmd.Println("Error signing up. " + err.Error())
		} else {
			cmd.Println("Successfully signed up!")
		}
	},
}

func init() {
	rootCmd.AddCommand(storePasswordCmd)
}
