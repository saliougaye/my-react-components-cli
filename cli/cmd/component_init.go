/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"

	"github.com/saliougaye/my-react-components/helpers"
	"github.com/saliougaye/my-react-components/services"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// componentInitCmd represents the componentInit command
var componentInitCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize a new react component",
	Long:  `component init command create a folder structure for a new component to develop `,
	Run:   runComponentInitCommand,
}

func runComponentInitCommand(cmd *cobra.Command, args []string) {

	repoPath := helpers.InputString(helpers.InputContentString{
		Label:    "Repository Path",
		Validate: helpers.ValidateRepoDir,
	})

	componentName := helpers.InputString(helpers.InputContentString{
		Label:    "Component Name",
		Validate: helpers.ValidateComponentName,
	})

	if !viper.IsSet("token") {
		helpers.CheckError(errors.New("initialize cli first"))
	}

	token := viper.GetString("token")

	cliService := services.NewCliService(token)

	cliService.ComponentInit(componentName, repoPath)
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
