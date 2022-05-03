/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// componentLsCmd represents the componentLs command
var componentLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list all components",
	Long:  `component ls command list all the components available`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("component ls called")
	},
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
