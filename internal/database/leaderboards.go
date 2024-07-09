package database

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/cufee/aftermath/internal/database/ent/db"
	"github.com/cufee/aftermath/internal/database/ent/db/leaderboardscore"
	"github.com/cufee/aftermath/internal/database/models"
)

func toScore(r *db.LeaderboardScore) models.LeaderboardScore {
	return models.LeaderboardScore{
		ID:            r.ID,
		Type:          r.Type,
		CreatedAt:     r.CreatedAt,
		UpdatedAt:     r.UpdatedAt,
		AccountID:     r.AccountID,
		ReferenceID:   r.ReferenceID,
		LeaderboardID: r.LeaderboardID,
		Score:         r.Score,
		Meta:          r.Meta,
	}
}

func (c *client) CreateLeaderboardScores(ctx context.Context, scores ...models.LeaderboardScore) (map[string]error, error) {
	if len(scores) < 1 {
		return nil, nil
	}

	var errors = make(map[string]error)
	for _, score := range scores {
		// make a transaction per write to avoid locking for too long
		err := c.withTx(ctx, func(tx *db.Tx) error {
			return c.db.LeaderboardScore.Create().
				SetMeta(score.Meta).
				SetType(score.Type).
				SetScore(score.Score).
				SetCreatedAt(score.CreatedAt).
				SetReferenceID(score.ReferenceID).
				SetLeaderboardID(score.LeaderboardID).
				Exec(ctx)
		})
		if err != nil {
			errors[score.ReferenceID] = err
		}
	}

	if len(errors) > 0 {
		return errors, nil
	}

	return nil, nil
}

func (c *client) GetLeaderboardScores(ctx context.Context, leaderboardID string, scoreType models.ScoreType, opts ...Query) ([]models.LeaderboardScore, error) {
	var query baseQueryOptions
	for _, apply := range opts {
		apply(&query)
	}

	var where []*sql.Predicate
	where = append(where, sql.EQ(leaderboardscore.FieldType, scoreType))
	where = append(where, sql.EQ(leaderboardscore.FieldLeaderboardID, leaderboardID))

	orderBy := sql.Desc(leaderboardscore.FieldCreatedAt)
	if query.createdAfter != nil {
		where = append(where, sql.GT(leaderboardscore.FieldCreatedAt, *query.createdAfter))
		orderBy = sql.Asc(leaderboardscore.FieldCreatedAt)
	}
	if query.createdBefore != nil {
		where = append(where, sql.LT(leaderboardscore.FieldCreatedAt, *query.createdBefore))
		orderBy = sql.Desc(leaderboardscore.FieldCreatedAt)
	}

	if in := query.refIDIn(); in != nil {
		where = append(where, sql.In(leaderboardscore.FieldReferenceID, in...))
	}
	if nin := query.refIDNotIn(); nin != nil {
		where = append(where, sql.NotIn(leaderboardscore.FieldReferenceID, nin...))
	}

	selectFields := leaderboardscore.Columns
	if fields := query.selectFields(leaderboardscore.FieldAccountID); fields != nil {
		selectFields = fields
	}

	q := sql.Select(selectFields...).From(sql.Table(leaderboardscore.Table))
	q = q.Where(sql.And(where...))
	q = q.OrderBy(orderBy)

	innerQueryString, queryArgs := q.Query()
	queryString, _ := sql.Select(selectFields...).FromExpr(wrap(innerQueryString)).GroupBy(leaderboardscore.FieldAccountID).Query()
	rows, err := c.db.LeaderboardScore.QueryContext(ctx, queryString, queryArgs...)
	if err != nil {
		return nil, err
	}

	records, err := rowsToRecords[*db.LeaderboardScore](rows, selectFields)
	if err != nil {
		return nil, err
	}

	var scores []models.LeaderboardScore
	for _, r := range records {
		scores = append(scores, toScore(r))
	}
	return scores, nil
}

func (c *client) DeleteExpiredLeaderboardScores(ctx context.Context, expiration time.Time, scoreType models.ScoreType) error {
	_, err := c.db.LeaderboardScore.Delete().Where(leaderboardscore.CreatedAtLT(expiration), leaderboardscore.TypeEQ(scoreType)).Exec(ctx)
	return err
}
