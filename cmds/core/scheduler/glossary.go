package scheduler

import (
	"context"
	"time"

	"github.com/cufee/aftermath/cmds/core"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/rs/zerolog/log"
	"golang.org/x/text/language"
)

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
