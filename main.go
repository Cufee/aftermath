package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cufee/aftermath/cmds/discord"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/external/blitzstars"
	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/localization"
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

	loadStaticAssets(static)
	// statsClient := fetchClientFromEnv()

	discordHandler, err := discord.NewRouterHandler(os.Getenv("DISCORD_TOKEN"), os.Getenv("DISCORD_PUBLIC_KEY"))
	if err != nil {
		panic(err)
	}

	http.Handle("/discord/callback", discordHandler)
	panic(http.ListenAndServe(":9092", nil))

	// test
	// localePrinter := func(s string) string { return "localized:" + s }

	// renderer := stats.NewRenderer(statsClient, language.English)

	// var days time.Duration = 60
	// img, meta, err := renderer.Period(context.Background(), "579553160", time.Now().Add(time.Hour*24*days*-1))
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("rendered in %v\n", meta.TotalTime())

	// bgImage, _ := assets.GetLoadedImage("bg-light")
	// finalImage := render.AddBackground(img, bgImage, render.Style{Blur: 10, BorderRadius: 30})

	// f, err := os.Create("tmp/test-image.png")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// err = png.Encode(f, finalImage)
	// if err != nil {
	// 	panic(err)
	// }

}

func loadStaticAssets(static fs.FS) {
	// Assets for rendering
	err := assets.LoadAssets(static, "static")
	if err != nil {
		log.Fatalf("assets#LoadAssets failed to load assets from static/ embed.FS %s", err)
	}
	err = render.InitLoadedAssets()
	if err != nil {
		log.Fatalf("render#InitLoadedAssets failed %s", err)
	}
	err = localization.LoadAssets(static, "static/localization")
	if err != nil {
		log.Fatalf("localization#LoadAssets failed %s", err)
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
