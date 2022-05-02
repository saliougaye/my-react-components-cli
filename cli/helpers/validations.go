package helpers

import (
	"errors"
	"net/url"
	"os"
	"regexp"
)

func ValidateToken(token string) error {
	if len(token) == 0 {
		return errors.New("please, provide the gh token")
	}

	return nil
}

func ValidateRepoUrl(repoUrl string) error {
	u, err := url.ParseRequestURI(repoUrl)

	if err != nil {
		return errors.New("url is not valid")
	}

	if u.Hostname() != "github.com" {
		return errors.New("only github repos are valid")
	}

	return nil
}

func ValidateComponentName(componentName string) error {

	matched, err := regexp.MatchString("^[a-zA-Z]+$", componentName)

	if err != nil || !matched {
		return errors.New("component name is not valid")
	}

	return nil
}

func ValidateRepoDir(inputPath string) error {

	f, err := os.Stat(inputPath)
	if os.IsNotExist(err) {
		return errors.New("path not exist")
	}

	if !f.IsDir() {
		return errors.New("path is not a directory")
	}

	return nil
}
