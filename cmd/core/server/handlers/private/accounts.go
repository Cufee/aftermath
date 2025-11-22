package private

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/cufee/aftermath/internal/json"

	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	"golang.org/x/sync/semaphore"
)

/*
This function is temporary and does not need to be good :D
*/
func LoadAccountsHandler(client core.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var accounts []string
		err := json.NewDecoder(r.Body).Decode(&accounts)
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
		log.Debug().Int("count", len(accounts)-len(existing)).Msg("importing accounts")

		go func(accounts []string, existing []models.Account) {
			existingMap := make(map[string]struct{}, len(existing))
			for _, a := range existing {
				existingMap[a.ID] = struct{}{}
			}

			accountsByRealm := make(map[types.Realm][]string, len(accounts))
			for _, a := range accounts {
				if _, ok := existingMap[a]; ok {
					continue
				}

				realm, ok := client.Wargaming().RealmFromID(a)
				if !ok {
					continue
				}

				accountsByRealm[realm] = append(accountsByRealm[realm], a)
			}

			batchSize := 50
			var wg sync.WaitGroup
			sem := semaphore.NewWeighted(5)
			for realm, accounts := range accountsByRealm {
				for i := 0; i < len(accounts); i += batchSize {
					end := i + batchSize
					if end > len(accounts) {
						end = len(accounts)
					}

					wg.Add(1)
					go func(accounts []string, realm types.Realm) {
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

						var inserts []*models.Account
						for id, account := range data {
							if id == "" && account.ID == 0 {
								log.Warn().Str("reason", "id is blank").Msg("wargaming returned a bad account")
								continue
							}

							var private bool
							if account.ID == 0 {
								account.ID, _ = strconv.Atoi(id)
								private = true
							}
							if account.Nickname == "" {
								account.Nickname = "@Hidden"
								private = true
							}

							update := fetch.WargamingToAccount(realm, account, types.ClanMember{}, private)
							inserts = append(inserts, &update)
						}

						err = client.Database().UpsertAccounts(ctx, inserts...)
						if err != nil {
							log.Err(err).Msg("failed to upsert accounts")
						}
					}(accounts[i:end], realm)
				}
			}
			wg.Wait()

			log.Debug().Int("count", len(accounts)-len(existing)).Msg("finished importing accounts")
		}(accounts, existing)
	}
}
