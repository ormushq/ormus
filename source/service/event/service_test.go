package event

import (
	"context"
	"sync"
	"testing"

	"github.com/ormushq/ormus/contract/go/destination"
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

func (m *MockRepository) CreateNewEvent(ctx context.Context, evt []event.CoreEvent, wg *sync.WaitGroup, queueName string) ([]string, error) {
	args := m.Called(ctx, evt, wg, queueName)
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockRepository) EventHasDelivered(ctx context.Context, evt *destination.DeliveredEventsList) error {
	args := m.Called(ctx, evt)
	return args.Error(0)
}

func TestCreateNewEvent_Success(T *testing.T) {
	var wg *sync.WaitGroup
	mockRepo := new(MockRepository)
	mockRepo.On("CreateNewEvent", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return([]string{"2sdf121fa-dfe21"}, nil)
	svc := New(mockRepo, source.Config{}, wg)

	resp, err := svc.CreateNewEvent(context.Background(), []params.TrackEventRequest{
		{
			Type: "test_type",
			Name: "test_name",
		},
	}, []string{"1234"})
	mockRepo.AssertCalled(T, "CreateNewEvent", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	assert.Nil(T, err)
	assert.NotNil(T, resp)
	assert.Equal(T, &params.TrackEventResponse{
		ID:               []string{"2sdf121fa-dfe21"},
		InvalidWriteKeys: []string{"1234"},
		FAIL:             1,
		Success:          1,
	}, resp)
}

func TestCreateNewEvent_Has_Error(T *testing.T) {
	var wg *sync.WaitGroup
	mockRepo := new(MockRepository)
	mockRepo.On("CreateNewEvent", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return([]string{}, richerror.New("sample"))
	svc := New(mockRepo, source.Config{}, wg)

	resp, err := svc.CreateNewEvent(context.Background(), []params.TrackEventRequest{
		{
			Type: "test_type",
			Name: "test_name",
		},
	}, []string{})
	assert.Nil(T, resp)
	assert.NotNil(T, err)
}
