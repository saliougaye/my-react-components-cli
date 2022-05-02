package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/spf13/viper"
)

type ghService struct {
	client HTTPClient
}

func NewGHService() ghService {
	return ghService{
		client: NewHTTPClient("https://github.com"),
	}
}

var ghHeader = "application/vnd.github.v3+json"

type ghIssue struct {
	title  string
	body   string
	labels []string
}

type ghIssueResponse struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}

func (gh ghService) CreateIssue(repoUrl, componentName string) (*ghIssueResponse, error) {

	urlParsed, err := url.ParseRequestURI(repoUrl)

	if err != nil {
		return nil, errors.New("url not valid")
	}

	owner, repo := getOwnerAndRepo(*urlParsed)

	issue := ghIssue{
		title:  "[CLI] Develop New Component: " + componentName,
		body:   "[CLI] Requested to develop new component " + componentName,
		labels: []string{"enhancement"},
	}

	path := fmt.Sprintf("/repos/%s/%s/issues", owner, repo)

	if !viper.IsSet("token") {
		return nil, errors.New("first init the cli with the GH token.\ntry ./my-react-components init")
	}

	res, err := gh.client.request(
		"POST",
		path,
		map[string]string{
			"Authorization": "Bearer " + viper.GetString("token"),
			"accept":        ghHeader,
		},
		issue,
	)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode == 401 {
		return nil, errors.New("not authorized")
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, errors.New("failed to read body response")
	}

	if res.StatusCode != 201 {
		return nil, errors.New("HTTP " + string(rune(res.StatusCode)) + ": " + string(body))
	}

	var issueCreated ghIssueResponse

	json.Unmarshal(body, &issueCreated)

	return &issueCreated, nil
}

func getOwnerAndRepo(inputUrl url.URL) (string, string) {
	escapedPath := inputUrl.EscapedPath()
	params := strings.Split(escapedPath, "/")

	owner := params[1]
	repo := params[2]

	return owner, repo
}

func (gh ghService) CreatePR() {}

func (gh ghService) AcceptPR() {}

func (g ghService) Merge() {}

func (gh ghService) Tag() {}

func (gh ghService) Read() {}
