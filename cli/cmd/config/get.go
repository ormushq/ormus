/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package config

import (
	"fmt"
	"log"

	"github.com/ormushq/ormus/cli/cmd"
	"github.com/spf13/cobra"
)

// getCmd represents the get command.
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a specific configuration setting.",
	Long:  `ormus config get --key <key>`,
	Run: func(cmdCobra *cobra.Command, args []string) {
		key, err := cmdCobra.Flags().GetString("key")
		if err != nil {
			log.Fatal(err)
		}
		if key == "" {
			log.Fatal("Key is empty")
		}
		value, err := cmd.Client.GetConfig(key)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s => %s\n", key, value)
	},
}

func init() {
	configCmd.AddCommand(getCmd)
	getCmd.Flags().String("key", "", "key")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
