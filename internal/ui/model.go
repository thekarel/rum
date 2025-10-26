package ui

import (
	"codeberg.org/thekarel/rum/internal/core"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	// the width of the name column in the list: the length of the longest name
	nameWidth int
	// path is the absolute path to the package.json
	path string
	// pm is the package manager, e.g. npm
	pm string
	// pj is the relevant content from package.json
	pj core.PackageJson
	// scriptList is the list of name-command pairs in a Bubble list
	scriptList list.Model
	// selected is the command selected by the user
	selected string
	// the width of the window
	winWidth int
}

func (m Model) GetSelected() string {
	return m.selected
}

func InitialModel(pj core.PackageJson, filePath, pm string) Model {
	nameWidth := 0
	for name := range pj.Scripts {
		if len(name) > nameWidth {
			nameWidth = len(name)
		}
	}

	items := newItems(pj.Scripts)
	// The width will be update once we get the window width in model.Update
	// In the height +5 is the help bar and other cruft
	scriptList := newList(items, newItemDelegate(nameWidth), 80, len(items)+4)

	return Model{
		nameWidth:  nameWidth,
		path:       filePath,
		pj:         pj,
		pm:         pm,
		scriptList: scriptList,
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
			sel := m.scriptList.SelectedItem()
			if sel == nil {
				return m, nil
			}

			m.selected = sel.(script).name
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.winWidth = msg.Width
		m.scriptList.SetWidth(msg.Width)
	}

	var cmd tea.Cmd
	m.scriptList, cmd = m.scriptList.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	s := "\n"
	s += Header(m.winWidth, m.pj, m.pm, m.path)
	s += "\n"
	s += m.scriptList.View()

	return s
}
