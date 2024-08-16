/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package source

import (
	"fmt"
	"github.com/ormushq/ormus/cli/cmd"

	"github.com/spf13/cobra"
)

// sourceCmd represents the source command
var sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Define and Manage Sources",
	Long: `Define and Manage Sources
ormus source create --project-id <project-id> --name <source-name> --type <source-type>: Create a new source within a project.
ormus source list --project-id <project-id>: List all sources for a specific project.
ormus source show --project-id <project-id> --source-id <source-id>: Show details of a specific source.
ormus source update --project-id <project-id> --source-id <source-id> --name <new-name>: Update a source's details.
ormus source delete --project-id <project-id> --source-id <source-id>: Delete a source.
ormus source enable --project-id <project-id> --source-id <source-id>: Enable a source
ormus source disable --project-id <project-id> --source-id <source-id>: Disable a source
ormus source get-write-key --project-id <project-id> --source-id <source-id>: Get write-key of a source
ormus source rotate-write-key --project-id <project-id> --source-id <source-id>: Rotate write-key for a source
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("source called")
	},
}

func init() {
	cmd.RootCmd.AddCommand(sourceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sourceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sourceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
