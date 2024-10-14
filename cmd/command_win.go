//go:build windows
// +build windows

package main

import (
	"fmt"

	"github.com/inconshreveable/mousetrap"
)

func showWindowsWarning() {
	if mousetrap.StartedByExplorer() {
		fmt.Println(`This is a command line tool.

You need to open cmd.exe and run it from there.

Press return to continue...
`)
		fmt.Scanln()
	}
}
