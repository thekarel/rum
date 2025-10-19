package core

import (
	"os"
	"path/filepath"
	"regexp"
)

// Find_package_manager returns the name of the JavaScript package manager used in the project
// pj is the parsed JSON
// filePath is the path to the JSON, so that lock files etc can be discovered.
func Find_package_manager(pj PackageJson, filePath string) string {
	if pj.PackageManager != "" {

		re := regexp.MustCompile(`@.*$`)
		pm := re.ReplaceAllString(pj.PackageManager, "")
		return pm
	}

	if pj.DevEngines.PackageManager.Name != "" {
		return pj.DevEngines.PackageManager.Name
	}

	dir := filepath.Dir(filePath)

	if _, err := os.Stat(filepath.Join(dir, "package-lock.json")); err == nil {
		return "npm"
	}

	if _, err := os.Stat(filepath.Join(dir, "pnpm-lock.yaml")); err == nil {
		return "pnpm"
	}

	if _, err := os.Stat(filepath.Join(dir, "yarn.lock")); err == nil {
		return "yarn"
	}

	if _, err := os.Stat(filepath.Join(dir, "bun.lockb")); err == nil {
		return "bun"
	}

	// npm is the default
	return "npm"
}
