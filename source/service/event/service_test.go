package event

import (
	"context"
	"sync"
	"testing"

	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/pkg/richerror"
	"github.com/ormushq/ormus/source"
	"github.com/ormushq/ormus/source/params"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateNewEvent(ctx context.Context, evt event.CoreEvent, wg *sync.WaitGroup, queueName string) (string, error) {
	args := m.Called(ctx, evt, wg, queueName)
	return args.String(0), args.Error(1)
}

func TestCreateNewEvent_Success(T *testing.T) {
	var wg *sync.WaitGroup
	mockRepo := new(MockRepository)
	mockRepo.On("CreateNewEvent", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return("2sdf121fa-dfe21", nil)
	svc := New(mockRepo, source.Config{}, wg)

	resp, err := svc.CreateNewEvent(context.Background(), params.TrackEventRequest{})
	assert.Nil(T, err)
	assert.NotNil(T, resp)
	assert.Equal(T, &params.TrackEventResponse{ID: "2sdf121fa-dfe21"}, resp)
}

func TestCreateNewEvent_Has_Error(T *testing.T) {
	var wg *sync.WaitGroup
	mockRepo := new(MockRepository)
	mockRepo.On("CreateNewEvent", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return("", richerror.New("sample"))
	svc := New(mockRepo, source.Config{}, wg)

	resp, err := svc.CreateNewEvent(context.Background(), params.TrackEventRequest{})
	assert.Nil(T, resp)
	assert.NotNil(T, err)
}
