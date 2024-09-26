/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package project

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ormushq/ormus/cli/cmd"
	"github.com/spf13/cobra"
)

// listCmd represents the list command.
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all projects associated with the user.",
	Long:  `ormus project list --last-token-id <last-token-id> --per-page <per-page>`,
	Run: func(cmdCobra *cobra.Command, args []string) {
		perPage, err := cmdCobra.Flags().GetString("per-page")
		if err != nil {
			fmt.Println("error on get per-page flag", err)

			return
		}
		lastTokenID, err := cmdCobra.Flags().GetString("last-token-id")
		if err != nil {
			fmt.Println("error on get last-token-id flag", err)

			return
		}
		resp, err := cmd.Client.SendRequest(cmd.Client.Project.List(perPage, lastTokenID))
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

		fmt.Printf("success response : \n%s\n", j)
	},
}

func init() {
	projectCmd.AddCommand(listCmd)

	listCmd.Flags().String("per-page", "10", "per-page")
	listCmd.Flags().String("last-token-id", "", "last-token-id")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
