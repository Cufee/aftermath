package fetch

import (
	"errors"
	"time"
)

type DataWithErr[T any] struct {
	Data T
	Err  error
}

func withRetry[T any](fn func() (T, error), tries int, sleepOnFail time.Duration) DataWithErr[T] {
	if tries < 1 {
		return DataWithErr[T]{Err: errors.New("invalid number of tries provided")}
	}

	tries -= 1
	data, err := fn()
	if err == nil || tries <= 0 {
		return DataWithErr[T]{Data: data, Err: err}
	}

	time.Sleep(sleepOnFail)
	return withRetry(fn, tries, sleepOnFail)
}
