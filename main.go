package main

import (
	"fmt"
	"github.com/ormushq/ormus/manager/delivery/httpserver/userhandler"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/config"
	authrepository "github.com/ormushq/ormus/manager/repository/auth"
	authservice "github.com/ormushq/ormus/manager/service/auth"
)

func main() {
	cfg := config.C()
	fmt.Println(cfg)

	e := echo.New()

	jwt := authservice.NewJWT(cfg.Manager.JWTConfig)

	userrepo := authrepository.StorageAdapter{}
	usersvc := authservice.NewService(jwt, userrepo)
	userhand := userhandler.New(usersvc)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, welcome to the user registration app!")
	})
	e.POST("/register", userhand.RegisterUser)
	e.POST("/login", userhand.UserLogin)
	e.Logger.Fatal(e.Start(":8080"))
}
