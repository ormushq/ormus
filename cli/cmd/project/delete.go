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

// deleteCmd represents the delete command.
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a specific project.",
	Long:  `ormus project delete --project-id <project-id>`,
	Run: func(cmdCobra *cobra.Command, args []string) {
		projectID, err := cmdCobra.Flags().GetString("project-id")
		if err != nil {
			fmt.Println("error on get project id flag", err)

			return
		}

		if projectID == "" {
			fmt.Println("project id is required")

			return
		}
		resp, err := cmd.Client.SendRequest(cmd.Client.Project.Delete(projectID))
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
	projectCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().String("project-id", "", "project-id")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
