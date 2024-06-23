package main

import (
	"embed"
	"io/fs"
	"path/filepath"

	"os"
	"strconv"
	"time"

	"github.com/cufee/aftermath/cmds/core"
	"github.com/cufee/aftermath/cmds/core/scheduler"
	"github.com/cufee/aftermath/cmds/core/scheduler/tasks"
	"github.com/cufee/aftermath/cmds/core/server"
	"github.com/cufee/aftermath/cmds/core/server/handlers/private"
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

//go:generate go generate ./internal/database/ent

//go:embed static/*
var static embed.FS

func main() {
	// Logger
	level, _ := zerolog.ParseLevel(os.Getenv("LOG_LEVEL"))
	zerolog.SetGlobalLevel(level)

	loadStaticAssets(static)

	liveCoreClient, cacheCoreClient := coreClientsFromEnv()
	startSchedulerFromEnvAsync(cacheCoreClient.Database(), cacheCoreClient.Wargaming())

	// Load some init options to registered admin accounts and etc
	logic.ApplyInitOptions(liveCoreClient.Database())

	discordHandler, err := discord.NewRouterHandler(liveCoreClient, os.Getenv("DISCORD_TOKEN"), os.Getenv("DISCORD_PUBLIC_KEY"))
	if err != nil {
		log.Fatal().Msgf("discord#NewRouterHandler failed %s", err)
	}

	if e := os.Getenv("PRIVATE_SERVER_ENABLED"); e == "true" {
		port := os.Getenv("PRIVATE_SERVER_PORT")
		servePrivate := server.NewServer(port, []server.Handler{
			{
				Path: "POST /tasks/restart",
				Func: private.RestartStaleTasks(cacheCoreClient),
			},
			{
				Path: "POST /accounts/import",
				Func: private.LoadAccountsHandler(cacheCoreClient),
			},
			{
				Path: "POST /snapshots/{realm}",
				Func: private.SaveRealmSnapshots(cacheCoreClient),
			},
		}...)
		log.Info().Str("port", port).Msg("starting a private server")
		go servePrivate()
	}

	port := os.Getenv("PORT")
	servePublic := server.NewServer(port, server.Handler{
		Path: "POST /discord/callback",
		Func: discordHandler,
	})
	log.Info().Str("port", port).Msg("starting a public server")
	servePublic()
}

func startSchedulerFromEnvAsync(dbClient database.Client, wgClient wargaming.Client) {
	if os.Getenv("SCHEDULER_ENABLED") != "true" {
		return
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
}

func coreClientsFromEnv() (core.Client, core.Client) {
	err := os.MkdirAll(os.Getenv("DATABASE_PATH"), os.ModePerm)
	if err != nil {
		log.Fatal().Msgf("os#MkdirAll failed %s", err)
	}

	// Dependencies
	dbClient, err := database.NewSQLiteClient(filepath.Join(os.Getenv("DATABASE_PATH"), os.Getenv("DATABASE_NAME")))
	if err != nil {
		log.Fatal().Msgf("database#NewClient failed %s", err)
	}
	bsClient, err := blitzstars.NewClient(os.Getenv("BLITZ_STARS_API_URL"), time.Second*10)
	if err != nil {
		log.Fatal().Msgf("failed to init a blitzstars client %s", err)
	}

	liveClient, cacheClient := wargamingClientsFromEnv()

	// Fetch client
	fetchClient, err := fetch.NewMultiSourceClient(cacheClient, bsClient, dbClient)
	if err != nil {
		log.Fatal().Msgf("fetch#NewMultiSourceClient failed %s", err)
	}

	return core.NewClient(fetchClient, liveClient, dbClient), core.NewClient(fetchClient, cacheClient, dbClient)
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

func wargamingClientsFromEnv() (wargaming.Client, wargaming.Client) {
	liveClient, err := wargaming.NewClientFromEnv(os.Getenv("WG_PRIMARY_APP_ID"), os.Getenv("WG_PRIMARY_APP_RPS"), os.Getenv("WG_PRIMARY_REQUEST_TIMEOUT_SEC"), os.Getenv("WG_PROXY_HOST_LIST"))
	if err != nil {
		log.Fatal().Msgf("wargamingClientsFromEnv#NewClientFromEnv failed %s", err)
	}

	// This wargaming client is using a different proxy as it needs a lot higher rps, but can be slow
	cacheClient, err := wargaming.NewClientFromEnv(os.Getenv("WG_CACHE_APP_ID"), os.Getenv("WG_CACHE_APP_RPS"), os.Getenv("WG_CACHE_REQUEST_TIMEOUT_SEC"), os.Getenv("WG_CACHE_PROXY_HOST_LIST"))
	if err != nil {
		log.Fatal().Msgf("wargamingClientsFromEnv#NewClientFromEnv failed %s", err)
	}

	return liveClient, cacheClient
}
