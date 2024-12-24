package writekey

import (
	"context"
	"testing"

	proto_source "github.com/ormushq/ormus/contract/go/source"
	"github.com/ormushq/ormus/pkg/richerror"
	"github.com/ormushq/ormus/source"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateNewWriteKey(ctx context.Context, writeKey *proto_source.NewSourceEvent, expirationTime uint) error {
	args := m.Called(ctx, writeKey, expirationTime)
	return args.Error(0)
}

func TestCreateNewWriteKey_Success(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("CreateNewWriteKey", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	service := New(mockRepo, source.Config{WriteKeyRedisExpiration: 3600})

	err := service.CreateNewWriteKey(context.Background(), "owner1", "project1", "writeKey123")

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestCreateNewWriteKey_Has_Error(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("CreateNewWriteKey", mock.Anything, mock.Anything, mock.Anything).
		Return(richerror.New("source.service").WithMessage("some error"))

	service := New(mockRepo, source.Config{WriteKeyRedisExpiration: 3600})

	err := service.CreateNewWriteKey(context.Background(), "owner1", "project1", "writeKey123")

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
