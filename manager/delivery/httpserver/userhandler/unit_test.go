package userhandler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/manager/delivery/httpserver/userhandler"
	"github.com/ormushq/ormus/manager/mock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_RegisterUser(t *testing.T) {
	testCases := []struct {
		name  string
		email string

		hasValidationErr bool
		hasServiceErr    bool

		expectedHttpStatusCode int
		expectedResponseBody   string
	}{
		{
			name:  "Successful Registration",
			email: "jon@labstack.com",

			expectedHttpStatusCode: http.StatusCreated,
			expectedResponseBody:   `{"email":"jon@labstack.com"}`,
		},
		{
			name:  "Validator Error",
			email: "jon@labstack.com",

			hasValidationErr: true,

			expectedHttpStatusCode: http.StatusBadRequest,
			expectedResponseBody:   `{"error":null,"message":"validation error"}`,
		},
		{
			name:  "Service Error",
			email: "jon@labstack.com",

			hasServiceErr: true,

			expectedHttpStatusCode: http.StatusBadRequest,
			expectedResponseBody:   `{"message":"Bad request"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. Setup
			mockService := usermock.NewMockService(tc.hasServiceErr, tc.email)
			mockValidator := usermock.NewMockValidator(tc.hasValidationErr)
			handler := userhandler.New(mockService, mockValidator)

			// Echo setup
			e := echo.New()
			requestBody := []byte(`{"email":"` + tc.email + `"}`)
			req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			// 2. Execution
			_ = handler.RegisterUser(ctx)

			// 3. Assertion
			assert.Equal(t, tc.expectedHttpStatusCode, rec.Code)
			assert.Equal(t, tc.expectedResponseBody, strings.TrimSpace(rec.Body.String()))
		})
	}
}

func TestHandler_UserLogin(t *testing.T) {
	testCases := []struct {
		name  string
		email string

		hasValidationErr bool
		hasServiceErr    bool

		expectedHttpStatusCode int
		expectedResponseBody   string
	}{
		{
			name:  "Successful Login",
			email: "jon@labstack.com",

			expectedHttpStatusCode: http.StatusOK,
			expectedResponseBody:   `{"user":{"ID":"new_id","CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"Email":"jon@labstack.com"},"token":{"access_token":"","refresh_token":""}}`,
		},
		{
			name:  "Validator Error",
			email: "jon@labstack.com",

			hasValidationErr: true,

			expectedHttpStatusCode: http.StatusBadRequest,
			expectedResponseBody:   `{"error":null,"message":"validation error"}`,
		},
		{
			name:  "Service Error",
			email: "jon@labstack.com",

			hasServiceErr: true,

			expectedHttpStatusCode: http.StatusBadRequest,
			expectedResponseBody:   `{"message":"Bad request"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. Setup
			mockService := usermock.NewMockService(tc.hasServiceErr, tc.email)
			mockValidator := usermock.NewMockValidator(tc.hasValidationErr)
			handler := userhandler.New(mockService, mockValidator)

			// Echo setup
			e := echo.New()
			requestBody := []byte(`{"email":"` + tc.email + `"}`)
			req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			// 2. Execution
			_ = handler.UserLogin(ctx)

			// 3. Assertion
			assert.Equal(t, tc.expectedHttpStatusCode, rec.Code)
			assert.Equal(t, tc.expectedResponseBody, strings.TrimSpace(rec.Body.String()))
		})
	}
}
