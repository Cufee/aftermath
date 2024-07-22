package api

import (
	"net/http"
	"strings"
	"time"

	"github.com/cufee/aftermath/cmd/frontend/handler"
	"github.com/cufee/aftermath/cmd/frontend/logic/auth"
	"github.com/cufee/aftermath/cmd/frontend/logic/discord"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/logic"
)

var Login handler.Endpoint = func(ctx *handler.Context) error {
	user, err := ctx.SessionUser()
	if err == nil && user.ID != "" {
		if from := strings.ToLower(ctx.Query("from")); from != "" && !strings.HasPrefix(from, "http://") && !strings.HasPrefix(from, "https://") {
			return ctx.Redirect(from, http.StatusTemporaryRedirect)
		}
		return ctx.Redirect("/app", http.StatusTemporaryRedirect)
	}

	nonceID, err := logic.RandomString(32)
	if err != nil {
		return ctx.Err(err, "failed to authenticate")
	}

	identifier, err := ctx.Identifier()
	if err != nil {
		return ctx.Err(err, "failed to extract an identifier")
	}
	log.Debug().Str("identifier", identifier).Msg("new login request")

	redirectURL := discord.NewOAuthURL(nonceID)
	meta := map[string]string{"redirectUrl": redirectURL.String()}
	if path := ctx.Query("from"); path != "" && !strings.HasPrefix(path, "http://") && !strings.HasPrefix(path, "https://") {
		meta["from"] = path
	}

	nonce, err := ctx.Database().CreateAuthNonce(ctx.Context, nonceID, identifier, time.Now().Add(time.Minute*5), meta)
	if err != nil {
		return ctx.Err(err, "failed to authenticate")
	}

	ctx.SetCookie(auth.NewNonceCookie(nonce.PublicID, nonce.ExpiresAt))
	return ctx.Redirect(redirectURL.String(), http.StatusTemporaryRedirect)
}
