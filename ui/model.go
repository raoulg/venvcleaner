package ui

import (
	"sort"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rgrouls/venvcleaner/model"
)

// Model represents the Bubbletea application state
type Model struct {
	repos           []model.VenvInfo
	cursor          int
	sortMode        model.SortMode
	state           model.UIState
	progress        progress.Model
	spinner         spinner.Model
	scanResults     <-chan *model.VenvInfo
	scanProgress    <-chan model.ScanProgress
	currentScanProg model.ScanProgress
	progressChan    chan model.Progress
	totalCleaned    int64
	cleanedCount    int
	err             error
	startPath       string
	version         string
}

// NewModel creates a new UI model
func NewModel(startPath string, scanResults <-chan *model.VenvInfo, scanProgress <-chan model.ScanProgress, version string) Model {
	s := spinner.New()
	s.Spinner = spinner.Dot

	p := progress.New(progress.WithDefaultGradient())

	return Model{
		repos:        []model.VenvInfo{},
		cursor:       0,
		sortMode:     model.SortByTime,
		state:        model.StateScanning,
		progress:     p,
		spinner:      s,
		scanResults:  scanResults,
		scanProgress: scanProgress,
		progressChan: make(chan model.Progress),
		startPath:    startPath,
		version:      version,
	}
}

// Init initializes the Bubbletea model
func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		waitForScanResult(m.scanResults),
		waitForScanProgress(m.scanProgress),
	)
}

// waitForScanResult waits for the next scan result from the channel
func waitForScanResult(scanResults <-chan *model.VenvInfo) tea.Cmd {
	return func() tea.Msg {
		result, ok := <-scanResults
		if !ok {
			// Channel closed, scanning is done
			return scanDoneMsg{}
		}
		return scanResultMsg{result}
	}
}

// waitForScanProgress waits for scan progress updates
func waitForScanProgress(scanProgress <-chan model.ScanProgress) tea.Cmd {
	return func() tea.Msg {
		progress, ok := <-scanProgress
		if !ok {
			// Channel closed
			return nil
		}
		return scanProgressMsg{progress}
	}
}

// waitForProgress waits for deletion progress updates
func waitForProgress(progressChan <-chan model.Progress) tea.Cmd {
	return func() tea.Msg {
		progress, ok := <-progressChan
		if !ok {
			// Channel closed, cleaning is done
			return cleanDoneMsg{}
		}
		return cleanProgressMsg{progress}
	}
}

// Messages
type scanResultMsg struct {
	result *model.VenvInfo
}

type scanProgressMsg struct {
	progress model.ScanProgress
}

type scanDoneMsg struct{}

type cleanProgressMsg struct {
	progress model.Progress
}

type cleanDoneMsg struct{}

// sortRepos sorts the repos based on the current sort mode
func (m *Model) sortRepos() {
	switch m.sortMode {
	case model.SortByTime:
		sort.Slice(m.repos, func(i, j int) bool {
			return m.repos[i].LastModified.After(m.repos[j].LastModified)
		})
	case model.SortBySize:
		sort.Slice(m.repos, func(i, j int) bool {
			return m.repos[i].Size > m.repos[j].Size
		})
	case model.SortByName:
		sort.Slice(m.repos, func(i, j int) bool {
			return m.repos[i].RepoPath < m.repos[j].RepoPath
		})
	}
}

// toggleSelection toggles the selection state of the current item
func (m *Model) toggleSelection() {
	if m.cursor < len(m.repos) {
		m.repos[m.cursor].Selected = !m.repos[m.cursor].Selected
	}
}

// selectedCount returns the number of selected repos
func (m *Model) selectedCount() int {
	count := 0
	for _, repo := range m.repos {
		if repo.Selected {
			count++
		}
	}
	return count
}

// selectedSize returns the total size of selected repos
func (m *Model) selectedSize() int64 {
	var size int64
	for _, repo := range m.repos {
		if repo.Selected {
			size += repo.Size
		}
	}
	return size
}
