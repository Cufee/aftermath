package retry

import (
	"time"

	"github.com/pkg/errors"
)

type DataWithErr[T any] struct {
	Data T
	Err  error
}

func Retry[T any](fn func() (T, error), tries int, sleepOnFail time.Duration, breakIf ...func(error) bool) DataWithErr[T] {
	if tries < 1 {
		return DataWithErr[T]{Err: errors.New("invalid number of tries provided")}
	}

	tries -= 1
	data, err := fn()
	if err == nil || tries <= 0 {
		return DataWithErr[T]{Data: data, Err: err}
	}

	for _, check := range breakIf {
		shouldBreak := check(err)
		if shouldBreak {
			return DataWithErr[T]{Err: err}
		}
	}

	time.Sleep(sleepOnFail)
	return Retry(fn, tries, sleepOnFail)
}
