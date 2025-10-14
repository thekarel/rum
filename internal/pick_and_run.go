package internal

import (
	"fmt"
	"log"

	"codeberg.org/thekarel/rum/internal/core"
)

func Pick_and_run() {
	scripts, err := core.Read_package_json("./package.json")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(scripts)
}
