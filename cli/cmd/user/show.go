/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package user

import (
	"fmt"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display details of the logged-in user.",
	Long:  `ormus user show`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("show called")
	},
}

func init() {
	userCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
