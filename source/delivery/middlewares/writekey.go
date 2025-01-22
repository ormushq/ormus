package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/logger"
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

			var req []params.TrackEventRequest
			if err := json.Unmarshal(body, &req); err != nil {
				logger.L().Error(err.Error())

				return echo.NewHTTPError(http.StatusBadRequest, "Failed to unmarshal body")
			}
			invalidWriteKeys := make([]string, 0)
			filteredRequests := make([]params.TrackEventRequest, 0, len(req))
			for _, r := range req {
				isValid, err := validator.ValidateWriteKey(c.Request().Context(), r.WriteKey)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, echo.Map{
						"message": "something went wrong",
					})
				}
				if isValid {
					filteredRequests = append(filteredRequests, r)
				} else {
					invalidWriteKeys = append(invalidWriteKeys, r.WriteKey)
				}
			}

			filteredBody, err := json.Marshal(filteredRequests)
			if err != nil {
				logger.L().Error(err.Error())

				return echo.NewHTTPError(http.StatusBadRequest, "Failed to marshal body")
			}
			c.Request().Body = io.NopCloser(bytes.NewReader(filteredBody))
			c.Set("invalid_write_keys", invalidWriteKeys)

			return next(c)
		}
	}
}
