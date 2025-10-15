package internal

import (
	"fmt"
	"log"

	"codeberg.org/thekarel/rum/internal/core"
)

func Pick_and_run(searchPath string) {
	path, err := core.Normalize_path(searchPath)
	if err != nil {
		log.Fatal(err)
	}

	scripts, err := core.Read_package_json(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(scripts)
}
