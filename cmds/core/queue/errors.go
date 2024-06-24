package queue

import "github.com/pkg/errors"

var ErrQueueLocked = errors.New("queue is locked")
