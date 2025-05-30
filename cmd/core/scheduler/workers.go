package scheduler

import (
	"archive/zip"
	"bytes"
	"context"
	"io"
	"time"

	"github.com/cufee/aftermath-assets/types"
	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/cmd/core/tasks"
	"golang.org/x/text/language"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/external/github"
	"github.com/cufee/aftermath/internal/glossary"
	"github.com/cufee/aftermath/internal/json"
	"github.com/cufee/aftermath/internal/log"
)

func CreateCleanupTaskWorker(client core.Client) func() {
	return func() {
		err := tasks.CreateCleanupTasks(client)
		if err != nil {
			log.Err(err).Msg("failed to schedule a cleanup tasks")
		}
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

		for id, err := range aErr {
			if err != nil {
				log.Err(err).Str("", id).Msg("failed to update some average cache")
			}
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

		// update vehicles
		vf, err := zr.Open("assets/vehicles.json")
		if err != nil {
			log.Err(err).Msg("failed to open vehicles.json in assets.zip")

		}
		var vehicles map[string]models.Vehicle
		err = json.NewDecoder(vf).Decode(&vehicles)
		if err != nil {
			log.Err(err).Msg("failed to decode vehicles.json")
			return
		}

		vctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()

		_, err = client.Database().UpsertVehicles(vctx, vehicles)
		if err != nil {
			log.Err(err).Msg("failed to save vehicle glossary")
			return
		}

		// load new glossary into the search memory cache
		cctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		err = glossary.RefreshVehicleCache(cctx, client.Database())
		if err != nil {
			log.Err(err).Msg("failed to refresh glossary cache")
			return
		}

		// updates maps
		mf, err := zr.Open("assets/maps.json")
		if err != nil {
			log.Err(err).Msg("failed to open maps.json in assets.zip")
			return
		}
		var maps map[string]types.Map
		err = json.NewDecoder(mf).Decode(&maps)
		if err != nil {
			log.Err(err).Msg("failed to decode maps.json")
			return
		}

		mctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		err = client.Database().UpsertMaps(mctx, maps)
		if err != nil {
			log.Err(err).Msg("failed save maps glossary")
			return
		}

		// update game modes
		gmf, err := zr.Open("assets/game_modes.json")
		if err != nil {
			log.Err(err).Msg("failed to open game_modes.json in assets.zip")
			return
		}
		var gameModes map[string]map[string]string
		err = json.NewDecoder(gmf).Decode(&gameModes)
		if err != nil {
			log.Err(err).Msg("failed to decode maps.json")
			return
		}
		var withTags = make(map[string]map[language.Tag]string)
		for key, value := range gameModes {
			localized := make(map[language.Tag]string)
			for locale, name := range value {
				localized[language.MustParse(locale)] = name
			}
			withTags[key] = localized
		}

		gmctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		_, err = client.Database().UpsertGameModes(gmctx, withTags)
		if err != nil {
			log.Err(err).Msg("failed save game modes glossary")
			return
		}

		log.Info().Msg("glossary cache updated")
	}
}

func CleanupPubSubWorker(client core.Client) func() {
	return func() {
		err := client.PubSub().Cleanup()
		if err != nil {
			log.Err(err).Msg("failed to cleanup pubsub")
		}
	}
}
