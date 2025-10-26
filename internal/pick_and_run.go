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

func PickAndRun(searchPath string) {
	path, err := core.NormalizePath(searchPath)
	if err != nil {
		log.Fatal(err)
	}

	packageJson, err := core.ReadPackageJson(path)
	if err != nil {
		log.Fatal(err)
	}

	pm := core.FindPackageManager(packageJson, path)

	p := tea.NewProgram(ui.InitialModel(packageJson, path, pm))
	modelOut, err := p.Run()

	if err != nil {
		log.Fatal("There's been an error", err)
	}

	model, ok := modelOut.(ui.Model)
	if !ok {
		log.Fatal("Unable to get the selected command")
	}

	name := model.GetSelected()
	if len(name) < 1 {
		os.Exit(0)
	}

	fmt.Printf("%s run %s\n", pm, name)

	cmd := exec.Command(pm, "run", name)
	cmd.Dir = filepath.Dir(path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

}
