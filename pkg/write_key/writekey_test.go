package writekey_test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"

	"github.com/oklog/ulid/v2"
	writekey "github.com/ormushq/ormus/pkg/write_key"
	"github.com/stretchr/testify/assert"
)

func TestWriteKeyUniqueness(t *testing.T) {
	var wg sync.WaitGroup
	ids := make([]string, 0)

	for i := 0; i < 1000000; i++ {

		wg.Add(1)
		go func() {
			id, err := writekey.GenerateNewWriteKey()
			if err != nil {
				t.Errorf("error while generating writekey")
				return
			}

			ids = append(ids, id)
			wg.Done()
		}()
	}

	wg.Wait()

	m := make(map[string]bool)
	for _, id := range ids {
		fmt.Println(id)
		if m[id] {
			t.Errorf("there is same write key")
			return
		}
	}
}

func TestWriteKeyValidation(t *testing.T) {
	var wg sync.WaitGroup
	ids := make([]string, 0)

	for i := 0; i < 1000000; i++ {

		wg.Add(1)
		go func() {
			id, err := writekey.GenerateNewWriteKey()
			if err != nil {
				t.Errorf("error while generating writekey")
				return
			}

			ids = append(ids, id)
			wg.Done()
		}()
	}

	wg.Wait()

	for _, id := range ids {
		if id != "" {
			_, err := ulid.Parse(id)
			if err != nil {
				t.Errorf("error while validation writekey %s, id: >%s<", err.Error(), id)
				return
			}
		}
	}
}

func TestWriteKeyInValidation(t *testing.T) {
	var wg sync.WaitGroup
	fakeIDs := make([]string, 0)

	for i := 0; i < 1000000; i++ {

		wg.Add(1)
		go func() {
			fakeID, err := generateFakeULID()
			if err != nil {
				t.Errorf("error while generating writekey")
				return
			}

			fakeIDs = append(fakeIDs, fakeID)
			wg.Done()
		}()
	}

	wg.Wait()

	for _, id := range fakeIDs {
		if id != "" {
			_, err := ulid.Parse(id)
			assert.NotNil(t, err)
			return
		}
	}
}

// Define a function that generates a fake ULID string
func generateFakeULID() (string, error) {
	chars := "0123456789ABCDEFGHJKMNPQRSTVWXYZ"
	runes := []rune(chars)
	result := make([]byte, 26)

	for i := 0; i < 26; i++ {
		r := runes[rand.Intn(len(runes))]
		result[i] = byte(r)
	}

	return string(result), nil
}
