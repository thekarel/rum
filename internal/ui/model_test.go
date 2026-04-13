package ui

import (
	"testing"

	"github.com/thekarel/rum/internal/core"
)

func TestCopyRunCommand(t *testing.T) {
	tests := []struct {
		name     string
		pm       string
		subDir   string
		pkgName  string
		script   string
		expected string
	}{
		{
			name:     "same directory, no subDir",
			pm:       "pnpm",
			subDir:   "",
			pkgName:  "my-pkg",
			script:   "build",
			expected: "pnpm run build",
		},
		{
			name:     "pnpm subfolder with package name",
			pm:       "pnpm",
			subDir:   "packages/web",
			pkgName:  "@acme/web",
			script:   "dev",
			expected: "pnpm -F @acme/web run dev",
		},
		{
			name:     "npm subfolder with package name",
			pm:       "npm",
			subDir:   "packages/api",
			pkgName:  "@acme/api",
			script:   "start",
			expected: "npm -w @acme/api run start",
		},
		{
			name:     "yarn subfolder with package name",
			pm:       "yarn",
			subDir:   "packages/lib",
			pkgName:  "my-lib",
			script:   "test",
			expected: "yarn workspace my-lib run test",
		},
		{
			name:     "bun subfolder with package name",
			pm:       "bun",
			subDir:   "apps/web",
			pkgName:  "web-app",
			script:   "lint",
			expected: "bun --filter web-app run lint",
		},
		{
			name:     "subfolder without package name falls back to cd",
			pm:       "pnpm",
			subDir:   "packages/web",
			pkgName:  "",
			script:   "build",
			expected: "cd packages/web && pnpm run build",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scripts := map[string]string{tt.script: "some-command"}
			m := InitialModel(ModelInitOpts{
				Pj:       core.PackageJson{Name: tt.pkgName, Scripts: scripts},
				FilePath: "/tmp/package.json",
				Pm:       tt.pm,
				SubDir:   tt.subDir,
				WinWidth: 80,
			})

			got := m.CopyRunCommand()
			if got != tt.expected {
				t.Errorf("CopyRunCommand() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestCopyScriptCommand(t *testing.T) {
	tests := []struct {
		name     string
		subDir   string
		script   string
		cmd      string
		expected string
	}{
		{
			name:     "same directory",
			subDir:   "",
			script:   "build",
			cmd:      "tsc --build",
			expected: "tsc --build",
		},
		{
			name:     "subfolder prefixes cd and pm",
			subDir:   "packages/web",
			script:   "build",
			cmd:      "tsc --build",
			expected: "cd packages/web && pnpm exec tsc --build",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scripts := map[string]string{tt.script: tt.cmd}
			m := InitialModel(ModelInitOpts{
				Pj:       core.PackageJson{Scripts: scripts},
				FilePath: "/tmp/package.json",
				Pm:       "pnpm",
				SubDir:   tt.subDir,
				WinWidth: 80,
			})

			got := m.CopyScriptCommand()
			if got != tt.expected {
				t.Errorf("CopyScriptCommand() = %q, want %q", got, tt.expected)
			}
		})
	}
}
