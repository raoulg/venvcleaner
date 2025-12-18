package model

import "time"

// VenvInfo represents a Python virtual environment found in a git repository
type VenvInfo struct {
	RepoPath     string    // Path to the git repository
	VenvPath     string    // Path to the .venv folder
	HasPyproject bool      // Whether pyproject.toml exists in the repo
	LastModified time.Time // Most recent modification time in .venv
	Size         int64     // Total size of .venv in bytes
	Selected     bool      // Whether this venv is selected for deletion
}

// UIState represents the current state of the UI
type UIState int

const (
	StateScanning UIState = iota
	StateSelecting
	StateConfirming
	StateCleaning
	StateDone
)

// SortMode represents how the list should be sorted
type SortMode int

const (
	SortByTime SortMode = iota
	SortBySize
	SortByName
)

// Progress represents deletion progress
type Progress struct {
	Current int
	Total   int
	Size    int64
}

// ScanProgress represents scanning progress
type ScanProgress struct {
	CurrentPath  string // Path currently being scanned
	ReposFound   int    // Number of repos with .venv found so far
	FoldersScanned int  // Total folders scanned
}
