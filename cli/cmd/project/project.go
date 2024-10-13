/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package project

import (
	"fmt"

	"github.com/ormushq/ormus/cli/cmd"
	"github.com/spf13/cobra"
)

// projectCmd represents the project command.
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Create, View, and Manage Projects",
	Long: `Create, View, and Manage Projects
ormus project create --name <new-name> --description <new-description>: Create a new project.
ormus project list: List all projects associated with the user.
ormus project show --project-id <project-id>: Display details of a specific project.
ormus project update --project-id <project-id> --name <new-name> --description <new-description>: Update project details.
ormus project delete --project-id <project-id>: Delete a specific project.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("project called")
	},
}

func init() {
	cmd.RootCmd.AddCommand(projectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
