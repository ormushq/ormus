/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package config

import (
	"fmt"
	"github.com/ormushq/ormus/cli/cmd"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: " List all configuration settings.",
	Long:  `ormus config list`,
	Run: func(cmdCobra *cobra.Command, args []string) {
		config, err := cmd.Client.ListConfig()
		if err != nil {
			panic(err)
		}
		for k, v := range config {
			fmt.Printf("%s => %s\n", k, v)
		}

	},
}

func init() {
	configCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
