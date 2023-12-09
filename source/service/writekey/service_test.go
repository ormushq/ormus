package writekey_test

import (
	"context"
	"fmt"
	"github.com/ormushq/ormus/source/service/writekey"
	"testing"
)

type mockRepo struct {
}

// TODO - use https://github.com/golang/mock
func (m mockRepo) IsValidWriteKey(ctx context.Context, writeKey string) (bool, error) {
	if writeKey == "" {
		return false, fmt.Errorf("writekey not found")
	}
	return true, nil
}

func TestIsValid(t *testing.T) {
	t.Run("writekey not found", func(t *testing.T) {
		m := new(mockRepo)
		service := writekey.New(m)
		ctx := context.Background()
		_, err := service.IsValid(ctx, "")
		if err == nil {
			t.Fatal("error is nil")
		}
	})

	t.Run("writekey is exists and valid", func(t *testing.T) {
		m := new(mockRepo)
		service := writekey.New(m)
		ctx := context.Background()
		isValid, err := service.IsValid(ctx, "asdfffg4g5g56d5s4s6s5sd8")
		if err != nil {
			t.Fatal("error is not nil")
		}
		if !isValid {
			t.Fatal("writekey is not valid")
		}
	})
}
