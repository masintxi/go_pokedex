//go:build windows

package main

import (
	"os"

	"golang.org/x/sys/windows"
)

func init() {
	// Get the handle (reference) to standard output
	stdout := windows.Handle(os.Stdout.Fd())

	// Get current console mode settings
	var mode uint32
	windows.GetConsoleMode(stdout, &mode)

	// Add ENABLE_VIRTUAL_TERMINAL_PROCESSING to existing settings
	// This tells Windows to properly process ANSI color codes
	windows.SetConsoleMode(stdout, mode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
}
