package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "faunamart",
	Short: "Lottery app for Fauna Mart",
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	cobra.EnableCommandSorting = false

}

func Execute() error {
	return rootCmd.Execute()
}
