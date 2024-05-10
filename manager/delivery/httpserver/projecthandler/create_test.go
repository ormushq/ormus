package projecthandler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/manager/delivery/httpserver/projecthandler"
	"github.com/ormushq/ormus/manager/mock/projectstub"
	"github.com/ormushq/ormus/manager/mock/usermock"
	"github.com/ormushq/ormus/manager/service/authservice"
	"github.com/ormushq/ormus/manager/service/projectservice"
	"github.com/ormushq/ormus/manager/service/userservice"
	"github.com/ormushq/ormus/manager/validator/projectvalidator"
	"github.com/ormushq/ormus/param"
	"github.com/stretchr/testify/assert"
)

func TestHandler_Create(t *testing.T) {
	testCases := []struct {
		name           string
		requestBody    interface{}
		expectedStatus int

		expectedResponse string
		err              bool
	}{}

	cfg := config.C()
	repo := projectstub.New(false)

	// TODO: these better to be automated to get a user and make login token for them
	userRepo := usermock.NewMockRepository(false)
	jwt := authservice.NewJWT(cfg.Manager.JWTConfig)
	userService := userservice.New(jwt, userRepo)

	service := projectservice.New(repo)
	validator := projectvalidator.New(userService)
	handler := projecthandler.New(service, validator)

	e := echo.New()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			requestBody, _ := json.Marshal(tc.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			// 2. execution
			_ = handler.Create(ctx)

			// 3. assertion
			assert.Equal(t, tc.expectedStatus, rec.Code)

			if tc.err {
				assert.JSONEq(t, tc.expectedResponse, rec.Body.String())
				return
			}

			response := &param.CreateProjectResponse{}
			err := json.Unmarshal(
				rec.Body.Bytes(),
				response,
			)
			assert.Nil(t, err, "error in deserializing the request")

			assert.NotEmpty(t, response.ID)
			assert.NotEmpty(t, response.Name)
		})
	}
}
