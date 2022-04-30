/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// componentCompletedCmd represents the componentCompleted command
var componentCompletedCmd = &cobra.Command{
	Use:   "completed",
	Short: "component completed",
	Long:  `component completed command execute the workflow when a react component develop its completed`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("component completed called")
	},
}

func init() {
	componentCmd.AddCommand(componentCompletedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// componentCompletedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// componentCompletedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
