package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/manager/delivery/httpserver/userhandler"
	userrepo "github.com/ormushq/ormus/manager/repository/user"
	service "github.com/ormushq/ormus/manager/service/authservice"
	"github.com/ormushq/ormus/manager/service/userservice"
	uservalidator "github.com/ormushq/ormus/manager/validator/user"
)

func main() {
	cfg := config.C()
	fmt.Println(cfg)

	e := echo.New()
	jwt := service.NewJWT(cfg.Manager.JWTConfig)

	unknownRepo := userrepo.StorageAdapter{}
	userSvc := userservice.NewService(jwt, unknownRepo)
	validateUserSvc := uservalidator.New(unknownRepo)
	userHand := userhandler.New(userSvc, validateUserSvc)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, welcome to the user registration app!")
	})
	e.POST("/register", userHand.RegisterUser)
	e.POST("/login", userHand.UserLogin)
	e.Logger.Fatal(e.Start(":8080"))
}
