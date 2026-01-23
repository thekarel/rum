package ui

import (
	"fmt"
	"os"

	"github.com/aymanbagabas/go-osc52/v2"
)

func CopyToClipboard(text string) {
	fmt.Fprint(os.Stderr, osc52.New(text))
}
