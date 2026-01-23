package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/thekarel/rum/internal/ui/tokens"
)

type clearFlashMsg struct{}

func clearFlashAfter(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return clearFlashMsg{}
	})
}

var flashStyle = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: tokens.Primary, Dark: tokens.Secondary})
