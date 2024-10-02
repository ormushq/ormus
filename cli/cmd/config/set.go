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

// setCmd represents the set command.
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a configuration setting.",
	Long:  `ormus config set --key <key> --value <value>`,
	Run: func(cmdCobra *cobra.Command, args []string) {
		key, err := cmdCobra.Flags().GetString("key")
		if err != nil {
			log.Fatal(err)
		}
		value, err := cmdCobra.Flags().GetString("value")
		if err != nil {
			log.Fatal(err)
		}
		if key == "" || value == "" {
			log.Fatal("Key or value is empty")
		}
		err = cmd.Client.SetConfig(key, value)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Config updated successfully %s => %s \n", key, value)
	},
}

func init() {
	configCmd.AddCommand(setCmd)

	setCmd.Flags().String("key", "", "key")
	setCmd.Flags().String("value", "", "value")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
