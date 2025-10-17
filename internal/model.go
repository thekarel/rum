package internal

import (
	"codeberg.org/thekarel/rum/internal/core"
	"codeberg.org/thekarel/rum/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	filePath    string
	packageJson core.PackageJson
}

func initialModel(packageJson core.PackageJson, filePath string) model {
	return model{
		filePath:    filePath,
		packageJson: packageJson,
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	s := "\n\n"
	s += ui.Title(m.packageJson.Name, m.filePath)
	s += "\n"

	s += ui.ScriptList(m.packageJson.Scripts)

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
