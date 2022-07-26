package cmd

import (
	"github.com/spf13/cobra"
)

var (
	version = "unknown"
	rootCmd = &cobra.Command{
		Use:     "pure",
		Short:   "pure admin cli",
		Version: version,
	}
)

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(newCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
