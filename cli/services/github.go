package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

type ghService struct {
	client HTTPClient
	token  string
}

func NewGHService(token string) ghService {
	return ghService{
		client: NewHTTPClient("https://api.github.com"),
		token:  token,
	}
}

var ghHeader = "application/vnd.github.v3+json"

type ghIssue struct {
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	Labels []string `json:"labels"`
}

type ghIssueResponse struct {
	Id  int    `json:"number"`
	Url string `json:"html_url"`
}

func (gh ghService) CreateIssue(repoUrl, componentName string) (*ghIssueResponse, error) {

	urlParsed, err := url.ParseRequestURI(repoUrl)

	if err != nil {
		return nil, errors.New("url not valid")
	}

	owner, repo := getOwnerAndRepo(*urlParsed)

	issue := ghIssue{
		Title:  "[CLI] Develop New Component: " + componentName,
		Body:   "[CLI] Requested to develop new component " + componentName,
		Labels: []string{"enhancement"},
	}

	path := fmt.Sprintf("/repos/%s/%s/issues", owner, repo)

	res, err := gh.client.request(
		"POST",
		path,
		map[string]string{
			"Authorization": "Bearer " + gh.token,
			"accept":        ghHeader,
			"Content-Type":  "application/json",
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
