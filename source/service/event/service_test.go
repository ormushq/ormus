package event

import (
	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/pkg/richerror"
	"github.com/ormushq/ormus/source/params"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateNewEvent(evt event.CoreEvent) (string, error) {
	args := m.Called(evt)
	return args.String(0), args.Error(1)
}

func TestCreateNewEvent_Success(T *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("CreateNewEvent", mock.Anything).Return("2sdf121fa-dfe21", nil)
	svc := New(mockRepo)

	resp, err := svc.CreateNewEvent(params.TrackEventRequest{})
	assert.Nil(T, err)
	assert.NotNil(T, resp)
	assert.Equal(T, &params.TrackEventResponse{ID: "2sdf121fa-dfe21"}, resp)
}

func TestCreateNewEvent_Has_Error(T *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("CreateNewEvent", mock.Anything).Return("", richerror.New("sample"))
	svc := New(mockRepo)

	resp, err := svc.CreateNewEvent(params.TrackEventRequest{})
	assert.Nil(T, resp)
	assert.NotNil(T, err)

}
