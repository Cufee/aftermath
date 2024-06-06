package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cufee/aftermath/cmds/core"
	"github.com/cufee/aftermath/cmds/core/scheduler"
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

//go:generate go run github.com/steebchen/prisma-client-go generate --schema ./internal/database/prisma/schema.prisma

//go:embed static/*
var static embed.FS

func main() {
	// Logger
	level, _ := zerolog.ParseLevel(os.Getenv("LOG_LEVEL"))
	zerolog.SetGlobalLevel(level)

	loadStaticAssets(static)

	coreClient := coreClientFromEnv()

	scheduler.UpdateAveragesWorker(coreClient)()

	discordHandler, err := discord.NewRouterHandler(coreClient, os.Getenv("DISCORD_TOKEN"), os.Getenv("DISCORD_PUBLIC_KEY"))
	if err != nil {
		panic(err)
	}

	http.Handle("/discord/callback", discordHandler)
	panic(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func coreClientFromEnv() core.Client {
	// Dependencies
	wgClient, err := wargaming.NewClientFromEnv(os.Getenv("WG_PRIMARY_APP_ID"), os.Getenv("WG_PRIMARY_APP_RPS"), os.Getenv("WG_REQUEST_TIMEOUT_SEC"), os.Getenv("WG_PROXY_HOST_LIST"))
	if err != nil {
		log.Fatalf("wargaming#NewClientFromEnv failed %s", err)
	}
	dbClient, err := database.NewClient()
	if err != nil {
		log.Fatalf("database#NewClient failed %s", err)
	}
	bsClient, err := blitzstars.NewClient(os.Getenv("BLITZ_STARS_API_URL"), time.Second*10)
	if err != nil {
		log.Fatalf("failed to init a blitzstars client %s", err)
	}

	// Fetch client
	client, err := fetch.NewMultiSourceClient(wgClient, bsClient, dbClient)
	if err != nil {
		log.Fatalf("fetch#NewMultiSourceClient failed %s", err)
	}

	return core.NewClient(client, dbClient)
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
