package writekey

import (
	"bytes"
	"crypto/rand"
	"sync"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/ormushq/ormus/logger"
)

type WriteKey string

var mu sync.Mutex

const lengthOfEntropySlice = 16

func New() (WriteKey, error) {
	t := time.Now()
	entropy, err := getEntropy()
	if err != nil {
		logger.L().WithGroup("write-key").Error(err.Error())

		return "", err
	}

	// we lock the process of creating new ULID because it's not Concurrent safe
	mu.Lock()
	ulidID, err := ulid.New(ulid.Timestamp(t), bytes.NewReader(entropy))
	mu.Unlock()
	if err != nil {
		logger.L().WithGroup("write-key").Error(err.Error())

		return "", err
	}

	return WriteKey(ulidID.String()), nil
}

func (w WriteKey) String() string {
	return string(w)
}

func getEntropy() ([]byte, error) {
	entropy := make([]byte, lengthOfEntropySlice)
	_, err := rand.Read(entropy)
	if err != nil {
		return nil, err
	}

	return entropy, nil
}
