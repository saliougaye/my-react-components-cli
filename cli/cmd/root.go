/*
Copyright © 2022 Saliou Gaye

*/
package cmd

import (
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/saliougaye/my-react-components/helpers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "my-react-components",
	Short: "My React Components CLI",
	Long: `My React Components CLI
My React Components CLI is a CLI written in Go to handle my re-usable react components.
The CLI permit to initalize, list, push to github and download components
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	home, err := homedir.Dir()

	if err != nil {
		helpers.PrintError(err)
		os.Exit(1)
	}

	configFile := ".myreactcomponents-config-cli.json"
	viper.SetConfigType("json")
	viper.SetConfigFile(home + "/" + configFile)

	if err := viper.ReadInConfig(); err != nil {
		helpers.PrintError(err)
		os.Exit(1)
	}
}
