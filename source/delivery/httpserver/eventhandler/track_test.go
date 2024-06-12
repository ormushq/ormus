package eventhandler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/source/delivery/httpserver/eventhandler"
	"github.com/ormushq/ormus/source/delivery/httpserver/middleware"
	"github.com/ormushq/ormus/source/mock"
	"github.com/ormushq/ormus/source/params"
	"github.com/ormushq/ormus/source/service/eventservice"
	"github.com/ormushq/ormus/source/validator/eventvalidator"
	"time"

	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_TrackEvent(t *testing.T) {
	type testCase struct {
		name           string
		requestBody    interface{}
		expectedStatus int

		expectedErrBody string
		err             bool
	}

	defaultEvent := mock.DefaultEvent()

	testCases := []testCase{
		{
			name:            "invalid payload",
			requestBody:     `invalid payload`,
			expectedStatus:  http.StatusBadRequest,
			expectedErrBody: `{"message": "Bad request"}`,
			err:             true,
		},
		{
			name: "Successful track",
			requestBody: params.TrackEventRequest{
				MessageID:         defaultEvent.MessageID,
				Type:              defaultEvent.Type,
				Name:              defaultEvent.Name,
				Properties:        nil,
				Integration:       nil,
				Ctx:               nil,
				SendAt:            time.Time{},
				ReceivedAt:        time.Time{},
				OriginalTimeStamp: time.Time{},
				Timestamp:         time.Time{},
				AnonymousID:       defaultEvent.AnonymousID,
				UserID:            defaultEvent.UserID,
				GroupID:           defaultEvent.GroupID,
				PreviousID:        defaultEvent.PreviousID,
				Event:             defaultEvent.Event,
				WriteKey:          defaultEvent.WriteKey,
				MetaData:          event.MetaData{},
				Options:           nil,
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name: "Validation Failed",
			requestBody: params.TrackEventRequest{
				MessageID:         defaultEvent.MessageID,
				Type:              defaultEvent.Type,
				Name:              "",
				Properties:        nil,
				Integration:       nil,
				Ctx:               nil,
				SendAt:            time.Time{},
				ReceivedAt:        time.Time{},
				OriginalTimeStamp: time.Time{},
				Timestamp:         time.Time{},
				AnonymousID:       defaultEvent.AnonymousID,
				UserID:            defaultEvent.UserID,
				GroupID:           defaultEvent.GroupID,
				PreviousID:        defaultEvent.PreviousID,
				Event:             defaultEvent.Event,
				WriteKey:          defaultEvent.WriteKey,
				MetaData:          event.MetaData{},
				Options:           nil,
			},
			err:             true,
			expectedStatus:  http.StatusUnprocessableEntity,
			expectedErrBody: `{"errors":{"Name":"cannot be blank"},"message":"invalid input"}`,
		},
	}

	repo := mock.NewMockRepository(false)
	eventSvc := eventservice.New(&repo)
	eventVld := eventvalidator.New(repo)
	eventHandler := eventhandler.New(eventSvc, eventVld)
	e := echo.New()
	r := e.Group("/v1/track-event")

	r.POST("/", eventHandler.TrackEvent, middleware.WritekeyValidation())

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			requestBody, _ := json.Marshal(tc.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/v1/track-event/", bytes.NewBuffer(requestBody))

			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()

			e.ServeHTTP(rec, req)

			// Assertion
			assert.Equal(t, tc.expectedStatus, rec.Code)
			fmt.Println(rec.Body.String())
			if tc.err {
				assert.JSONEq(t, tc.expectedErrBody, rec.Body.String())
				return
			}
			assert.Equal(t, tc.expectedStatus, rec.Code)
		})
	}
}
