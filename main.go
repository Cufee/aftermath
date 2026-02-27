package main

import (
	"bytes"
	"context"
	"embed"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"os/signal"
	"strings"
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
	_ "github.com/cufee/aftermath/cmd/discord/commands/private"
	"github.com/cufee/aftermath/cmd/discord/commands/public"
	"github.com/cufee/aftermath/cmd/discord/cta"
	"github.com/cufee/aftermath/cmd/discord/gateway"
	"github.com/cufee/aftermath/cmd/discord/middleware"
	"github.com/cufee/aftermath/cmd/discord/rest"
	"github.com/cufee/aftermath/cmd/discord/router"
	"github.com/cufee/aftermath/cmd/frontend"
	"github.com/nao1215/imaging"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/cufee/aftermath/cmd/core/server"
	"github.com/cufee/aftermath/cmd/core/server/handlers/private"
	"github.com/cufee/aftermath/internal/constants"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/external/blitzkit"
	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/localization"
	"github.com/cufee/aftermath/internal/logic"
	"github.com/cufee/aftermath/internal/metrics"
	"github.com/cufee/aftermath/internal/realtime"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"

	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/render/assets"
	"github.com/cufee/aftermath/internal/render/common"
	render "github.com/cufee/aftermath/internal/render/v1"

	_ "github.com/joho/godotenv/autoload"

	"net/http"
	"net/http/pprof"
)

//go:generate go tool templ generate
//go:generate go generate ./internal/assets
//go:generate go generate ./internal/database
//go:generate go generate ./cmd/frontend/assets/generate
//go:generate task build-widget

//go:embed static/*
var static embed.FS

func main() {
	loadStaticAssets(static)

	loadCtx, cancelLoadCtx := context.WithTimeout(context.Background(), time.Second*5)
	db, err := newDatabaseClientFromEnv(loadCtx)
	if err != nil {
		log.Fatal().Err(err).Msg("newDatabaseClientFromEnv failed")
	}
	cancelLoadCtx()

	// Logger
	level, _ := zerolog.ParseLevel(os.Getenv("LOG_LEVEL"))
	log.SetupGlobalLogger(func(l zerolog.Logger) zerolog.Logger {
		return l.Level(level)
	})

	errorTracker := metrics.NewErrorTracker(metrics.DefaultErrorTrackerConfig())
	liveCoreClient, cacheCoreClient := coreClientsFromEnv(db, errorTracker)

	// Metrics instrument
	instrument, err := metrics.NewInstrument()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create a new metrics instrument")
	}

	var alertClient alerts.DiscordAlertClient
	if constants.DiscordAlertsEnabled {
		c, err := alerts.NewClient(constants.DiscordPrimaryToken, constants.DiscordAlertsWebhookURL)
		if err != nil {
			log.Fatal().Err(err).Msg("alerts#NewClient failed")
		}
		alertClient = c

		pr, pw := io.Pipe()
		log.SetupGlobalLogger(func(l zerolog.Logger) zerolog.Logger {
			return l.Output(zerolog.MultiLevelWriter(os.Stderr, pw))
		})
		c.Reader(pr, zerolog.ErrorLevel)
	}

	globalCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	restObserver := rest.WithObserver(errorTracker)

	// Discord Gateway - commands public commands are handled through the gateway
	gw, err := discordGatewayFromEnv(globalCtx, liveCoreClient, instrument, errorTracker, restObserver)
	if err != nil {
		log.Fatal().Err(err).Msg("discordGatewayFromEnv failed")
	}
	registerHealthTransitionSinks(gw, errorTracker, alertClient)
	go errorTracker.Start(globalCtx)

	stopQueue, err := startQueueFromEnv(globalCtx, cacheCoreClient)
	if err != nil {
		log.Fatal().Err(err).Msg("startQueueFromEnv failed")
	}
	defer stopQueue()

	stopScheduler, err := startSchedulerFromEnv(globalCtx, cacheCoreClient)
	if err != nil {
		log.Fatal().Err(err).Msg("startSchedulerFromEnv failed")
	}

	go scheduler.UpdateGlossaryWorker(liveCoreClient)()

	// Load some init options to registered admin accounts and etc
	logic.ApplyInitOptions(liveCoreClient.Database())

	if constants.ServePrivateEndpointsEnabled {
		servePrivate := server.NewServer(constants.ServePrivateEndpointsPort, []server.Handler{
			{
				Path: "GET /metrics",
				Func: instrument.Handler(),
			},
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
		}, log.NewMiddleware(log.Logger(), "/metrics"))
		go servePrivate()
	}

	// will handle all GET routes with a wildcard
	frontendHandlers, err := frontend.Handlers(liveCoreClient)
	if err != nil {
		log.Fatal().Err(err).Msg("frontend#Handlers failed")
	}

	var handlers []server.Handler
	handlers = append(handlers, frontendHandlers...)
	handlers = append(handlers, redirectHandlersFromEnv()...)

	handlers = append(handlers, discordPublicHandlersFromEnv(liveCoreClient, instrument, restObserver)...)   // POST /discord/public/callback
	handlers = append(handlers, discordInternalHandlersFromEnv(liveCoreClient, instrument, restObserver)...) // POST /discord/internal/callback

	port := os.Getenv("PORT")
	servePublic := server.NewServer(port, handlers, log.NewMiddleware(log.Logger(), "/discord/public/callback"))
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

func discordGatewayFromEnv(globalCtx context.Context, core core.Client, instrument metrics.Instrument, tracker *metrics.ErrorTracker, restOpts ...rest.ClientOption) (gateway.Client, error) {
	collector := metrics.NewDiscordGatewayCollector("public")
	instrument.Register(collector)

	// discord gateway
	gw, err := gateway.NewClient(core, constants.DiscordPrimaryToken, discordgo.IntentDirectMessages|discordgo.IntentGuildMessages|discordgo.IntentGuildMessageReactions|discordgo.IntentDirectMessageReactions, restOpts...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create a gateway client")
	}
	gw.LoadMiddleware(collector.Middleware())

	helpImage, ok := assets.GetLoadedImage("discord-help")
	if !ok {
		return nil, errors.New("discord-help image is not loaded")
	}

	var buf bytes.Buffer
	err = imaging.Encode(&buf, helpImage, imaging.JPEG)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encode help image")
	}
	gw.Handler(gateway.ToGatewayHandler(gw, public.AutomodHandler(buf.Bytes())))
	gw.Handler(gateway.ToGatewayHandler(gw, public.MentionHandler(buf.Bytes())))
	gw.Handler(gateway.ToGatewayHandler(gw, public.ReplayFileHandler))
	gw.Handler(gateway.ToGatewayHandler(gw, public.ReplayInteractionHandler))

	err = gw.Connect()
	if err != nil {
		return nil, errors.Wrap(err, "gateway client failed to connect")
	}

	setHealthyStatus := func() {
		if tracker.OverallState() != metrics.HealthStateHealthy {
			return
		}
		if err := gw.SetStatus(gateway.StatusListening, "/help", nil); err != nil {
			log.Err(err).Msg("failed to update gateway status")
		}
	}

	setHealthyStatus()
	go func(ctx context.Context) {
		// the status update gets stuck sometimes, probably due to some Discord cache
		time.Sleep(time.Second * 15)
		setHealthyStatus()

		time.Sleep(time.Second * 45)
		setHealthyStatus()
	}(globalCtx)

	go func(ctx context.Context) {
		// update the status every hour
		ticker := time.NewTicker(time.Hour * 1)
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				setHealthyStatus()
			}
		}
	}(globalCtx)

	return gw, nil
}

func discordPublicHandlersFromEnv(coreClient core.Client, instrument metrics.Instrument, restOpts ...rest.ClientOption) []server.Handler {
	collector := metrics.NewDiscordInteractionCollector("public")
	instrument.Register(collector)

	var handlers []server.Handler

	publicKey, err := hex.DecodeString(constants.DiscordPrimaryPublicKey)
	if err != nil {
		log.Fatal().Msg("invalid discord primary public key")
	}

	log.Debug().Msg("setting up a public commands router handlers")
	router, err := router.NewRouter(coreClient, constants.DiscordPrimaryToken, router.NewPublicKeyValidator(publicKey), router.WithRestClientOptions(restOpts...))
	if err != nil {
		log.Fatal().Msgf("router#NewRouterHandler failed %s", err)
	}

	router.LoadMiddleware(cta.Middleware(cta.DefaultCooldownServer, cta.DefaultCooldownDM)) // cta messages
	router.LoadMiddleware(middleware.MarkUsersVerified())                                   // automod verify users using the bot
	router.LoadMiddleware(collector.Middleware())                                           // metrics

	router.LoadCommands(public.Help().Build())
	router.LoadCommands(commands.LoadedPublic.Compose()...)

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		err = router.UpdateLoadedCommands(ctx)
		if err != nil {
			log.Fatal().Msgf("router#UpdateLoadedCommands failed %s", err)
		}
	}()

	fn, err := router.HTTPHandler()
	if err != nil {
		log.Fatal().Msgf("router#HTTPHandler failed %s", err)
	}
	handlers = append(handlers, server.Handler{
		Path: "POST /discord/public/callback",
		Func: fn,
	})

	return handlers
}

func discordInternalHandlersFromEnv(coreClient core.Client, _ metrics.Instrument, restOpts ...rest.ClientOption) []server.Handler {
	var handlers []server.Handler
	// secondary Discord bot with mod/admin commands - permissions are still checked
	if constants.DiscordPrivateBotEnabled {
		log.Debug().Msg("setting up an internal commands router")

		publicKey, err := hex.DecodeString(constants.DiscordPrivatePublicKey)
		if err != nil {
			log.Fatal().Msg("invalid discord private public key")
		}

		router, err := router.NewRouter(coreClient, constants.DiscordPrivateToken, router.NewPublicKeyValidator(publicKey), router.WithRestClientOptions(restOpts...))
		if err != nil {
			log.Fatal().Msgf("discord#NewHTTPRouter failed %s", err)
		}

		router.LoadCommands(commands.LoadedInternal.Compose()...)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		err = router.UpdateLoadedCommands(ctx)
		if err != nil {
			log.Fatal().Msgf("router#UpdateLoadedCommands failed %s", err)
		}

		fn, err := router.HTTPHandler()
		if err != nil {
			log.Fatal().Msgf("router#HTTPHandler failed %s", err)
		}

		handlers = append(handlers, server.Handler{
			Path: "POST /discord/internal/callback",
			Func: fn,
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

func startQueueFromEnv(ctx context.Context, coreClient core.Client) (func(), error) {
	q := queue.New(constants.SchedulerConcurrentWorkers, func() (core.Client, error) {
		return coreClient, nil
	})

	q.SetHandlers(tasks.DefaultHandlers())
	return q.Start(ctx)
}

func coreClientsFromEnv(db database.Client, observer metrics.ErrorObserver) (core.Client, core.Client) {
	bkClient, err := blitzkit.NewClient(time.Second * 10)
	if err != nil {
		log.Fatal().Msgf("failed to init a blitzkit client %s", err)
	}

	liveClient, cacheClient := wargamingClientsFromEnv(observer)

	// Fetch client
	liveFetchClient, err := fetch.NewMultiSourceClient(liveClient, bkClient, db)
	if err != nil {
		log.Fatal().Msgf("fetch#NewMultiSourceClient failed %s", err)
	}

	cacheFetchClient, err := fetch.NewMultiSourceClient(cacheClient, bkClient, db)
	if err != nil {
		log.Fatal().Msgf("fetch#NewMultiSourceClient failed %s", err)
	}

	// PubSub client - shared across fetch clients
	pubsub := realtime.NewClient()

	return core.NewClient(liveFetchClient, liveClient, db, pubsub), core.NewClient(cacheFetchClient, cacheClient, db, pubsub)
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
	err = common.InitLoadedAssets()
	if err != nil {
		log.Fatal().Msgf("common#InitLoadedAssets failed %s", err)
	}
	err = localization.LoadAssets(static, "static/localization")
	if err != nil {
		log.Fatal().Msgf("localization#LoadAssets failed %s", err)
	}
}

func wargamingClientsFromEnv(observer metrics.ErrorObserver) (wargaming.Client, wargaming.Client) {
	liveClient, err := wargaming.NewClientFromEnv(constants.WargamingPrimaryAppID, constants.WargamingPrimaryAppRPS, constants.WargamingPrimaryAppRequestTimeout, constants.WargamingPrimaryAppProxyHostList)
	if err != nil {
		log.Fatal().Msgf("wargamingClientsFromEnv#NewClientFromEnv failed %s", err)
	}

	// This wargaming client is using a different proxy as it needs a lot higher rps, but can be slow
	cacheClient, err := wargaming.NewClientFromEnv(constants.WargamingCacheAppID, constants.WargamingCacheAppRPS, constants.WargamingCacheAppRequestTimeout, constants.WargamingCacheAppProxyHostList)
	if err != nil {
		log.Fatal().Msgf("wargamingClientsFromEnv#NewClientFromEnv failed %s", err)
	}

	liveClient = wargaming.NewTrackedClient(liveClient, observer)

	return liveClient, cacheClient
}

func newDatabaseClientFromEnv(ctx context.Context) (database.Client, error) {
	client, err := database.NewPostgresClient(ctx, constants.DatabaseConnString)
	if err != nil {
		return nil, fmt.Errorf("database#NewClient failed %w", err)
	}

	return client, nil
}

func registerHealthTransitionSinks(gw gateway.Client, tracker *metrics.ErrorTracker, alertClient alerts.DiscordAlertClient) {
	tracker.RegisterSink(func(t metrics.Transition) {
		setGatewayStatusFromHealthState(gw, t.To, t.Sources)
	})

	if alertClient == nil {
		return
	}

	tracker.RegisterSink(func(t metrics.Transition) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		header := fmt.Sprintf("**[HEALTH]** %s -> %s", strings.ToUpper(string(t.From)), strings.ToUpper(string(t.To)))
		err := alertClient.Error(ctx, header, formatTransitionSources(t.Sources))
		if err != nil {
			log.Err(err).Msg("failed to send health transition alert")
		}
	})
}

func setGatewayStatusFromHealthState(gw gateway.Client, state metrics.HealthState, sources []metrics.SourceHealth) {
	var err error
	switch state {
	case metrics.HealthStateCritical:
		err = gw.SetStatus(gateway.StatusRed, healthStatusText(sources), nil)
	case metrics.HealthStateWarning:
		err = gw.SetStatus(gateway.StatusYellow, healthStatusText(sources), nil)
	default:
		err = gw.SetStatus(gateway.StatusListening, "/help", nil)
	}

	if err != nil {
		log.Err(err).Msg("failed to set gateway status")
	}
}

func healthStatusText(sources []metrics.SourceHealth) string {
	var hasDiscord bool
	var hasWargaming bool

	for _, source := range sources {
		if source.State == metrics.HealthStateHealthy {
			continue
		}
		if source.Source == "discord" {
			hasDiscord = true
		}
		if source.Source == "wargaming" {
			hasWargaming = true
		}
	}

	if hasDiscord {
		return "Service Disrupted - Discord"
	}
	if hasWargaming {
		return "Service Disrupted - Wargaming"
	}

	return "Service Disrupted"
}

func formatTransitionSources(sources []metrics.SourceHealth) string {
	if len(sources) < 1 {
		return "no data points in rolling window"
	}

	var b strings.Builder
	for _, source := range sources {
		rate := source.ErrorRate * 100
		_, _ = fmt.Fprintf(&b, "%s state=%s failed=%d total=%d rate=%.2f%%\n", source.Source, source.State, source.Failed, source.Total, rate)
	}
	return strings.TrimSpace(b.String())
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
