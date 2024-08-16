/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package source

import (
	"fmt"

	"github.com/spf13/cobra"
)

// enableCmd represents the enable command
var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable a source",
	Long:  `ormus source enable --project-id <project-id> --source-id <source-id>`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("enable called")
	},
}

func init() {
	sourceCmd.AddCommand(enableCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// enableCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// enableCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
