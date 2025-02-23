package retry

import (
	"fmt"
	"time"

	"github.com/ormushq/ormus/logger"
)

type RetryableFunc func() error

func Do(fn RetryableFunc, maxAttempts int) error {
	var err error

	for i := 0; i < maxAttempts; i++ {
		logger.L().Info(fmt.Sprintf("Retrying function, attempt %d/%d", i+1, maxAttempts))
		err = fn()
		if err == nil {
			logger.L().Info("Function succeeded")
			break
		}
		logger.L().Error(fmt.Sprintf("Function failed on attempt %d: %v. Retrying...", i+1, err))
		time.Sleep(time.Second)
	}

	return err
}
