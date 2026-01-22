package ui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/thekarel/rum/internal/ui/tokens"
)

var lineStyle = lipgloss.NewStyle().
	Bold(false)

var selectedLineStyle = lipgloss.NewStyle().
	Background(lipgloss.AdaptiveColor{Light: tokens.PrimaryBg, Dark: tokens.PrimaryBg}).
	Foreground(lipgloss.AdaptiveColor{Light: tokens.Primary, Dark: tokens.Secondary}).
	Bold(true)

type itemDelegate struct {
	nameWidth int
	highlight bool
}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	s, ok := listItem.(script)
	if !ok {
		return
	}

	lineFn := lineStyle.Render
	if index == m.Index() && d.highlight {
		lineFn = selectedLineStyle.Render
	}

	namePadded := fmt.Sprintf("%-*s", d.nameWidth+3, s.name)
	cmdPadding := m.Width() - len(namePadded) - 0
	cmdPadded := fmt.Sprintf("%-*s", cmdPadding, s.cmd)

	fmt.Fprint(w,
		lineFn(fmt.Sprintf("%s %s", namePadded, cmdPadded)),
	)
}

func newItemDelegate(nameWidth int, highlight bool) itemDelegate {
	return itemDelegate{
		nameWidth: nameWidth,
		highlight: highlight,
	}
}
