package queue

import "github.com/pkg/errors"

var ErrQueueLocked = errors.New("queue is locked")
var ErrQueueFull = errors.New("queue is full")
