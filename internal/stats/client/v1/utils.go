package client

import (
	"context"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/logic"
)

func recordAccountSnapshots(wargaming wargaming.Client, database database.Client, accountID, referenceID string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	realm, ok := wargaming.RealmFromID(accountID)
	if !ok {
		log.Error().Str("accountId", accountID).Msg("invalid account realm")
		return
	}

	_, err := logic.RecordAccountSnapshots(ctx, wargaming, database, realm, logic.WithReference(accountID, referenceID))
	if err != nil {
		log.Err(err).Str("accountId", accountID).Msg("failed to record account snapshot")
	}
}

func getSubscriptions(db database.UsersClient, requestor models.User, accountID string) ([]models.UserSubscription, error) {
	return nil, nil
}
