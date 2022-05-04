package helpers

import (
	"encoding/json"

	"github.com/saliougaye/my-react-components/cli_types"
)

func CreateInitConfigFile(component string) string {

	config := cli_types.GetInitConfigFile(component)

	bytes, err := json.MarshalIndent(config, "", "\t")

	CheckError(err)

	return string(bytes)
}

func GetComponentReadmeInitContent(component string) string {

	return `# ` + component
}

func GetCliReadmeInitContent() string {
	return `
# My React Components
My Re-usable components, created from my-react-components CLI
`
}

func GetChangelogInitContent() string {

	return `
# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
`
}

func GetCliGitIgnoreContent() string {
	return `
node_modules
test_env
`
}
