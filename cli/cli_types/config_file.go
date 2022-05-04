package cli_types

import (
	"encoding/json"
)

type ConfigFile struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

func GetConfigFileFromJson(s string) (ConfigFile, error) {

	var configFile ConfigFile

	err := json.Unmarshal([]byte(s), &configFile)

	return configFile, err
}

func GetInitConfigFile(component string) ConfigFile {
	config := ConfigFile{
		Name:            component,
		Version:         "0.0.1",
		Dependencies:    map[string]string{},
		DevDependencies: map[string]string{},
	}

	return config
}
