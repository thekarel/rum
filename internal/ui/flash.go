package ui

import (
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/thekarel/rum/internal/ui/tokens"
)

type clearFlashMsg struct{}

func clearFlashAfter(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return clearFlashMsg{}
	})
}

var flashStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color(tokens.Secondary))
