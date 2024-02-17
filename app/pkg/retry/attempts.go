package retry

import (
	"time"
)

func WithAttempts(maxAttempts int, delay time.Duration, fn func() error) error {
	var err error

	for ma := 0; ma < maxAttempts; ma++ {
		err = fn()
		if err != nil {
			time.Sleep(delay)

			continue
		}

		return nil
	}

	return err
}
