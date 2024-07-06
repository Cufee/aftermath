//go:build ignore

package main

import (
	"embed"
	"net/http"
	"os"

	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/cmd/core/server"
	"github.com/cufee/aftermath/cmd/frontend"
	"github.com/cufee/aftermath/internal/constants"
	"github.com/cufee/aftermath/internal/localization"
	"github.com/cufee/aftermath/internal/stats/render/assets"
	render "github.com/cufee/aftermath/internal/stats/render/common/v1"
	"github.com/cufee/aftermath/tests"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

//go:generate go generate ./cmd/frontend/assets

//go:embed static/*
var static embed.FS

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	// Assets for rendering
	err = assets.LoadAssets(static, "static")
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

	coreClient := core.NewClient(tests.StaticTestingFetch(), nil, tests.StaticTestingDatabase())
	handlers, err := frontend.Handlers(coreClient)
	if err != nil {
		panic(err)
	}

	handlers = append(handlers, redirectHandlersFromEnv()...)

	listen := server.NewServer(os.Getenv("PORT"), handlers...)
	listen()
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
