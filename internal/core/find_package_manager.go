package core

import (
	"os"
	"path/filepath"
	"strings"
)

// FindPackageManager returns the name of the JavaScript package manager used in the project
// pj is the parsed JSON
// filePath is the path to the JSON, so that lock files etc can be discovered.
func FindPackageManager(pj PackageJson, filePath string) string {
	if pj.PackageManager != "" {

		pm := strings.Split(pj.PackageManager, "@")[0]
		if pm != "" {
			return pm
		}
	}

	if pj.DevEngines.PackageManager.Name != "" {
		return pj.DevEngines.PackageManager.Name
	}

	lockFiles := map[string]string{
		"package-lock.json": "npm",
		"pnpm-lock.yaml":   "pnpm",
		"yarn.lock":         "yarn",
		"bun.lockb":         "bun",
	}

	// Walk up the directory tree to find a lock file.
	// This handles monorepo subfolders where the lock file is at the root.
	dir := filepath.Dir(filePath)
	for {
		for lockFile, pm := range lockFiles {
			if _, err := os.Stat(filepath.Join(dir, lockFile)); err == nil {
				return pm
			}
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	// npm is the default
	return "npm"
}
