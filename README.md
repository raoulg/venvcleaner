# VenvCleaner

A beautiful, colorful terminal UI for finding and removing Python virtual environments (.venv folders) in your git repositories.

![Version](https://img.shields.io/badge/version-1.0.0-blue)
![Platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux%20%7C%20Windows-green)

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

### Cross-compilation using Makefile

The included Makefile makes it easy to build for all platforms:

```bash
# Build for all platforms (creates binaries in dist/ directory)
make build-all

# Build for specific platform
make build-darwin-amd64   # macOS Intel
make build-darwin-arm64   # macOS Apple Silicon
make build-linux-amd64    # Linux
make build-windows-amd64  # Windows

# Install to GOPATH/bin
make install

# Clean build artifacts
make clean

# Show all available commands
make help
```

### Manual cross-compilation

Or build manually for different platforms:

```bash
# macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -o venvcleaner-darwin-amd64

# macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o venvcleaner-darwin-arm64

# Linux (amd64)
GOOS=linux GOARCH=amd64 go build -o venvcleaner-linux-amd64

# Windows (amd64)
GOOS=windows GOARCH=amd64 go build -o venvcleaner-windows-amd64.exe
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

## Requirements

- Go 1.21 or higher
- Works on: macOS, Linux, Windows, WSL
- Optional: [rip](https://github.com/nivekuil/rip) for safer deletion (Unix/macOS only)

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

## Distribution

### How to Distribute Your App

Once you've built your app, here's how to share it with others:

#### Option 1: Go Install (Recommended)

For users to install via `go install github.com/raoulg/venvcleaner@latest`:

1. **Push your code to GitHub**:
   ```bash
   git add .
   git commit -m "Initial release"
   git push origin main
   ```

2. **Tag a release**:
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```

That's it! Users can now run:
```bash
go install github.com/raoulg/venvcleaner@latest
```

#### Option 2: GitHub Releases (For Binary Distribution)

To provide pre-built binaries for users who don't have Go installed:

1. **Build binaries for all platforms**:
   ```bash
   make build-all
   ```
   This creates binaries in the `dist/` directory.

2. **Commit and tag your release**:
   ```bash
   git add .
   git commit -m "Release v1.0.0"
   git tag v1.0.0
   git push origin main
   git push origin v1.0.0
   ```

3. **Create a GitHub Release**:
   - Go to your GitHub repository
   - Click on "Releases" (right sidebar)
   - Click "Create a new release"
   - Select your tag (v1.0.0)
   - Fill in the release title and description
   - **Drag and drop the binaries** from `dist/` folder
   - Click "Publish release"

4. **Users can now download**:
   - They visit your releases page
   - Download the appropriate binary for their platform
   - Run it directly (no Go installation needed)

#### What's the Difference?

- **`go install`**: Users need Go installed. Simple for Go developers. Always gets latest code.
- **GitHub Releases**: Pre-built binaries. Anyone can download and run. Good for non-developers.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - feel free to use this however you'd like!

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

## Tips

- Run periodically to keep your disk clean
- Sort by time to find old, unused environments
- Sort by size to find the biggest space hogs
- Use with `rip` for peace of mind (deleted folders go to trash)
- Run from your main projects directory for best results

Enjoy your newly freed disk space!
