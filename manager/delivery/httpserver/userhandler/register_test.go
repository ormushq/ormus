package userhandler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/manager/delivery/httpserver/userhandler"
	"github.com/ormushq/ormus/manager/mockRepo/projectstub"
	"github.com/ormushq/ormus/manager/mockRepo/usermock"
	"github.com/ormushq/ormus/manager/service/authservice"
	"github.com/ormushq/ormus/manager/service/projectservice"
	"github.com/ormushq/ormus/manager/service/userservice"
	"github.com/ormushq/ormus/manager/validator/projectvalidator"
	"github.com/ormushq/ormus/manager/validator/uservalidator"
	"github.com/ormushq/ormus/manager/workers"
	"github.com/ormushq/ormus/param"
	"github.com/ormushq/ormus/pkg/channel"
	"github.com/ormushq/ormus/pkg/channel/adapter/simple"
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
	cfg := config.C().Manager
	done := make(chan bool)
	wg := sync.WaitGroup{}
	internalBroker := simple.New(done, &wg)
	internalBroker.NewChannel("CreateDefaultProject", channel.BothMode,
		cfg.InternalBrokerConfig.ChannelSize, cfg.InternalBrokerConfig.MaxRetryPolicy)
	repo := usermock.NewMockRepository(false)
	jwt := authservice.New(cfg.AuthConfig)
	validator := uservalidator.New(repo)
	service := userservice.New(jwt, repo, internalBroker, validator)
	RepoPr := projectstub.New(false)
	val := projectvalidator.New(&RepoPr)
	ProjectSvc := projectservice.New(&RepoPr, internalBroker, val)
	handler := userhandler.New(service, ProjectSvc)
	e := echo.New()

	workers.New(ProjectSvc, internalBroker).Run(done, &wg)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			requestBody, _ := json.Marshal(tc.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			// Execution
			_ = handler.RegisterUser(ctx)

			// Waiting for the project persist to channel and project service to create the project
			time.Sleep(15 * time.Second)
			// Assertion
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if http.StatusCreated == rec.Code {
				exist := RepoPr.IsCreated("new-id")
				assert.True(t, exist)
			}
			assert.JSONEq(t, tc.expectedBody, strings.TrimSpace(rec.Body.String()))
		})
	}
	outChannel, _ := internalBroker.GetInputChannel("CreateDefaultProject")
	close(outChannel)
}
