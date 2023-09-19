/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// storePasswordCmd represents the storePassword command
var storePasswordCmd = &cobra.Command{
	Use:   "storePassword name/website username password",
	Short: "stores your password to anything",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(storePasswordCmd)
}
