package scheduler

import (
	"github.com/cufee/aftermath/cmds/core"
	"github.com/cufee/aftermath/cmds/core/scheduler/tasks"
)

func rotateBackgroundPresetsWorker(client core.Client) func() {
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

func createSessionTasksWorker(client core.Client, realm string) func() {
	return func() {
		// err := tasks.CreateSessionUpdateTasks(realm)
		// if err != nil {
		// 	log.Err(err).Msg("failed to create session update tasks")
		// }
	}
}

func runTasksWorker(queue *tasks.Queue) func() {
	return func() {
		// 	if tasks.DefaultQueue.ActiveWorkers() > 0 {
		// 		return
		// 	}

		// 	activeTasks, err := tasks.StartScheduledTasks(nil, 50)
		// 	if err != nil {
		// 		log.Err(err).Msg("failed to start scheduled tasks")
		// 		return
		// 	}
		// 	if len(activeTasks) == 0 {
		// 		return
		// 	}

		// 	tasks.DefaultQueue.Process(func(err error) {
		// 		if err != nil {
		// 			log.Err(err).Msg("failed to process tasks")
		// 			return
		// 		}

		// 		// If the queue is now empty, we can run the next batch of tasks right away
		// 		runTasksWorker()

		// 	}, activeTasks...)
	}
}

func restartTasksWorker(queue *tasks.Queue) func() {
	return func() {
		// _, err := tasks.RestartAbandonedTasks(nil)
		// if err != nil {
		// 	log.Err(err).Msg("failed to start scheduled tasks")
		// 	return
		// }
	}
}
