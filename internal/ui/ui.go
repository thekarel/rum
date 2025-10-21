package ui

import (
	"fmt"

	"codeberg.org/thekarel/rum/internal/ui/tokens"
	"github.com/charmbracelet/lipgloss"
)

var HeaderStyle = lipgloss.NewStyle().
	Background(lipgloss.Color(tokens.PrimaryBg)).
	Padding(2).
	Width(80)

var titleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color(tokens.Primary))

var subTitleStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color(tokens.Tertiary))

// Header renders the app's header
// main: the main title, the name of the package
// sub: the subtitle, the path to the json file
// pm: the detected package manager
func Header(main, sub, pm string) string {
	return HeaderStyle.Render(
		fmt.Sprintf(
			"%s\n%s",
			titleStyle.Render(main),
			subTitleStyle.Render(fmt.Sprintf("%s | %s", sub, pm)),
		),
	)
}
