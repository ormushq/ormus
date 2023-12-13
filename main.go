package main

import (
	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/manager/service/auth"
	"github.com/ormushq/ormus/manager/service/user"
	"github.com/ormushq/ormus/source/delivery/httpserver/userhandler"
	"net/http"
	"time"
)

var (
	defultSignKey               = "Ormus_jwt"
	defultAccessExpirationTime  = time.Hour * 24 * 7
	defultRefreshExpirationTime = time.Hour * 24 * 7 * 4
	defultAccessSubject         = "ac"
	defultRefreshSubject        = "rt"
)

func main() {

	e := echo.New()
	authsvc := auth.New(defultSignKey, defultAccessSubject,
		defultRefreshSubject, defultAccessExpirationTime, defultAccessExpirationTime)
	usersvc := user.New(authsvc)
	userhand := userhandler.New(usersvc)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, welcome to the user registration app!")
	})
	e.POST("/register", userhand.RegisterUser)
	e.POST("/login", userhand.UserLogin)
	e.Logger.Fatal(e.Start(":8080"))

}