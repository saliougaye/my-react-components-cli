/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"time"

	"github.com/saliougaye/my-react-components/helpers"
	"github.com/saliougaye/my-react-components/services"
	"github.com/spf13/cobra"
)

// componentInitCmd represents the componentInit command
var componentInitCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize a new react component",
	Long:  `component init command create a folder structure for a new component to develop `,
	Run:   runComponentInitCommand,
}

func runComponentInitCommand(cmd *cobra.Command, args []string) {

	// ANCHOR input required
	// - repo url
	// - component name

	repoUrl := helpers.InputString(helpers.InputContent{
		Label:    "Github Repo URL",
		Validate: helpers.ValidateRepoUrl,
	})

	componentName := helpers.InputString(helpers.InputContent{
		Label:    "Component Name",
		Validate: helpers.ValidateComponentName,
	})

	loading := helpers.Loading("Creating Issue for "+componentName+" ", "Issue Created -> ")

	loading.Start()
	time.Sleep(4 * time.Second)

	ghService := services.NewGHService()
	// TODO create the issue
	ghIssue, err := ghService.CreateIssue(repoUrl, componentName)

	helpers.CheckError(err)

	loading.Stop()

	fmt.Printf("#%b ✅\n", ghIssue.Id)
	fmt.Printf("Issue Url: %s\n", ghIssue.Url)

	// TODO create the branch `feature/component name`

	// TODO initialize folder structure
}

func init() {
	componentCmd.AddCommand(componentInitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// componentInitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// componentInitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
