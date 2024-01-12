package userhandler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/manager/delivery/httpserver/userhandler"
	usermock "github.com/ormushq/ormus/manager/mock"
	"github.com/ormushq/ormus/manager/service/authservice"
	"github.com/ormushq/ormus/manager/service/userservice"
	"github.com/ormushq/ormus/manager/validator/uservalidator"
	"github.com/ormushq/ormus/param"
	"github.com/ormushq/ormus/pkg/cryption"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationHandler_Register(t *testing.T) {
	type testCase struct {
		name           string
		requestBody    interface{}
		expectedStatus int
		expectedBody   string
	}

	testCases := []testCase{
		{
			name: "Successful Registration",
			requestBody: param.RegisterRequest{
				Name:     "test_user",
				Email:    "test@gmail.com",
				Password: "HeavYPasS123!",
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   `{"email":"test@gmail.com", "id":"new_id"}`,
		},
		{
			name: "Validation Failed",
			requestBody: param.RegisterRequest{
				Email:    "testgmail.com",
				Password: "weak",
			},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody: `{
				"errors":{
					"email":"email is not valid",
					"name":"cannot be blank",
					"password":"the length must be between 8 and 32"
				},
				"message":"invalid input"
			}`,
		},
	}

	cfg := config.C()
	repo := usermock.NewMockRepository(false)
	crypt := cryption.New(cryption.CryptConfing{})
	jwt := authservice.NewJWT(cfg.Manager.JWTConfig, crypt)
	service := userservice.New(jwt, repo)
	validator := uservalidator.New(repo)
	handler := userhandler.New(service, validator)

	e := echo.New()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			requestBody, _ := json.Marshal(tc.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			// Execution
			_ = handler.RegisterUser(ctx)

			// Assertion
			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.JSONEq(t, tc.expectedBody, strings.TrimSpace(rec.Body.String()))
		})
	}
}
