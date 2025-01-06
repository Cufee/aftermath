//go:build ignore

package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-co-op/gocron"
	_ "github.com/joho/godotenv/autoload"

	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/json"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/rs/zerolog"
)

func main() {
	// Logger
	level, _ := zerolog.ParseLevel(os.Getenv("LOG_LEVEL"))
	log.SetupGlobalLogger(func(l zerolog.Logger) zerolog.Logger {
		return l.Level(level)
	})

	backendApi := flag.String("backend", "", "Aftermath api domain")
	flag.Parse()
	if *backendApi == "" {
		log.Fatal().Msg("backend is a required argument")
	}

	scheduler := gocron.NewScheduler(time.UTC)

	scheduler.Cron("0 3 * * 2").Do(collectRealmIDs, *backendApi, types.RealmAsia)
	scheduler.Cron("0 7 * * 2").Do(collectRealmIDs, *backendApi, types.RealmEurope)
	scheduler.Cron("0 11 * * 2").Do(collectRealmIDs, *backendApi, types.RealmNorthAmerica)

	log.Info().Msg("started a cron scheduler")
	scheduler.StartBlocking()
}

func collectRealmIDs(backendApi string, realm types.Realm) {
	client, err := wargaming.NewRatingLeaderboardClient()
	if err != nil {
		log.Err(err).Msg("failed to create a leaderboard client")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	season, err := client.CurrentSeason(ctx, realm)
	if err != nil {
		log.Err(err).Msg("failed to get current season")
		return
	}

	if len(season.Leagues) < 1 {
		log.Err(err).Msg("season contains no leagues")
		return
	}

	players, err := client.LeagueTop(ctx, realm, season.Leagues[0].ID)
	if err != nil {
		log.Err(err).Msg("failed to get league top players")
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
			log.Debug().Int("count", len(ids)).Int("total", total).Str("realm", realm.String()).Msg("collected player ids")
		}
	}()

	err = client.CollectPlayerIDs(context.Background(), types.RealmNorthAmerica, collector, players[len(players)-1].AccountID)
	if err != nil {
		log.Err(err).Msg("failed to complete player id collection")
		return
	}

	log.Info().Int("total", total).Msg("finished collecting player ids")
}

var client = http.Client{
	Timeout: time.Second * 5,
}

func savePlayerIDs(apiDomain string, data []int) {
	log.Info().Int("count", len(data)).Msg("saving player ids")

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
