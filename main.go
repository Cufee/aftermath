package main

import (
	"context"
	"embed"
	"fmt"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/external/blitzstars"
	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/stats"
	"github.com/cufee/aftermath/internal/stats/fetch"

	"github.com/cufee/aftermath/internal/stats/render/assets"
	render "github.com/cufee/aftermath/internal/stats/render/common"
	"github.com/rs/zerolog"

	_ "github.com/joho/godotenv/autoload"
)

//go:generate go run github.com/steebchen/prisma-client-go generate

//go:embed static/*
var static embed.FS

func main() {
	// Logger
	level, _ := zerolog.ParseLevel(os.Getenv("LOG_LEVEL"))
	zerolog.SetGlobalLevel(level)

	// Assets for rendering
	err := assets.LoadAssets(static)
	if err != nil {
		log.Fatalf("assets#LoadAssets failed to load assets from static/ embed.FS %s", err)
	}
	err = render.InitLoadedAssets()
	if err != nil {
		log.Fatalf("render#InitLoadedAssets failed %s", err)
	}

	statsClient := fetchClientFromEnv()

	// test
	// localePrinter := func(s string) string { return "localized:" + s }

	renderer := stats.NewRenderer(statsClient)

	var days time.Duration = 60
	img, meta, err := renderer.Period(context.Background(), "579553160", time.Now().Add(time.Hour*24*days*-1))
	if err != nil {
		panic(err)
	}

	fmt.Printf("rendered in %v\n", meta.TotalTime())

	f, err := os.Create("tmp/test-image.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}

}

func fetchClientFromEnv() fetch.Client {
	// Dependencies
	wgClient, err := wargaming.NewClientFromEnv(os.Getenv("WG_PRIMARY_APP_ID"), os.Getenv("WG_PRIMARY_APP_RPS"), os.Getenv("WG_REQUEST_TIMEOUT_SEC"), os.Getenv("WG_PROXY_HOST_LIST"))
	if err != nil {
		log.Fatalf("wargaming#NewClientFromEnv failed %s", err)
	}
	dbClient, err := database.NewClient()
	if err != nil {
		log.Fatalf("database#NewClient failed %s", err)
	}
	bsClient := blitzstars.NewClient(os.Getenv("BLITZ_STARS_API_URL"), time.Second*3)

	// Fetch client
	client, err := fetch.NewMultiSourceClient(wgClient, bsClient, dbClient)
	if err != nil {
		log.Fatalf("fetch#NewMultiSourceClient failed %s", err)
	}

	return client
}
