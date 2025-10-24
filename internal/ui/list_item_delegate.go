package ui

import (
	"fmt"
	"io"

	"codeberg.org/thekarel/rum/internal/ui/tokens"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var nameStyle = lipgloss.NewStyle().
	PaddingLeft(2).
	Bold(false)

var selectedNameStyle = lipgloss.NewStyle().
	// Background(lipgloss.AdaptiveColor{Light: tokens.Primary, Dark: tokens.Secondary}).
	// Border(lipgloss.ThickBorder(), false, false, false, true).
	// BorderForeground(lipgloss.AdaptiveColor{Light: tokens.Primary, Dark: tokens.Secondary}).
	// PaddingLeft(1).
	Bold(true)

var cmdStyle = lipgloss.NewStyle()

var selectedCmdStyle = lipgloss.NewStyle().
	// Background(lipgloss.AdaptiveColor{Light: tokens.Primary, Dark: tokens.Secondary}).
	// Foreground(lipgloss.AdaptiveColor{Light: tokens.Primary, Dark: tokens.Secondary}).
	Bold(true)

var lineStyle = lipgloss.NewStyle().
	// Background(lipgloss.AdaptiveColor{Light: tokens.PrimaryBg, Dark: tokens.PrimaryBg}).
	Bold(false)

var selectedLineStyle = lipgloss.NewStyle().
	Background(lipgloss.AdaptiveColor{Light: tokens.PrimaryBg, Dark: tokens.PrimaryBg}).
	Foreground(lipgloss.AdaptiveColor{Light: tokens.Primary, Dark: tokens.Secondary}).
	Bold(true)

type itemDelegate struct {
	nameWidth int
}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	s, ok := listItem.(script)
	if !ok {
		return
	}

	// nameFn := nameStyle.Render
	// cmdFn := cmdStyle.Render
	lineFn := lineStyle.Render
	if index == m.Index() {
		// nameFn = selectedNameStyle.Render
		// cmdFn = selectedCmdStyle.Render
		lineFn = selectedLineStyle.Render
	}

	namePadded := fmt.Sprintf("%-*s", d.nameWidth+3, s.name)
	cmdPadding := m.Width() - len(namePadded) - 0
	cmdPadded := fmt.Sprintf("%-*s", cmdPadding, s.cmd)

	fmt.Fprint(w,
		lineFn(fmt.Sprintf("%s %s", namePadded, cmdPadded)),
	)
}

func newItemDelegate(nameWidth int) itemDelegate {
	return itemDelegate{
		nameWidth: nameWidth,
	}
}
