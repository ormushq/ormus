/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package user

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ormushq/ormus/cli/cmd"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command.
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new user with an email.",
	Long:  `ormus user register --email <email> --password <password> --name <name>`,
	Run: func(cmdCobra *cobra.Command, args []string) {
		email, err := cmdCobra.Flags().GetString("email")
		if err != nil {
			fmt.Println("error on get email flag", err)

			return
		}

		password, err := cmdCobra.Flags().GetString("password")
		if err != nil {
			fmt.Println("error on get password flag", err)

			return
		}

		name, err := cmdCobra.Flags().GetString("name")
		if err != nil {
			fmt.Println("error on get name flag", err)

			return
		}

		if password == "" || email == "" || name == "" {
			fmt.Println("password and email and name is required")

			return
		}
		resp, err := cmd.Client.SendRequest(cmd.Client.User.Register(name, email, password))
		if err != nil {
			log.Fatal(err)
		}
		j, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode != http.StatusCreated {
			log.Fatalf("status not created ,status code %d, body: %s", resp.StatusCode, j)
		}

		fmt.Printf("success response : \n%s\n", j)
	},
}

func init() {
	userCmd.AddCommand(registerCmd)

	registerCmd.Flags().String("name", "", "name")
	registerCmd.Flags().String("email", "", "email")
	registerCmd.Flags().String("password", "", "password")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
