package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/saliougaye/my-react-components/cli_types"
	"github.com/saliougaye/my-react-components/helpers"
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

func (gh ghService) CreateIssue(repoUrl, componentName string) (*cli_types.GhIssueResponse, error) {

	urlParsed, err := url.ParseRequestURI(repoUrl)

	if err != nil {
		return nil, errors.New("url not valid")
	}

	owner, repo := helpers.GetOwnerAndRepo(*urlParsed)

	issue := cli_types.GhIssue{
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

	var issueCreated cli_types.GhIssueResponse

	json.Unmarshal(body, &issueCreated)

	return &issueCreated, nil
}

func (gh ghService) GetRepoFileList(repoUrl string) []cli_types.GhTree {
	urlParsed, err := url.ParseRequestURI(repoUrl)

	helpers.CheckError(err)

	owner, repo := helpers.GetOwnerAndRepo(*urlParsed)

	path := fmt.Sprintf("/repos/%s/%s/git/trees/main?recursive=1", owner, repo)

	res, err := gh.client.request(
		"GET",
		path,
		map[string]string{
			"Authorization": "Bearer " + gh.token,
			"accept":        ghHeader,
			"Content-Type":  "application/json",
		},
		map[string]string{},
	)

	helpers.CheckError(err)

	defer res.Body.Close()

	if res.StatusCode == 401 {
		helpers.CheckError(errors.New("not authorized"))
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		helpers.CheckError(errors.New("failed to read body response"))
	}

	if res.StatusCode != 200 {
		helpers.CheckError(errors.New("HTTP " + string(rune(res.StatusCode)) + ": " + string(body)))
	}

	var repoFileList cli_types.GhFileTreeResponse

	json.Unmarshal(body, &repoFileList)

	return repoFileList.Tree

}

func (gh ghService) GetFile(repoUrl string, file cli_types.GhTree) cli_types.GhFile {
	urlParsed, err := url.ParseRequestURI(repoUrl)

	helpers.CheckError(err)

	owner, repo := helpers.GetOwnerAndRepo(*urlParsed)

	path := fmt.Sprintf("/repos/%s/%s/git/blobs/%s", owner, repo, file.SHA)

	res, err := gh.client.request(
		"GET",
		path,
		map[string]string{
			"Authorization": "Bearer " + gh.token,
			"accept":        ghHeader,
			"Content-Type":  "application/json",
		},
		map[string]string{},
	)

	helpers.CheckError(err)

	defer res.Body.Close()

	if res.StatusCode == 401 {
		helpers.CheckError(errors.New("not authorized"))
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		helpers.CheckError(errors.New("failed to read body response"))
	}

	if res.StatusCode != 200 {
		helpers.CheckError(errors.New("HTTP " + string(rune(res.StatusCode)) + ": " + string(body)))
	}

	var repoFile cli_types.GhFile

	json.Unmarshal(body, &repoFile)

	return repoFile
}

func (gh ghService) CreatePR() {}

func (gh ghService) AcceptPR() {}

func (g ghService) Merge() {}

func (gh ghService) Tag() {}

func (gh ghService) Read() {}
