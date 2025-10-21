package ui

import (
	"codeberg.org/thekarel/rum/internal/core"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	// path is the absolute path to the package.json
	path string
	// pj is the relevant content from package.json
	pj core.PackageJson
	// scripts is the list of name-command pairs in a bullbe list
	scripts list.Model
	// pm is the package manager, e.g. npm
	pm string
	// selected is the command selected by the user
	selected string
}

func (m Model) GetSelected() string {
	return m.selected
}

func InitialModel(packageJson core.PackageJson, filePath, pm string) Model {
	scripts := []list.Item{}

	for name, cmd := range packageJson.Scripts {
		scripts = append(scripts, script{name: name, cmd: cmd})
	}

	delegate := newScriptListDelegate()
	scriptList := list.New(scripts, delegate, 80, 20)
	scriptList.SetShowTitle(false)
	scriptList.SetShowStatusBar(false)

	return Model{
		path:    filePath,
		pj:      packageJson,
		scripts: scriptList,
		pm:      pm,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "enter" {
			sel := m.scripts.SelectedItem()
			m.selected = sel.(script).name
			return m, tea.Quit
		}
		// case tea.WindowSizeMsg:
		// 	h, v := listStyle.GetFrameSize()
		// 	m.scriptList.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.scripts, cmd = m.scripts.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	s := "\n\n"
	s += Header(m.pj.Name, m.path, m.pm)
	s += "\n"
	s += m.scripts.View()

	return s
}
