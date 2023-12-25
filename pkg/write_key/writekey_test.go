package writekey_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/oklog/ulid/v2"
	writekey "github.com/ormushq/ormus/pkg/write_key"
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
