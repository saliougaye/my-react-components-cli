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
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize cli",
	Long:  `Initialize CLI with Github Authenthication`,
	Run:   runCommand,
}

var as = services.CreateAuthService()

func runCommand(cmd *cobra.Command, args []string) {

	token := helpers.InputString(helpers.InputContent{
		Label:    "GH Token:",
		Validate: validateToken,
	})

	err := as.IsTokenValid(token)

	if err != nil {
		helpers.PrintError(err)

		return
	}

	fmt.Printf("Authenthicated Successfully\n")

	helpers.SaveInConfigFile("token", token)
}

func validateToken(token string) error {

	if len(token) == 0 {
		return errors.New("please, provide the gh token")
	}

	return nil

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
