package ui

import (
	"fmt"
	"log"
	"os"
	"strings"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
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
	// subDir is the relative path from cwd to the package.json directory, empty if same
	subDir string
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

// CopyRunCommand returns a command string suitable for pasting from the cwd.
// When running in a subfolder, it uses the package manager's workspace filter
// flag (e.g. "pnpm -F pkg run script") so the command works from the project root.
// Falls back to "cd dir && pm run script" when the package name is not available.
func (m Model) CopyRunCommand() string {
	sel := m.scriptList.SelectedItem()
	if sel == nil {
		return ""
	}

	name := sel.(script).name
	baseCmd := fmt.Sprintf("%s run %s", m.pm, name)

	if m.subDir == "" {
		return baseCmd
	}

	if m.pj.Name != "" {
		switch m.pm {
		case "pnpm":
			return fmt.Sprintf("pnpm -F %s run %s", m.pj.Name, name)
		case "npm":
			return fmt.Sprintf("npm -w %s run %s", m.pj.Name, name)
		case "yarn":
			return fmt.Sprintf("yarn workspace %s run %s", m.pj.Name, name)
		case "bun":
			return fmt.Sprintf("bun --filter %s run %s", m.pj.Name, name)
		}
	}

	return fmt.Sprintf("cd %s && %s", m.subDir, baseCmd)
}

// CopyScriptCommand returns the raw script command, prefixed with "cd dir &&"
// when running in a subfolder so it can be pasted from the project root.
// The command is run through the package manager so that it resolves correctly.
func (m Model) CopyScriptCommand() string {
	sel := m.scriptList.SelectedItem()
	if sel == nil {
		return ""
	}

	cmd := sel.(script).cmd

	if m.subDir != "" {
		return fmt.Sprintf("cd %s && %s exec %s", m.subDir, m.pm, cmd)
	}

	return cmd
}

type ModelInitOpts struct {
	Pj       core.PackageJson
	FilePath string
	Pm       string
	SubDir   string
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
		subDir:     opts.SubDir,
		winWidth:   opts.WinWidth,
		scriptList: scriptList,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
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
			CopyToClipboard(m.CopyRunCommand())
			return m, tea.Quit
		}

		// Copy and quit on C
		if msg.String() == "C" {
			CopyToClipboard(m.CopyScriptCommand())
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

func (m Model) View() tea.View {
	s := "\n"
	s += Header(m.winWidth, m.pj, m.pm, m.path, m.flash)
	s += "\n"
	s += m.scriptList.View()

	return tea.NewView(s)
}
