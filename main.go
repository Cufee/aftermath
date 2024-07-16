package main

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"os/signal"
	"path/filepath"
	"syscall"

	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/cmd/core/queue"
	"github.com/cufee/aftermath/cmd/core/scheduler"
	"github.com/cufee/aftermath/cmd/core/tasks"
	"github.com/cufee/aftermath/cmd/discord/alerts"
	"github.com/cufee/aftermath/cmd/discord/commands"
	"github.com/cufee/aftermath/cmd/discord/gateway"
	"github.com/cufee/aftermath/cmd/frontend"
	"github.com/disintegration/imaging"
	"github.com/pkg/errors"

	"github.com/cufee/aftermath/cmd/core/server"
	"github.com/cufee/aftermath/cmd/core/server/handlers/private"
	"github.com/cufee/aftermath/cmd/discord"
	"github.com/cufee/aftermath/internal/constants"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/external/blitzstars"
	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/localization"
	"github.com/cufee/aftermath/internal/logic"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"

	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/stats/render/assets"
	render "github.com/cufee/aftermath/internal/stats/render/common/v1"
	"github.com/rs/zerolog"

	_ "github.com/joho/godotenv/autoload"

	"net/http"
	"net/http/pprof"
)

//go:generate go generate ./internal/assets
//go:generate go generate ./cmd/frontend/assets/generate

//go:embed static/*
var static embed.FS

func main() {
	loadStaticAssets(static)
	db, err := newDatabaseClientFromEnv()
	if err != nil {
		log.Fatal().Err(err).Msg("newDatabaseClientFromEnv failed")
	}

	// Logger
	level, _ := zerolog.ParseLevel(os.Getenv("LOG_LEVEL"))
	log.SetupGlobalLogger(func(l zerolog.Logger) zerolog.Logger {
		return l.Level(level)
	})

	if constants.DiscordAlertsEnabled {
		alertClient, err := alerts.NewClient(constants.DiscordPrimaryToken, constants.DiscordAlertsWebhookURL)
		if err != nil {
			log.Fatal().Err(err).Msg("alerts#NewClient failed")
		}

		pr, pw := io.Pipe()
		log.SetupGlobalLogger(func(l zerolog.Logger) zerolog.Logger {
			return l.Output(zerolog.MultiLevelWriter(os.Stderr, pw))
		})
		alertClient.Reader(pr, zerolog.ErrorLevel)
	}

	// Discord Gateway
	gw, err := discordGatewayFromEnv()
	if err != nil {
		log.Fatal().Err(err).Msg("discordGatewayFromEnv failed")
	}

	globalCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

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

	// Load some init options to registered admin accounts and etc
	logic.ApplyInitOptions(liveCoreClient.Database())

	if constants.ServePrivateEndpointsEnabled {
		servePrivate := server.NewServer(constants.ServePrivateEndpointsPort, []server.Handler{
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
				Path: "POST /v1/tasks/averages/refresh",
				Func: private.RefreshAverages(cacheCoreClient),
			},
			{
				Path: "POST /v1/tasks/glossary/refresh",
				Func: private.RefreshGlossary(cacheCoreClient),
			},
			{
				Path: "POST /v1/accounts/import",
				Func: private.LoadAccountsHandler(cacheCoreClient),
			},
			{
				Path: "POST /v1/snapshots/{realm}",
				Func: private.SaveRealmSnapshots(cacheCoreClient),
			},
		}, log.NewMiddleware(log.Logger()))
		log.Info().Str("port", constants.ServePrivateEndpointsPort).Msg("starting a private server")
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
	servePublic := server.NewServer(port, handlers, log.NewMiddleware(log.Logger()))
	log.Info().Str("port", port).Msg("starting a public server")
	go servePublic()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	sig := <-c
	log.Info().Msgf("received %s, exiting after cleanup", sig.String())
	gw.SetStatus(gateway.StatusYellow, "🔄 Updating...", nil)
	cancel()
	stopScheduler()
	log.Info().Msg("finished cleanup tasks")
}

func discordGatewayFromEnv() (gateway.Client, error) {
	// discord gateway
	gw, err := gateway.NewClient(constants.DiscordPrimaryToken, discordgo.IntentDirectMessages|discordgo.IntentGuildMessages)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create a gateway client")
	}

	helpImage, ok := assets.GetLoadedImage("discord-help")
	if !ok {
		return nil, errors.New("discord-help image is not loaded")
	}

	var buf bytes.Buffer
	err = imaging.Encode(&buf, helpImage, imaging.JPEG)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encode help image")
	}
	_ = gw.Handler(commands.MentionHandler(buf.Bytes()))

	err = gw.Connect()
	if err != nil {
		return nil, errors.Wrap(err, "gateway client failed to connect")
	}
	gw.SetStatus(gateway.StatusListening, "/help", nil)

	return gw, nil
}

func discordHandlersFromEnv(coreClient core.Client) []server.Handler {
	var handlers []server.Handler

	// main Discord with all the user-facing command
	{
		mainDiscordHandler, err := discord.NewRouterHandler(coreClient, constants.DiscordPrimaryToken, constants.DiscordPrimaryPublicKey, commands.LoadedPublic.Compose()...)
		if err != nil {
			log.Fatal().Msgf("discord#NewRouterHandler failed %s", err)
		}
		handlers = append(handlers, server.Handler{
			Path: "POST /discord/public/callback",
			Func: mainDiscordHandler,
		})
	}

	// secondary Discord bot with mod/admin commands - permissions are still checked
	if constants.DiscordPrivateBotEnabled {
		internalDiscordHandler, err := discord.NewRouterHandler(coreClient, constants.DiscordPrivateToken, constants.DiscordPrivatePublicKey, commands.LoadedInternal.Compose()...)
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
	if !constants.SchedulerEnabled {
		return func() {}, nil
	}
	s := scheduler.New()
	scheduler.RegisterDefaultTasks(s, coreClient)
	return s.Start(ctx)
}

func startQueueFromEnv(ctx context.Context, db database.Client, wgClient wargaming.Client) (func(), error) {
	bsClient, err := blitzstars.NewClient(constants.BlitzStarsApiURL, time.Second*10)
	if err != nil {
		log.Fatal().Msgf("failed to init a blitzstars client %s", err)
	}

	// Fetch client
	client, err := fetch.NewMultiSourceClient(wgClient, bsClient, db)
	if err != nil {
		log.Fatal().Msgf("fetch#NewMultiSourceClient failed %s", err)
	}

	// Queue - pulls tasks from database and runs the logic
	coreClient := core.NewClient(client, wgClient, db)
	q := queue.New(constants.SchedulerConcurrentWorkers, func() (core.Client, error) {
		return coreClient, nil
	})

	q.SetHandlers(tasks.DefaultHandlers())
	return q.Start(ctx)
}

func coreClientsFromEnv(db database.Client) (core.Client, core.Client) {
	bsClient, err := blitzstars.NewClient(constants.BlitzStarsApiURL, time.Second*10)
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
	liveClient, err := wargaming.NewClientFromEnv(constants.WargamingPrimaryAppID, constants.WargamingPrimaryAppRPS, constants.WargamingPrimaryAppRequestTimeout, constants.WargamingPrimaryAppProxyHostList)
	if err != nil {
		log.Fatal().Msgf("wargamingClientsFromEnv#NewClientFromEnv failed %s", err)
	}

	// This wargaming client is using a different proxy as it needs a lot higher rps, but can be slow
	cacheClient, err := wargaming.NewClientFromEnv(constants.WargamingCacheAppID, constants.WargamingCacheAppRPS, constants.WargamingCacheAppRequestTimeout, constants.WargamingCacheAppProxyHostList)
	if err != nil {
		log.Fatal().Msgf("wargamingClientsFromEnv#NewClientFromEnv failed %s", err)
	}

	return liveClient, cacheClient
}

func newDatabaseClientFromEnv() (database.Client, error) {
	err := os.MkdirAll(constants.DatabasePath, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("os#MkdirAll failed %w", err)
	}

	client, err := database.NewSQLiteClient(filepath.Join(constants.DatabasePath, constants.DatabaseName))
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
				http.Redirect(w, r, constants.DiscordBotInviteURL, http.StatusTemporaryRedirect)
			},
		},
		{
			Path: "GET /join/{$}",
			Func: func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, constants.DiscordPrimaryGuildInviteURL, http.StatusTemporaryRedirect)
			},
		},
	}
}
