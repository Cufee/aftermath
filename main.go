package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cufee/aftermath/internal/external/blitzstars"
	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/stats/fetch"
	"github.com/cufee/aftermath/internal/stats/prepare"
	"github.com/cufee/aftermath/internal/stats/prepare/period"
	"github.com/cufee/aftermath/internal/stats/render/assets"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog"
)

//go:generate go run github.com/steebchen/prisma-client-go generate

//go:embed static/*
var static embed.FS

func fetchClientFromEnv() fetch.Client {
	// Dependencies
	wgClient, err := wargaming.NewClientFromEnv(os.Getenv("WG_PRIMARY_APP_ID"), os.Getenv("WG_PRIMARY_APP_RPS"), os.Getenv("WG_REQUEST_TIMEOUT_SEC"), os.Getenv("WG_PROXY_HOST_LIST"))
	if err != nil {
		log.Fatalf("wargaming#NewClientFromEnv failed %s", err)
	}
	bsClient := blitzstars.NewClient(os.Getenv("BLITZ_STARS_API_URL"), time.Second*3)

	// Fetch client
	client, err := fetch.NewMultiSourceClient(wgClient, bsClient)
	if err != nil {
		log.Fatalf("fetch#NewMultiSourceClient failed %s", err)
	}

	return client
}

func main() {
	// Logger
	level, _ := zerolog.ParseLevel(os.Getenv("LOG_LEVEL"))
	zerolog.SetGlobalLevel(level)

	// Assets for rendering
	err := assets.LoadAssets(static)
	if err != nil {
		log.Fatalf("assets#LoadAssets failed to load assets from static/ embed.FS %s", err)
	}

	statsClient := fetchClientFromEnv()

	// test
	var days time.Duration = 30
	stats, err := statsClient.PeriodStats("579553160", time.Now().Add(time.Hour*24*days*-1))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v - %v\n", stats.Account.Nickname, stats.RegularBattles.Battles)

	localePrinter := func(s string) string { return "localized:" + s }

	cards, err := prepare.Period(stats, period.WithPrinter(localePrinter))
	if err != nil {
		panic(err)
	}

	data, _ := json.MarshalIndent(cards, "", "  ")
	println(string(data))
}
