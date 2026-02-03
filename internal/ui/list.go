package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// listItem represents a profile in the list
type listItem struct {
	name   string
	active bool
}

func (i listItem) Title() string {
	if i.active {
		return i.name + " (active)"
	}
	return i.name
}

func (i listItem) Description() string {
	return ""
}

func (i listItem) FilterValue() string {
	return i.name
}

// listStyles contains styles for the list panel
var listStyles = struct {
	normalTitle, activeTitle, dimmedTitle lipgloss.Style
	normal, selected, dimmed lipgloss.Style
}{
	normalTitle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Bold(true),
	activeTitle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#FAFAFA")).Bold(true).Background(lipgloss.Color("#F25D94")),
	normal:       lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("#FAFAFA")),
	selected:     lipgloss.NewStyle().PaddingLeft(0).Foreground(lipgloss.Color("#F25D94")).Background(lipgloss.Color("#FAFAFA")),
	dimmed:       lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("#626262")),
	dimmedTitle:  lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("#626262")),
}

// newListModel creates a new list model
func newListModel() list.Model {
	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = false
	delegate.Styles.SelectedTitle = listStyles.selected
	delegate.Styles.SelectedDesc = listStyles.selected

	l := list.New(nil, delegate, 0, 0)
	l.Title = "Profiles"
	l.Styles.Title = listStyles.normalTitle
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	return l
}

// ListPanel represents the profile list panel
type ListPanel struct {
	list   list.Model
	width  int
	height int
}

// NewListPanel creates a new list panel
func NewListPanel() *ListPanel {
	return &ListPanel{
		list: newListModel(),
	}
}

// SetItems sets the items in the list
func (p *ListPanel) SetItems(names []string, current string) {
	items := make([]list.Item, len(names))
	for i, name := range names {
		items[i] = listItem{
			name:   name,
			active: name == current,
		}
	}
	p.list.SetItems(items)
}

// GetSelected returns the selected profile name
func (p *ListPanel) GetSelected() string {
	if p.list.SelectedItem() == nil {
		return ""
	}
	return p.list.SelectedItem().(listItem).name
}

// SetSize sets the size of the panel
func (p *ListPanel) SetSize(width, height int) {
	p.width = width
	p.height = height
	p.list.SetSize(width, height-2) // Leave room for title
}

// Update updates the list panel
func (p *ListPanel) Update(msg tea.Msg) (*ListPanel, tea.Cmd) {
	var cmd tea.Cmd
	p.list, cmd = p.list.Update(msg)
	return p, cmd
}

// View returns the view of the list panel
func (p *ListPanel) View() string {
	return p.list.View()
}
