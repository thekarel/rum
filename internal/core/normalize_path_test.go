package core

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNormalize_path(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: filepath.Join(cwd, "package.json"),
		},
		{
			name:     "current directory",
			input:    ".",
			expected: filepath.Join(cwd, "package.json"),
		},
		{
			name:     "relative path",
			input:    "subfolder",
			expected: filepath.Join(cwd, "subfolder", "package.json"),
		},
		{
			name:     "parent directory",
			input:    "..",
			expected: filepath.Join(filepath.Dir(cwd), "package.json"),
		},
		{
			name:     "absolute path",
			input:    "/tmp/myproject",
			expected: "/tmp/myproject/package.json",
		},
		{
			name:     "already package.json relative",
			input:    "myproject/package.json",
			expected: filepath.Join(cwd, "myproject", "package.json"),
		},
		{
			name:     "already package.json absolute",
			input:    "/tmp/project/package.json",
			expected: "/tmp/project/package.json",
		},
		{
			name:     "nested relative path",
			input:    "src/components",
			expected: filepath.Join(cwd, "src", "components", "package.json"),
		},
		{
			name:     "multiple parent directories",
			input:    "../../projects/myapp",
			expected: filepath.Join(filepath.Dir(filepath.Dir(cwd)), "projects", "myapp", "package.json"),
		},
		{
			name:     "path to a file is treated as a folder",
			input:    "/cat.jpg",
			expected: "/cat.jpg/package.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NormalizePath(tt.input)

			if err != nil {
				t.Errorf("Normalize_path() unexpected error: %v", err)
				return
			}

			if result != tt.expected {
				t.Errorf("Normalize_path() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
