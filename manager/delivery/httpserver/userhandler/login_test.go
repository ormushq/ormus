package userhandler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/manager/delivery/httpserver/userhandler"
	"github.com/ormushq/ormus/manager/mockRepo/usermock"
	"github.com/ormushq/ormus/manager/service/authservice"
	"github.com/ormushq/ormus/manager/service/userservice"
	"github.com/ormushq/ormus/manager/validator/uservalidator"
	"github.com/ormushq/ormus/param"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationHandler_Login(t *testing.T) {
	type testCase struct {
		name           string
		requestBody    interface{}
		expectedStatus int

		expectedErrBody string
		err             bool
	}

	defaultUser := usermock.DefaultUser()

	testCases := []testCase{
		{
			name: "Successful Login",
			requestBody: param.LoginRequest{
				Email:    defaultUser.Email,
				Password: defaultUser.Password,
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "not existing user",
			requestBody: param.LoginRequest{
				Email:    "test@gmail.com",
				Password: "wRonG123!",
			},
			expectedStatus:  http.StatusUnauthorized,
			err:             true,
			expectedErrBody: fmt.Sprintf(`{"message":"%s"}`, errmsg.ErrWrongCredentials),
		},
		{
			name: "wrong credentials",
			requestBody: param.LoginRequest{
				Email:    defaultUser.Email,
				Password: "wRonG123!",
			},
			expectedStatus:  http.StatusUnauthorized,
			err:             true,
			expectedErrBody: fmt.Sprintf(`{"message":"%s"}`, errmsg.ErrWrongCredentials),
		},
		{
			name: "Validation Failed",
			requestBody: param.LoginRequest{
				Email:    "testgmail.com",
				Password: "weak",
			},
			err:            true,
			expectedStatus: http.StatusUnprocessableEntity,
			expectedErrBody: `{
				"errors":{
					"email":"email is not valid",
					"password":"the length must be between 8 and 32"
				},
				"message":"invalid input"
			}`,
		},
	}

	cfg := config.C()
	repo := usermock.NewMockRepository(false)
	jwt := authservice.NewJWT(cfg.Manager.JWTConfig)
	service := userservice.New(jwt, repo, nil)
	validator := uservalidator.New(repo)
	handler := userhandler.New(service, validator, nil)

	e := echo.New()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			requestBody, _ := json.Marshal(tc.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			// 2. execution
			_ = handler.UserLogin(ctx)

			// 3. assertion
			assert.Equal(t, tc.expectedStatus, rec.Code)

			if tc.err {
				assert.JSONEq(t, tc.expectedErrBody, rec.Body.String())
				return
			}

			response := &param.LoginResponse{}
			err := json.Unmarshal(
				rec.Body.Bytes(),
				response,
			)
			assert.Nil(t, err, "error in deserializing the request")

			assert.NotEmpty(t, response.Tokens)
			assert.NotEmpty(t, response.User)
			assert.Equal(t, response.User.Email, tc.requestBody.(param.LoginRequest).Email)
		})
	}
}
