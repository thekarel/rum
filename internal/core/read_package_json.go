package core

import (
	"encoding/json"
	"fmt"
	"os"
)

type packageJson struct {
	Scripts map[string]string
}

func Read_package_json(filePath string) (map[string]string, error) {
	fmt.Println("Reading package.json", filePath)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error reading the package.json at %s: %v", filePath, err)
	}

	var data packageJson

	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, fmt.Errorf("Parse error at %s: %v", filePath, err)
	}

	return data.Scripts, nil
}
