package main

import (
	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/source/api"
)

func main() {
	e := echo.New()

	api.RegisterRoutes(e)

	e.Start(":8080")
}