package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// ClaudeSettings represents Claude Code's settings.json structure
type ClaudeSettings struct {
	Env map[string]string `json:"env"`
}

// backup creates a backup of the current Claude settings
func backupClaudeSettings() error {
	claudePath := getClaudeConfigPath()
	if _, err := os.Stat(claudePath); err != nil {
		if os.IsNotExist(err) {
			return nil // No file to backup
		}
		return err
	}

	// Create backup directory
	backupDir := filepath.Join(filepath.Dir(getProfilesPath()), "backups")
	if err := os.MkdirAll(backupDir, 0755); err != nil {
		return err
	}

	// Create backup filename with timestamp
	timestamp := time.Now().Format("20060102-150405")
	backupPath := filepath.Join(backupDir, "settings-"+timestamp+".json")

	// Read current settings
	data, err := os.ReadFile(claudePath)
	if err != nil {
		return err
	}

	// Write backup
	if err := os.WriteFile(backupPath, data, 0644); err != nil {
		return err
	}

	// Clean up old backups (keep last 5)
	rotateBackups(backupDir, 5)

	return nil
}

// rotateBackups keeps only the most recent n backups
func rotateBackups(backupDir string, keep int) {
	entries, err := os.ReadDir(backupDir)
	if err != nil {
		return
	}

	var files []os.FileInfo
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		info, err := entry.Info()
		if err != nil {
			continue
		}
		files = append(files, info)
	}

	// Sort by modification time (newest first)
	for i := 0; i < len(files); i++ {
		for j := i + 1; j < len(files); j++ {
			if files[i].ModTime().Before(files[j].ModTime()) {
				files[i], files[j] = files[j], files[i]
			}
		}
	}

	// Delete old backups
	for i := keep; i < len(files); i++ {
		path := filepath.Join(backupDir, files[i].Name())
		os.Remove(path)
	}
}

// ApplyToClaude applies the profile to Claude's settings.json
func (p *Profile) ApplyToClaude() error {
	// Backup current settings first
	if err := backupClaudeSettings(); err != nil {
		return fmt.Errorf("failed to backup settings: %w", err)
	}

	// Read current Claude settings
	claudePath := getClaudeConfigPath()
	var settings ClaudeSettings

	if data, err := os.ReadFile(claudePath); err == nil {
		if err := json.Unmarshal(data, &settings); err != nil {
			// If parse fails, start fresh
			settings = ClaudeSettings{Env: make(map[string]string)}
		}
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("failed to read Claude settings: %w", err)
	} else {
		settings = ClaudeSettings{Env: make(map[string]string)}
	}

	// Initialize Env if nil
	if settings.Env == nil {
		settings.Env = make(map[string]string)
	}

	// Apply profile settings
	for key, value := range p.Env {
		settings.Env[key] = value
	}

	// Ensure directory exists
	claudeDir := filepath.Dir(claudePath)
	if err := os.MkdirAll(claudeDir, 0755); err != nil {
		return fmt.Errorf("failed to create Claude config directory: %w", err)
	}

	// Write to temp file first for atomicity
	tmpPath := claudePath + ".tmp"
	data, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal settings: %w", err)
	}

	if err := os.WriteFile(tmpPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write temp file: %w", err)
	}

	// Atomic rename
	if err := os.Rename(tmpPath, claudePath); err != nil {
		return fmt.Errorf("failed to rename temp file: %w", err)
	}

	return nil
}

// ClearFromClaude removes the profile's env vars from Claude's settings.json
func (p *Profile) ClearFromClaude() error {
	claudePath := getClaudeConfigPath()

	// Read current settings
	var settings ClaudeSettings

	data, err := os.ReadFile(claudePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // No file to clear
		}
		return fmt.Errorf("failed to read Claude settings: %w", err)
	}

	if err := json.Unmarshal(data, &settings); err != nil {
		return fmt.Errorf("failed to parse Claude settings: %w", err)
	}

	// Remove profile's env vars
	for key := range p.Env {
		delete(settings.Env, key)
	}

	// Write to temp file first for atomicity
	tmpPath := claudePath + ".tmp"
	data, err = json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal settings: %w", err)
	}

	if err := os.WriteFile(tmpPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write temp file: %w", err)
	}

	// Atomic rename
	if err := os.Rename(tmpPath, claudePath); err != nil {
		return fmt.Errorf("failed to rename temp file: %w", err)
	}

	return nil
}

// GetCurrentClaudeSettings reads the current Claude settings
func GetCurrentClaudeSettings() (*ClaudeSettings, error) {
	claudePath := getClaudeConfigPath()
	data, err := os.ReadFile(claudePath)
	if err != nil {
		if os.IsNotExist(err) {
			return &ClaudeSettings{Env: make(map[string]string)}, nil
		}
		return nil, fmt.Errorf("failed to read Claude settings: %w", err)
	}

	var settings ClaudeSettings
	if err := json.Unmarshal(data, &settings); err != nil {
		return nil, fmt.Errorf("failed to parse Claude settings: %w", err)
	}

	if settings.Env == nil {
		settings.Env = make(map[string]string)
	}

	return &settings, nil
}
