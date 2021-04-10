package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "remote",
	Short: "A parser for mRemoteNG configs",
	Long: "'remote' is an easily extendable CLI tool to access mRemoteNG configurations " +
		"from non-Windows operating systems.",
}

func init() {
	rootCmd.AddCommand(startCmd)
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}
