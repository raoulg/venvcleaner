package main

import (
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rgrouls/venvcleaner/scanner"
	"github.com/rgrouls/venvcleaner/ui"
)

func main() {
	// Parse command line arguments
	startPath := "."
	if len(os.Args) > 1 {
		startPath = os.Args[1]
	}

	// Convert to absolute path
	absPath, err := filepath.Abs(startPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error resolving path: %v\n", err)
		os.Exit(1)
	}

	// Check if path exists
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Path does not exist: %s\n", absPath)
		os.Exit(1)
	}

	// Start scanning in background
	scanResults, scanProgress := scanner.ScanForVenvs(absPath)

	// Initialize Bubbletea program with full-screen mode
	model := ui.NewModel(absPath, scanResults, scanProgress, Version)
	p := tea.NewProgram(model, tea.WithAltScreen())

	// Run the program
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running program: %v\n", err)
		os.Exit(1)
	}
}
