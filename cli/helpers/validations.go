package helpers

import (
	"errors"
	"net/url"
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
