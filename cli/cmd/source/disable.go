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

// disableCmd represents the disable command.
var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable a source",
	Long:  `ormus source disable --source-id <source-id>`,
	Run: func(cmdCobra *cobra.Command, args []string) {
		sourceID, err := cmdCobra.Flags().GetString("source-id")
		if err != nil {
			fmt.Println("error on get source id flag", err)

			return
		}

		if sourceID == "" {
			fmt.Println("source id is required")

			return
		}
		resp, err := cmd.Client.SendRequest(cmd.Client.Source.Disable(sourceID))
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
	sourceCmd.AddCommand(disableCmd)

	disableCmd.Flags().String("source-id", "", "source-id")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// disableCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// disableCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
