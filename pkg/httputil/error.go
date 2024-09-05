package httputil

import "github.com/labstack/echo/v4"

func NewError(ctx echo.Context, status int, msg string) error {
	er := HTTPError{
		Message: msg,
	}

	return ctx.JSON(status, er)
}

// HTTPError example.
type HTTPError struct {
	Message string `json:"message" example:"status bad request"`
}
