package ui

import "github.com/charmbracelet/bubbles/list"

// script is a single record in the script list
type script struct {
	name, cmd string
}

func (i script) Title() string       { return i.name }
func (i script) Description() string { return i.cmd }
func (i script) FilterValue() string { return i.name+i.cmd }

func newItems(packagJsonScripts map[string]string) []list.Item {
	scripts := []list.Item{}

	for name, cmd := range packagJsonScripts {
		scripts = append(scripts, script{name: name, cmd: cmd})
	}

	return scripts
}
