package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

const (
	// AppName is name of the application
	AppName = "ccs"
)

var (
	once         sync.Once
	profilesPath string
)

// getProfilesPath returns the path to the profiles.json file
func getProfilesPath() string {
	once.Do(func() {
		home, _ := os.UserHomeDir()
		profilesPath = filepath.Join(home, ".ccs", "profiles.json")
	})
	return profilesPath
}

// getClaudeConfigPath returns the path to Claude's settings.json
func getClaudeConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".claude", "settings.json")
}

// getClaudeJSONPath returns the path to ~/.claude.json
// This file contains MCP servers and hasCompletedOnboarding flag
func getClaudeJSONPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".claude.json")
}

// GetClaudeJSONPath returns the path to ~/.claude.json (exported)
func GetClaudeJSONPath() string {
	return getClaudeJSONPath()
}

// Profile represents a Claude Code configuration profile
type Profile struct {
	Env map[string]string `json:"env"`
}

// NewProfile creates a new profile with the given configuration
func NewProfile() *Profile {
	return &Profile{
		Env: make(map[string]string),
	}
}

// SetEnv sets an environment variable in the profile
func (p *Profile) SetEnv(key, value string) {
	p.Env[key] = value
}

// GetEnv gets an environment variable from the profile
func (p *Profile) GetEnv(key string) (string, bool) {
	val, ok := p.Env[key]
	return val, ok
}

// Store represents the profiles storage
type Store struct {
	Current string             `json:"current"`
	Profiles map[string]*Profile `json:"profiles"`
}

// NewStore creates a new store
func NewStore() *Store {
	return &Store{
		Profiles: make(map[string]*Profile),
	}
}

// validateProfileName checks if a profile name is valid
func validateProfileName(name string) error {
	// Remove any invalid characters, check if something remains
	invalidChars := regexp.MustCompile(`[^a-zA-Z0-9_-]`)
	trimmed := invalidChars.ReplaceAllString(name, "")
	if trimmed != name {
		return fmt.Errorf("profile name must contain only letters, numbers, hyphens, and underscores")
	}
	if strings.TrimSpace(name) == "" {
		return fmt.Errorf("profile name cannot be empty")
	}
	return nil
}

// Load loads the profiles from disk
func Load() (*Store, error) {
	path := getProfilesPath()
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			// Return empty store if file doesn't exist
			return NewStore(), nil
		}
		return nil, fmt.Errorf("failed to read profiles file: %w", err)
	}

	var store Store
	if err := json.Unmarshal(data, &store); err != nil {
		return nil, fmt.Errorf("failed to parse profiles file: %w", err)
	}

	// Initialize profiles map if nil
	if store.Profiles == nil {
		store.Profiles = make(map[string]*Profile)
	}

	return &store, nil
}

// Save saves the profiles to disk
func (s *Store) Save() error {
	path := getProfilesPath()
	dir := filepath.Dir(path)

	// Create directory if it doesn't exist
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// Write to temp file first for atomicity
	tmpPath := path + ".tmp"
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal profiles: %w", err)
	}

	if err := os.WriteFile(tmpPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write temp file: %w", err)
	}

	// Atomic rename
	if err := os.Rename(tmpPath, path); err != nil {
		return fmt.Errorf("failed to rename temp file: %w", err)
	}

	return nil
}

// AddProfile adds a new profile
func (s *Store) AddProfile(name string, profile *Profile) error {
	if err := validateProfileName(name); err != nil {
		return err
	}

	if _, exists := s.Profiles[name]; exists {
		return fmt.Errorf("profile '%s' already exists", name)
	}

	s.Profiles[name] = profile

	// Set as current if it's the first profile
	if s.Current == "" {
		s.Current = name
	}

	return nil
}

// RemoveProfile removes a profile
func (s *Store) RemoveProfile(name string) error {
	if _, exists := s.Profiles[name]; !exists {
		return fmt.Errorf("profile '%s' not found", name)
	}

	delete(s.Profiles, name)

	// Update current if we removed the active profile
	if s.Current == name {
		if len(s.Profiles) > 0 {
			// Set to first available profile
			for k := range s.Profiles {
				s.Current = k
				break
			}
		} else {
			s.Current = ""
		}
	}

	return nil
}

// GetProfile gets a profile by name
func (s *Store) GetProfile(name string) (*Profile, error) {
	profile, exists := s.Profiles[name]
	if !exists {
		return nil, fmt.Errorf("profile '%s' not found", name)
	}
	return profile, nil
}

// ListProfiles returns a list of all profile names
func (s *Store) ListProfiles() []string {
	names := make([]string, 0, len(s.Profiles))
	for name := range s.Profiles {
		names = append(names, name)
	}
	return names
}

// SetCurrent sets the current profile
func (s *Store) SetCurrent(name string) error {
	if _, exists := s.Profiles[name]; !exists {
		return fmt.Errorf("profile '%s' not found", name)
	}
	s.Current = name
	return nil
}

// GetProfileNames returns profile names in a sorted order
func (s *Store) GetProfileNames() []string {
	return s.ListProfiles()
}
