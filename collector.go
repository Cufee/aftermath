//go:build ignore

package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/go-co-op/gocron/v2"
	_ "github.com/joho/godotenv/autoload"

	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/json"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/rs/zerolog"
)

func main() {
	if os.Getenv("COLLECTOR_ENABLED") == "false" {
		os.Exit(0)
	}

	// Logger
	level, _ := zerolog.ParseLevel(os.Getenv("COLLECTOR_LOG_LEVEL"))
	log.SetupGlobalLogger(func(l zerolog.Logger) zerolog.Logger {
		return l.Level(level)
	})

	backendApi := os.Getenv("COLLECTOR_BACKEND_URL")
	if backendApi == "" {
		log.Fatal().Msg("backend is a required argument")
	}

	go func() {
		realms := strings.Split(os.Getenv("COLLECTOR_RUN_ON_START"), ",")
		if slices.Contains(realms, types.RealmNorthAmerica.String()) {
			collectRealmIDs(backendApi, types.RealmNorthAmerica)
		}
		if slices.Contains(realms, types.RealmEurope.String()) {
			collectRealmIDs(backendApi, types.RealmEurope)
		}
		if slices.Contains(realms, types.RealmAsia.String()) {
			collectRealmIDs(backendApi, types.RealmAsia)
		}
	}()

	scheduler, err := gocron.NewScheduler(gocron.WithLocation(time.UTC))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create collector scheduler")
	}
	for _, job := range []struct {
		cron  string
		realm types.Realm
	}{
		{"0 0 * * 6", types.RealmAsia},
		{"0 0 * * 0", types.RealmNorthAmerica},
		{"0 0 * * 1", types.RealmEurope},
	} {
		if _, err := scheduler.NewJob(gocron.CronJob(job.cron, false),
			gocron.NewTask(collectRealmIDs, backendApi, job.realm)); err != nil {
			_ = scheduler.Shutdown()
			log.Fatal().Err(err).Str("cron", job.cron).Msg("failed to register collector job")
		}
	}
	log.Info().Msg("started a cron scheduler")
	scheduler.Start()
	select {}
}

func collectRealmIDs(backendApi string, realm types.Realm) {
	log.Info().Str("realm", realm.String()).Msg("started collecting player ids")

	client, err := wargaming.NewRatingLeaderboardClient()
	if err != nil {
		log.Err(err).Str("realm", realm.String()).Msg("failed to create a leaderboard client")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	season, err := client.CurrentSeason(ctx, realm)
	if err != nil {
		log.Err(err).Str("realm", realm.String()).Msg("failed to get current season")
		return
	}

	if len(season.Leagues) < 1 {
		log.Error().Str("realm", realm.String()).Msg("season contains no leagues")
		return
	}

	players, err := client.LeagueTop(ctx, realm, season.Leagues[0].ID)
	if err != nil {
		log.Err(err).Str("realm", realm.String()).Msg("failed to get league top players")
		return
	}

	var initialIDs []int
	for _, p := range players {
		initialIDs = append(initialIDs, p.AccountID)
	}
	go savePlayerIDs(backendApi, initialIDs)

	var total int
	collector := make(chan []int, 1)
	go func() {
		for ids := range collector {
			total += len(ids)
			go savePlayerIDs(backendApi, ids)
			log.Debug().Str("realm", realm.String()).Int("count", len(ids)).Int("total", total).Str("realm", realm.String()).Msg("collected player ids")
		}
	}()
	defer func() {
		close(collector)
	}()

	err = client.CollectPlayerIDs(context.Background(), realm, collector, players[len(initialIDs)-1].AccountID)
	if err != nil {
		log.Err(err).Str("realm", realm.String()).Msg("failed to complete player id collection")
		return
	}

	log.Info().Str("realm", realm.String()).Int("total", total).Msg("finished collecting player ids")
}

var client = http.Client{
	Timeout: time.Second * 5,
}

func savePlayerIDs(apiDomain string, data []int) {
	log.Debug().Int("count", len(data)).Msg("saving player ids")

	var accounts []string
	for _, a := range data {
		accounts = append(accounts, strconv.Itoa(a))
	}

	encoded, err := json.Marshal(accounts)
	if err != nil {
		log.Err(err).Msg("failed to marshal accounts payload")
		return
	}

	res, err := client.Post(fmt.Sprintf("http://%s/v1/accounts/import", apiDomain), "application/json", bytes.NewBuffer(encoded))
	if err != nil {
		log.Err(err).Msg("failed to post accounts import")
		return
	}
	defer res.Body.Close()
}
