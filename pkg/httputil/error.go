package httputil

import (
	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/pkg/httpmsg"
)

func NewError(ctx echo.Context, status int, msg string) error {
	er := HTTPError{
		Message: msg,
	}

	return ctx.JSON(status, er)
}

func NewErrorWithError(ctx echo.Context, err error) error {
	message, code := httpmsg.Error(err)

	return NewError(ctx, code, message)
}

// HTTPError example.
type HTTPError struct {
	Message string `json:"message" example:"status bad request"`
}
