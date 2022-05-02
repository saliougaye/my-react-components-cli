package services

import "github.com/go-git/go-git/v5"

type gitService struct {
}

func NewGitService() gitService {
	return gitService{}
}

func (g gitService) CreateBranch() {}

func (g gitService) Add() {}

func (g gitService) Commit() {}

func (g gitService) Push() {}

func (g gitService) Clone(repoUrl, directory string) (*git.Repository, error) {

	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:               repoUrl,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	return r, err

}

func (g gitService) Checkout() {}
