package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/source/params"
	"github.com/ormushq/ormus/source/validator/eventvalidator/eventvalidator"
)

func WriteKeyMiddleware(validator eventvalidator.Validator) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			body, err := io.ReadAll(c.Request().Body)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "Failed to read body")
			}

			c.Request().Body = io.NopCloser(bytes.NewReader(body))

			var req params.TrackEventRequest
			if err := json.Unmarshal(body, &req); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "Failed to unmarshal body")
			}

			isValid, err := validator.ValidateWriteKey(c.Request().Context(), req.WriteKey)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": "something went wrong",
				})
			}
			if !isValid {
				return c.JSON(http.StatusForbidden, "the write key is invalid")
			}
			c.Set("body", req)

			return next(c)
		}
	}
}
