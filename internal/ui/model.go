package ui

import (
	"fmt"
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
	// flash message
	flash string
}

func (m Model) GetSelected() string {
	return m.selected
}

// RunCommand returns the run command as a string.
// Example: "pnpm run db:overload"
func (m Model) RunCommand() string {
	sel := m.scriptList.SelectedItem()
	if sel == nil {
		return ""
	}

	name := sel.(script).name

	return fmt.Sprintf("%s run %s", m.pm, name)
}

type ModelInitOpts struct {
	Pj       core.PackageJson
	FilePath string
	Pm       string
	WinWidth int
	Readonly bool
}

func InitialModel(opts ModelInitOpts) Model {
	nameWidth := 0
	for name := range opts.Pj.Scripts {
		if len(name) > nameWidth {
			nameWidth = len(name)
		}
	}

	items := newItems(opts.Pj.Scripts)
	// Width will be update once we get the window width in model.Update
	// Height: +N is the header, this will also be updated on window height change.
	scriptList := newList(items, newItemDelegate(nameWidth, !opts.Readonly), 80, len(items)+2)

	if opts.Readonly {
		scriptList.SetShowPagination(false)
		scriptList.SetShowStatusBar(false)
		scriptList.SetShowHelp(false)
		scriptList.SetHeight(len(items))
	}

	path := opts.FilePath
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	path = strings.Replace(path, homeDir, "~", 1)
	path = strings.Replace(path, "package.json", "", 1)

	return Model{
		nameWidth:  nameWidth,
		path:       path,
		pj:         opts.Pj,
		pm:         opts.Pm,
		winWidth:   opts.WinWidth,
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

		// The keys below don't have effect while filtering
		if m.scriptList.FilterState() == list.Filtering {
			break
		}

		// The keys below assume there is a selected command
		sel := m.scriptList.SelectedItem()
		if sel == nil {
			break
		}

		// Handle pressing enter, but do nothing if the list is being filtered,
		// in order to allow the user to
		// 1. Press /
		// 2. Type a search term
		// 3. See the filtered list, move up or down
		// 4. Press enter -> This stops the filtering and selects a command without executing it.
		// The user then can decide what to do with it (run or copy, for example).
		if msg.String() == "enter" {
			m.selected = sel.(script).name
			return m, tea.Quit
		}

		// Copy the command to the clipboard
		if msg.String() == "c" {
			CopyToClipboard(m.RunCommand())
			return m, tea.Quit
		}

		// Copy and quit on C
		if msg.String() == "C" {
			CopyToClipboard(sel.(script).cmd)
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.winWidth = msg.Width
		m.scriptList.SetWidth(msg.Width)
		m.scriptList.SetHeight(msg.Height - 2)

	case clearFlashMsg:
		m.flash = ""
	}

	var cmd tea.Cmd
	m.scriptList, cmd = m.scriptList.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	s := "\n"
	s += Header(m.winWidth, m.pj, m.pm, m.path, m.flash)
	s += "\n"
	s += m.scriptList.View()

	return s
}
