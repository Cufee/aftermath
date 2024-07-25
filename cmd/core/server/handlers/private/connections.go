package private

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/log"
)

type connectionsPayload struct {
	AccountID string `json:"accountId"`
	UserID    string `json:"userId"`
}

/*
This function is temporary and does not need to be good :D
*/
func ImportConnections(client core.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var imports []connectionsPayload
		err := json.NewDecoder(r.Body).Decode(&imports)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, connection := range imports {
			user, err := client.Database().GetOrCreateUserByID(context.Background(), connection.UserID, database.WithConnections())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			existing, ok := user.Connection(models.ConnectionTypeWargaming, nil)
			if !ok {
				existing.UserID = connection.UserID
				existing.Type = models.ConnectionTypeWargaming
			}
			existing.Metadata = map[string]any{"default": true}
			existing.ReferenceID = connection.AccountID
			_, err = client.Database().UpsertUserConnection(context.Background(), existing)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		log.Info().Int("count", len(imports)).Msg("finished importing connections")
	}
}
