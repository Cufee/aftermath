package handler

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/a-h/templ"
	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/cmd/frontend/logic/auth"
	"github.com/pkg/errors"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/logic"
	"github.com/rs/zerolog/log"
)

var devMode = os.Getenv("AUTH_DEV_MODE") == "true"

type Servable interface {
	Serve(ctx *Context) error
}

type Middleware func(ctx *Context, next func(ctx *Context) error) func(ctx *Context) error

type Context struct {
	core.Client
	context.Context

	user *models.User

	w http.ResponseWriter
	r *http.Request
}

type Layout func(ctx *Context, children ...templ.Component) (templ.Component, error)
type Page func(ctx *Context) (Layout, templ.Component, error)
type Partial func(ctx *Context) (templ.Component, error)
type Endpoint func(ctx *Context) error

var (
	_ Servable = new(Page)
	_ Servable = new(Partial)
	_ Servable = new(Endpoint)
)

func (ctx *Context) Cookie(key string) (*http.Cookie, error) {
	return ctx.r.Cookie(key)
}
func (ctx *Context) SetCookie(cookie *http.Cookie) {
	http.SetCookie(ctx.w, cookie)
}
func (ctx *Context) Query(key string) string {
	return ctx.r.URL.Query().Get(key)
}
func (ctx *Context) Form(key string) string {
	return ctx.r.Form.Get(key)
}
func (ctx *Context) Path(key string) string {
	return ctx.r.PathValue(key)
}
func (ctx *Context) URL() *url.URL {
	return ctx.r.URL
}
func (ctx *Context) RealIP() (string, bool) {
	if ip := ctx.r.Header.Get("X-Forwarded-For"); ip != "" {
		return ip, true
	}
	if ip := ctx.r.RemoteAddr; ip != "" {
		return ip, true
	}
	return "", false
}

/*
Returns a stable identifier based on requestor ip address
*/
func (ctx *Context) Identifier() (string, error) {
	if ip, ok := ctx.RealIP(); ok {
		return logic.HashString(ip), nil
	}
	return "", errors.New("failed to extract ip address")
}

func (ctx *Context) SessionUser() (*models.User, error) {
	if ctx.user != nil {
		return ctx.user, nil
	}
	if devMode {
		user, _ := ctx.Database().UserFromSession(ctx.Context, "dev-user")
		return &user, nil
	}

	cookie, err := ctx.Cookie(auth.SessionCookieName)
	if err != nil || cookie == nil {
		return nil, ErrSessionNotFound
	}
	if cookie.Value == "" {
		return nil, ErrSessionNotFound
	}

	user, err := ctx.Database().UserFromSession(ctx.Context, cookie.Value)
	if err != nil {
		if database.IsNotFound(err) {
			return nil, ErrSessionNotFound
		}
		return nil, err
	}

	ctx.user = &user
	return ctx.user, nil
}

/*
Redirects a user to /error with an error message set as query param
*/
func (ctx *Context) Error(err error, context ...string) error {
	query := make(url.Values)
	if err != nil {
		query.Set("message", err.Error())
	}
	if len(context) > 1 {
		log.Err(err).Msg("error while serving") // end user does not get any context, so we log the error instead
		query.Set("message", strings.Join(context, ", "))
	}

	http.Redirect(ctx.w, ctx.r, "/error?"+query.Encode(), http.StatusTemporaryRedirect)
	return nil // this would never cause an error
}

func (ctx *Context) Redirect(path string, code int) error {
	http.Redirect(ctx.w, ctx.r, path, code)
	return nil // this would never cause an error
}

func (ctx *Context) SetStatus(code int) {
	ctx.w.WriteHeader(code)
}

func newContext(core core.Client, w http.ResponseWriter, r *http.Request) *Context {
	return &Context{core, r.Context(), nil, w, r}
}

func (partial Partial) Serve(ctx *Context) error {
	content, err := partial(ctx)
	if err != nil {
		return ctx.Error(err)
	}
	if content == nil {
		return nil
	}

	err = content.Render(ctx.Context, ctx.w)
	if err != nil {
		return ctx.Error(err)
	}

	return nil
}

func (page Page) Serve(ctx *Context) error {
	layout, body, err := page(ctx)
	if err != nil {
		return ctx.Error(err, "failed to render the page")
	}
	if layout == nil && body == nil {
		return nil
	} else if layout == nil {
		return body.Render(ctx.Context, ctx.w)
	}

	withLayout, err := layout(ctx, body)
	if err != nil {
		return ctx.Error(err, "failed to render the layout")
	}
	if withLayout == nil {
		return nil
	}

	err = withLayout.Render(ctx.Context, ctx.w)
	if err != nil {
		return ctx.Error(err, "failed to render content")
	}

	return nil
}

func (endpoint Endpoint) Serve(ctx *Context) error {
	err := endpoint(ctx)
	if err != nil {
		return ctx.Error(err, "internal server error")
	}
	return nil
}

func Chain(core core.Client, serve Servable, middleware ...Middleware) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := newContext(core, w, r)
		chain := serve.Serve
		for i := len(middleware) - 1; i >= 0; i-- {
			chain = middleware[i](ctx, chain)
		}
		err := chain(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Err(err).Msg("unhandled error in handler chain, this should never happen")
		}
	}
}
