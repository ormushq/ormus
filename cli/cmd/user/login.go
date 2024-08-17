/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package user

import (
	"encoding/json"
	"fmt"
	"github.com/ormushq/ormus/cli/cmd"
	"github.com/ormushq/ormus/param"
	"github.com/spf13/cobra"
	"io"
	"net/http"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in as a user.",
	Long:  `ormus user login --email <email> --password <password>`,
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

		if password == "" || email == "" {
			fmt.Println("password and email is required")
			return
		}
		resp, err := cmd.Client.SendRequest(cmd.Client.User.Login(email, password))
		if err != nil {
			panic(err)
		}
		j, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		if resp.StatusCode != http.StatusOK {
			panic(fmt.Sprintf("status not OK ,status code %d, body: %s", resp.StatusCode, j))
		}

		var lRsp param.LoginResponse
		err = json.Unmarshal(j, &lRsp)
		if err != nil {
			panic(err)
		}
		cmd.Client.StoreToken(lRsp.Tokens.AccessToken)
		fmt.Println("Token stored successfully")

	},
}

func init() {
	userCmd.AddCommand(loginCmd)

	loginCmd.Flags().String("email", "", "email")
	loginCmd.Flags().String("password", "", "password")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
