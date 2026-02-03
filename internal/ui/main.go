package ui

import (
	"fmt"
	"os"

	"github.com/bytedance/ccs/internal/config"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// mainStyles contains styles for the main TUI
var mainStyles = struct {
	title       lipgloss.Style
	help        lipgloss.Style
	divider     lipgloss.Style
	border       lipgloss.Style
}{
	title:   lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Bold(true).PaddingLeft(1),
	help:    lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")),
	divider: lipgloss.NewStyle().Foreground(lipgloss.Color("#F25D94")),
	border:  lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")),
}

// Model represents the main TUI model
type Model struct {
	store     *config.Store
	listPanel *ListPanel
	preview   *PreviewPanel
	width     int
	height    int
	quitting   bool
}

// Init initializes the main model
func (m *Model) Init() tea.Cmd {
	return nil
}

// Update updates the main model
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter", " ":
			// Switch to selected profile
			selected := m.listPanel.GetSelected()
			if selected != "" {
				profile, err := m.store.GetProfile(selected)
				if err == nil {
					// Clear old profile if different
					if m.store.Current != "" && m.store.Current != selected {
						if oldProfile, err := m.store.GetProfile(m.store.Current); err == nil {
							_ = oldProfile.ClearFromClaude()
						}
					}

					// Apply new profile
					if err := profile.ApplyToClaude(); err != nil {
						return m, nil // Error handled silently in TUI
					}

					// Update store
					m.store.SetCurrent(selected)
					_ = m.store.Save()

					// Refresh list to show new active
					m.listPanel.SetItems(m.store.GetProfileNames(), selected)
					m.preview.SetProfile(profile)
				}
			}

		case "r":
			// Remove selected profile
			selected := m.listPanel.GetSelected()
			if selected != "" {
				if m.store.Current == selected {
					if profile, err := m.store.GetProfile(selected); err == nil {
						_ = profile.ClearFromClaude()
					}
				}
				_ = m.store.RemoveProfile(selected)
				_ = m.store.Save()

				// Refresh list
				profiles := m.store.GetProfileNames()
				current := m.store.Current
				m.listPanel.SetItems(profiles, current)
				if current != "" {
					if profile, err := m.store.GetProfile(current); err == nil {
						m.preview.SetProfile(profile)
					} else {
						m.preview.SetProfile(nil)
					}
				} else {
					m.preview.SetProfile(nil)
				}
			}

		case "?":
			// Show help (could expand this)
			return m, nil
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.resize()
	}

	// Update list panel
	var cmd tea.Cmd
	m.listPanel, cmd = m.listPanel.Update(msg)

	// Update preview based on selection
	selected := m.listPanel.GetSelected()
	if profile, err := m.store.GetProfile(selected); err == nil {
		m.preview.SetProfile(profile)
	} else {
		m.preview.SetProfile(nil)
	}

	return m, cmd
}

// resize handles window resize
func (m *Model) resize() {
	listWidth := m.width / 2
	previewWidth := m.width - listWidth
	if listWidth < 30 {
		listWidth = 30
		previewWidth = m.width - 30
	}

	m.listPanel.SetSize(listWidth, m.height-2)
	m.preview.SetSize(previewWidth, m.height-2)
}

// View returns the view of the main model
func (m *Model) View() string {
	if m.quitting {
		return ""
	}

	// Title bar
	title := mainStyles.title.Render("CCS - Claude Code Switcher")

	// Content - split layout
	content := lipgloss.JoinHorizontal(
		lipgloss.Left,
		mainStyles.border.Render(m.listPanel.View()),
		mainStyles.border.Render(m.preview.View()),
	)

	// Help bar
	help := mainStyles.help.Render(" [Enter] Switch  [r] Remove  [q] Quit")

	// Combine all
	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		content,
		help,
	)
}

// NewModel creates a new main TUI model
func NewModel() (*Model, error) {
	store, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load profiles: %w", err)
	}

	profiles := store.GetProfileNames()
	current := store.Current

	listPanel := NewListPanel()
	listPanel.SetItems(profiles, current)

	previewPanel := NewPreviewPanel()
	if current != "" {
		if profile, err := store.GetProfile(current); err == nil {
			previewPanel.SetProfile(profile)
		}
	}

	return &Model{
		store:     store,
		listPanel: listPanel,
		preview:   previewPanel,
		quitting:   false,
	}, nil
}

// Run launches the TUI
func Run() error {
	model, err := NewModel()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return err
	}

	p := tea.NewProgram(
		model,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running TUI: %v\n", err)
		return err
	}

	return nil
}
