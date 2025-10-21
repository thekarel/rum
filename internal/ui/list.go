package ui

import (
	"codeberg.org/thekarel/rum/internal/ui/tokens"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

// script is a record in the script list
type script struct {
	name, cmd string
}

func (i script) Title() string       { return i.name }
func (i script) Description() string { return i.cmd }
func (i script) FilterValue() string { return i.name }

var ListItemTitleStyle = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: tokens.PrimaryBg, Dark: tokens.Secondary}).
	Padding(0, 0, 0, 2)

var ListItemActiveTitleStyle = lipgloss.NewStyle().
	Inherit(ListItemTitleStyle).
	Border(lipgloss.NormalBorder(), false, false, false, true).
	BorderForeground(lipgloss.AdaptiveColor{Light: tokens.Secondary, Dark: tokens.Secondary}).
	Padding(0, 0, 0, 1)

var ListItemDescriptionStyle = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: tokens.Tertiary, Dark: tokens.Tertiary}).
	Padding(0, 0, 0, 2)

var ListItemActiveDescriptionStyle = lipgloss.NewStyle().
	Inherit(ListItemDescriptionStyle).
	Border(lipgloss.NormalBorder(), false, false, false, true).
	BorderForeground(lipgloss.AdaptiveColor{Light: tokens.Secondary, Dark: tokens.Secondary}).
	Padding(0, 0, 0, 1)

func newScriptListDelegate() list.DefaultDelegate {
	delegate := list.NewDefaultDelegate()
	delegate.Styles.NormalTitle = ListItemTitleStyle
	delegate.Styles.SelectedTitle = ListItemActiveTitleStyle
	delegate.Styles.NormalDesc = ListItemDescriptionStyle
	delegate.Styles.SelectedDesc = ListItemActiveDescriptionStyle

	return delegate
}
