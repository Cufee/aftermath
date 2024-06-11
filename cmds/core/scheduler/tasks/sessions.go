package tasks

import (
	"context"
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/cufee/aftermath/cmds/core"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/retry"
	"github.com/cufee/aftermath/internal/stats/fetch"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/rs/zerolog/log"
)

type vehicleResponseData map[string][]types.VehicleStatsFrame

func init() {
	defaultHandlers[database.TaskTypeRecordSessions] = TaskHandler{
		Process: func(client core.Client, task database.Task) (string, error) {
			if task.Data == nil {
				return "no data provided", errors.New("no data provided")
			}
			realm, ok := task.Data["realm"].(string)
			if !ok {
				task.Data["triesLeft"] = int(0) // do not retry
				return "invalid realm", errors.New("invalid realm")
			}
			if len(task.Targets) > 100 {
				task.Data["triesLeft"] = int(0) // do not retry
				return "cannot process 100+ accounts at a time", errors.New("invalid targets length")
			}
			if len(task.Targets) < 1 {
				task.Data["triesLeft"] = int(0) // do not retry
				return "targed ids cannot be left blank", errors.New("invalid targets length")
			}
			forceUpdate, _ := task.Data["force"].(bool)

			log.Debug().Str("taskId", task.ID).Any("targets", task.Targets).Msg("started working on a session refresh task")

			createdAt := time.Now()
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
			defer cancel()

			accounts, err := client.Wargaming().BatchAccountByID(ctx, realm, task.Targets)
			if err != nil {
				return "failed to fetch accounts", err
			}

			// Make a new slice just in case some accounts were not returned/are private
			var validAccouts []string
			for _, id := range task.Targets {
				data, ok := accounts[id]
				if !ok {
					go func(id string) {
						log.Debug().Str("accountId", id).Str("taskId", task.ID).Msg("account is private")
						// update account cache (if it exists) to set account as private
						ctx, cancel := context.WithTimeout(context.Background(), time.Second)
						defer cancel()
						err := client.Database().AccountSetPrivate(ctx, id, true)
						if err != nil {
							log.Err(err).Str("accountId", id).Msg("failed to set account status as private")
						}
					}(id)
					continue
				}
				if !forceUpdate && data.LastBattleTime < int(createdAt.Add(time.Hour*-25).Unix()) {
					// if the last battle was played 25+ hours ago, there is nothing for us to update
					log.Debug().Str("accountId", id).Str("taskId", task.ID).Msg("account played no battles")
					continue
				}
				validAccouts = append(validAccouts, id)
			}

			if len(validAccouts) == 0 {
				return "no updates required due to last battle or private accounts status", nil
			}

			// clans are options-ish
			clans, err := client.Wargaming().BatchAccountClan(ctx, realm, validAccouts)
			if err != nil {
				log.Err(err).Msg("failed to get batch account clans")
				clans = make(map[string]types.ClanMember)
			}

			vehicleCh := make(chan retry.DataWithErr[vehicleResponseData], len(validAccouts))
			var group sync.WaitGroup
			group.Add(len(validAccouts))
			for _, id := range validAccouts {
				go func(id string) {
					defer group.Done()
					data, err := client.Wargaming().AccountVehicles(ctx, realm, id)
					vehicleCh <- retry.DataWithErr[vehicleResponseData]{Data: vehicleResponseData{id: data}, Err: err}
				}(id)
			}
			group.Wait()
			close(vehicleCh)

			var withErrors []string
			var accountUpdates []database.Account
			var snapshots []database.AccountSnapshot
			var vehicleSnapshots []database.VehicleSnapshot
			for result := range vehicleCh {
				// there is only 1 key in this map
				for id, vehicles := range result.Data {
					if result.Err != nil {
						withErrors = append(withErrors, id)
						continue
					}

					stats := fetch.WargamingToStats(realm, accounts[id], clans[id], vehicles)

					accountUpdates = append(accountUpdates, database.Account{
						Realm:    stats.Realm,
						ID:       stats.Account.ID,
						Nickname: stats.Account.Nickname,

						Private:        false,
						CreatedAt:      stats.Account.CreatedAt,
						LastBattleTime: stats.LastBattleTime,

						ClanID:  stats.Account.ClanID,
						ClanTag: stats.Account.ClanTag,
					})

					snapshots = append(snapshots, database.AccountSnapshot{
						Type:           database.SnapshotTypeDaily,
						CreatedAt:      createdAt,
						AccountID:      stats.Account.ID,
						ReferenceID:    stats.Account.ID,
						LastBattleTime: stats.LastBattleTime,
						RatingBattles:  stats.RatingBattles.StatsFrame,
						RegularBattles: stats.RegularBattles.StatsFrame,
					})
					if len(vehicles) < 1 {
						continue
					}

					vehicleStats := fetch.WargamingVehiclesToFrame(vehicles)
					for _, vehicle := range vehicleStats {
						if !forceUpdate && vehicle.LastBattleTime.Before(createdAt.Add(time.Hour*-25)) {
							// if the last battle was played 25+ hours ago, there is nothing for us to update
							log.Debug().Str("accountId", id).Str("vehicleId", vehicle.VehicleID).Str("taskId", task.ID).Msg("vehicle played no battles")
							continue
						}
						vehicleSnapshots = append(vehicleSnapshots, database.VehicleSnapshot{
							CreatedAt:      createdAt,
							Type:           database.SnapshotTypeDaily,
							LastBattleTime: vehicle.LastBattleTime,
							AccountID:      stats.Account.ID,
							VehicleID:      vehicle.VehicleID,
							ReferenceID:    stats.Account.ID,
							Stats:          *vehicle.StatsFrame,
						})
					}
				}
			}

			err = client.Database().CreateAccountSnapshots(ctx, snapshots...)
			if err != nil {
				return "failed to save account snapshots to database", err
			}

			err = client.Database().CreateVehicleSnapshots(ctx, vehicleSnapshots...)
			if err != nil {
				return "failed to save vehicle snapshots to database", err
			}

			if len(withErrors) == 0 {
				return "finished session update on all accounts", nil
			}

			// Retry failed accounts
			task.Targets = withErrors
			return "retrying failed accounts", errors.New("some accounts failed")
		},
		ShouldRetry: func(task *database.Task) bool {
			triesLeft, ok := task.Data["triesLeft"].(int)
			if !ok {
				return false
			}
			if triesLeft <= 0 {
				return false
			}

			triesLeft -= 1
			task.Data["triesLeft"] = triesLeft
			task.ScheduledAfter = time.Now().Add(5 * time.Minute) // Backoff for 5 minutes to avoid spamming
			return true
		},
	}
}

func CreateSessionUpdateTasks(client core.Client, realm string) error {
	realm = strings.ToUpper(realm)
	task := database.Task{
		Type:           database.TaskTypeRecordSessions,
		ReferenceID:    "realm_" + realm,
		ScheduledAfter: time.Now(),
		Data: map[string]any{
			"realm":     realm,
			"triesLeft": int(3),
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	accounts, err := client.Database().GetRealmAccounts(ctx, realm)
	if err != nil {
		return err
	}
	if len(accounts) < 1 {
		return nil
	}
	for _, account := range accounts {
		task.Targets = append(task.Targets, account.ID)
	}

	// This update requires (2 + n) requests per n players
	return client.Database().CreateTasks(ctx, splitTaskByTargets(task, 90)...)
}
