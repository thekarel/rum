package internal

import (
	"fmt"
	"log"
	"os"

	"codeberg.org/thekarel/rum/internal/core"

	tea "github.com/charmbracelet/bubbletea"
)

func Pick_and_run(searchPath string) {
	path, err := core.Normalize_path(searchPath)
	if err != nil {
		log.Fatal(err)
	}

	packageJson, err := core.Read_package_json(path)
	if err != nil {
		log.Fatal(err)
	}

	p := tea.NewProgram(initialModel(packageJson, path))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
