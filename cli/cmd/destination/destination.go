/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package destination

import (
	"fmt"
	"github.com/ormushq/ormus/cli/cmd"

	"github.com/spf13/cobra"
)

// destinationCmd represents the destination command
var destinationCmd = &cobra.Command{
	Use:   "destination",
	Short: "Set Up and Manage Destinations",
	Long: `Set Up and Manage Destinations
ormus destination create --project-id <project-id> --name <destination-name> --type <destination-type>: Create a new destination within a project.
ormus destination list --project-id <project-id>: List all destinations for a specific project.
ormus destination show --project-id <project-id> --destination-id <destination-id>: Show details of a specific destination.
ormus destination update --project-id <project-id> --destination-id <destination-id> --name <new-name>: Update a destination's details.
ormus destination delete --project-id <project-id> --destination-id <destination-id>: Delete a destination.
ormus destination enable --project-id <project-id> --destination-id <destination-id>: Enable a destination
ormus destination disable --project-id <project-id> --destination-id <destination-id>: Disable a destination
ormus destination test-connection --project-id <project-id> --destination-id <destination-id>: Test the connection of a destination`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("destination called")
	},
}

func init() {
	cmd.RootCmd.AddCommand(destinationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// destinationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// destinationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
