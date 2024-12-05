//go:build ignore

package main

import (
	"embed"
	"time"
)

//go:generate go generate ./cmd/frontend/assets/generate

//go:embed static/*
var static embed.FS

func main() {

	t, err := time.Parse(time.RFC3339Nano, "2024-07-28 19:04:52.920131689+02:00")
	if err != nil {
		panic(err)
	}
	println(t.String())

	return

	// err := godotenv.Load(".env")
	// if err != nil {
	// 	panic(err)
	// }

	// // Assets for rendering
	// err = assets.LoadAssets(static, "static")
	// if err != nil {
	// 	log.Fatal().Msgf("assets#LoadAssets failed to load assets from static/ embed.FS %s", err)
	// }
	// err = render.InitLoadedAssets()
	// if err != nil {
	// 	log.Fatal().Msgf("render#InitLoadedAssets failed %s", err)
	// }
	// err = localization.LoadAssets(static, "static/localization")
	// if err != nil {
	// 	log.Fatal().Msgf("localization#LoadAssets failed %s", err)
	// }

	// pubsub := realtime.NewClient()

	// coreClient := core.NewClient(tests.StaticTestingFetch(), nil, tests.StaticTestingDatabase(), pubsub)
	// handlers, err := frontend.Handlers(coreClient)
	// if err != nil {
	// 	panic(err)
	// }

	// handlers = append(handlers, redirectHandlersFromEnv()...)

	// listen := server.NewServer(os.Getenv("PORT"), handlers)
	// listen()
}

// func redirectHandlersFromEnv() []server.Handler {
// 	return []server.Handler{
// 		{
// 			Path: "GET /invite/{$}",
// 			Func: func(w http.ResponseWriter, r *http.Request) {
// 				http.Redirect(w, r, constants.DiscordBotInviteURL, http.StatusTemporaryRedirect)
// 			},
// 		},
// 		{
// 			Path: "GET /join/{$}",
// 			Func: func(w http.ResponseWriter, r *http.Request) {
// 				http.Redirect(w, r, constants.DiscordPrimaryGuildInviteURL, http.StatusTemporaryRedirect)
// 			},
// 		},
// 	}
// }
