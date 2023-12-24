package writekey

import (
	"crypto/rand"
	"fmt"
	"sync"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/ormushq/ormus/logger"
)

var entropyPool = sync.Pool{
	New: func() any {
		entropy := ulid.Monotonic(rand.Reader, 0)

		return entropy
	},
}

func newULID() (string, error) {
	e := entropyPool.Get()

	entropy, ok := e.(*ulid.MonotonicEntropy)
	if !ok {
		logger.L().Info("Unable to Convert Interface to ULID")

		return "", fmt.Errorf("unable to Convert Interface to ULID")
	}

	s := ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
	entropyPool.Put(e)

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
