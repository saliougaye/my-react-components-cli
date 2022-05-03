/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
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

	token, repoUrl, repoDir := initInput()

	cliService := services.NewCliService(token)

	cliService.Init(repoUrl, repoDir)

}

func initInput() (string, string, string) {
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

	return token, repoUrl, repoDir
}

func init() {
	rootCmd.AddCommand(initCmd)
}
