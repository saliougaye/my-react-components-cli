/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// componentDownloadCmd represents the componentDownload command
var componentDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "use a component in the project",
	Long:  `clone one of my react component `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("component download called")
	},
}

func init() {
	componentCmd.AddCommand(componentDownloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// componentDownloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// componentDownloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
