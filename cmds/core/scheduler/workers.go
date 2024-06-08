package scheduler

import (
	"context"
	"time"

	"github.com/cufee/aftermath/cmds/core"
	"github.com/cufee/aftermath/cmds/core/scheduler/tasks"
	"github.com/rs/zerolog/log"
)

func RotateBackgroundPresetsWorker(client core.Client) func() {
	return func() {
		// // We just run the logic directly as it's not a heavy task and it doesn't matter if it fails due to the app failing
		// log.Info().Msg("rotating background presets")
		// images, err := content.PickRandomBackgroundImages(3)
		// if err != nil {
		// 	log.Err(err).Msg("failed to pick random background images")
		// 	return
		// }
		// err = database.UpdateAppConfiguration[[]string]("backgroundImagesSelection", images, nil, true)
		// if err != nil {
		// 	log.Err(err).Msg("failed to update background images selection")
		// }
	}
}

func CreateSessionTasksWorker(client core.Client, realm string) func() {
	return func() {
		err := tasks.CreateSessionUpdateTasks(client, realm)
		if err != nil {
			log.Err(err).Str("realm", realm).Msg("failed to schedule session update tasks")
		}
	}
}

func RunTasksWorker(queue *Queue) func() {
	return func() {
		activeWorkers := queue.ActiveWorkers()
		if activeWorkers >= queue.concurrencyLimit {
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		// each task worked handles 1 task at a time, but tasks might be very fast
		// for now, we queue up 10 tasks per worker, this can be adjusted later/smarter
		batchSize := queue.concurrencyLimit - activeWorkers
		tasks, err := queue.core.Database().GetAndStartTasks(ctx, batchSize*10)
		if err != nil {
			log.Err(err).Msg("failed to start scheduled tasks")
			return
		}

		queue.Process(func(err error) {
			if err != nil {
				log.Err(err).Msg("failed to process tasks")
				return
			}
			// if the queue is now empty, we can run the next batch of tasks right away
			RunTasksWorker(queue)
		}, tasks...)
	}
}

func RestartTasksWorker(queue *Queue) func() {
	return func() {
		// _, err := tasks.RestartAbandonedTasks(nil)
		// if err != nil {
		// 	log.Err(err).Msg("failed to start scheduled tasks")
		// 	return
		// }
	}
}
