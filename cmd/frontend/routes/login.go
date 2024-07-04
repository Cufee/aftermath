package routes

import (
	"net/http"
	"time"

	"github.com/cufee/aftermath/cmd/frontend/handler"
	"github.com/cufee/aftermath/cmd/frontend/logic/auth"
	"github.com/cufee/aftermath/cmd/frontend/logic/discord"
	"github.com/cufee/aftermath/internal/logic"
	"github.com/rs/zerolog/log"
)

var Login handler.Endpoint = func(ctx *handler.Context) error {
	user, err := ctx.SessionUser()
	if err == nil && user.ID != "" {
		return ctx.Redirect("/app", http.StatusTemporaryRedirect)
	}

	nonceID, err := logic.RandomString(32)
	if err != nil {
		return ctx.Error(err, "failed to authenticate")
	}

	identifier, err := ctx.Identifier()
	if err != nil {
		return ctx.Error(err, "failed to extract an identifier")
	}
	log.Debug().Str("identifier", identifier).Msg("new login request")

	redirectURL := discord.NewOAuthURL(nonceID)
	meta := map[string]string{"from": ctx.Query("from"), "redirectUrl": redirectURL.String()}

	nonce, err := ctx.Database().CreateAuthNonce(ctx.Context, nonceID, identifier, time.Now().Add(time.Minute*5), meta)
	if err != nil {
		return ctx.Error(err, "failed to authenticate")
	}

	ctx.SetCookie(auth.NewNonceCookie(nonce.PublicID, nonce.ExpiresAt))
	return ctx.Redirect(redirectURL.String(), http.StatusTemporaryRedirect)
}
