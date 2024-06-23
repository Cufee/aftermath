package private

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/cufee/aftermath/cmds/core"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/stats/fetch"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/semaphore"
)

/*
This function is temporary and does not need to be good :D
*/
func LoadAccountsHandler(client core.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var accounts []string
		err = json.Unmarshal(payload, &accounts)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		existing, err := client.Database().GetAccounts(context.Background(), accounts)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(fmt.Sprintf("working on %d accounts", len(accounts)-len(existing))))
		log.Info().Int("count", len(accounts)-len(existing)).Msg("importing accounts")

		go func() {
			existingMap := make(map[string]struct{})
			for _, a := range existing {
				existingMap[a.ID] = struct{}{}
			}

			accountsByRealm := make(map[string][]string)
			for _, a := range accounts {
				if _, ok := existingMap[a]; ok {
					continue
				}

				id := client.Wargaming().RealmFromAccountID(a)
				accountsByRealm[id] = append(accountsByRealm[id], a)
			}

			batchSize := 50
			var wg sync.WaitGroup
			sem := semaphore.NewWeighted(5)
			errors := make(map[string]error)
			var errorsMx sync.Mutex

			for realm, accounts := range accountsByRealm {
				for i := 0; i < len(accounts); i += batchSize {
					end := i + batchSize
					if end > len(accounts) {
						end = len(accounts)
					}

					wg.Add(1)
					go func(accounts []string, realm string) {
						defer wg.Done()

						ctx, cancel := context.WithCancel(context.Background())
						defer cancel()

						err := sem.Acquire(ctx, 1)
						if err != nil {
							log.Err(err).Msg("failed to acquire a semaphore")
							return
						}
						defer sem.Release(1)

						data, err := client.Wargaming().BatchAccountByID(ctx, realm, accounts)
						if err != nil {
							log.Err(err).Msg("failed to get accounts from wg")
							return
						}

						var inserts []models.Account
						for _, account := range data {
							inserts = append(inserts, fetch.WargamingToAccount(realm, account, types.ClanMember{}, false))
						}

						accErr, err := client.Database().UpsertAccounts(ctx, inserts)
						if err != nil {
							log.Err(err).Msg("failed to upsert accounts")
						}
						if len(accErr) > 0 {
							errorsMx.Lock()
							for id, err := range accErr {
								errors[id] = err
							}
							errorsMx.Unlock()
						}

					}(accounts[i:end], realm)
				}
			}
			wg.Wait()

			if len(errors) > 0 {
				log.Error().Any("errors", errors).Msg("some account imports failed")
			}
			log.Info().Int("count", len(accounts)-len(existing)).Msg("finished importing accounts")
		}()
	}
}
