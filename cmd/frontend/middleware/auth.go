package middleware

import (
	"net/http"
	"time"

	"github.com/cufee/aftermath/cmd/frontend/handler"
	"github.com/cufee/aftermath/cmd/frontend/logic/auth"
)

var SessionCheck handler.Middleware = func(ctx *handler.Context, next func(ctx *handler.Context) error) func(ctx *handler.Context) error {
	user, err := ctx.SessionUser()
	if err == nil && user.ID != "" {
		return next
	}

	return func(ctx *handler.Context) error {
		ctx.SetCookie(auth.NewSessionCookie("", time.Time{}))
		return ctx.Redirect("/login?from="+ctx.URL().Path, http.StatusTemporaryRedirect)
	}
}
