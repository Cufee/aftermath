package database

import (
	"context"
	"time"

	"github.com/cufee/aftermath/internal/database/prisma/db"
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/rs/zerolog/log"
	"github.com/steebchen/prisma-client-go/runtime/transaction"
)

type snapshotType string

const (
	SnapshotTypeDaily snapshotType = "daily"
)

// model VehicleSnapshot {
//   id        String   @id @default(cuid())
//   createdAt DateTime

//   type           String
//   lastBattleTime DateTime

//   accountId   String
//   vehicleId   String
//   referenceId String

//   frameEncoded Bytes

//   @@index([createdAt])
//   @@index([accountId, vehicleId, lastBattleTime])
//   @@map("vehicle_snapshots")
// }

type VehicleSnapshot struct {
	ID        string
	CreatedAt time.Time

	Type           snapshotType
	LastBattleTime time.Time

	AccountID   string
	VehicleID   string
	ReferenceID string

	Stats frame.StatsFrame
}

func (s *VehicleSnapshot) FromModel(model db.VehicleSnapshotModel) error {
	s.ID = model.ID
	s.Type = snapshotType(model.Type)
	s.CreatedAt = model.CreatedAt
	s.LastBattleTime = model.LastBattleTime

	s.AccountID = model.AccountID
	s.VehicleID = model.VehicleID
	s.ReferenceID = model.ReferenceID

	stats, err := frame.DecodeStatsFrame(model.FrameEncoded)
	if err != nil {
		return err
	}
	s.Stats = stats
	return nil
}

type AccountSnapshot struct {
	ID             string
	Type           snapshotType
	CreatedAt      time.Time
	AccountID      string
	ReferenceID    string
	LastBattleTime time.Time
	RatingBattles  frame.StatsFrame
	RegularBattles frame.StatsFrame
}

func (s *AccountSnapshot) FromModel(model db.AccountSnapshotModel) error {
	s.ID = model.ID
	s.Type = snapshotType(model.Type)
	s.CreatedAt = model.CreatedAt
	s.AccountID = model.AccountID
	s.ReferenceID = model.ReferenceID
	s.LastBattleTime = model.LastBattleTime

	rating, err := frame.DecodeStatsFrame(model.RatingFrameEncoded)
	if err != nil {
		return err
	}
	s.RatingBattles = rating

	regular, err := frame.DecodeStatsFrame(model.RegularFrameEncoded)
	if err != nil {
		return err
	}
	s.RegularBattles = regular

	return nil
}

func (c *client) CreateVehicleSnapshots(ctx context.Context, snapshots ...VehicleSnapshot) error {
	if len(snapshots) < 1 {
		return nil
	}

	var transactions []transaction.Transaction
	for _, data := range snapshots {
		encoded, err := data.Stats.Encode()
		if err != nil {
			log.Err(err).Str("accountId", data.AccountID).Str("vehicleId", data.VehicleID).Msg("failed to encode a stats frame for vehicle snapthsot")
			continue
		}

		transactions = append(transactions, c.prisma.VehicleSnapshot.
			CreateOne(
				db.VehicleSnapshot.CreatedAt.Set(data.CreatedAt),
				db.VehicleSnapshot.Type.Set(string(data.Type)),
				db.VehicleSnapshot.LastBattleTime.Set(data.LastBattleTime),
				db.VehicleSnapshot.AccountID.Set(data.AccountID),
				db.VehicleSnapshot.VehicleID.Set(data.VehicleID),
				db.VehicleSnapshot.ReferenceID.Set(data.ReferenceID),
				db.VehicleSnapshot.Battles.Set(int(data.Stats.Battles)),
				db.VehicleSnapshot.FrameEncoded.Set(encoded),
			).Tx(),
		)
	}

	return c.prisma.Prisma.Transaction(transactions...).Exec(ctx)
}

func (c *client) CreateAccountSnapshots(ctx context.Context, snapshots ...AccountSnapshot) error {
	if len(snapshots) < 1 {
		return nil
	}

	var transactions []transaction.Transaction
	for _, data := range snapshots {
		ratingEncoded, err := data.RatingBattles.Encode()
		if err != nil {
			log.Err(err).Str("accountId", data.AccountID).Msg("failed to encode rating stats frame for account snapthsot")
			continue
		}
		regularEncoded, err := data.RegularBattles.Encode()
		if err != nil {
			log.Err(err).Str("accountId", data.AccountID).Msg("failed to encode regular stats frame for account snapthsot")
			continue
		}

		transactions = append(transactions, c.prisma.AccountSnapshot.
			CreateOne(
				db.AccountSnapshot.CreatedAt.Set(data.CreatedAt),
				db.AccountSnapshot.Type.Set(string(data.Type)),
				db.AccountSnapshot.LastBattleTime.Set(data.LastBattleTime),
				db.AccountSnapshot.AccountID.Set(data.AccountID),
				db.AccountSnapshot.ReferenceID.Set(data.ReferenceID),
				db.AccountSnapshot.RatingBattles.Set(int(data.RatingBattles.Battles)),
				db.AccountSnapshot.RatingFrameEncoded.Set(ratingEncoded),
				db.AccountSnapshot.RegularBattles.Set(int(data.RegularBattles.Battles)),
				db.AccountSnapshot.RegularFrameEncoded.Set(regularEncoded),
			).Tx(),
		)
	}

	return c.prisma.Prisma.Transaction(transactions...).Exec(ctx)
}
