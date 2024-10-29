package writekey_test

import (
	"context"
	"fmt"
)

type mockRepo struct{}

// TODO - use https://github.com/golang/mock
func (m mockRepo) IsValidWriteKey(ctx context.Context, writeKey string) (bool, error) {
	if writeKey == "" {
		return false, fmt.Errorf("writekey not found")
	}
	return true, nil
}
