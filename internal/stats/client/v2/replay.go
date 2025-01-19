package client

import (
	"context"

	"github.com/cufee/aftermath/internal/stats/client/common"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/replay/v1"
)

func (r *client) ReplayCards(ctx context.Context, replayURL string, o ...common.RequestOption) (prepare.Cards, common.Metadata, error) {
	return r.v1.ReplayCards(ctx, replayURL, o...)
}

func (r *client) ReplayImage(ctx context.Context, replayURL string, o ...common.RequestOption) (common.Image, common.Metadata, error) {
	return r.v1.ReplayImage(ctx, replayURL, o...)
}
