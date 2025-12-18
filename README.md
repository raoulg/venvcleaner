# VenvCleaner

A beautiful, colorful terminal UI for finding and removing Python virtual environments (.venv folders) in your git repositories.

![Version](https://img.shields.io/badge/version-1.0.1-blue)
![Platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux%20%7C%20Windows-green)

## Why This Tool?

Python virtual environments can take up significant disk space, especially if you work on many projects. It's easy to forget about old .venv folders sitting in abandoned or rarely-used projects. This tool makes it easy to:

- Find all virtual environments across your entire projects directory
- See which ones are old and unused
- Safely remove them in bulk
- Reclaim potentially gigabytes of disk space

## Example Output

```
🔍 VenvCleaner

Select .venv folders to remove:
Sorted by: Time (newest first)

  [ ] ./project1       │ 2 days ago    │ 234.5 MB
→ [✓] ./old-project   │ 3 months ago  │ 512.1 MB
  [✓] ./test-app      │ 1 year ago    │ 189.3 MB

Selected: 2/3 | Total size: 701.4 MB

↑/↓: navigate | space: toggle | enter: confirm | t/s/n: sort | a: select all | d: deselect all | q: quit
```

## Features

- **Full-screen TUI**: Immersive terminal experience that takes over your screen
- **Live scanning progress**: Watch in real-time as folders are scanned with live counters
- **Recursive scanning**: Finds all git repositories with .venv folders
- **Interactive selection**: Multi-select with visual feedback and smooth navigation
- **Smart sorting**: Sort by last modified time, size, or name with a single key press
- **Vibrant colors**: Color-coded by age (green=recent, yellow=old, red=very old) and size
- **Aligned table view**: Clean, professional table layout with proper column alignment
- **Safe deletion**: Confirmation screen showing exactly what will be deleted
- **Progress tracking**: Real-time progress bar and space freed counter
- **Cross-platform**: Works on macOS, Linux, and Windows
- **Smart removal**: Uses `rip` (trash) on Unix, `rm` fallback, or native Go on Windows

## Installation
## Requirements

- **Go 1.21 or higher** - [How to install Go](#installing-go)
- Works on: macOS, Linux, Windows, WSL
- Optional: [rip](https://github.com/nivekuil/rip) for safer deletion (Unix/macOS only)

### Installing Go

If you don't have Go installed:

**Official Installation:**
- Visit [go.dev/dl](https://go.dev/dl/)
- Download the installer for your platform
- Follow the installation instructions

### Using go install

```bash
go install github.com/raoulg/venvcleaner@latest
```

### Build from source

```bash
git clone https://github.com/raoulg/venvcleaner.git
cd venvcleaner
make build
```

Or using Go directly:

```bash
go build
```


## Usage

### Basic usage

Scan the current directory:

```bash
venvcleaner
```

### Scan a specific directory

```bash
venvcleaner ~/projects
```

### Keyboard Controls

#### Selection Mode
- `↑/↓` or `k/j`: Navigate up/down
- `space`: Toggle selection on current item
- `enter`: Proceed to confirmation (if any selected)
- `t`: Sort by time (newest first)
- `s`: Sort by size (largest first)
- `n`: Sort by name (alphabetical)
- `a`: Select all
- `d`: Deselect all
- `q`: Quit

#### Confirmation Mode
- `y` or `enter`: Confirm deletion
- `n` or `q`: Cancel and return to selection

#### Done Mode
- Any key: Exit

## How It Works

1. **Scanning**: Recursively walks the filesystem from the starting path, looking for directories containing a `.git` folder
2. **Filtering**: For each git repository, checks if a `.venv` folder exists
3. **Analysis**: Calculates the size and last modification time of each .venv folder
4. **Interactive Selection**: Presents a colorful list with sorting options
5. **Confirmation**: Shows a summary before deletion
6. **Deletion**: Uses `rip` if available (safer), otherwise falls back to `rm -rf`
7. **Progress**: Shows real-time progress and total space freed

## Color Coding

- **Green**: Recently modified (< 7 days ago)
- **Yellow**: Moderately old (7-30 days ago)
- **Red**: Very old (> 30 days ago)

This helps you identify which virtual environments are actively used vs. abandoned.

## Safety Features

- Only scans git repositories (prevents accidental deletion of system folders)
- Confirmation screen before deletion
- Shows exactly what will be deleted and how much space will be freed
- Graceful error handling (continues if one deletion fails)
- Uses `rip` (trash/recycle bin) when available

## Dependencies

- [charmbracelet/bubbletea](https://github.com/charmbracelet/bubbletea) - Terminal UI framework
- [charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss) - Style definitions
- [charmbracelet/bubbles](https://github.com/charmbracelet/bubbles) - Common Bubbletea components


**Quick Install (Platform-Specific):**

```bash
# macOS (using Homebrew)
brew install go

# Linux (Ubuntu/Debian)
sudo apt update
sudo apt install golang-go

# Windows (using Chocolatey)
choco install golang

# Or use winget
winget install GoLang.Go
```

**Verify Installation:**
```bash
go version
# Should show: go version go1.21.x or higher
```

**For `go install` to work**, make sure your `GOPATH/bin` is in your PATH:
```bash
# Add to your ~/.bashrc, ~/.zshrc, or equivalent
export PATH="$PATH:$(go env GOPATH)/bin"
```

## Platform-Specific Notes

### macOS / Linux / WSL
- Supports `rip` for safer deletion (sends files to trash)
- Falls back to `rm -rf` if `rip` is not installed
- Full terminal color support

### Windows
- Uses Go's built-in file removal (os.RemoveAll)
- Direct deletion (no trash/recycle bin integration)
- Works in PowerShell, CMD, or Windows Terminal
- Full terminal color support in Windows Terminal (limited in CMD)


## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - feel free to use this however you'd like!

## Tips

- Run periodically to keep your disk clean
- Sort by time to find old, unused environments
- Sort by size to find the biggest space hogs
- Use with `rip` for peace of mind (deleted folders go to trash)
- Run from your main projects directory for best results

Enjoy your newly freed disk space!
