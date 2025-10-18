package internal

import (
	"codeberg.org/thekarel/rum/internal/core"
	"codeberg.org/thekarel/rum/internal/ui"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	filePath    string
	packageJson core.PackageJson
	scriptList  list.Model
}

// item is for the script list
type item struct {
	name, cmd string
}

func (i item) Title() string {
	return i.name
}
func (i item) Description() string {
	return i.cmd
}
func (i item) FilterValue() string {
	return i.name
}

func initialModel(packageJson core.PackageJson, filePath string) model {
	listItems := []list.Item{}

	for name, cmd := range packageJson.Scripts {
		listItems = append(listItems, item{name: name, cmd: cmd})
	}

	delegate := list.NewDefaultDelegate()
	delegate.Styles.NormalTitle = ui.ListItemTitleStyle
	delegate.Styles.SelectedTitle = ui.ListItemActiveTitleStyle
	delegate.Styles.NormalDesc = ui.ListItemDescriptionStyle
	delegate.Styles.SelectedDesc = ui.ListItemActiveDescriptionStyle
	scriptList := list.New(listItems, delegate, 80, 20)
	scriptList.SetShowTitle(false)
	scriptList.SetShowStatusBar(false)

	return model{
		filePath:    filePath,
		packageJson: packageJson,
		scriptList:  scriptList,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		// case tea.WindowSizeMsg:
		// 	h, v := listStyle.GetFrameSize()
		// 	m.scriptList.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.scriptList, cmd = m.scriptList.Update(msg)
	return m, cmd
}

func (m model) View() string {
	s := "\n\n"
	s += ui.Header(m.packageJson.Name, m.filePath)
	s += "\n"
	s += m.scriptList.View()

	return s
}
