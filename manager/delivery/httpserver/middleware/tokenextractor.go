package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (m Middleware) GetTokenFromCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authHeader, err := ctx.Cookie("jwtToken")
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, EchoErrorMessage(err.Error()))
		}

		if authHeader.Value == "" {
			return echo.NewHTTPError(http.StatusBadRequest, EchoErrorMessage("jwt token cookie is nil"))
		}

		claims, err := m.jwtSvc.ParseToken(authHeader.Value)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, EchoErrorMessage(err.Error()))
		}

		enUserEmail := claims.EnUserEmail
		// TODO : what is best practice to give encryption key from jwt confing because confing not exported feild
		userEmail, err := m.jwtSvc.Cryption.Decrypt(enUserEmail)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, EchoErrorMessage(err.Error()))
		}

		ctx.Set("userEmail", userEmail)

		return next(ctx)
	}
}
