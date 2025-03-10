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

// updateCmd represents the update command.
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a source's details.",
	Long:  `ormus source update --source-id <source-id> --name <new-name> --description <new-description>`,
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

		sourceID, err := cmdCobra.Flags().GetString("source-id")
		if err != nil {
			fmt.Println("error on get source id flag", err)

			return
		}

		if name == "" || description == "" || sourceID == "" {
			fmt.Println("name and description and source id is required")

			return
		}
		resp, err := cmd.Client.SendRequest(cmd.Client.Source.Update(sourceID, name, description))
		if err != nil {
			log.Fatal(err)
		}
		j, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode != http.StatusOK {
			log.Fatal(fmt.Errorf("status not Ok ,status code %d, body: %s", resp.StatusCode, j))
		}

		fmt.Printf("success response : \n%s\n", j)
	},
}

func init() {
	sourceCmd.AddCommand(updateCmd)

	updateCmd.Flags().String("name", "", "name")
	updateCmd.Flags().String("description", "", "description")
	updateCmd.Flags().String("source-id", "", "source-id")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
