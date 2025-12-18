package ui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/raoulg/venvcleaner/cleaner"
	"github.com/raoulg/venvcleaner/model"
)

// Update handles all UI events and state transitions
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch m.state {
		case model.StateScanning:
			// Allow quitting during scanning
			if msg.String() == "q" || msg.String() == "ctrl+c" {
				return m, tea.Quit
			}

		case model.StateSelecting:
			switch msg.String() {
			case "q", "ctrl+c":
				return m, tea.Quit

			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}

			case "down", "j":
				if m.cursor < len(m.repos)-1 {
					m.cursor++
				}

			case " ":
				m.toggleSelection()

			case "enter":
				// Only proceed if something is selected
				if m.selectedCount() > 0 {
					m.state = model.StateConfirming
				}

			case "t":
				m.sortMode = model.SortByTime
				m.sortRepos()
				m.cursor = 0

			case "s":
				m.sortMode = model.SortBySize
				m.sortRepos()
				m.cursor = 0

			case "n":
				m.sortMode = model.SortByName
				m.sortRepos()
				m.cursor = 0

			case "a":
				// Select all
				for i := range m.repos {
					m.repos[i].Selected = true
				}

			case "d":
				// Deselect all
				for i := range m.repos {
					m.repos[i].Selected = false
				}
			}

		case model.StateConfirming:
			switch msg.String() {
			case "y", "Y", "enter":
				// Start deletion
				m.state = model.StateCleaning
				return m, tea.Batch(
					startCleaning(m.repos, m.progressChan),
					waitForProgress(m.progressChan),
				)

			case "n", "N", "q", "ctrl+c":
				// Go back to selection
				m.state = model.StateSelecting
			}

		case model.StateDone:
			// Any key quits
			return m, tea.Quit
		}

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case scanResultMsg:
		// Add new repo to list
		m.repos = append(m.repos, *msg.result)
		m.sortRepos()
		// Wait for next result
		return m, waitForScanResult(m.scanResults)

	case scanProgressMsg:
		// Update current scan progress
		m.currentScanProg = msg.progress
		// Wait for next progress update
		return m, waitForScanProgress(m.scanProgress)

	case scanDoneMsg:
		// Scanning complete
		if len(m.repos) == 0 {
			// No repos found, go to done state with message
			m.state = model.StateDone
		} else {
			m.state = model.StateSelecting
		}

	case cleanProgressMsg:
		m.cleanedCount = msg.progress.Current
		m.totalCleaned = msg.progress.Size
		// Wait for next progress update
		return m, waitForProgress(m.progressChan)

	case cleanDoneMsg:
		m.state = model.StateDone
	}

	return m, nil
}

// startCleaning begins the deletion process in a goroutine
func startCleaning(repos []model.VenvInfo, progressChan chan model.Progress) tea.Cmd {
	return func() tea.Msg {
		go cleaner.DeleteSelected(repos, progressChan)
		return nil
	}
}
