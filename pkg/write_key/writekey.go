package writekey

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func newULID() (string, error) {
	entropy := ulid.Monotonic(rand.Reader, 0)

	s := ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()

	return s, nil
}

func GenerateNewWriteKey() (string, error) {
	return newULID()
}

func ValidateWriteKey(writeKey string) error {
	_, err := ulid.Parse(writeKey)
	if err != nil {
		return err
	}

	return nil
}
