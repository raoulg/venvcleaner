package scanner

import (
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/rgrouls/venvcleaner/model"
)

// FindGitRepos walks the filesystem starting from rootPath and finds all git repositories
func FindGitRepos(rootPath string) ([]string, error) {
	var repos []string

	err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			// Skip directories we can't access
			return nil
		}

		// Skip hidden directories except .git
		if d.IsDir() && d.Name()[0] == '.' && d.Name() != ".git" {
			return filepath.SkipDir
		}

		// If we find a .git directory, record its parent as a repo
		if d.IsDir() && d.Name() == ".git" {
			repoPath := filepath.Dir(path)
			repos = append(repos, repoPath)
			// Don't descend into .git directories
			return filepath.SkipDir
		}

		return nil
	})

	return repos, err
}

// CheckVenv checks if a repository has a .venv folder and returns info about it
func CheckVenv(repoPath string) (*model.VenvInfo, error) {
	venvPath := filepath.Join(repoPath, ".venv")

	// Check if .venv exists and is a directory
	info, err := os.Stat(venvPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil // No .venv, not an error
		}
		return nil, err
	}

	if !info.IsDir() {
		return nil, nil // .venv exists but is not a directory
	}

	// Check for pyproject.toml
	pyprojectPath := filepath.Join(repoPath, "pyproject.toml")
	_, err = os.Stat(pyprojectPath)
	hasPyproject := err == nil

	// Get venv size
	size, err := GetVenvSize(venvPath)
	if err != nil {
		size = 0 // If we can't calculate size, default to 0
	}

	// Get last modified time
	lastModified, err := GetLastModified(venvPath)
	if err != nil {
		lastModified = info.ModTime() // Fallback to venv dir modification time
	}

	return &model.VenvInfo{
		RepoPath:     repoPath,
		VenvPath:     venvPath,
		HasPyproject: hasPyproject,
		LastModified: lastModified,
		Size:         size,
		Selected:     false,
	}, nil
}

// GetVenvSize calculates the total size of a .venv directory
func GetVenvSize(venvPath string) (int64, error) {
	var size int64

	err := filepath.WalkDir(venvPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			// Skip files/dirs we can't access
			return nil
		}

		if !d.IsDir() {
			info, err := d.Info()
			if err == nil {
				size += info.Size()
			}
		}

		return nil
	})

	return size, err
}

// GetLastModified finds the most recently modified file in a .venv directory
func GetLastModified(venvPath string) (time.Time, error) {
	var lastModified time.Time

	err := filepath.WalkDir(venvPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			// Skip files/dirs we can't access
			return nil
		}

		info, err := d.Info()
		if err == nil {
			if info.ModTime().After(lastModified) {
				lastModified = info.ModTime()
			}
		}

		return nil
	})

	return lastModified, err
}

// ScanForVenvs scans a root path for git repos with .venv folders
// Returns two channels: one for results and one for progress updates
func ScanForVenvs(rootPath string) (<-chan *model.VenvInfo, <-chan model.ScanProgress) {
	results := make(chan *model.VenvInfo)
	progress := make(chan model.ScanProgress)

	go func() {
		defer close(results)
		defer close(progress)

		foldersScanned := 0
		reposFound := 0

		err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				// Skip directories we can't access
				return nil
			}

			// Skip hidden directories except .git
			if d.IsDir() && len(d.Name()) > 0 && d.Name()[0] == '.' && d.Name() != ".git" {
				return filepath.SkipDir
			}

			// Send progress update for each directory we scan
			if d.IsDir() {
				foldersScanned++
				progress <- model.ScanProgress{
					CurrentPath:    path,
					ReposFound:     reposFound,
					FoldersScanned: foldersScanned,
				}
			}

			// If we find a .git directory, check the parent for .venv
			if d.IsDir() && d.Name() == ".git" {
				repoPath := filepath.Dir(path)

				venvInfo, err := CheckVenv(repoPath)
				if err == nil && venvInfo != nil {
					// Found a repo with .venv
					reposFound++
					results <- venvInfo

					// Send updated progress
					progress <- model.ScanProgress{
						CurrentPath:    repoPath,
						ReposFound:     reposFound,
						FoldersScanned: foldersScanned,
					}
				}

				// Don't descend into .git directories
				return filepath.SkipDir
			}

			return nil
		})

		if err != nil {
			return
		}
	}()

	return results, progress
}
