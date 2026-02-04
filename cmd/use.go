package cmd

import (
	"fmt"
	"os"

	"github.com/bytedance/ccs/internal/config"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use <name>",
	Short: "Switch to a profile",
	Long:  `Switch to the specified Claude Code configuration profile.`,
	Args:  cobra.ExactArgs(1),
	Run:   runUse,
}

func runUse(cmd *cobra.Command, args []string) {
	name := args[0]

	store, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading profiles: %v\n", err)
		os.Exit(1)
	}

	profile, err := store.GetProfile(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// If there was a previous active profile, clear its settings first
	if store.Current != "" && store.Current != name {
		if oldProfile, err := store.GetProfile(store.Current); err == nil {
			_ = oldProfile.ClearFromClaude() // Ignore errors, continue anyway
		}
	}

	// Apply the new profile
	if err := profile.ApplyToClaude(); err != nil {
		fmt.Fprintf(os.Stderr, "Error applying profile: %v\n", err)
		os.Exit(1)
	}

	// Set hasCompletedOnboarding to skip Claude Code's first-time setup
	if changed, err := config.SetHasCompletedOnboarding(); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: failed to set onboarding flag: %v\n", err)
	} else if changed {
		fmt.Println("Onboarding flag set: first-time setup will be skipped.")
	}

	// Update store current
	if err := store.SetCurrent(name); err != nil {
		fmt.Fprintf(os.Stderr, "Error updating active profile: %v\n", err)
		os.Exit(1)
	}

	if err := store.Save(); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving profiles: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Switched to profile '%s'.\n", name)
	fmt.Println("Restart your terminal or Claude Code to apply changes.")
}
