package core

import (
	"encoding/json"
	"fmt"
	"os"
)

type PackageJson struct {
	Scripts map[string]string
	Name    string
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
