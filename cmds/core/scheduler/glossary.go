package scheduler

import (
	"context"
	"time"

	"github.com/cufee/aftermath/cmds/core"
	"github.com/rs/zerolog/log"
)

// CurrentTankAverages

func UpdateAveragesWorker(client core.Client) func() {
	return func() {
		log.Info().Msg("updating tank averages cache")

		ctx, cancel := context.WithTimeout(context.Background(), time.Minute*1)
		defer cancel()

		// we just run the logic directly as it's not a heavy task and it doesn't matter if it fails
		averages, err := client.Fetch().CurrentTankAverages(ctx)
		if err != nil {
			log.Err(err).Msg("failed to update averages cache")
			return
		}

		err = client.Database().UpsertVehicleAverages(ctx, averages)
		if err != nil {
			log.Err(err).Msg("failed to update averages cache")
			return
		}

		log.Info().Msg("averages cache updated")
	}
}

func UpdateGlossaryWorker(client core.Client) func() {
	return func() {
		// // We just run the logic directly as it's not a heavy task and it doesn't matter if it fails due to the app failing
		// log.Info().Msg("updating glossary cache")
		// err := cache.UpdateGlossaryCache()
		// if err != nil {
		// 	log.Err(err).Msg("failed to update glossary cache")
		// } else {
		// 	log.Info().Msg("glossary cache updated")
		// }
	}
}
