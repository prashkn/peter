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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// storePasswordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// storePasswordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
