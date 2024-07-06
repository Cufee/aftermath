package middleware

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/cufee/aftermath/cmd/frontend/handler"
	"github.com/cufee/aftermath/cmd/frontend/logic/auth"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/rs/zerolog/log"
)

var SessionCheck handler.Middleware = func(ctx *handler.Context, next func(ctx *handler.Context) error) func(ctx *handler.Context) error {
	user, err := ctx.SessionUser()
	if err == nil && user.ID != "" {
		go func(session *models.Session, err error) {
			if err != nil {
				log.Err(err).Msg("failed to retrieve a session after retrieving the user from it")
				return
			}
			if session.ExpiresAt.After(time.Now().Add(time.Hour * 24)) {
				// we only update the session once it is about to expire
				return
			}
			if slices.Contains([]string{"wargaming-redirect"}, session.Meta["flow"]) {
				// some session types should not be refreshed, even though this handler should never check them in the first place
				return
			}

			c, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			err = ctx.Database().SetSessionExpiresAt(c, session.ID, time.Now().Add(time.Hour*24*7))
			if err != nil {
				log.Err(err).Msg("failed to set session expiration")
				return
			}
		}(ctx.Session())

		return next
	}

	return func(ctx *handler.Context) error {
		ctx.SetCookie(auth.NewSessionCookie("", time.Time{}))
		return ctx.Redirect("/login?from="+ctx.URL().Path, http.StatusTemporaryRedirect)
	}
}
