package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ccs",
	Short: "CCS - Claude Code Switcher",
	Long:  `A CLI tool to manage and switch between Claude Code configuration profiles.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(useCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(uiCmd)
}
