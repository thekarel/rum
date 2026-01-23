package internal

import (
	"fmt"
	"log"

	"github.com/thekarel/rum/internal/core"
	"github.com/thekarel/rum/internal/ui"
)

func ListScripts(searchPath string) {
	path, err := core.NormalizePath(searchPath)
	if err != nil {
		log.Fatal(err)
	}

	packageJson, err := core.ReadPackageJson(path)
	if err != nil {
		log.Fatal(err)
	}

	pm := core.FindPackageManager(packageJson, path)

	model := ui.InitialModel(ui.ModelInitOpts{
		Pj:       packageJson,
		FilePath: path,
		Pm:       pm,
		WinWidth: 80,
		Readonly: true,
	})

	fmt.Println(model.View())
}
