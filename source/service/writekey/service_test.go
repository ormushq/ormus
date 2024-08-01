package writekey_test

import (
	"context"
	"fmt"
	"github.com/ormushq/ormus/manager/entity"
	"testing"

	"github.com/ormushq/ormus/source/service/writekey"
)

type mockRepo struct{}

func (m mockRepo) IsValidWriteKey(ctx context.Context, writeKey string) (*entity.WriteKeyMetaData, error) {
	if writeKey == "" {
		return nil, fmt.Errorf("writekey not found")
	}
	return &entity.WriteKeyMetaData{}, nil
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
		metaData, err := service.IsValid(ctx, "asdfffg4g5g56d5s4s6s5sd8")
		if err != nil {
			t.Fatal("error is not nil")
		}
		if metaData == nil {
			t.Fatal("writekey is not valid")
		}
	})
}
