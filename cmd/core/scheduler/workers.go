package scheduler

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"time"

	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/cmd/core/tasks"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/external/github"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/search"
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

		dctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()

		rc, size, err := github.GetLatestGameAssets(dctx)
		if err != nil {
			log.Err(err).Msg("failed to download latest game assets")
			return
		}

		data, err := io.ReadAll(rc)
		rc.Close()
		if err != nil {
			log.Err(err).Msg("failed to read assets data")
			return
		}

		zr, err := zip.NewReader(bytes.NewReader(data), size)
		if err != nil {
			log.Err(err).Msg("failed to read assets as zip")
			return
		}

		vf, err := zr.Open("assets/vehicles.json")
		if err != nil {
			log.Err(err).Msg("failed to open vehicles.json in assets.zip")
			return
		} else {
			var vehicles map[string]models.Vehicle
			err := json.NewDecoder(vf).Decode(&vehicles)
			if err != nil {
				log.Err(err).Msg("failed to decode vehicles.json")
				return
			}

			// load new glossary into the search memory cache
			search.LoadVehicleCache(vehicles)

			vctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
			defer cancel()

			_, err = client.Database().UpsertVehicles(vctx, vehicles)
			if err != nil {
				log.Err(err).Msg("failed to save vehicle glossary")
				return
			}
		}
		log.Info().Msg("glossary cache updated")
	}
}

func CreateLeaderboardTasksWorker(client core.Client, realm string, scoreType models.ScoreType) func() {
	return func() {
		err := tasks.CreateUpdateLeaderboardsTasks(client, realm, scoreType)
		if err != nil {
			log.Err(err).Str("realm", realm).Msg("failed to schedule leaderboard update tasks")
		}
	}
}
