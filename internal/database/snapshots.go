package database

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/cufee/aftermath/internal/database/prisma/db"
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/rs/zerolog/log"
	"github.com/steebchen/prisma-client-go/runtime/transaction"
)

type getSnapshotQuery struct {
	vehicleIDs []string

	createdAfter  *time.Time
	createdBefore *time.Time
}

func (s getSnapshotQuery) accountParams(accountID, referenceID string, kind snapshotType) []db.AccountSnapshotWhereParam {
	var params []db.AccountSnapshotWhereParam
	params = append(params, db.AccountSnapshot.Type.Equals(string(kind)))
	params = append(params, db.AccountSnapshot.AccountID.Equals(accountID))
	params = append(params, db.AccountSnapshot.ReferenceID.Equals(referenceID))

	if s.createdAfter != nil {
		params = append(params, db.AccountSnapshot.CreatedAt.After(*s.createdAfter))
	}
	if s.createdBefore != nil {
		params = append(params, db.AccountSnapshot.CreatedAt.Before(*s.createdBefore))
	}

	return params
}

func (s getSnapshotQuery) vehiclesQuery(accountID, referenceID string, kind snapshotType) (query string, params []interface{}) {
	var conditions []string
	var args []interface{}

	// Mandatory conditions
	conditions = append(conditions, "type = ?")
	args = append(args, kind)

	conditions = append(conditions, "accountId = ?")
	args = append(args, accountID)

	conditions = append(conditions, "referenceId = ?")
	args = append(args, referenceID)

	// Optional conditions
	if s.createdAfter != nil {
		conditions = append(conditions, "createdAt > ?")
		args = append(args, *s.createdAfter)
	}
	if s.createdBefore != nil {
		conditions = append(conditions, "createdAt < ?")
		args = append(args, *s.createdBefore)
	}

	// Filter by vehicle IDs if provided
	if len(s.vehicleIDs) > 0 {
		placeholders := make([]string, len(s.vehicleIDs))
		for i, id := range s.vehicleIDs {
			placeholders[i] = "?"
			args = append(args, id)
		}
		conditions = append(conditions, fmt.Sprintf("vehicleId IN (%s)", strings.Join(placeholders, ",")))
	}

	// Determine the order by clause
	var orderBy string = "createdAt ASC"
	if s.createdBefore != nil {
		orderBy = "createdAt DESC"
	}

	// Base query
	query = `
		SELECT
			id, createdAt, type, lastBattleTime, accountId, vehicleId, referenceId, battles, frameEncoded
		FROM
			vehicle_snapshots
		WHERE
			%s
		ORDER BY
			%s
	`

	// Combine conditions into a single string
	conditionsStr := strings.Join(conditions, " AND ")
	query = fmt.Sprintf(query, conditionsStr, orderBy)

	// Wrap the query to select the latest or earliest snapshot per vehicleId
	wrappedQuery := `
		SELECT * FROM (
			%s
		) AS ordered_snapshots
		GROUP BY vehicleId
	`

	query = fmt.Sprintf(wrappedQuery, query)
	return query, args
}

func (s getSnapshotQuery) accountOrder() []db.AccountSnapshotOrderByParam {
	var order []db.AccountSnapshotOrderByParam

	switch {
	case s.createdAfter != nil:
		order = append(order, db.AccountSnapshot.CreatedAt.Order(db.ASC))
	default:
		order = append(order, db.AccountSnapshot.CreatedAt.Order(db.DESC))
	}

	return order
}

type SnapshotQuery func(*getSnapshotQuery)

func WithVehicleIDs(ids []string) SnapshotQuery {
	return func(q *getSnapshotQuery) {
		q.vehicleIDs = ids
	}
}
func WithCreatedAfter(after time.Time) SnapshotQuery {
	return func(q *getSnapshotQuery) {
		q.createdAfter = &after
	}
}
func WithCreatedBefore(before time.Time) SnapshotQuery {
	return func(q *getSnapshotQuery) {
		q.createdBefore = &before
	}
}

type snapshotType string

const (
	SnapshotTypeLive  snapshotType = "live"
	SnapshotTypeDaily snapshotType = "daily"
)

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

func (s VehicleSnapshot) FromModel(model db.VehicleSnapshotModel) (VehicleSnapshot, error) {
	s.ID = model.ID
	s.Type = snapshotType(model.Type)
	s.CreatedAt = model.CreatedAt
	s.LastBattleTime = model.LastBattleTime

	s.AccountID = model.AccountID
	s.VehicleID = model.VehicleID
	s.ReferenceID = model.ReferenceID

	stats, err := frame.DecodeStatsFrame(model.FrameEncoded)
	if err != nil {
		return VehicleSnapshot{}, err
	}
	s.Stats = stats
	return s, nil
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

func (c *client) GetVehicleSnapshots(ctx context.Context, accountID, referenceID string, kind snapshotType, options ...SnapshotQuery) ([]VehicleSnapshot, error) {
	var query getSnapshotQuery
	for _, apply := range options {
		apply(&query)
	}

	var models []db.VehicleSnapshotModel
	raw, args := query.vehiclesQuery(accountID, referenceID, kind)
	err := c.prisma.Prisma.Raw.QueryRaw(raw, args...).Exec(ctx, &models)
	if err != nil {
		return nil, err
	}

	var snapshots []VehicleSnapshot
	for _, model := range models {
		vehicle, err := VehicleSnapshot{}.FromModel(model)
		if err != nil {
			return nil, err
		}
		snapshots = append(snapshots, vehicle)
	}

	return snapshots, nil
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

func (s AccountSnapshot) FromModel(model *db.AccountSnapshotModel) (AccountSnapshot, error) {
	s.ID = model.ID
	s.Type = snapshotType(model.Type)
	s.CreatedAt = model.CreatedAt
	s.AccountID = model.AccountID
	s.ReferenceID = model.ReferenceID
	s.LastBattleTime = model.LastBattleTime

	rating, err := frame.DecodeStatsFrame(model.RatingFrameEncoded)
	if err != nil {
		return AccountSnapshot{}, err
	}
	s.RatingBattles = rating

	regular, err := frame.DecodeStatsFrame(model.RegularFrameEncoded)
	if err != nil {
		return AccountSnapshot{}, err
	}
	s.RegularBattles = regular

	return s, nil
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

func (c *client) GetAccountSnapshot(ctx context.Context, accountID, referenceID string, kind snapshotType, options ...SnapshotQuery) (AccountSnapshot, error) {
	var query getSnapshotQuery
	for _, apply := range options {
		apply(&query)
	}

	model, err := c.prisma.AccountSnapshot.FindFirst(query.accountParams(accountID, referenceID, kind)...).OrderBy(query.accountOrder()...).Exec(ctx)
	if err != nil {
		return AccountSnapshot{}, err
	}

	return AccountSnapshot{}.FromModel(model)
}
