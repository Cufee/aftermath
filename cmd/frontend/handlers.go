package frontend

import (
	"embed"
	"io/fs"
	"net/http"
	"path/filepath"

	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/cmd/core/server"
	"github.com/cufee/aftermath/cmd/frontend/handler"
	"github.com/cufee/aftermath/cmd/frontend/routes"
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
			Path: "GET /assets/{$}",
			Func: redirect("/"),
		},
		{
			Path: "GET /assets/{_...}",
			Func: http.StripPrefix("/assets/", http.FileServerFS(assetsFS)).ServeHTTP,
		},
		// wildcard to catch all invalid requests
		{
			Path: "GET /{pathname...}",
			Func: handler.Handler(core, routes.ErrorNotFound),
		},
		// common routes
		{
			Path: get("/"),
			Func: handler.Handler(core, routes.Index),
		},
		{
			Path: get("/error"),
			Func: handler.Handler(core, routes.GenericError),
		},
		// widget
		{
			Path: get("/widget/{accountId}"),
			Func: handler.Handler(core, widget.ConfigureWidget),
		},
		{
			Path: get("/widget/{accountId}/live"),
			Func: handler.Handler(core, widget.LiveWidget),
		},
		// app routes
		//
	}, nil
}

func get(path string) string {
	// we add the suffix in order to route past the wildcard handler correctly
	return "GET " + filepath.Join(path, "{$}")
}

func redirect(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { http.Redirect(w, r, path, http.StatusPermanentRedirect) }
}
