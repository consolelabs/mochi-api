package util

import (
	"time"
)

func RetryRequest(handler func() error) error {
	count := 0
	err := handler()
	for err != nil && !IsAcceptableErr(err) && count < 10 {
		err = handler()
		time.Sleep(time.Second)
		count++
	}
	if IsAcceptableErr(err) {
		return err
	}

	// backup for case server cannot handle func more than 10 times
	if err != nil && count >= 10 {
		return err
	}
	return nil
}
