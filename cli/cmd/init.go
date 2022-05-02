/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"path"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/saliougaye/my-react-components/helpers"
	"github.com/saliougaye/my-react-components/services"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize cli",
	Long:  `Initialize CLI with Github Authenthication`,
	Run:   runInitCommand,
}

var as = services.CreateAuthService()
var gitService = services.NewGitService()
var fsService = services.NewFsService()

func runInitCommand(cmd *cobra.Command, args []string) {

	token := helpers.InputString(helpers.InputContent{
		Label:    "GH Token:",
		Validate: helpers.ValidateToken,
	})

	repoUrl := helpers.InputString(helpers.InputContent{
		Label:    "Repo Url:",
		Validate: helpers.ValidateRepoUrl,
	})

	repoDir := helpers.InputString(helpers.InputContent{
		Label:    "Where to clone? ",
		Validate: helpers.ValidateRepoDir,
	})

	err := as.IsTokenValid(token)

	if err != nil {
		helpers.PrintError(err)

		return
	}

	fmt.Printf("Authenthicated Successfully\n")

	helpers.SaveInConfigFile("token", token)

	r, err := gitService.Clone(repoUrl, repoDir)

	if err != nil {
		helpers.PrintError(errors.New("failed to clone repo: " + err.Error()))

		return
	}

	createRepoStructure(repoDir)

	workTree, err := r.Worktree()

	if err != nil {
		helpers.PrintError(errors.New("failed to git add: " + err.Error()))

		return
	}

	err = workTree.AddGlob(".")

	if err != nil {
		helpers.PrintError(errors.New("failed to git add: " + err.Error()))

		return
	}

	_, err = workTree.Commit("[CLI] Initialize Repo", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "My React Components CLI",
			Email: "https://github.com/saliougaye/my-react-components-cli",
			When:  time.Now(),
		},
	})

	if err != nil {
		helpers.PrintError(errors.New("failed to git commit: " + err.Error()))

		return
	}

	err = workTree.Checkout(&git.CheckoutOptions{
		Branch: "main",
		Create: true,
	})

	if err != nil {
		helpers.PrintError(errors.New("failed to create branch main: " + err.Error()))

		return
	}

	err = r.Push(&git.PushOptions{
		RemoteName: "main",
	})

	if err != nil {
		helpers.PrintError(errors.New("failed to push branch main: " + err.Error()))

		return
	}

}

func createRepoStructure(repoDir string) {

	files := map[string]string{
		"README.md": `
# My React Components
My Re-usable components, created from my-react-components CLI
`,
		".gitignore": `
node_modules
test_env
`,
	}

	folders := []string{"components"}

	for i := range folders {
		err := fsService.CreateFolder(folders[i])

		if err != nil {
			errMsg := fmt.Sprintf("failed to create %s directory: %s", folders[i], err.Error())

			helpers.PrintError(errors.New(errMsg))
		}
	}

	for k, v := range files {

		err := fsService.CreateFileWithContent(
			path.Join(repoDir, k),
			v,
		)

		if err != nil {
			errMsg := fmt.Sprintf("failed to create %s: %s", k, err.Error())

			helpers.PrintError(errors.New(errMsg))
		}

	}

}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
