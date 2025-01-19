package client

import (
	"context"
	"time"

	"github.com/cufee/aftermath/internal/stats/client/common"
	"github.com/cufee/aftermath/internal/stats/prepare/session/v1"
)

func (r *client) SessionCards(ctx context.Context, accountId string, from time.Time, opts ...common.RequestOption) (session.Cards, common.Metadata, error) {
	return r.v1.SessionCards(ctx, accountId, from, opts...)
}
func (r *client) SessionImage(ctx context.Context, accountId string, from time.Time, opts ...common.RequestOption) (common.Image, common.Metadata, error) {
	return r.v1.SessionImage(ctx, accountId, from, opts...)
}
func (r *client) EmptySessionCards(ctx context.Context, accountId string) (session.Cards, common.Metadata, error) {
	return r.v1.EmptySessionCards(ctx, accountId)
}
