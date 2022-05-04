package helpers

import (
	"net/url"
	"strings"
)

func GetOwnerAndRepo(inputUrl url.URL) (string, string) {
	escapedPath := inputUrl.EscapedPath()
	params := strings.Split(escapedPath, "/")

	owner := params[1]
	repo := params[2]

	return owner, repo
}
