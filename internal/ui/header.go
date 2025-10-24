package ui

import (
	"fmt"

	"codeberg.org/thekarel/rum/internal/core"
	"codeberg.org/thekarel/rum/internal/ui/tokens"
	"github.com/charmbracelet/lipgloss"
)

var headerStyle = lipgloss.NewStyle().
	Background(lipgloss.Color(tokens.PrimaryBg)).
	Padding(1, 2)

var titleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color(tokens.Primary))

var subTitleStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color(tokens.Tertiary))

// Header renders the app's header
// width: the width of the header
// pj: the package.json content
// pm: the detected package manager's name
// path: the file path
func Header(width int, pj core.PackageJson, pm string, path string) string {
	return headerStyle.Width(width).Render(
		fmt.Sprintf(
			"%s\n%s\n%s",
			titleStyle.Render(pj.Name),
			subTitleStyle.Render(path),
			subTitleStyle.Render(pm),
		),
	)
}
