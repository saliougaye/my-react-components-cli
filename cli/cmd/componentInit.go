/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// componentInitCmd represents the componentInit command
var componentInitCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize a new react component",
	Long:  `component init command create a folder structure for a new component to develop `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("componentInit called")
	},
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
