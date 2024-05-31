package main

import (
	"fmt"
	"os"
	"time"

	"github.com/cufee/aftermath/internal/external/blitzstars"
	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/stats/fetch"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog"
)

//go:generate go run github.com/steebchen/prisma-client-go generate

func main() {
	level, _ := zerolog.ParseLevel(os.Getenv("LOG_LEVEL"))
	zerolog.SetGlobalLevel(level)

	wgClient, err := wargaming.NewClientFromEnv(os.Getenv("WG_PRIMARY_APP_ID"), os.Getenv("WG_PRIMARY_APP_RPS"), os.Getenv("WG_REQUEST_TIMEOUT_SEC"), os.Getenv("WG_PROXY_HOST_LIST"))
	if err != nil {
		panic(err)
	}

	bsClient := blitzstars.NewClient(os.Getenv("BLITZ_STARS_API_URL"), time.Second*3)

	client, err := fetch.NewMultiSourceClient(wgClient, bsClient)
	if err != nil {
		panic(err)
	}

	var days time.Duration = 30
	stats, err := client.PeriodStats("579553160", time.Now().Add(time.Hour*24*days*-1))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v - %v\n", stats.Account.Nickname, stats.RegularBattles.Battles)
}
