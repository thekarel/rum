package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
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

var titleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color(tokens.primary)).
	Background(lipgloss.Color(tokens.primaryBg)).
	PaddingLeft(2).
	PaddingRight(2).
	Width(80)

var subTitleStyle = titleStyle.
	Bold(false).
	Foreground(lipgloss.Color(tokens.tertiary))

func Title(main, sub string) string {
	return fmt.Sprintf(
		"%s\n%s\n%s\n%s",
		titleStyle.Render(" "),
		titleStyle.Render(main),
		subTitleStyle.Render(sub),
		titleStyle.Render(" "),
	)
}

func ScriptList(scripts map[string]string) string {
	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Width(78).
		Padding(1)

	var items []string

	for name, command := range scripts {
		items = append(items, fmt.Sprintf("%s: %s", lipgloss.NewStyle().Bold(true).Render(name), command))
	}

	return fmt.Sprint(boxStyle.Render(list.New(items).String()))
}
