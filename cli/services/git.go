package services

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/saliougaye/my-react-components/helpers"
)

type gitService struct {
}

func NewGitService() gitService {
	return gitService{}
}

func (g gitService) Init(directory string) (*git.Repository, error) {
	return git.PlainInit(directory, false)
}

func (g gitService) GetRemotes(directory string) []*git.Remote {

	r, err := git.PlainOpen(directory)

	helpers.CheckError(err)

	list, err := r.Remotes()

	helpers.CheckError(err)

	return list

}

func (g gitService) GetRepo(path string) *git.Repository {

	r, err := git.PlainOpen(path)

	helpers.CheckError(err)

	return r
}

func (g gitService) Clone(token, repoUrl, directory string) (*git.Repository, error) {

	r, err := git.PlainClone(directory, true, &git.CloneOptions{
		URL:               repoUrl,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Auth: &http.BasicAuth{
			Username: "cli-user",
			Password: token,
		},
		Progress: os.Stdout,
	})

	return r, err

}
