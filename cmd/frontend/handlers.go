package frontend

import (
	"embed"
	"fmt"
	"io/fs"

	"net/http"
	"net/url"

	"strings"

	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/cmd/core/server"
	"github.com/cufee/aftermath/cmd/frontend/handler"
	"github.com/cufee/aftermath/cmd/frontend/middleware"
	"github.com/cufee/aftermath/cmd/frontend/routes"
	"github.com/cufee/aftermath/cmd/frontend/routes/api"
	"github.com/cufee/aftermath/cmd/frontend/routes/api/auth"
	"github.com/cufee/aftermath/cmd/frontend/routes/api/realtime"
	wa "github.com/cufee/aftermath/cmd/frontend/routes/api/widget"
	a "github.com/cufee/aftermath/cmd/frontend/routes/app"
	"github.com/cufee/aftermath/cmd/frontend/routes/app/widgets"
	r "github.com/cufee/aftermath/cmd/frontend/routes/redirect"
	w "github.com/cufee/aftermath/cmd/frontend/routes/widget"
	"github.com/cufee/aftermath/internal/log"
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
	srv := App(core)

	root := srv.Group("/")
	root.GET("/", routes.Index)
	root.GET("/login", routes.LoginStatic)
	root.GET("/error", routes.GenericError)
	root.GET("/linked", routes.AccountLinked)
	root.GET("/{pathname...}", routes.ErrorNotFound)

	assets := srv.Group("/assets")
	assets.GET("/", handler.Redirect("/", http.StatusMovedPermanently))
	assets.GET("/{_...}", NewAssetsHandler(assetsFS))

	legal := srv.Group("/legal")
	legal.GET("/privacy-policy", routes.PrivacyPolicy)
	legal.GET("/terms-of-service", routes.TermsOfService)

	widget := srv.Group("/widget")
	widget.GET("/", w.WidgetHome)
	widget.GET("/account/{accountId}", w.WidgetPreview)
	widget.GET("/account/{accountId}/live", w.LiveWidget)
	widget.GET("/custom/{widgetId}/live", w.CustomLiveWidget)
	widget.GET("/custom", handler.Redirect("/app/widget", http.StatusMovedPermanently))

	app := srv.Group("/app", middleware.SessionCheck)
	app.GET("/", a.Index)
	app.GET("/widgets/new", widgets.NewCustom)
	app.GET("/widgets/{widgetId}", widgets.EditSettings)

	secureApi := srv.Group("/api/s", middleware.SessionCheck)
	secureApi.PATCH("/widget/custom", wa.CreateCustomWidget)
	secureApi.PATCH("/widget/custom/{widgetId}", wa.UpdateCustomWidget)
	secureApi.PATCH("/widget/custom/{widgetId}/action", wa.QuickAction)
	secureApi.PATCH("/widget/custom/{widgetId}/session", wa.ResetSession)

	secureApi.DELETE("/connections/{connectionId}", api.RemoveConnection)
	secureApi.PATCH("/connections/{connectionId}/default", api.SetDefaultConnection)

	publicApi := srv.Group("/api/p")
	publicApi.GET("/login", api.Login)
	publicApi.GET("/auth/discord", auth.DiscordRedirect)
	publicApi.GET("/auth/wargaming/login/{realm}", auth.WargamingBegin)
	publicApi.GET("/auth/wargaming/redirect/{token}", auth.WargamingRedirect)
	publicApi.GET("/widget/{accountId}", wa.AccountWidget)
	publicApi.GET("/widget/mock", wa.MockWidget)
	publicApi.GET("/realtime/widget/custom/{widgetId}", realtime.WidgetSettings)

	redirect := srv.Group("/r")
	redirect.GET("/verify/{realm}", r.VerifyFromDiscord)

	return srv.Handlers(), nil
}

type app struct {
	groups []*group
	core   core.Client
}

func App(core core.Client) *app {
	return &app{core: core}
}

func (a *app) Group(prefix string, middleware ...handler.Middleware) *group {
	g := &group{core: a.core, prefix: prefix, middleware: middleware}
	a.groups = append(a.groups, g)
	return g
}

func (a *app) Handlers() []server.Handler {
	var handlers []server.Handler
	for _, g := range a.groups {
		handlers = append(handlers, g.Handlers...)
	}
	return handlers
}

type group struct {
	core       core.Client
	prefix     string
	Handlers   []server.Handler
	middleware []handler.Middleware
}

func (g *group) buildPath(m, p string) string {
	elms := []string{p}
	if !strings.HasSuffix(p, "...}") {
		elms = append(elms, "{$}")
	}

	path, err := url.JoinPath(g.prefix, elms...)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to build a path")
	}
	path, _ = url.PathUnescape(path)

	if m == "" {
		return path
	}
	return fmt.Sprintf("%s %s", strings.ToUpper(m), path)
}

func (g *group) POST(path string, h handler.Servable, middleware ...handler.Middleware) {
	g.Handlers = append(g.Handlers, server.Handler{Path: g.buildPath("POST", path), Func: handler.Chain(g.core, h, append(middleware, g.middleware...)...)})
}
func (g *group) GET(path string, h handler.Servable, middleware ...handler.Middleware) {
	g.Handlers = append(g.Handlers, server.Handler{Path: g.buildPath("GET", path), Func: handler.Chain(g.core, h, append(middleware, g.middleware...)...)})
}
func (g *group) PUT(path string, h handler.Servable, middleware ...handler.Middleware) {
	g.Handlers = append(g.Handlers, server.Handler{Path: g.buildPath("PUT", path), Func: handler.Chain(g.core, h, append(middleware, g.middleware...)...)})
}
func (g *group) PATCH(path string, h handler.Servable, middleware ...handler.Middleware) {
	g.Handlers = append(g.Handlers, server.Handler{Path: g.buildPath("PATCH", path), Func: handler.Chain(g.core, h, append(middleware, g.middleware...)...)})
}
func (g *group) DELETE(path string, h handler.Servable, middleware ...handler.Middleware) {
	g.Handlers = append(g.Handlers, server.Handler{Path: g.buildPath("DELETE", path), Func: handler.Chain(g.core, h, append(middleware, g.middleware...)...)})
}
