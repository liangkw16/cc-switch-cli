package cmd

import (
	"os"

	"github.com/bytedance/ccs/internal/ui"
	"github.com/spf13/cobra"
)

var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "Launch interactive TUI",
	Long:  `Launch the terminal UI for interactive profile management.`,
	Run:   runUI,
}

func runUI(cmd *cobra.Command, args []string) {
	if err := ui.Run(); err != nil {
		os.Exit(1)
	}
}
