package tests

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWriteKeyValidation(t *testing.T) {
	c := GetConfigs()
	c.redisAdapter.Client().Set(context.Background(), "valid_fake_writekey", "sample", time.Minute)
	isValid, err := c.writeKeyRepo.IsWriteKeyValid(context.Background(), "valid_fake_writekey", 3)
	assert.Nil(t, err)
	assert.True(t, isValid)
	// Assume that manager service is not available
	isValid, err = c.writeKeyRepo.IsWriteKeyValid(context.Background(), "inivalid_fake_writekey", 3)
	assert.Nil(t, err)
	assert.False(t, isValid)
}
