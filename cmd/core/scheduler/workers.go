package scheduler

import (
	"context"
	"time"

	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/cmd/core/tasks"
	"golang.org/x/text/language"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/rs/zerolog/log"
)

func CreateCleanupTaskWorker(client core.Client) func() {
	return func() {
		err := tasks.CreateCleanupTasks(client)
		if err != nil {
			log.Err(err).Msg("failed to schedule a cleanup tasks")
		}
	}
}

func RotateBackgroundPresetsWorker(client core.Client) func() {
	return func() {
		// // We just run the logic directly as it's not a heavy task and it doesn't matter if it fails due to the app failing
		// log.Info().Msg("rotating background presets")
		// images, err := content.PickRandomBackgroundImages(3)
		// if err != nil {
		// 	log.Err(err).Msg("failed to pick random background images")
		// 	return
		// }
		// err = models.UpdateAppConfiguration[[]string]("backgroundImagesSelection", images, nil, true)
		// if err != nil {
		// 	log.Err(err).Msg("failed to update background images selection")
		// }
	}
}

func CreateSnapshotTasksWorker(client core.Client, realm string) func() {
	return func() {
		err := tasks.CreateRecordSnapshotsTasks(client, realm)
		if err != nil {
			log.Err(err).Str("realm", realm).Msg("failed to schedule session update tasks")
		}
	}
}

func RestartTasksWorker(core core.Client) func() {
	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()

		staleTasks, err := core.Database().GetStaleTasks(ctx, 100)
		if err != nil {
			if database.IsNotFound(err) {
				log.Debug().Msg("no stale tasks found")
				return
			}
			log.Err(err).Msg("failed to reschedule stale tasks")
			return
		}
		if len(staleTasks) < 1 {
			log.Debug().Msg("no stale tasks found")
			return
		}
		log.Debug().Int("count", len(staleTasks)).Msg("fetched stale tasks from database")

		now := time.Now()
		for i, task := range staleTasks {
			task.Status = models.TaskStatusScheduled
			task.ScheduledAfter = now
			staleTasks[i] = task
		}

		log.Debug().Int("count", len(staleTasks)).Msg("updating stale tasks")
		err = core.Database().UpdateTasks(ctx, staleTasks...)
		if err != nil {
			log.Err(err).Msg("failed to update stale tasks")
			return
		}
		log.Debug().Int("count", len(staleTasks)).Msg("rescheduled stale tasks")
	}
}

func UpdateAveragesWorker(client core.Client) func() {
	// we just run the logic directly as it's not a heavy task and it doesn't matter if it fails
	return func() {
		log.Info().Msg("updating tank averages cache")

		ctx, cancel := context.WithTimeout(context.Background(), time.Minute*1)
		defer cancel()

		averages, err := client.Fetch().CurrentTankAverages(ctx)
		if err != nil {
			log.Err(err).Msg("failed to update averages cache")
			return
		}

		aErr, err := client.Database().UpsertVehicleAverages(ctx, averages)
		if err != nil {
			log.Err(err).Msg("failed to update averages cache")
			return
		}
		if len(aErr) > 0 {
			event := log.Error()
			for id, err := range aErr {
				event.Str(id, err.Error())
			}
			event.Msg("failed to update some average cache")
			return
		}

		log.Info().Msg("averages cache updated")
	}
}

func UpdateGlossaryWorker(client core.Client) func() {
	// we just run the logic directly as it's not a heavy task and it doesn't matter if it fails
	return func() {
		log.Info().Msg("updating glossary cache")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()

		glossary, err := client.Wargaming().CompleteVehicleGlossary(ctx, "eu", "en")
		if err != nil {
			log.Err(err).Msg("failed to get vehicle glossary")
			return
		}

		vehicles := make(map[string]models.Vehicle)
		for id, data := range glossary {
			vehicles[id] = models.Vehicle{
				ID:             id,
				Tier:           data.Tier,
				LocalizedNames: map[string]string{language.English.String(): data.Name},
			}
		}

		vErr, err := client.Database().UpsertVehicles(ctx, vehicles)
		if err != nil {
			log.Err(err).Msg("failed to save vehicle glossary")
			return
		}
		if len(vErr) > 0 {
			event := log.Error()
			for id, err := range vErr {
				event.Str(id, err.Error())
			}
			event.Msg("failed to save some vehicle glossaries")
			return
		}

		log.Info().Msg("glossary cache updated")
	}
}

func CreateAchievementLeaderboardTasksWorker(client core.Client, realm string) func() {
	return func() {
		err := tasks.CreateRecordSnapshotsTasks(client, realm)
		if err != nil {
			log.Err(err).Str("realm", realm).Msg("failed to schedule session update tasks")
		}
	}
}
