package cmd

import (
	"fmt"
	"os"

	"github.com/bytedance/ccs/internal/config"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List all profiles",
	Long:    `List all Claude Code configuration profiles.`,
	Aliases: []string{"list"},
	Run:     runList,
}

func runList(cmd *cobra.Command, args []string) {
	store, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading profiles: %v\n", err)
		os.Exit(1)
	}

	profiles := store.GetProfileNames()
	if len(profiles) == 0 {
		fmt.Println("No profiles found. Use 'ccs add <name>' to create one.")
		return
	}

	fmt.Println("Profiles:")
	for _, name := range profiles {
		if name == store.Current {
			fmt.Printf("  %s (active)\n", name)
		} else {
			fmt.Printf("  %s\n", name)
		}
	}
}
