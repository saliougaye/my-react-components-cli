package services

import (
	"errors"
	"fmt"
	"path/filepath"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	githttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/saliougaye/my-react-components/helpers"
)

type cliService struct {
	token string
	auth  authService
	git   gitService
	fs    fsService
}

func NewCliService(token string) cliService {
	return cliService{
		token: token,
		auth:  *CreateAuthService(),
		git:   NewGitService(),
		fs:    NewFsService(),
	}
}

func (c cliService) Init(repoUrl, repoDir string) {
	c.initVerifyToken()

	r := c.initInitRepo(repoDir)

	c.initCreateRemote(r, repoUrl)

	workTree, err := r.Worktree()

	helpers.CheckError(err)

	c.createRepoStructure(repoDir)

	c.initGitAdd(workTree)

	c.initGitCommit(workTree)

	c.initGitCheckout(workTree)

	c.initDeleteOtherBranches(r, workTree)

	c.initGitPush(r)

}

func (c cliService) initVerifyToken() {
	loading := helpers.Loading("Check token if is valid...", "Token Verified ✅\n")

	loading.Start()

	time.Sleep(2 * time.Second)

	err := c.auth.IsTokenValid(c.token)

	helpers.CheckError(err)

	err = helpers.SaveInConfigFile("token", c.token)

	helpers.CheckError(err)

	loading.Stop()

}

func (c cliService) initCreateRemote(r *git.Repository, repoUrl string) {

	loading := helpers.Loading("Creating origin remote...", "Remote Created ✅\n")

	loading.Start()

	time.Sleep(2 * time.Second)

	_, err := r.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{repoUrl},
	})

	helpers.CheckError(err)

	loading.Stop()
}

func (c cliService) initInitRepo(repoDir string) *git.Repository {

	loading := helpers.Loading("Init Repository...", "Repository Initialized ✅\n")

	loading.Start()

	time.Sleep(2 * time.Second)

	r, err := c.git.Init(repoDir)

	helpers.CheckError(err)

	loading.Stop()

	return r

}

func (c cliService) createRepoStructure(repoDir string) {

	loading := helpers.Loading("Create Repository Structure...", "Repository Structure Created ✅\n")

	loading.Start()

	time.Sleep(2 * time.Second)

	folders := []string{"components"}

	files := map[string]string{
		"README.md": `
# My React Components
My Re-usable components, created from my-react-components CLI
`,
		".gitignore": `
node_modules
test_env
`,
		filepath.Join("components", "README.md"): `
# Components Folder
`,
	}

	for i := range folders {
		err := c.fs.CreateFolder(
			filepath.Join(repoDir, folders[i]),
		)

		if err != nil {
			errMsg := fmt.Sprintf("failed to create %s directory: %s", folders[i], err.Error())

			helpers.CheckError(errors.New(errMsg))
		}

	}

	for k, v := range files {

		err := c.fs.CreateFileWithContent(
			filepath.Join(repoDir, k),
			v,
		)

		if err != nil {
			errMsg := fmt.Sprintf("failed to create %s: %s", k, err.Error())

			helpers.CheckError(errors.New(errMsg))
		}

	}

	loading.Stop()
}

func (c cliService) initGitAdd(workTree *git.Worktree) {

	loading := helpers.Loading("Adding changes to staging area...", "Adding changes to staging area ✅\n")

	loading.Start()

	time.Sleep(2 * time.Second)

	err := workTree.AddGlob(".")

	helpers.CheckError(err)

	loading.Stop()
}

func (c cliService) initGitCommit(workTree *git.Worktree) {
	loading := helpers.Loading("Commit changes...", "Commit changes ✅\n")

	loading.Start()

	time.Sleep(2 * time.Second)

	_, err := workTree.Commit("[CLI] Initialize Repo", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "My React Components CLI",
			Email: "https://github.com/saliougaye/my-react-components-cli",
			When:  time.Now(),
		},
	})

	helpers.CheckError(err)

	loading.Stop()
}

func (c cliService) initGitCheckout(workTree *git.Worktree) {
	loading := helpers.Loading("Rename branch...", "Rename branch ✅\n")

	loading.Start()

	time.Sleep(2 * time.Second)

	err := workTree.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName("main"),
		Create: true,
	})

	helpers.CheckError(err)

	loading.Stop()
}

func (c cliService) initDeleteOtherBranches(r *git.Repository, workTree *git.Worktree) {

	loading := helpers.Loading("Delete other branches main ...", "Branches Deleted ✅\n")

	loading.Start()

	time.Sleep(2 * time.Second)

	b, err := r.Branches()

	helpers.CheckError(err)

	b.ForEach(func(b *plumbing.Reference) error {

		if b.Name().Short() != "main" {

			err := r.Storer.RemoveReference(b.Name())

			helpers.CheckError(err)
		}

		return nil
	})

	loading.Stop()
}

func (c cliService) initGitPush(r *git.Repository) {

	loading := helpers.Loading("Pushing changes to origin...", "Pushing changes to origin ✅\n")

	loading.Start()

	time.Sleep(2 * time.Second)

	err := r.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth: &githttp.BasicAuth{
			Username: "user",
			Password: c.token,
		},
	})

	helpers.CheckError(err)

	loading.Stop()
}
