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

Usage:

  rum [flags] <path to folder or package.json>`

func help() {
	fmt.Fprintf(os.Stdout, "%s \n", helpText)
}

func main() {
	var helpFlag = flag.Bool("help", false, "display help")

	flag.Parse()

	if *helpFlag {
		help()
		os.Exit(0)
	}

	path := "."
	if len(flag.Args()) > 1 {
		fmt.Fprintf(os.Stdout, "Too many arguments")
		os.Exit(1)
	}

	if len(flag.Args()) == 1 {
		path = flag.Arg(0)
	}

	internal.PickAndRun(path)
}
