package services

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	githttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/saliougaye/my-react-components/cli_types"
	"github.com/saliougaye/my-react-components/helpers"
)

type cliService struct {
	token string
	auth  authService
	git   gitService
	fs    fsService
	gh    ghService
}

func NewCliService(token string) cliService {
	return cliService{
		token: token,
		auth:  *CreateAuthService(),
		git:   NewGitService(),
		fs:    NewFsService(),
		gh:    NewGHService(token),
	}
}

func (c cliService) Init(repoUrl, repoDir string) {

	// TODO get username of token user and use it in commit description
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

func (c cliService) ComponentInit(component, localRepoPath string) {

	remoteSelected := c.componentInitSelectRemote(localRepoPath)

	c.componentInitCreateIssue(component, remoteSelected)

	c.componentInitCreateFeatureBranch(component, localRepoPath, remoteSelected)

	c.componentInitCreateFolderStructure(component, localRepoPath, remoteSelected)

}

func (c cliService) ListComponents(localRepoPath string) []cli_types.ConfigFile {

	remoteSelected := c.componentInitSelectRemote(localRepoPath)

	loading := helpers.Loading("Fetching components", "Components fetched ✅\n")

	loading.Start()

	time.Sleep(2 * time.Second)

	repoUrl := remoteSelected.Config().URLs[0]

	file_tree := c.gh.GetRepoFileList(repoUrl)

	configs_file_in_components_folder := []cli_types.GhTree{}

	for _, v := range file_tree {

		if strings.HasPrefix(v.Path, "components/") && strings.HasSuffix(v.Path, "config.json") {
			configs_file_in_components_folder = append(configs_file_in_components_folder, v)
		}
	}

	components_configs := []cli_types.ConfigFile{}

	for _, v := range configs_file_in_components_folder {
		file := c.gh.GetFile(repoUrl, v)

		contentJson := helpers.ConvertBase64ToString(file.Content)

		component_config, err := cli_types.GetConfigFileFromJson(contentJson)

		helpers.CheckError(err)

		components_configs = append(components_configs, component_config)

	}

	loading.Stop()

	return components_configs

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

	folders := []string{
		filepath.Join(repoDir, "components"),
	}

	files := map[string]string{
		filepath.Join(repoDir, "README.md"):               helpers.GetCliReadmeInitContent(),
		filepath.Join(repoDir, ".gitignore"):              helpers.GetCliGitIgnoreContent(),
		filepath.Join(repoDir, "components", "README.md"): helpers.GetComponentReadmeInitContent("Components Folder"),
	}

	c.fs.CreateFolders(folders)

	c.fs.CreateFilesWithContent(helpers.GetFileInputFromMap(files))

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

func (c cliService) componentInitSelectRemote(localRepoPath string) *git.Remote {

	remotes := c.git.GetRemotes(localRepoPath)

	items := []helpers.InputContentSelectItem{}

	for _, v := range remotes {
		items = append(items, helpers.InputContentSelectItem{
			Name:   v.Config().Name,
			Detail: v.Config().URLs[0],
		})
	}

	remoteKeySelected := helpers.InputSelect(helpers.InputContentSelect{
		Label: "Choose the remote: ",
		Items: items,
	})

	return remotes[remoteKeySelected]
}

func (c cliService) componentInitCreateIssue(component string, remoteSelected *git.Remote) {
	loading := helpers.Loading("Creating Issue for "+component+" ", "Issue Created: ")

	loading.Start()
	time.Sleep(2 * time.Second)

	ghIssue, err := c.gh.CreateIssue(remoteSelected.Config().URLs[0], component)

	helpers.CheckError(err)

	loading.Stop()

	fmt.Printf("#%d -> %s %s\n", ghIssue.Id, ghIssue.Url, " ✅")

}

func (c cliService) componentInitCreateFeatureBranch(component, localRepoPath string, remoteSelected *git.Remote) {
	loading := helpers.Loading("Creating feature/"+component+" branch", "Creating feature/"+component+" branch ✅")

	loading.Start()
	time.Sleep(2 * time.Second)

	featureBranchName := fmt.Sprintf("feature/%s", component)
	r := c.git.GetRepo(localRepoPath)

	workTree, err := r.Worktree()

	helpers.CheckError(err)

	err = workTree.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(featureBranchName),
		Create: true,
	})

	helpers.CheckError(err)

	loading.Stop()

}

func (c cliService) componentInitCreateFolderStructure(component, localRepoPath string, remoteSelected *git.Remote) {
	loading := helpers.Loading("Creating "+component+" folder structure", "Creating "+component+" folder structure ✅")

	loading.Start()
	time.Sleep(2 * time.Second)

	folders := []string{
		filepath.Join(localRepoPath, "components", component),
		filepath.Join(localRepoPath, "components", component, "src"),
		filepath.Join(localRepoPath, "components", component, "test_env"),
	}

	files := map[string]string{
		filepath.Join(localRepoPath, "components", component, "CHANGELOG.md"): helpers.GetChangelogInitContent(),
		filepath.Join(localRepoPath, "components", component, "README.md"):    helpers.GetComponentReadmeInitContent(component),
		filepath.Join(localRepoPath, "components", component, "config.json"):  helpers.CreateInitConfigFile(component),
	}

	c.fs.CreateFolders(folders)

	c.fs.CreateFilesWithContent(
		helpers.GetFileInputFromMap(files),
	)

	// TODO create ts files

	// TODO run npx create-react-app to test_env folder

	loading.Stop()
}
