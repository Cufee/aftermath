package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/cufee/aftermath/cmd/frontend/handler"
	"github.com/cufee/aftermath/cmd/frontend/logic/auth"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/logic"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

var WargamingRedirect handler.Endpoint = func(ctx *handler.Context) error {
	accountID := ctx.Query("account_id")
	if accountID == "" {
		return ctx.Redirect("/error?message=invalid response received from wargaming", http.StatusTemporaryRedirect)
	}
	if ctx.Query("status") == "error" {
		return ctx.Redirect("/error?message="+ctx.Query("message"), http.StatusTemporaryRedirect)
	}

	token := ctx.Path("token")
	if token == "" {
		return ctx.Redirect("/error?message=this verification link has expired", http.StatusTemporaryRedirect)
	}

	user, err := ctx.Database().UserFromSession(ctx.Context, token, database.WithConnections())
	if err != nil {
		return ctx.Redirect("/error?message=this verification link has expired", http.StatusTemporaryRedirect)
	}

	err = ctx.Database().SetSessionExpiresAt(ctx.Context, token, time.Time{})
	if err != nil {
		log.Err(err).Msg("failed to set session expiration")
		return ctx.Redirect("/error?message=this verification link has expired", http.StatusTemporaryRedirect)
	}

	connections, _ := user.FilterConnections(models.ConnectionTypeWargaming, nil)
	for _, conn := range connections {
		if conn.ReferenceID == accountID {
			conn.Metadata["default"] = true
			conn.Metadata["verified"] = true
			_, err := ctx.Database().UpdateConnection(ctx.Context, conn)
			if err != nil {
				return ctx.Error(err, "failed to update user connection")
			}
			return ctx.Redirect("/app#links", http.StatusTemporaryRedirect)
		}
	}

	conn := models.UserConnection{
		Type:        models.ConnectionTypeWargaming,
		UserID:      user.ID,
		ReferenceID: accountID,
		Metadata:    map[string]any{"verified": true, "default": true},
	}
	_, err = ctx.Database().UpsertConnection(ctx.Context, conn)
	if err != nil {
		return ctx.Error(err, "failed to update user connection")
	}
	return ctx.Redirect("/app#links", http.StatusTemporaryRedirect)
}

var WargamingBegin handler.Endpoint = func(ctx *handler.Context) error {
	user, err := ctx.SessionUser()
	if err != nil {
		return ctx.Redirect("/login", http.StatusTemporaryRedirect)
	}

	baseWgURL, err := authURLFromRealm(ctx.Path("realm"))
	if err != nil {
		return ctx.Redirect("/app", http.StatusTemporaryRedirect)
	}

	verifySessionID, err := logic.RandomString(32)
	if err != nil {
		return ctx.Redirect("/error?message=failed to start a new session", http.StatusTemporaryRedirect)
	}

	verifySession, err := ctx.Database().CreateSession(ctx.Context, verifySessionID, user.ID, time.Now().Add(time.Minute), map[string]string{"flow": "wargaming-redirect"})
	if err != nil {
		log.Err(err).Str("userId", user.ID).Msg("failed to create a new user session for wargaming auth flow")
		return ctx.Redirect("/error?message=failed to start a new session", http.StatusTemporaryRedirect)
	}

	redirectURL := fmt.Sprintf("%s?application_id=%s&redirect_uri=%s", baseWgURL, wargaming.PublicAppID, auth.NewWargamingAuthRedirectURL(verifySession.ID))
	return ctx.Redirect(redirectURL, http.StatusTemporaryRedirect)
}

func authURLFromRealm(realm string) (string, error) {
	switch strings.ToUpper(realm) {
	case "EU":
		return "https://api.worldoftanks.eu/wot/auth/login/", nil
	case "NA":
		return "https://api.worldoftanks.com/wot/auth/login/", nil
	case "AS":
		return "https://api.worldoftanks.asia/wot/auth/login/", nil
	default:
		return "", errors.New("unknown realm")
	}
}
