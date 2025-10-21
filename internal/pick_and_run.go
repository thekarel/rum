package internal

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"codeberg.org/thekarel/rum/internal/core"
	"codeberg.org/thekarel/rum/internal/ui"

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

	pm := core.Find_package_manager(packageJson, path)

	p := tea.NewProgram(ui.InitialModel(packageJson, path, pm))
	modelOut, err := p.Run()

	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	model, ok := modelOut.(ui.Model)
	if ok != true {
		fmt.Printf("Unable to get the selected command")
		os.Exit(1)
	}

	name := model.GetSelected()
	if len(name) < 1 {
		os.Exit(0)
	}

	fmt.Printf("%s run %s\n", pm, name)

	cmd := exec.Command(pm, "run", name)
	cmd.Dir = filepath.Dir(path)
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

}
