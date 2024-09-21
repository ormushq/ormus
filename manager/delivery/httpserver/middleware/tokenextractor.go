package middleware

import (
	"encoding/base64"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	mw "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/manager/service/authservice"
)

func GetTokenFromCookie(js authservice.Service) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			authHeader, err := ctx.Cookie("jwtToken")
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, EchoErrorMessage(err.Error()))
			}

			if authHeader.Value == "" {
				return echo.NewHTTPError(http.StatusBadRequest, EchoErrorMessage("jwt token cookie is nil"))
			}

			claims, err := js.ParseToken(authHeader.Value)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, EchoErrorMessage(err.Error()))
			}

			enUserID := claims.UserID
			userID, err := base64.StdEncoding.DecodeString(enUserID)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, EchoErrorMessage(err.Error()))
			}

			ctx.Set("userID", userID)

			return next(ctx)
		}
	}
}

func GetTokenFromHeader(authSvc authservice.Service) echo.MiddlewareFunc {
	return mw.WithConfig(mw.Config{
		ContextKey: authSvc.GetConfig().ContextKey,
		SigningKey: []byte(authSvc.GetConfig().SecretKey),
		// TODO - as sign method string to config...
		SigningMethod: jwt.SigningMethodHS256.Name,
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			claims, err := authSvc.ParseToken(auth)
			if err != nil {
				return nil, err
			}

			return claims, nil
		},
	})
}

func EchoErrorMessage(message string) echo.Map {
	return echo.Map{"message": message}
}
