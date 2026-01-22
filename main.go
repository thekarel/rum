package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/thekarel/rum/internal"
)

var helpText = `TUI to list, filter and run package.json scripts.

To list the scripts in the current folder:
  rum

You can also pass relative or absolute paths either to a folder or a file:
  rum ./modules/thing/
  rum /code/project/package.json

To show a non iteractive list:
  run -l /path/to/somewhere

Usage:

  rum [flags] <path to folder or package.json>

Flags:
  -h        display help
  -l        list scripts`

func help() {
	fmt.Fprintf(os.Stdout, "%s \n", helpText)
}

func main() {
	var helpFlag = flag.Bool("h", false, "display help")
	var listFlag = flag.Bool("l", false, "list scripts")

	flag.Parse()

	if *helpFlag {
		help()
		os.Exit(0)
	}

	path := "."

	if len(flag.Args()) == 1 {
		path = flag.Arg(0)
	}

	if *listFlag {
		internal.ListScripts(path)
	} else {
		internal.PickAndRun(path)
	}
}
