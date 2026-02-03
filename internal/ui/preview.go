package ui

import (
	"fmt"

	"github.com/bytedance/ccs/internal/config"
	"github.com/charmbracelet/lipgloss"
)

// previewStyles contains styles for the preview panel
var previewStyles = struct {
	title   lipgloss.Style
	key     lipgloss.Style
	value   lipgloss.Style
	empty   lipgloss.Style
}{
	title:   lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Bold(true),
	key:     lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Width(30),
	value:   lipgloss.NewStyle().Foreground(lipgloss.Color("#FAFAFA")).Faint(true),
	empty:   lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")),
}

// PreviewPanel represents the config preview panel
type PreviewPanel struct {
	profile *config.Profile
	width   int
	height  int
}

// NewPreviewPanel creates a new preview panel
func NewPreviewPanel() *PreviewPanel {
	return &PreviewPanel{
		profile: config.NewProfile(),
	}
}

// SetProfile sets the profile to preview
func (p *PreviewPanel) SetProfile(profile *config.Profile) {
	if profile != nil {
		p.profile = profile
	} else {
		p.profile = config.NewProfile()
	}
}

// SetSize sets the size of the panel
func (p *PreviewPanel) SetSize(width, height int) {
	p.width = width
	p.height = height
}

// View returns the view of the preview panel
func (p *PreviewPanel) View() string {
	if p.profile == nil || len(p.profile.Env) == 0 {
		return p.renderEmpty()
	}

	lines := []string{previewStyles.title.Render("Profile Details")}

	// Display all env vars in order
	for _, key := range config.EnvKeys {
		value, exists := p.profile.GetEnv(key)
		if exists {
			label, ok := config.EnvLabels[key]
			if ok {
				lines = append(lines, fmt.Sprintf("%s%s",
					previewStyles.key.Render(label),
					previewStyles.value.Render(maskValue(key, value)),
				))
			}
		}
	}

	return lipgloss.NewStyle().Width(p.width).Height(p.height).Render(
		lipgloss.JoinVertical(lipgloss.Left, lines...),
	)
}

// renderEmpty returns the empty state view
func (p *PreviewPanel) renderEmpty() string {
	return lipgloss.NewStyle().Width(p.width).Height(p.height).Align(
		lipgloss.Center, lipgloss.Center,
	).Render(previewStyles.empty.Render("No profile selected"))
}

// maskValue masks sensitive values for display
func maskValue(key, value string) string {
	// Mask API tokens
	if key == config.EnvAuthToken {
		if len(value) <= 8 {
			return "***"
		}
		return value[:4] + "***" + value[len(value)-4:]
	}
	return value
}
