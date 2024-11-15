/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package source

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ormushq/ormus/cli/cmd"
	"github.com/spf13/cobra"
)

// createCmd represents the create command.
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new source within a project.",
	Long:  `ormus source create --project-id <project-id> --name <source-name> --description <description>`,
	Run: func(cmdCobra *cobra.Command, args []string) {
		projectID, err := cmdCobra.Flags().GetString("project-id")
		if err != nil {
			fmt.Println("error on get project-id flag", err)

			return
		}
		name, err := cmdCobra.Flags().GetString("name")
		if err != nil {
			fmt.Println("error on get name flag", err)

			return
		}

		description, err := cmdCobra.Flags().GetString("description")
		if err != nil {
			fmt.Println("error on get description flag", err)

			return
		}

		if name == "" || description == "" || projectID == "" {
			fmt.Println("name and description and project id is required")

			return
		}

		resp, err := cmd.Client.SendRequest(cmd.Client.Source.Create(name, description, projectID))
		if err != nil {
			log.Fatal(err)
		}
		j, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode != http.StatusCreated {
			log.Fatal(fmt.Errorf("status not Created ,status code %d, body: %s", resp.StatusCode, j))
		}

		fmt.Printf("success response : \n%s\n", j)
	},
}

func init() {
	sourceCmd.AddCommand(createCmd)

	createCmd.Flags().String("name", "", "name")
	createCmd.Flags().String("description", "", "description")
	createCmd.Flags().String("project-id", "", "project-id")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
