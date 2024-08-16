/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package destination

import (
	"fmt"

	"github.com/spf13/cobra"
)

// testConnectionCmd represents the testConnection command
var testConnectionCmd = &cobra.Command{
	Use:   "testConnection",
	Short: "Test the connection of a destination",
	Long:  `ormus destination test-connection --project-id <project-id> --destination-id <destination-id>`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("testConnection called")
	},
}

func init() {
	destinationCmd.AddCommand(testConnectionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testConnectionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testConnectionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
