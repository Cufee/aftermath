package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"

	"os"
	"time"

	"github.com/cufee/aftermath/cmds/core"
	"github.com/cufee/aftermath/cmds/core/queue"
	"github.com/cufee/aftermath/cmds/core/scheduler"

	"github.com/cufee/aftermath/cmds/core/server"
	"github.com/cufee/aftermath/cmds/core/server/handlers/private"
	"github.com/cufee/aftermath/cmds/discord"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/external/blitzstars"
	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/localization"
	"github.com/cufee/aftermath/internal/logic"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"

	"github.com/cufee/aftermath/internal/stats/render/assets"
	render "github.com/cufee/aftermath/internal/stats/render/common/v1"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	_ "github.com/joho/godotenv/autoload"
)

//go:generate go generate ./internal/database/ent

//go:embed static/*
var static embed.FS

func main() {
	globalCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Logger
	level, _ := zerolog.ParseLevel(os.Getenv("LOG_LEVEL"))
	zerolog.SetGlobalLevel(level)

	loadStaticAssets(static)

	liveCoreClient, cacheCoreClient := coreClientsFromEnv()
	stopQueue, err := startQueueFromEnv(globalCtx, cacheCoreClient.Wargaming())
	if err != nil {
		log.Fatal().Err(err).Msg("startQueueFromEnv failed")
	}
	defer stopQueue()

	stopScheduler, err := startSchedulerFromEnv(globalCtx, cacheCoreClient)
	if err != nil {
		log.Fatal().Err(err).Msg("startSchedulerFromEnv failed")
	}
	defer stopScheduler()

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
				Path: "POST /v1/tasks/restart",
				Func: private.RestartStaleTasks(cacheCoreClient),
			},
			{
				Path: "POST /v1/accounts/import",
				Func: private.LoadAccountsHandler(cacheCoreClient),
			},
			{
				Path: "POST /v1/snapshots/{realm}",
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
	go servePublic()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	sig := <-c
	cancel()
	log.Info().Msgf("received %s, exiting", sig.String())
}

func startSchedulerFromEnv(ctx context.Context, coreClient core.Client) (func(), error) {
	if os.Getenv("SCHEDULER_ENABLED") != "true" {
		return func() {}, nil
	}
	s := scheduler.New()
	scheduler.RegisterDefaultTasks(s, coreClient)
	return s.Start(ctx)
}

func startQueueFromEnv(ctx context.Context, wgClient wargaming.Client) (func(), error) {
	bsClient, err := blitzstars.NewClient(os.Getenv("BLITZ_STARS_API_URL"), time.Second*10)
	if err != nil {
		log.Fatal().Msgf("failed to init a blitzstars client %s", err)
	}

	// Fetch client
	fetchDBClient, err := newDatabaseClientFromEnv()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create a database client")
	}
	client, err := fetch.NewMultiSourceClient(wgClient, bsClient, fetchDBClient)
	if err != nil {
		log.Fatal().Msgf("fetch#NewMultiSourceClient failed %s", err)
	}

	// Queue - pulls tasks from database and runs the logic
	queueWorkerLimit := 10
	if v, err := strconv.Atoi(os.Getenv("SCHEDULER_CONCURRENT_WORKERS")); err == nil {
		queueWorkerLimit = v
	}
	q := queue.New(queueWorkerLimit, func() (core.Client, error) {
		// make a new database client and re-use the core client deps from before
		dbClient, err := newDatabaseClientFromEnv()
		if err != nil {
			return nil, err
		}
		return core.NewClient(client, wgClient, dbClient), nil
	})
	// q.SetHandlers(tasks.DefaultHandlers())

	return q.Start(ctx)
}

func coreClientsFromEnv() (core.Client, core.Client) {
	bsClient, err := blitzstars.NewClient(os.Getenv("BLITZ_STARS_API_URL"), time.Second*10)
	if err != nil {
		log.Fatal().Msgf("failed to init a blitzstars client %s", err)
	}

	liveClient, cacheClient := wargamingClientsFromEnv()

	// Fetch client
	liveDBClient, err := newDatabaseClientFromEnv()
	if err != nil {
		log.Fatal().Err(err).Msg("newDatabaseClientFromEnv failed")
	}
	liveFetchClient, err := fetch.NewMultiSourceClient(liveClient, bsClient, liveDBClient)
	if err != nil {
		log.Fatal().Msgf("fetch#NewMultiSourceClient failed %s", err)
	}

	cacheDBClient, err := newDatabaseClientFromEnv()
	if err != nil {
		log.Fatal().Err(err).Msg("newDatabaseClientFromEnv failed")
	}
	cacheFetchClient, err := fetch.NewMultiSourceClient(cacheClient, bsClient, cacheDBClient)
	if err != nil {
		log.Fatal().Msgf("fetch#NewMultiSourceClient failed %s", err)
	}

	return core.NewClient(liveFetchClient, liveClient, liveDBClient), core.NewClient(cacheFetchClient, cacheClient, cacheDBClient)
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

func newDatabaseClientFromEnv() (database.Client, error) {
	err := os.MkdirAll(os.Getenv("DATABASE_PATH"), os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("os#MkdirAll failed %w", err)
	}

	client, err := database.NewSQLiteClient(filepath.Join(os.Getenv("DATABASE_PATH"), os.Getenv("DATABASE_NAME")))
	if err != nil {

		return nil, fmt.Errorf("database#NewClient failed %w", err)
	}

	return client, nil
}
