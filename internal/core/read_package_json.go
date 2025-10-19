package core

import (
	"encoding/json"
	"fmt"
	"os"
)

type PackageJson struct {
	// Scripts are the scripts to run
	Scripts map[string]string
	// Name of the project
	Name string
	// PackageManager might be set, example: "pnpm@9.15.6"
	PackageManager string
	// devEngines.packageManager.name example: "yarn" - this is the NPM standard
	DevEngines struct {
		PackageManager struct {
			Name string
		}
	}
}

func Read_package_json(filePath string) (PackageJson, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return PackageJson{}, fmt.Errorf("Error reading the package.json at %s: %v", filePath, err)
	}

	var data PackageJson

	err = json.Unmarshal(content, &data)
	if err != nil {
		return PackageJson{}, fmt.Errorf("Parse error at %s: %v", filePath, err)
	}

	return data, nil
}
