package tasks

import (
	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/internal/database/models"
)

type TaskHandler struct {
	Process     func(client core.Client, task models.Task) (string, error)
	ShouldRetry func(task *models.Task) bool
}

var defaultHandlers = make(map[models.TaskType]TaskHandler)

func DefaultHandlers() map[models.TaskType]TaskHandler {
	return defaultHandlers
}
