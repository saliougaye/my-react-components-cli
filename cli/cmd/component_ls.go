/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/saliougaye/my-react-components/helpers"
	"github.com/saliougaye/my-react-components/services"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// componentLsCmd represents the componentLs command
var componentLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list all components",
	Long:  `component ls command list all the components available`,
	Run:   component_ls_run,
}

func component_ls_run(cmd *cobra.Command, args []string) {
	repo_path := helpers.InputString(helpers.InputContentString{
		Label:    "Repository Path",
		Validate: helpers.ValidateRepoDir,
	})

	if !viper.IsSet("token") {
		helpers.CheckError(errors.New("initialize cli first"))
	}

	token := viper.GetString("token")

	cli_service := services.NewCliService(token)

	configs := cli_service.ListComponents(repo_path)

	fmt.Printf("Total Components: %d\n", len(configs))

	for _, v := range configs {
		fmt.Printf("%s@%s\n", v.Name, v.Version)
	}

}

func init() {
	componentCmd.AddCommand(componentLsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// componentLsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// componentLsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
