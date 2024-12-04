package eventvalidator

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) IsWriteKeyValid(ctx context.Context, writeKey string) (bool, error) {
	args := m.Called(ctx, writeKey)
	return args.Bool(0), args.Error(1)
}

func TestValidateWriteKey_Success(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("IsWriteKeyValid", mock.Anything, mock.Anything).Return(true, nil)
	validator := New(mockRepo)
	valid, err := validator.ValidateWriteKey(context.Background(), "valid_writekey")
	assert.NoError(t, err)
	assert.True(t, valid)
	mockRepo.AssertCalled(t, "IsWriteKeyValid", mock.Anything, mock.Anything)
}

func TestValidateWriteKey_Failure(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("IsWriteKeyValid", mock.Anything, mock.Anything).Return(false, nil)
	validator := New(mockRepo)
	valid, err := validator.ValidateWriteKey(context.Background(), "invalid_writekey")
	assert.NoError(t, err)
	assert.False(t, valid)
	mockRepo.AssertCalled(t, "IsWriteKeyValid", mock.Anything, mock.Anything)
}
