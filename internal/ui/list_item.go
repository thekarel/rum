package ui

import (
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/list"
)

// script is a single record in the script list
type script struct {
	name, cmd string
}

func (i script) Title() string       { return i.name }
func (i script) Description() string { return i.cmd }
func (i script) FilterValue() string { return i.name+i.cmd }

func newItems(pjs map[string]string) []list.Item {
	// This will be the final list
	scripts := []list.Item{}

	// Pick out the names only...
	names := make([]string, 0, len(pjs))
	for name := range pjs{
	  names = append(names, name)
	}
	// sort the names...
	slices.SortFunc(names, func(a, b string) int {
		return strings.Compare(strings.ToLower(a), strings.ToLower(b))
	})
	// Use this sorted name slice to populate the final list
	for _, name := range names {
		scripts = append(scripts, script{name: name, cmd: pjs[name]})
	}

	return scripts
}
