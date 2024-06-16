package main

import (
	"embed"
	"io/fs"

	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/cufee/aftermath/cmds/core"
	"github.com/cufee/aftermath/cmds/core/scheduler"
	"github.com/cufee/aftermath/cmds/core/scheduler/tasks"
	"github.com/cufee/aftermath/cmds/discord"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/external/blitzstars"
	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/localization"
	"github.com/cufee/aftermath/internal/logic"
	"github.com/cufee/aftermath/internal/stats/fetch"

	"github.com/cufee/aftermath/internal/stats/render/assets"
	render "github.com/cufee/aftermath/internal/stats/render/common"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

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
	startSchedulerFromEnvAsync()

	// Load some init options to registed admin accounts and etc
	logic.ApplyInitOptions(coreClient.Database())

	discordHandler, err := discord.NewRouterHandler(coreClient, os.Getenv("DISCORD_TOKEN"), os.Getenv("DISCORD_PUBLIC_KEY"))
	if err != nil {
		panic(err)
	}

	http.Handle("/discord/callback", discordHandler)
	panic(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func startSchedulerFromEnvAsync() {
	if os.Getenv("SCHEDULER_ENABLED") != "true" {
		return
	}

	// This wargaming client is using a different proxy as it needs a lot higher rps, but can be slow
	wgClient, err := wargaming.NewClientFromEnv(os.Getenv("WG_CACHE_APP_ID"), os.Getenv("WG_CACHE_APP_RPS"), os.Getenv("WG_CACHE_REQUEST_TIMEOUT_SEC"), os.Getenv("WG_CACHE_PROXY_HOST_LIST"))
	if err != nil {
		log.Fatal().Msgf("wgClient: wargaming#NewClientFromEnv failed %s", err)
	}
	dbClient, err := database.NewClient()
	if err != nil {
		log.Fatal().Msgf("database#NewClient failed %s", err)
	}
	bsClient, err := blitzstars.NewClient(os.Getenv("BLITZ_STARS_API_URL"), time.Second*10)
	if err != nil {
		log.Fatal().Msgf("failed to init a blitzstars client %s", err)
	}

	// Fetch client
	client, err := fetch.NewMultiSourceClient(wgClient, bsClient, dbClient)
	if err != nil {
		log.Fatal().Msgf("fetch#NewMultiSourceClient failed %s", err)
	}
	coreClient := core.NewClient(client, wgClient, dbClient)

	// Queue
	taskQueueConcurrency := 10
	if v, err := strconv.Atoi(os.Getenv("SCHEDULER_CONCURRENT_WORKERS")); err == nil {
		taskQueueConcurrency = v
	}
	queue := scheduler.NewQueue(coreClient, tasks.DefaultHandlers(), taskQueueConcurrency)

	defer func() {
		if r := recover(); r != nil {
			log.Error().Interface("error", r).Stack().Msg("scheduler panic")
		}
	}()

	queue.StartCronJobsAsync()
	// Some tasks should run on startup
	// scheduler.UpdateGlossaryWorker(coreClient)()
	// scheduler.UpdateAveragesWorker(coreClient)()
	// scheduler.CreateSessionTasksWorker(coreClient, "AS")()
}

func coreClientFromEnv() core.Client {
	// Dependencies
	wgClient, err := wargaming.NewClientFromEnv(os.Getenv("WG_PRIMARY_APP_ID"), os.Getenv("WG_PRIMARY_APP_RPS"), os.Getenv("WG_PRIMARY_REQUEST_TIMEOUT_SEC"), os.Getenv("WG_PROXY_HOST_LIST"))
	if err != nil {
		log.Fatal().Msgf("wgClient: wargaming#NewClientFromEnv failed %s", err)
	}
	dbClient, err := database.NewClient()
	if err != nil {
		log.Fatal().Msgf("database#NewClient failed %s", err)
	}
	bsClient, err := blitzstars.NewClient(os.Getenv("BLITZ_STARS_API_URL"), time.Second*10)
	if err != nil {
		log.Fatal().Msgf("failed to init a blitzstars client %s", err)
	}

	// Fetch client
	client, err := fetch.NewMultiSourceClient(wgClient, bsClient, dbClient)
	if err != nil {
		log.Fatal().Msgf("fetch#NewMultiSourceClient failed %s", err)
	}

	return core.NewClient(client, wgClient, dbClient)
}

func loadStaticAssets(static fs.FS) {
	// Assets for rendering
	err := assets.LoadAssets(static, "static")
	if err != nil {
		log.Fatal().Msgf("assets#LoadAssets failed to load assets from static/ embed.FS %s", err)
	}
	err = render.InitLoadedAssets()
	if err != nil {
		log.Fatal().Msgf("render#InitLoadedAssets failed %s", err)
	}
	err = localization.LoadAssets(static, "static/localization")
	if err != nil {
		log.Fatal().Msgf("localization#LoadAssets failed %s", err)
	}
}
