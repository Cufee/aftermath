package tasks

import (
	"github.com/cufee/aftermath/cmds/core"
	"github.com/cufee/aftermath/internal/database"
)

type TaskHandler struct {
	Process     func(client core.Client, task database.Task) (string, error)
	ShouldRetry func(task *database.Task) bool
}

var defaultHandlers = make(map[database.TaskType]TaskHandler)

func DefaultHandlers() map[database.TaskType]TaskHandler {
	return defaultHandlers
}
