package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bytedance/ccs/internal/config"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <name>",
	Short: "Add a new profile",
	Long:  `Add a new Claude Code configuration profile interactively.`,
	Args:  cobra.ExactArgs(1),
	Run:   runAdd,
}

func runAdd(cmd *cobra.Command, args []string) {
	name := args[0]

	store, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading profiles: %v\n", err)
		os.Exit(1)
	}

	profile := config.NewProfile()
	reader := bufio.NewReader(os.Stdin)

	// Prompt for each config value
	inputs := []struct {
		key   string
		prompt string
	}{
		{config.EnvAuthToken, "ANTHROPIC_AUTH_TOKEN"},
		{config.EnvBaseURL, "ANTHROPIC_BASE_URL"},
		{config.EnvDefaultHaikuModel, "ANTHROPIC_DEFAULT_HAIKU_MODEL"},
		{config.EnvDefaultOpusModel, "ANTHROPIC_DEFAULT_OPUS_MODEL"},
		{config.EnvDefaultSonnetModel, "ANTHROPIC_DEFAULT_SONNET_MODEL"},
		{config.EnvModel, "ANTHROPIC_MODEL"},
	}

	for _, input := range inputs {
		fmt.Printf("%s: ", input.prompt)
		value, _ := reader.ReadString('\n')
		value = strings.TrimSpace(value)
		if value != "" {
			profile.SetEnv(input.key, value)
		}
	}

	// Check if profile is empty (all fields empty)
	if len(profile.Env) == 0 {
		fmt.Fprintln(os.Stderr, "Error: Profile cannot be empty. Provide at least one configuration value.")
		os.Exit(1)
	}

	if err := store.AddProfile(name, profile); err != nil {
		fmt.Fprintf(os.Stderr, "Error adding profile: %v\n", err)
		os.Exit(1)
	}

	if err := store.Save(); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving profiles: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Profile '%s' added successfully.\n", name)
}
