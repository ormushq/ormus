/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package source

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rotateWriteKeyCmd represents the rotateWriteKey command.
var rotateWriteKeyCmd = &cobra.Command{
	Use:   "rotateWriteKey",
	Short: "Rotate write-key for a source",
	Long:  `ormus source rotate-write-key --project-id <project-id> --source-id <source-id>`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rotateWriteKey called")
	},
}

func init() {
	sourceCmd.AddCommand(rotateWriteKeyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rotateWriteKeyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rotateWriteKeyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
