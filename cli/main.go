/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/ormushq/ormus/cli/cmd"
	_ "github.com/ormushq/ormus/cli/cmd/config"
	_ "github.com/ormushq/ormus/cli/cmd/destination"
	_ "github.com/ormushq/ormus/cli/cmd/project"
	_ "github.com/ormushq/ormus/cli/cmd/source"
	_ "github.com/ormushq/ormus/cli/cmd/user"
)

func main() {
	cmd.Execute()
}
