package ui

import (
	"codeberg.org/thekarel/rum/internal/ui/tokens"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

func newList(scripts []list.Item, delegate list.ItemDelegate, w, h int) list.Model {
	scriptList := list.New(scripts, delegate, w, h)
	scriptList.SetShowTitle(false)
	scriptList.SetShowStatusBar(false)

	// Remove left padding from TitleBar (which wraps the filter)
	scriptList.Styles.TitleBar = lipgloss.NewStyle()
	// Remove padding from help
	scriptList.Styles.HelpStyle = lipgloss.NewStyle()

	// Filter prompt and input text style
	filterStyle :=  lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: tokens.Secondary, Dark: tokens.Secondary})
	scriptList.FilterInput.PromptStyle = filterStyle
	scriptList.FilterInput.TextStyle = filterStyle

	return scriptList
}
