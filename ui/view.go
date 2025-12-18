package ui

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/raoulg/venvcleaner/model"
)

// Synthwave/Cyberpunk Color Palette - Cohesive and Beautiful!
var (
	// Core palette colors
	neonPink    = lipgloss.Color("205") // Hot pink
	neonPurple  = lipgloss.Color("141") // Medium purple
	deepPurple  = lipgloss.Color("99")  // Deep purple
	neonCyan    = lipgloss.Color("51")  // Electric cyan
	neonTeal    = lipgloss.Color("87")  // Turquoise
	electricBlue = lipgloss.Color("81") // Electric blue
	neonYellow  = lipgloss.Color("227") // Neon yellow/gold
	darkBg      = lipgloss.Color("53")  // Dark purple background

	// Title - Neon pink gradient feel
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(neonPink).
			MarginBottom(1)

	// Headers - Electric cyan
	headerStyle = lipgloss.NewStyle().
			Foreground(neonCyan).
			Bold(true).
			MarginBottom(1)

	// Subheaders - Turquoise
	subheaderStyle = lipgloss.NewStyle().
			Foreground(neonTeal).
			Italic(true)

	// Selected items - Hot pink on dark purple
	selectedStyle = lipgloss.NewStyle().
			Foreground(neonPink).
			Background(darkBg).
			Bold(true)

	// Cursor - Electric cyan
	cursorStyle = lipgloss.NewStyle().
			Foreground(neonCyan).
			Bold(true)

	// Date colors - Purple to Pink gradient
	recentStyle = lipgloss.NewStyle().
			Foreground(neonTeal) // Teal for recent

	oldStyle = lipgloss.NewStyle().
			Foreground(neonPurple) // Purple for medium

	veryOldStyle = lipgloss.NewStyle().
			Foreground(neonPink) // Pink for old

	// Size colors - Teal to Yellow to Pink gradient
	sizeSmallStyle = lipgloss.NewStyle().
			Foreground(neonTeal). // Teal for small
			Bold(true)

	sizeMediumStyle = lipgloss.NewStyle().
			Foreground(electricBlue). // Blue for medium
			Bold(true)

	sizeLargeStyle = lipgloss.NewStyle().
			Foreground(neonPurple). // Purple for large
			Bold(true)

	sizeHugeStyle = lipgloss.NewStyle().
			Foreground(neonPink). // Pink for huge
			Bold(true)

	// Counters - Neon yellow (stands out)
	counterStyle = lipgloss.NewStyle().
			Foreground(neonYellow).
			Bold(true)

	// Paths - Electric blue
	pathStyle = lipgloss.NewStyle().
			Foreground(electricBlue)

	// Help text - Muted purple
	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("141")).
			Italic(true).
			MarginTop(1)

	// Footer - Neon yellow
	footerStyle = lipgloss.NewStyle().
			Foreground(neonYellow).
			Bold(true).
			MarginTop(1)

	// Progress bar - Cyan gradient
	progressBarStyle = lipgloss.NewStyle().
			Foreground(neonCyan)

	// Warning - Neon yellow
	warningStyle = lipgloss.NewStyle().
			Foreground(neonYellow).
			Bold(true)

	// Success - Teal
	successStyle = lipgloss.NewStyle().
			Foreground(neonTeal).
			Bold(true)

	// Separators - Deep purple
	separatorStyle = lipgloss.NewStyle().
			Foreground(deepPurple)

	// Icons - Hot pink
	iconStyle = lipgloss.NewStyle().
			Foreground(neonPink)

	// Box borders - Deep purple
	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(deepPurple).
			Padding(1, 2)

	// Accent styles matching palette
	accentCyan = lipgloss.NewStyle().
			Foreground(neonCyan).
			Bold(true)

	accentPink = lipgloss.NewStyle().
			Foreground(neonPink).
			Bold(true)

	accentYellow = lipgloss.NewStyle().
			Foreground(neonYellow).
			Bold(true)

	accentPurple = lipgloss.NewStyle().
			Foreground(neonPurple).
			Bold(true)
)

// View renders the UI
func (m Model) View() string {
	switch m.state {
	case model.StateScanning:
		return m.renderScanning()
	case model.StateSelecting:
		return m.renderSelecting()
	case model.StateConfirming:
		return m.renderConfirming()
	case model.StateCleaning:
		return m.renderCleaning()
	case model.StateDone:
		return m.renderDone()
	default:
		return "Unknown state"
	}
}

func (m Model) renderScanning() string {
	var s strings.Builder

	// Colorful title with gradient effect
	s.WriteString(titleStyle.Render(fmt.Sprintf("🔍 VenvCleaner v%s", m.version)))
	s.WriteString("\n")
	s.WriteString(accentCyan.Render("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"))
	s.WriteString("\n\n")

	s.WriteString(headerStyle.Render(fmt.Sprintf("%s Scanning for .venv folders...", m.spinner.View())))
	s.WriteString("\n\n")

	// Show current path being scanned with box
	currentPath := m.currentScanProg.CurrentPath
	if currentPath != "" {
		// Make path relative to start if possible
		if rel, err := filepath.Rel(m.startPath, currentPath); err == nil && !strings.HasPrefix(rel, "..") {
			currentPath = "./" + rel
		}

		// Truncate if too long
		maxLen := 60
		if len(currentPath) > maxLen {
			currentPath = "..." + currentPath[len(currentPath)-maxLen+3:]
		}

		s.WriteString(accentPink.Render("📂 ") + subheaderStyle.Render("Currently scanning:"))
		s.WriteString("\n")
		s.WriteString("   " + pathStyle.Render(currentPath))
		s.WriteString("\n\n")
	}

	// Show progress counters with colorful icons and numbers
	s.WriteString(accentYellow.Render("📁 ") +
		headerStyle.Render("Folders scanned: ") +
		counterStyle.Render(fmt.Sprintf("%d", m.currentScanProg.FoldersScanned)))
	s.WriteString("\n")
	s.WriteString(accentCyan.Render("✅ ") +
		headerStyle.Render("Repos with .venv: ") +
		successStyle.Render(fmt.Sprintf("%d", m.currentScanProg.ReposFound)))

	s.WriteString("\n\n")
	s.WriteString(accentCyan.Render("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"))
	s.WriteString("\n")
	s.WriteString(helpStyle.Render("💡 Press q to quit"))

	return s.String()
}

func (m Model) renderSelecting() string {
	if len(m.repos) == 0 {
		return titleStyle.Render(fmt.Sprintf("🔍 VenvCleaner v%s", m.version)) + "\n\n" +
			subheaderStyle.Render("No repositories with .venv folders found.") + "\n\n" +
			helpStyle.Render("💡 Press any key to exit")
	}

	var s strings.Builder

	// Title with decorative line
	s.WriteString(titleStyle.Render(fmt.Sprintf("🔍 VenvCleaner v%s", m.version)))
	s.WriteString("\n")
	s.WriteString(accentPink.Render("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"))
	s.WriteString("\n\n")
	s.WriteString(accentCyan.Render("🗂️  ") + headerStyle.Render("Select .venv folders to remove:"))
	s.WriteString("\n")

	// Sort mode indicator
	sortModeStr := ""
	switch m.sortMode {
	case model.SortByTime:
		sortModeStr = "Sorted by: Time (newest first)"
	case model.SortBySize:
		sortModeStr = "Sorted by: Size (largest first)"
	case model.SortByName:
		sortModeStr = "Sorted by: Name (A-Z)"
	}
	s.WriteString(headerStyle.Render(sortModeStr))
	s.WriteString("\n\n")

	// Calculate column widths for alignment
	pathWidth, dateWidth := m.calculateColumnWidths()

	// Render list of repos (with scrolling if needed)
	start, end := m.getVisibleRange()
	for i := start; i < end; i++ {
		s.WriteString(m.renderRepoLine(i, pathWidth, dateWidth))
		s.WriteString("\n")
	}

	// Footer with controls and summary
	s.WriteString("\n")
	s.WriteString(accentYellow.Render("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"))
	s.WriteString("\n")
	s.WriteString(accentPink.Render("📊 ") + footerStyle.Render(fmt.Sprintf(
		"Selected: %s/%s | Total size: %s",
		counterStyle.Render(fmt.Sprintf("%d", m.selectedCount())),
		counterStyle.Render(fmt.Sprintf("%d", len(m.repos))),
		successStyle.Render(formatSize(m.selectedSize())),
	)))
	s.WriteString("\n")
	s.WriteString(helpStyle.Render(
		"💡 ↑/↓: navigate | ⎵: toggle | ↵: confirm | t/s/n: sort | a/d: all/none | q: quit",
	))

	return s.String()
}

func (m Model) renderConfirming() string {
	var s strings.Builder

	s.WriteString(warningStyle.Render("⚠️  Confirm Deletion"))
	s.WriteString("\n\n")

	s.WriteString(headerStyle.Render("You are about to delete the following .venv folders:"))
	s.WriteString("\n\n")

	for _, repo := range m.repos {
		if repo.Selected {
			sizeColored := formatSize(repo.Size)
			if repo.Size >= 1024*1024*1024 {
				sizeColored = sizeHugeStyle.Render(sizeColored)
			} else if repo.Size >= 500*1024*1024 {
				sizeColored = sizeLargeStyle.Render(sizeColored)
			} else if repo.Size >= 50*1024*1024 {
				sizeColored = sizeMediumStyle.Render(sizeColored)
			} else {
				sizeColored = sizeSmallStyle.Render(sizeColored)
			}
			s.WriteString(fmt.Sprintf("  • %s (%s)\n", pathStyle.Render(repo.RepoPath), sizeColored))
		}
	}

	s.WriteString("\n")
	s.WriteString(footerStyle.Render(fmt.Sprintf(
		"Total: %s folders | %s",
		counterStyle.Render(fmt.Sprintf("%d", m.selectedCount())),
		warningStyle.Render(formatSize(m.selectedSize())),
	)))
	s.WriteString("\n\n")
	s.WriteString(helpStyle.Render("Are you sure? (y/N): "))

	return s.String()
}

func (m Model) renderCleaning() string {
	var s strings.Builder

	s.WriteString(headerStyle.Render("🧹 Cleaning..."))
	s.WriteString("\n\n")

	total := m.selectedCount()
	if total > 0 {
		percent := float64(m.cleanedCount) / float64(total)
		s.WriteString(m.progress.ViewAs(percent))
		s.WriteString("\n\n")
		s.WriteString(fmt.Sprintf(
			"%s %s/%s folders removed\n",
			successStyle.Render("Progress:"),
			counterStyle.Render(fmt.Sprintf("%d", m.cleanedCount)),
			counterStyle.Render(fmt.Sprintf("%d", total)),
		))
		s.WriteString(fmt.Sprintf("%s %s\n",
			successStyle.Render("Freed:"),
			footerStyle.Render(formatSize(m.totalCleaned))))
	}

	return s.String()
}

func (m Model) renderDone() string {
	var s strings.Builder

	// Big celebration header
	s.WriteString(successStyle.Render("✨ ✅ Done! ✅ ✨"))
	s.WriteString("\n")
	s.WriteString(accentCyan.Render("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"))
	s.WriteString("\n\n")

	if len(m.repos) == 0 {
		s.WriteString(accentYellow.Render("ℹ️  ") + subheaderStyle.Render("No repositories with .venv folders were found."))
		s.WriteString("\n")
	} else if m.cleanedCount > 0 {
		// Success message with colors
		s.WriteString(accentPink.Render("🎯 ") + headerStyle.Render(fmt.Sprintf(
			"Successfully removed %s .venv folders!",
			successStyle.Render(fmt.Sprintf("%d", m.cleanedCount)),
		)))
		s.WriteString("\n\n")

		// Big space freed announcement
		s.WriteString(accentYellow.Render("💾 ") + footerStyle.Render("Total space freed: "))
		s.WriteString(successStyle.Render(formatSize(m.totalCleaned)))
		s.WriteString("\n\n")

		// Celebration emojis
		s.WriteString(accentCyan.Render("🎉 🚀 ✨ 🎊 "))
		s.WriteString(headerStyle.Render("Your disk is cleaner!"))
		s.WriteString(accentCyan.Render(" 🎊 ✨ 🚀 🎉"))
	} else {
		s.WriteString(accentYellow.Render("ℹ️  ") + subheaderStyle.Render("No folders were removed."))
		s.WriteString("\n")
	}

	s.WriteString("\n\n")
	s.WriteString(accentPink.Render("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"))
	s.WriteString("\n")
	s.WriteString(helpStyle.Render("💡 Press any key to exit"))

	return s.String()
}

func (m Model) renderRepoLine(index int, pathWidth, dateWidth int) string {
	repo := m.repos[index]

	// Checkbox
	checkbox := "[ ]"
	if repo.Selected {
		checkbox = "[✓]"
	}

	// Cursor
	cursor := "  "
	if index == m.cursor {
		cursor = "→ "
	}

	// Path (shortened if needed)
	path := repo.RepoPath
	if home, err := filepath.Abs(m.startPath); err == nil {
		if rel, err := filepath.Rel(home, path); err == nil && !strings.HasPrefix(rel, "..") {
			path = "./" + rel
		}
	}

	// Truncate path if too long
	if len(path) > pathWidth {
		path = path[:pathWidth-3] + "..."
	}

	// Format date with color based on age
	dateStr := formatDate(repo.LastModified)
	plainDateStr := dateStr // Keep uncolored version for padding calculation
	age := time.Since(repo.LastModified)
	if age < 7*24*time.Hour {
		dateStr = recentStyle.Render(dateStr)
	} else if age < 30*24*time.Hour {
		dateStr = oldStyle.Render(dateStr)
	} else {
		dateStr = veryOldStyle.Render(dateStr)
	}

	// Size with color based on magnitude
	sizeStr := formatSize(repo.Size)

	const MB = 1024 * 1024
	const GB = 1024 * MB

	if repo.Size < 50*MB {
		sizeStr = sizeSmallStyle.Render(sizeStr)
	} else if repo.Size < 500*MB {
		sizeStr = sizeMediumStyle.Render(sizeStr)
	} else if repo.Size < GB {
		sizeStr = sizeLargeStyle.Render(sizeStr)
	} else {
		sizeStr = sizeHugeStyle.Render(sizeStr)
	}

	// Pad fields for alignment (using plain strings for width calculation)
	pathPadded := path + strings.Repeat(" ", pathWidth-len(path))
	datePadded := plainDateStr + strings.Repeat(" ", dateWidth-len(plainDateStr))
	// Replace plain date with colored version
	datePadded = dateStr + strings.Repeat(" ", dateWidth-len(plainDateStr))

	// Combine with aligned columns and colored separators
	separator := separatorStyle.Render(" │ ")
	line := fmt.Sprintf("%s%s %s%s%s%s%s",
		cursor,
		checkbox,
		pathPadded,
		separator,
		datePadded,
		separator,
		sizeStr,
	)

	// Highlight if selected or current
	if repo.Selected {
		// Apply selection style to the entire line except the colored parts
		parts := []string{
			selectedStyle.Render(cursor + checkbox + " " + pathPadded),
			separator,
			datePadded,
			separator,
			sizeStr,
		}
		line = strings.Join(parts, "")
	} else if index == m.cursor {
		// Apply cursor style to checkbox and path
		line = cursorStyle.Render(cursor+checkbox) + " " + pathPadded + separator + datePadded + separator + sizeStr
	}

	return line
}

func (m Model) getVisibleRange() (int, int) {
	// For now, show all repos. Could add pagination later.
	return 0, len(m.repos)
}

// calculateColumnWidths calculates the maximum width needed for each column
func (m Model) calculateColumnWidths() (pathWidth, dateWidth int) {
	pathWidth = 20  // minimum width
	dateWidth = 15  // minimum width

	for _, repo := range m.repos {
		// Calculate path width
		path := repo.RepoPath
		if rel, err := filepath.Rel(m.startPath, path); err == nil && !strings.HasPrefix(rel, "..") {
			path = "./" + rel
		}
		if len(path) > pathWidth {
			pathWidth = len(path)
		}

		// Calculate date width
		dateStr := formatDate(repo.LastModified)
		if len(dateStr) > dateWidth {
			dateWidth = len(dateStr)
		}
	}

	// Cap the path width to avoid overly long lines
	if pathWidth > 60 {
		pathWidth = 60
	}

	return pathWidth, dateWidth
}

// formatSize converts bytes to human-readable format
func formatSize(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

// formatDate formats a time in a human-readable way
func formatDate(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	if diff < 24*time.Hour {
		return "today"
	} else if diff < 48*time.Hour {
		return "yesterday"
	} else if diff < 7*24*time.Hour {
		return fmt.Sprintf("%d days ago", int(diff.Hours()/24))
	} else if diff < 30*24*time.Hour {
		return fmt.Sprintf("%d weeks ago", int(diff.Hours()/24/7))
	} else if diff < 365*24*time.Hour {
		return fmt.Sprintf("%d months ago", int(diff.Hours()/24/30))
	} else {
		return fmt.Sprintf("%d years ago", int(diff.Hours()/24/365))
	}
}
