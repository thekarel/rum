package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var tokens = struct {
	primary   string
	secondary string
	tertiary  string
	primaryBg string
}{
	primary:   "#6FFFE9",
	secondary: "#5BC0BE",
	tertiary:  "#3A506B",
	primaryBg: "#0B132B",
}

var ListItemTitleStyle = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: tokens.primaryBg, Dark: tokens.secondary}).
	Padding(0, 0, 0, 2)

var ListItemActiveTitleStyle = lipgloss.NewStyle().
	Inherit(ListItemTitleStyle).
	Border(lipgloss.NormalBorder(), false, false, false, true).
	BorderForeground(lipgloss.AdaptiveColor{Light: tokens.secondary, Dark: tokens.secondary}).
	Padding(0, 0, 0, 1)

var ListItemDescriptionStyle = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: tokens.tertiary, Dark: tokens.tertiary}).
	Padding(0, 0, 0, 2)

var ListItemActiveDescriptionStyle = lipgloss.NewStyle().
	Inherit(ListItemDescriptionStyle).
	Border(lipgloss.NormalBorder(), false, false, false, true).
	BorderForeground(lipgloss.AdaptiveColor{Light: tokens.secondary, Dark: tokens.secondary}).
	Padding(0, 0, 0, 1)

var HeaderStyle = lipgloss.NewStyle().
	Background(lipgloss.Color(tokens.primaryBg)).
	Padding(2).
	Width(80)

var titleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color(tokens.primary))

var subTitleStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color(tokens.tertiary))

func Header(main, sub string) string {
	return HeaderStyle.Render(
		fmt.Sprintf("%s\n%s", titleStyle.Render(main), subTitleStyle.Render(sub)),
	)
}
