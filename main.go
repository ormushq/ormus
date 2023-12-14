package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/manager/service/auth"
	"github.com/ormushq/ormus/source/delivery/httpserver/userhandler"
	"net/http"
)

func main() {
	cfg := config.C()
	fmt.Println(cfg)

	e := echo.New()

	jwt := auth.NewJWT(cfg.Manager.JWTConfig)

	usersvc := auth.New(jwt)
	userhand := userhandler.New(usersvc)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, welcome to the user registration app!")
	})
	e.POST("/register", userhand.RegisterUser)
	e.POST("/login", userhand.UserLogin)
	e.Logger.Fatal(e.Start(":8080"))

}
