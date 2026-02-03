package cmd

import (
	"fmt"
	"os"

	"github.com/bytedance/ccs/internal/config"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "rm <name>",
	Short: "Remove a profile",
	Long:  `Remove a Claude Code configuration profile.`,
	Args:  cobra.ExactArgs(1),
	Run:   runRemove,
}

func runRemove(cmd *cobra.Command, args []string) {
	name := args[0]

	store, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading profiles: %v\n", err)
		os.Exit(1)
	}

	// If removing the active profile, clear its settings first
	if store.Current == name {
		if profile, err := store.GetProfile(name); err == nil {
			_ = profile.ClearFromClaude() // Ignore errors, continue anyway
		}
	}

	if err := store.RemoveProfile(name); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if err := store.Save(); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving profiles: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Profile '%s' removed successfully.\n", name)
}
