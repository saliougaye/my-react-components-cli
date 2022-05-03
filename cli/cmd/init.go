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
	Long:  `Initialize CLI with Github Authentication`,
	Run:   runInitCommand,
}

func runInitCommand(cmd *cobra.Command, args []string) {

	token, repoUrl, repoDir := initInput()

	cliService := services.NewCliService(token)

	cliService.Init(repoUrl, repoDir)

}

func initInput() (string, string, string) {
	token := helpers.InputString(helpers.InputContentString{
		Label:    "Github Access Token:",
		Validate: helpers.ValidateToken,
	})

	repoUrl := helpers.InputString(helpers.InputContentString{
		Label:    "Github Repository Url:",
		Validate: helpers.ValidateRepoUrl,
	})

	repoDir := helpers.InputString(helpers.InputContentString{
		Label:    "Where to clone? ",
		Validate: helpers.ValidateRepoDir,
	})

	return token, repoUrl, repoDir
}

func init() {
	rootCmd.AddCommand(initCmd)
}
