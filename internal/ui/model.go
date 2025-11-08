package ui

import (
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/thekarel/rum/internal/core"
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
	// Width will be update once we get the window width in model.Update
	// Height: +N is the header, this will also be updated on window height change.
	scriptList := newList(items, newItemDelegate(nameWidth), 80, len(items)+2)

	path := filePath
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	path = strings.Replace(path, homeDir, "~", 1)
	path = strings.Replace(path, "package.json", "", 1)

	return Model{
		nameWidth:  nameWidth,
		path:       path,
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
		m.scriptList.SetHeight(msg.Height - 2)
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
