package eventvalidator

import (
	"context"
	"testing"

	"github.com/ormushq/ormus/source"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) IsWriteKeyValid(ctx context.Context, writeKey string, expirationTime uint) (bool, error) {
	args := m.Called(ctx, writeKey, expirationTime)
	return args.Bool(0), args.Error(1)
}

func TestValidateWriteKey_Success(t *testing.T) {
	cfg := source.Config{}
	mockRepo := new(MockRepository)
	mockRepo.On("IsWriteKeyValid", mock.Anything, mock.Anything, mock.Anything).Return(true, nil)
	validator := New(mockRepo, cfg)
	valid, err := validator.ValidateWriteKey(context.Background(), "valid_writekey")
	assert.NoError(t, err)
	assert.True(t, valid)
	mockRepo.AssertCalled(t, "IsWriteKeyValid", mock.Anything, mock.Anything, mock.Anything)
}

func TestValidateWriteKey_Failure(t *testing.T) {
	mockRepo := new(MockRepository)
	cfg := source.Config{}
	mockRepo.On("IsWriteKeyValid", mock.Anything, mock.Anything, mock.Anything).Return(false, nil)
	validator := New(mockRepo, cfg)
	valid, err := validator.ValidateWriteKey(context.Background(), "invalid_writekey")
	assert.NoError(t, err)
	assert.False(t, valid)
	mockRepo.AssertCalled(t, "IsWriteKeyValid", mock.Anything, mock.Anything, mock.Anything)
}
