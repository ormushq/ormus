package writekey_test

import (
	"sync"
	"testing"

	"github.com/ormushq/ormus/pkg/writekey"
)

func TestGenerateWriteKey(t *testing.T) {
	t.Run("Uniqueness", testUniqueness)
	t.Run("ConcurrencySafety", testConcurrency)
}

func testUniqueness(t *testing.T) {
	const count = 1000000

	ids := make(map[string]bool)

	for i := 0; i < count; i++ {
		id, err := writekey.New()
		if err != nil {
			t.Fatalf("generate write key: %s", err)
		}
		if _, ok := ids[id.String()]; ok {
			t.Errorf("duplicate write key")
		}

		ids[id.String()] = true
	}
}

func testConcurrency(t *testing.T) {
	const numOfGoroutines = 100
	const ulidPerGoroutine = 100

	var wg sync.WaitGroup
	wg.Add(numOfGoroutines)

	for i := 0; i < numOfGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < ulidPerGoroutine; j++ {
				if _, err := writekey.New(); err != nil {
					t.Fatalf("generate write key: %s", err)
				}
			}
		}()
	}

	wg.Wait()
}
