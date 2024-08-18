/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package destination

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd represents the create command.
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new destination within a project.",
	Long:  `ormus destination create --project-id <project-id> --name <destination-name> --type <destination-type>`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

func init() {
	destinationCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
