package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/thekarel/rum/internal/core"
)

var headerStyle = lipgloss.NewStyle()

var titleStyle = lipgloss.NewStyle().Bold(true)

var subTitleStyle = lipgloss.NewStyle()

// Header renders the app's header
// width: the width of the header
// pj: the package.json content
// pm: the detected package manager's name
// path: the file path
func Header(width int, pj core.PackageJson, pm string, path string, flash string) string {
	pm = pm + " ⌁"

	text := fmt.Sprintf(
		"%s %s %s %s",
		subTitleStyle.Render(pm),
		titleStyle.Render(pj.Name),
		subTitleStyle.Render(path),
		flashStyle.Render(flash),
	)

	if len(text) > width {
		text = fmt.Sprintf(
			"%s %s",
			subTitleStyle.Render(pm),
			titleStyle.Render(pj.Name),
		)
	}

	if len(text) > width {
		text = titleStyle.Render(pj.Name)
	}

	return headerStyle.Width(width).Render(text)
}
