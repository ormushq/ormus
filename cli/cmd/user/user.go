/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package user

import (
	"fmt"

	"github.com/ormushq/ormus/cli/cmd"
	"github.com/spf13/cobra"
)

// userCmd represents the user command.
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "User Registration and Authentication",
	Long: `User Registration and Authentication

ormus user register --email <email>: Register a new user with an email.
ormus user login --email <email>: Log in as a user.
ormus user logout: Log out the current user.
User Account Management

ormus user show: Display details of the logged-in user.
ormus user update --email <new-email> [--name <name>]: Update user details.
ormus user delete: Delete the logged-in user's account.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("user called")
	},
}

func init() {
	cmd.RootCmd.AddCommand(userCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
