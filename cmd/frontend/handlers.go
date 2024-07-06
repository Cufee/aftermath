package frontend

import (
	"embed"
	"io/fs"
	"net/http"
	"path/filepath"

	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/cmd/core/server"
	"github.com/cufee/aftermath/cmd/frontend/handler"
	"github.com/cufee/aftermath/cmd/frontend/middleware"
	"github.com/cufee/aftermath/cmd/frontend/routes"
	"github.com/cufee/aftermath/cmd/frontend/routes/api/auth"
	aWidget "github.com/cufee/aftermath/cmd/frontend/routes/api/widget"
	"github.com/cufee/aftermath/cmd/frontend/routes/app"
	r "github.com/cufee/aftermath/cmd/frontend/routes/redirect"
	"github.com/cufee/aftermath/cmd/frontend/routes/widget"
	"github.com/pkg/errors"
)

//go:embed public
var publicFS embed.FS

/*
Returns a slice of all handlers registered by the frontend
*/
func Handlers(core core.Client) ([]server.Handler, error) {
	assetsFS, err := fs.Sub(publicFS, "public")
	if err != nil {
		return nil, errors.Wrap(err, "failed to get embedded assets")
	}

	// https://go.dev/blog/routing-enhancements
	return []server.Handler{
		// assets
		{
			Path: get("/assets"),
			Func: redirect("/"),
		},
		{
			Path: "GET /assets/{_...}",
			Func: http.StripPrefix("/assets/", http.FileServerFS(assetsFS)).ServeHTTP,
		},
		// wildcard to catch all invalid requests
		{
			Path: "GET /{pathname...}",
			Func: handler.Chain(core, routes.ErrorNotFound),
		},
		// common routes
		{
			Path: get("/"),
			Func: handler.Chain(core, routes.Index),
		},
		{
			Path: get("/error"),
			Func: handler.Chain(core, routes.GenericError),
		},
		{
			Path: get("/login"),
			Func: handler.Chain(core, routes.Login),
		},
		{
			Path: get("/linked"),
			Func: handler.Chain(core, routes.AccountLinked),
		},
		// widget
		{
			Path: get("/widget"),
			Func: handler.Chain(core, widget.WidgetHome),
		},
		{
			Path: get("/widget/{accountId}"),
			Func: handler.Chain(core, widget.WidgetPreview),
		},
		{
			Path: get("/widget/{accountId}/live"),
			Func: handler.Chain(core, widget.LiveWidget),
		},
		{
			Path: get("/widget/personal"),
			Func: redirect("/app/widget"),
		},
		{
			Path: get("/widget/personal/{widgetId}/live"),
			Func: handler.Chain(core, widget.PersonalLiveWidget),
		},
		// app routes
		{
			Path: get("/app"),
			Func: handler.Chain(core, app.Index, middleware.SessionCheck),
		},
		{
			Path: get("/app/widgets"),
			Func: handler.Chain(core, app.Widgets, middleware.SessionCheck),
		},
		// api routes
		{
			Path: get("/api/auth/discord"),
			Func: handler.Chain(core, auth.DiscordRedirect),
		},
		{
			Path: get("/api/auth/wargaming/login/{realm}"),
			Func: handler.Chain(core, auth.WargamingBegin),
		},
		{
			Path: get("/api/auth/wargaming/redirect/{token}"),
			Func: handler.Chain(core, auth.WargamingRedirect),
		},
		{
			Path: get("/api/widget/mock"),
			Func: handler.Chain(core, aWidget.MockWidget),
		},
		{
			Path: get("/api/widget/{accountId}"),
			Func: handler.Chain(core, aWidget.AccountWidget),
		},
		// redirects
		{
			Path: get("/r/verify/{realm}"),
			Func: handler.Chain(core, r.VerifyFromDiscord),
		},
	}, nil
}

func get(path string) string {
	// we add the suffix in order to route past the wildcard handler correctly
	return "GET " + filepath.Join(path, "{$}")
}

func redirect(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { http.Redirect(w, r, path, http.StatusPermanentRedirect) }
}
