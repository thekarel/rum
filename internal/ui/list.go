package ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/lipgloss"
	"github.com/thekarel/rum/internal/ui/tokens"
)

const (
	bullet = "●"
)

func newList(scripts []list.Item, delegate list.ItemDelegate, w, h int) list.Model {
	scriptList := list.New(scripts, delegate, w, h)
	scriptList.SetShowTitle(false)
	scriptList.SetShowStatusBar(false)

	// Remove left padding from TitleBar (which wraps the filter)
	scriptList.Styles.TitleBar = lipgloss.NewStyle()
	// Remove padding from help
	scriptList.Styles.HelpStyle = lipgloss.NewStyle()

	// Add custom key binding for enter/run command
	enterKey := key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "Run selected command"),
	)
	// Add custom key binding for copy command
	copyKey := key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "Copy run command and quit"),
	)
	copyQuitKey := key.NewBinding(
		key.WithKeys("C"),
		key.WithHelp("C", "Copy script and quit"),
	)
	scriptList.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{copyKey, copyQuitKey, enterKey}
	}
	scriptList.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{copyKey, copyQuitKey, enterKey}
	}

	// Filter prompt and input text style
	filterStyle := lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: tokens.Secondary, Dark: tokens.Secondary})
	scriptList.FilterInput.PromptStyle = filterStyle
	scriptList.FilterInput.TextStyle = filterStyle

	// Navigation dots
	p := paginator.New()
	p.Type = paginator.Dots
	p.ActiveDot = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: tokens.Primary, Dark: tokens.Secondary}).
		SetString(bullet).
		PaddingLeft(1).
		String()
	p.InactiveDot = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: tokens.Secondary, Dark: tokens.Tertiary}).
		SetString(bullet).
		PaddingLeft(1).
		String()
	scriptList.Paginator = p
	scriptList.Styles.PaginationStyle = lipgloss.NewStyle().PaddingBottom(1)

	return scriptList
}
