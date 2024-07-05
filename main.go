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

	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/cmd/core/queue"
	"github.com/cufee/aftermath/cmd/core/scheduler"
	"github.com/cufee/aftermath/cmd/core/tasks"
	"github.com/cufee/aftermath/cmd/discord/commands"
	"github.com/cufee/aftermath/cmd/frontend"

	"github.com/cufee/aftermath/cmd/core/server"
	"github.com/cufee/aftermath/cmd/core/server/handlers/private"
	"github.com/cufee/aftermath/cmd/discord"
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

	"net/http"
	"net/http/pprof"
)

//go:generate go generate ./cmd/frontend/assets

//go:embed static/*
var static embed.FS

func main() {
	globalCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Logger
	level, _ := zerolog.ParseLevel(os.Getenv("LOG_LEVEL"))
	zerolog.SetGlobalLevel(level)

	loadStaticAssets(static)
	db, err := newDatabaseClientFromEnv()
	if err != nil {
		log.Fatal().Err(err).Msg("newDatabaseClientFromEnv failed")
	}

	liveCoreClient, cacheCoreClient := coreClientsFromEnv(db)
	stopQueue, err := startQueueFromEnv(globalCtx, db, cacheCoreClient.Wargaming())
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

	if e := os.Getenv("PRIVATE_SERVER_ENABLED"); e == "true" {
		port := os.Getenv("PRIVATE_SERVER_PORT")
		servePrivate := server.NewServer(port, []server.Handler{
			{
				Path: "GET /debug/profile/{name}",
				Func: func(w http.ResponseWriter, r *http.Request) {
					pprof.Handler(r.PathValue("name")).ServeHTTP(w, r)
				},
			},
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

	// will handle all GET routes with a wildcard
	frontendHandlers, err := frontend.Handlers(liveCoreClient)
	if err != nil {
		log.Fatal().Err(err).Msg("frontend#Handlers failed")
	}

	discordHandlers := discordHandlersFromEnv(liveCoreClient)
	// POST /discord/public/callback
	// POST /discord/internal/callback

	var handlers []server.Handler
	handlers = append(handlers, discordHandlers...)
	handlers = append(handlers, frontendHandlers...)
	handlers = append(handlers, redirectHandlersFromEnv()...)

	port := os.Getenv("PORT")
	servePublic := server.NewServer(port, handlers...)
	log.Info().Str("port", port).Msg("starting a public server")
	go servePublic()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	sig := <-c
	cancel()
	log.Info().Msgf("received %s, exiting", sig.String())
}

func discordHandlersFromEnv(coreClient core.Client) []server.Handler {
	var handlers []server.Handler

	// main Discord with all the user-facing command
	{
		if os.Getenv("DISCORD_TOKEN") == "" || os.Getenv("DISCORD_PUBLIC_KEY") == "" {
			log.Fatal().Msg("DISCORD_TOKEN and DISCORD_PUBLIC_KEY are required")

		}
		mainDiscordHandler, err := discord.NewRouterHandler(coreClient, os.Getenv("DISCORD_TOKEN"), os.Getenv("DISCORD_PUBLIC_KEY"), commands.LoadedPublic.Compose()...)
		if err != nil {
			log.Fatal().Msgf("discord#NewRouterHandler failed %s", err)
		}
		handlers = append(handlers, server.Handler{
			Path: "POST /discord/public/callback",
			Func: mainDiscordHandler,
		})
	}

	// secondary Discord bot with mod/admin commands - permissions are still checked
	if token := os.Getenv("INTERNAL_DISCORD_TOKEN"); token != "" {
		publicKey := os.Getenv("INTERNAL_DISCORD_PUBLIC_KEY")
		if publicKey == "" {
			log.Fatal().Msg("discordHandlersFromEnv failed missing INTERNAL_DISCORD_PUBLIC_KEY")
		}

		internalDiscordHandler, err := discord.NewRouterHandler(coreClient, token, publicKey, commands.LoadedInternal.Compose()...)
		if err != nil {
			log.Fatal().Msgf("discord#NewRouterHandler failed %s", err)
		}
		handlers = append(handlers, server.Handler{
			Path: "POST /discord/internal/callback",
			Func: internalDiscordHandler,
		})
	}

	return handlers
}

func startSchedulerFromEnv(ctx context.Context, coreClient core.Client) (func(), error) {
	if os.Getenv("SCHEDULER_ENABLED") != "true" {
		return func() {}, nil
	}
	s := scheduler.New()
	scheduler.RegisterDefaultTasks(s, coreClient)
	return s.Start(ctx)
}

func startQueueFromEnv(ctx context.Context, db database.Client, wgClient wargaming.Client) (func(), error) {
	bsClient, err := blitzstars.NewClient(os.Getenv("BLITZ_STARS_API_URL"), time.Second*10)
	if err != nil {
		log.Fatal().Msgf("failed to init a blitzstars client %s", err)
	}

	// Fetch client
	client, err := fetch.NewMultiSourceClient(wgClient, bsClient, db)
	if err != nil {
		log.Fatal().Msgf("fetch#NewMultiSourceClient failed %s", err)
	}

	// Queue - pulls tasks from database and runs the logic
	queueWorkerLimit := 10
	if v, err := strconv.Atoi(os.Getenv("SCHEDULER_CONCURRENT_WORKERS")); err == nil {
		queueWorkerLimit = v
	}

	coreClient := core.NewClient(client, wgClient, db)
	q := queue.New(queueWorkerLimit, func() (core.Client, error) {
		return coreClient, nil
	})

	q.SetHandlers(tasks.DefaultHandlers())
	return q.Start(ctx)
}

func coreClientsFromEnv(db database.Client) (core.Client, core.Client) {
	bsClient, err := blitzstars.NewClient(os.Getenv("BLITZ_STARS_API_URL"), time.Second*10)
	if err != nil {
		log.Fatal().Msgf("failed to init a blitzstars client %s", err)
	}

	liveClient, cacheClient := wargamingClientsFromEnv()

	// Fetch client
	liveFetchClient, err := fetch.NewMultiSourceClient(liveClient, bsClient, db)
	if err != nil {
		log.Fatal().Msgf("fetch#NewMultiSourceClient failed %s", err)
	}

	cacheFetchClient, err := fetch.NewMultiSourceClient(cacheClient, bsClient, db)
	if err != nil {
		log.Fatal().Msgf("fetch#NewMultiSourceClient failed %s", err)
	}

	return core.NewClient(liveFetchClient, liveClient, db), core.NewClient(cacheFetchClient, cacheClient, db)
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

func redirectHandlersFromEnv() []server.Handler {
	return []server.Handler{
		{
			Path: "GET /invite/{$}",
			Func: func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, os.Getenv("BOT_INVITE_LINK"), http.StatusTemporaryRedirect)
			},
		},
		{
			Path: "GET /join/{$}",
			Func: func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, os.Getenv("PRIMARY_GUILD_INVITE_LINK"), http.StatusTemporaryRedirect)
			},
		},
	}
}
