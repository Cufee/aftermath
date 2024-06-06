package database

import (
	"context"
	"time"

	"github.com/cufee/aftermath/internal/database/prisma/db"
	"github.com/cufee/aftermath/internal/localization"
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/rs/zerolog/log"
	"github.com/steebchen/prisma-client-go/runtime/transaction"
)

type GlossaryVehicle struct {
	db.VehicleModel
}

func (v GlossaryVehicle) Name(printer localization.Printer) string {
	return printer("name")
}

func (c *client) UpsertVehicleAverages(ctx context.Context, averages map[string]frame.StatsFrame) error {
	if len(averages) < 1 {
		return nil
	}

	var transactions []transaction.Transaction
	for id, data := range averages {
		encoded, err := data.Encode()
		if err != nil {
			log.Err(err).Str("id", id).Msg("failed to encode a stats frame for vehicle averages")
			continue
		}

		transactions = append(transactions, c.Raw.VehicleAverage.
			UpsertOne(db.VehicleAverage.ID.Equals(id)).
			Create(
				db.VehicleAverage.ID.Set(id),
				db.VehicleAverage.DataEncoded.Set(encoded),
			).
			Update(
				db.VehicleAverage.DataEncoded.Set(encoded),
			).Tx(),
		)
	}

	return c.Raw.Prisma.Transaction(transactions...).Exec(ctx)
}

func (c *client) GetVehicleAverages(ctx context.Context, ids []string) (map[string]frame.StatsFrame, error) {
	if len(ids) < 1 {
		return nil, nil
	}

	qCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	records, err := c.Raw.VehicleAverage.FindMany(db.VehicleAverage.ID.In(ids)).Exec(qCtx)
	if err != nil {
		return nil, err
	}

	averages := make(map[string]frame.StatsFrame)
	var badRecords []string
	var lastErr error

	for _, record := range records {
		parsed, err := frame.DecodeStatsFrame(record.DataEncoded)
		lastErr = err
		if err != nil {
			badRecords = append(badRecords, record.ID)
			continue
		}
		averages[record.ID] = parsed
	}

	if len(badRecords) == len(ids) || len(badRecords) == 0 {
		return averages, lastErr
	}

	go func() {
		// one bad record should not break the whole query since this data is optional
		// we can just delete the record and move on
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		_, err := c.Raw.VehicleAverage.FindMany(db.VehicleAverage.ID.In(badRecords)).Delete().Exec(ctx)
		if err != nil {
			log.Err(err).Strs("ids", badRecords).Msg("failed to delete a bad vehicle average records")
		}
	}()

	return averages, nil
}
