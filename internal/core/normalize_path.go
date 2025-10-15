package core

import (
	"path/filepath"
	"strings"
)

// Normalize_path takes an absolute or relative path to a folder or a package.json file and returns the absolute
// path that points to a package.json.
//
// It does not verify if the file exists.
//
// Examples:
//
// ../.. -> /path/to/package.json
// /project -> /project/package.json
// "" -> /current/folder/package.json
// "../package.json" -> /path/above/package.json
func Normalize_path(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	if strings.HasSuffix(absPath, "package.json") {
		return absPath, nil
	}

	return filepath.Join(absPath, "package.json"), nil
}
