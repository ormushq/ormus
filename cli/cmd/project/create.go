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

// createCmd represents the create command.
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project.",
	Long:  `ormus project create --name <project-name> --description <project-description>`,
	Run: func(cmdCobra *cobra.Command, args []string) {
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

		if name == "" || description == "" {
			fmt.Println("name and description is required")

			return
		}

		resp, err := cmd.Client.SendRequest(cmd.Client.Project.Create(name, description))
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
	projectCmd.AddCommand(createCmd)

	createCmd.Flags().String("name", "", "name")
	createCmd.Flags().String("description", "", "description")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
