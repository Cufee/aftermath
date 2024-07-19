package auth

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

var DiscordRedirect handler.Endpoint = func(ctx *handler.Context) error {
	log.Debug().Msg("started handling a discord redirect")

	code := ctx.Query("code")
	state := ctx.Query("state")
	if code == "" || state == "" {
		log.Debug().Msg("discord redirect missing state or code")
		return ctx.Redirect("/login?message=session not found", http.StatusTemporaryRedirect)
	}

	cookie, err := ctx.Cookie(auth.AuthNonceCookieName)
	if err != nil || cookie == nil {
		return ctx.Redirect("/login?message=session not found", http.StatusTemporaryRedirect)
	}
	if cookie.Value == "" {
		return ctx.Redirect("/login?message=session not found", http.StatusTemporaryRedirect)
	}
	log.Debug().Str("identifier", cookie.Value).Msg("handling a discord redirect")

	nonce, err := ctx.Database().FindAuthNonce(ctx.Context, state)
	if err != nil || !nonce.Active || nonce.ExpiresAt.Before(time.Now()) {
		log.Debug().Msg("discord redirect missing or invalid nonce")
		return ctx.Redirect("/login?message=session expired", http.StatusTemporaryRedirect)
	}
	err = ctx.Database().SetAuthNonceActive(ctx.Context, nonce.ID, false)
	if err != nil {
		log.Err(err).Msg("failed to update nonce active status")
		return ctx.Redirect("/login?message=session expired", http.StatusTemporaryRedirect)
	}

	identifier, err := ctx.Identifier()
	if err != nil {
		log.Err(err).Msg("failed to extract an identifier from a request")
		return ctx.Redirect("/login?message=session expired", http.StatusTemporaryRedirect)
	}

	if nonce.Identifier != identifier {
		log.Debug().Msg("discord redirect invalid identifier")
		return ctx.Redirect("/login?message=session expired", http.StatusTemporaryRedirect)
	}

	token, err := discord.ExchangeOAuthCode(code, nonce.Meta["redirectUrl"])
	if err != nil {
		log.Err(err).Msg("failed to exchange code for token")
		return ctx.Redirect("/login?message=failed to create a session", http.StatusTemporaryRedirect)
	}

	user, err := discord.GetUserFromToken(token)
	if err != nil {
		log.Err(err).Msg("failed to get user from token")
		return ctx.Redirect("/login?message=failed to create a session", http.StatusTemporaryRedirect)
	}
	if user.ID == "" { // just in case
		log.Error().Msg("blank user id received from discord")
		return ctx.Redirect("/login?message=failed to create a session", http.StatusTemporaryRedirect)
	}

	sessionID, err := logic.RandomString(32)
	if err != nil {
		log.Err(err).Msg("failed to generate a session id")
		return ctx.Redirect("/login?message=failed to create a session", http.StatusTemporaryRedirect)
	}

	sess, err := ctx.Database().CreateSession(ctx.Context, sessionID, user.ID, time.Now().Add(time.Hour*24*7), nil)
	if err != nil {
		log.Err(err).Msg("failed to create a new user session")
		return ctx.Redirect("/login?message=failed to create a session", http.StatusTemporaryRedirect)
	}

	ctx.SetCookie(auth.NewSessionCookie(sess.PublicID, sess.ExpiresAt))

	defer log.Debug().Msg("finished handling a discord redirect")
	if path, ok := nonce.Meta["from"]; ok && !strings.HasPrefix(path, "http://") && !strings.HasPrefix(path, "https://") {
		return ctx.Redirect(path, http.StatusTemporaryRedirect)
	}
	return ctx.Redirect("/app", http.StatusTemporaryRedirect)
}
