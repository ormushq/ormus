/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package project

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command.
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a specific project.",
	Long:  `ormus project delete --project-id <project-id>`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
	},
}

func init() {
	projectCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
