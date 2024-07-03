package handler

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/a-h/templ"
	"github.com/cufee/aftermath/cmd/core"
	"github.com/rs/zerolog/log"
)

type Renderable interface {
	Render(ctx *Context) error
}

type Context struct {
	core.Client
	context.Context

	w http.ResponseWriter
	r *http.Request
}

func (ctx *Context) Req() *http.Request {
	return ctx.r
}

func (ctx *Context) Error(err error, context ...string) error {

	query := make(url.Values)
	if err != nil {
		query.Set("message", err.Error())
	}
	if len(context) > 1 {
		log.Err(err).Msg("error while rendering") // end user does not get any context, so we log the error instead
		query.Set("message", strings.Join(context, ", "))
	}

	http.Redirect(ctx.w, ctx.r, "/error?"+query.Encode(), http.StatusTemporaryRedirect)
	return nil // this would never cause an error
}

func (ctx *Context) SetStatus(code int) {
	ctx.w.WriteHeader(code)
}

func newContext(core core.Client, w http.ResponseWriter, r *http.Request) *Context {
	return &Context{core, r.Context(), w, r}
}

type Layout func(ctx *Context, children ...templ.Component) (templ.Component, error)

type Partial func(ctx *Context) (templ.Component, error)

func (partial Partial) Render(ctx *Context) error {
	content, err := partial(ctx)
	if err != nil {
		return ctx.Error(err)
	}

	err = content.Render(ctx.Context, ctx.w)
	if err != nil {
		return ctx.Error(err)
	}

	return nil
}

type Page func(ctx *Context) (Layout, templ.Component, error)

func (page Page) Render(ctx *Context) error {
	layout, body, err := page(ctx)
	if err != nil {
		return ctx.Error(err, "failed to render the page")
	}

	content, err := layout(ctx, body)
	if err != nil {
		return ctx.Error(err, "failed to render the layout")
	}

	err = content.Render(ctx.Context, ctx.w)
	if err != nil {
		return ctx.Error(err, "failed to render content")
	}

	return nil
}

func Handler(core core.Client, content Renderable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := content.Render(newContext(core, w, r))
		if err != nil {
			// this should never be the case, we return an error to make early returns more convenient
			log.Err(err).Msg("handler failed to render, this should never happen")
		}
	}
}
