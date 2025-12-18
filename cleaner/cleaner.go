package cleaner

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/raoulg/venvcleaner/model"
)

// DetectRemovalTool checks if 'rip' is available, otherwise falls back to 'rm' or 'native'
func DetectRemovalTool() string {
	// On Windows, use native Go removal
	if runtime.GOOS == "windows" {
		return "native"
	}

	// On Unix-like systems, check if rip is available
	_, err := exec.LookPath("rip")
	if err == nil {
		return "rip"
	}
	return "rm"
}

// DeleteVenv removes a .venv directory using the specified tool
func DeleteVenv(venvPath string, tool string) error {
	switch tool {
	case "native":
		// Use Go's built-in removal (cross-platform, works on Windows)
		err := os.RemoveAll(venvPath)
		if err != nil {
			return fmt.Errorf("failed to delete %s: %w", venvPath, err)
		}
		return nil

	case "rip":
		// rip handles directories without flags (safer, sends to trash)
		cmd := exec.Command("rip", venvPath)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("failed to delete %s: %w (output: %s)", venvPath, err, string(output))
		}
		return nil

	case "rm":
		// rm requires -rf for recursive directory removal (Unix/Linux/macOS)
		cmd := exec.Command("rm", "-rf", venvPath)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("failed to delete %s: %w (output: %s)", venvPath, err, string(output))
		}
		return nil

	default:
		return fmt.Errorf("unknown removal tool: %s", tool)
	}
}

// DeleteSelected removes all selected .venv folders and sends progress updates
func DeleteSelected(repos []model.VenvInfo, progressChan chan<- model.Progress) error {
	tool := DetectRemovalTool()

	// Filter only selected repos
	var selected []model.VenvInfo
	for _, repo := range repos {
		if repo.Selected {
			selected = append(selected, repo)
		}
	}

	if len(selected) == 0 {
		close(progressChan)
		return nil
	}

	total := len(selected)
	var totalSize int64

	for i, repo := range selected {
		// Delete the .venv
		err := DeleteVenv(repo.VenvPath, tool)
		if err != nil {
			// Log error but continue with remaining deletions
			fmt.Fprintf(os.Stderr, "Error deleting %s: %v\n", repo.VenvPath, err)
			continue
		}

		totalSize += repo.Size

		// Send progress update
		progressChan <- model.Progress{
			Current: i + 1,
			Total:   total,
			Size:    totalSize,
		}
	}

	close(progressChan)
	return nil
}
