package writekey

import (
	"bytes"
	"crypto/rand"
	"sync"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/ormushq/ormus/logger"
)

var mu sync.Mutex

const lengthOfEntropySlice = 16

func GenerateWriteKey() (string, error) {
	mu.Lock()
	defer mu.Unlock()

	t := time.Now()
	entropy, err := getEntropy()
	if err != nil {
		logger.L().WithGroup("write-key").Error(err.Error())

		return "", err
	}

	ulidID, err := ulid.New(ulid.Timestamp(t), bytes.NewReader(entropy))
	if err != nil {
		logger.L().WithGroup("write-key").Error(err.Error())

		return "", err
	}

	return ulidID.String(), nil
}

func getEntropy() ([]byte, error) {
	entropy := make([]byte, lengthOfEntropySlice)
	_, err := rand.Read(entropy)
	if err != nil {
		return nil, err
	}

	return entropy, nil
}
