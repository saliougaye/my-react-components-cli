package services

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

type gitService struct {
}

func NewGitService() gitService {
	return gitService{}
}

func (g gitService) Init(directory string) (*git.Repository, error) {
	return git.PlainInit(directory, false)
}

func (g gitService) CreateBranch() {}

func (g gitService) Add() {}

func (g gitService) Commit() {}

func (g gitService) Push() {}

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

func (g gitService) Checkout() {}
