/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package project

import (
	"fmt"
	"github.com/ormushq/ormus/cli/cmd"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
)

// listCmd represents the list command.
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all projects associated with the user.",
	Long:  `ormus project list`,
	Run: func(cmdCobra *cobra.Command, args []string) {
		perPage, _ := cmdCobra.Flags().GetString("per-page")

		lastTokenId, _ := cmdCobra.Flags().GetString("last-token-id")

		resp, err := cmd.Client.SendRequest(cmd.Client.Project.List(perPage, lastTokenId))
		if err != nil {
			log.Fatal(err)
		}
		j, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode != http.StatusOK {
			log.Fatal(fmt.Errorf("status not OK ,status code %d, body: %s", resp.StatusCode, j))
		}

		fmt.Printf("success response : \n %s\n", j)

	},
}

func init() {
	projectCmd.AddCommand(listCmd)

	createCmd.Flags().String("per-page", "10", "per-page")
	createCmd.Flags().String("last-token-id", "", "last-token-id")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
