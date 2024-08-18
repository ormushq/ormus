/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package source

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getWriteKeyCmd represents the getWriteKey command.
var getWriteKeyCmd = &cobra.Command{
	Use:   "getWriteKey",
	Short: "Get write-key of a source",
	Long:  `ormus source get-write-key --project-id <project-id> --source-id <source-id>`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getWriteKey called")
	},
}

func init() {
	sourceCmd.AddCommand(getWriteKeyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getWriteKeyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getWriteKeyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
