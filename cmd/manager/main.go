package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/manager/delivery/httpserver/userhandler"
	usermock "github.com/ormushq/ormus/manager/mock"
	"github.com/ormushq/ormus/manager/service/authservice"
	"github.com/ormushq/ormus/manager/service/userservice"
	"github.com/ormushq/ormus/manager/validator/uservalidator"
)

func main() {
	cfg := config.C()

	e := echo.New()

	// TODO: implement the repository for user
	jwt := authservice.NewJWT(cfg.Manager.JWTConfig)
	unknownRepo := usermock.NewMockRepository(false)
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
