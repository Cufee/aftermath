package tasks

import (
	"context"

	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/internal/database/models"
)

type TaskHandler struct {
	Process func(context.Context, core.Client, *models.Task) error
}

var defaultHandlers = make(map[models.TaskType]TaskHandler)

func DefaultHandlers() map[models.TaskType]TaskHandler {
	return defaultHandlers
}
