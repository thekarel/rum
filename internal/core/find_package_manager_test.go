package core

import (
	"testing"
)

func Test_Find_package_manager(t *testing.T) {
	tests := []struct {
		name     string
		dir      string
		expected string
	}{
		{
			name:     "falls back to npm",
			dir:      "testing/empty",
			expected: "npm",
		},
		{
			name:     "reports npm when package-lock.json is present",
			dir:      "testing/package-lock",
			expected: "npm",
		},
		{
			name:     "reports pnpm when pnpm-lock.yaml is present",
			dir:      "testing/pnpm-lock",
			expected: "pnpm",
		},
		{
			name:     "reports yarn when yarn.lock is present",
			dir:      "testing/yarn-lock",
			expected: "yarn",
		},
		{
			name:     "reports bun when bun.lockb is present",
			dir:      "testing/bun-lockb",
			expected: "bun",
		},
		{
			name:     "packageManager field: pnpm",
			expected: "pnpm",
			dir:      "testing/package-manager-field-pnpm",
		},
		{
			name:     "packageManager field: yarn@1.2.3",
			expected: "yarn",
			dir:      "testing/package-manager-field-yarn",
		},
		{
			name:     "devEngines.packageManager.name field: bun",
			expected: "bun",
			dir:      "testing/dev-engines-bun",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pjPath, err := Normalize_path(tt.dir)
			if err != nil {
				t.Errorf("Normalize failed: %v", err)
			}

			pj, err := Read_package_json(pjPath)
			if err != nil {
				t.Errorf("Reading JSON failed: %v", err)
			}

			got := Find_package_manager(pj, pjPath)
			if got != tt.expected {
				t.Errorf("Find_package_manager(JSON: %+v, filePath: %v) = %v, want %v", pj, tt.dir, got, tt.expected)
			}
		})
	}
}
