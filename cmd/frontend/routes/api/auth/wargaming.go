package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/cufee/aftermath/cmd/frontend/handler"
	"github.com/cufee/aftermath/cmd/frontend/logic/auth"
	"github.com/cufee/aftermath/internal/constants"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/logic"
	"github.com/pkg/errors"
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
		log.Debug().Msg("missing token")
		return ctx.Redirect("/error?message=this verification link has expired", http.StatusTemporaryRedirect)
	}

	session, err := ctx.Database().FindSession(ctx.Context, token)
	if err != nil {
		log.Debug().Err(err).Msg("failed to find session")
		return ctx.Redirect("/error?message=this verification link has expired", http.StatusTemporaryRedirect)
	}
	if session.Meta["flow"] != "wargaming-redirect" {
		return ctx.Redirect("/error?message=this verification link has expired", http.StatusTemporaryRedirect)
	}

	user, err := ctx.Database().GetUserByID(ctx.Context, session.UserID, database.WithConnections())
	if err != nil {
		log.Debug().Err(err).Msg("failed to find user")
		return ctx.Redirect("/error?message=this verification link has expired", http.StatusTemporaryRedirect)
	}

	err = ctx.Database().SetSessionExpiresAt(ctx.Context, session.ID, time.Time{})
	if err != nil {
		log.Err(err).Msg("failed to set session expiration")
		return ctx.Redirect("/error?message=this verification link has expired", http.StatusTemporaryRedirect)
	}

	connections, _ := user.FilterConnections(models.ConnectionTypeWargaming, nil)
	var found bool
	for _, conn := range connections {
		conn.Metadata["default"] = conn.ReferenceID == accountID
		if conn.ReferenceID == accountID {
			conn.Metadata["verified"] = true
			found = true
		}

		_, err := ctx.Database().UpdateConnection(ctx.Context, conn)
		if err != nil {
			return ctx.Error(err, "failed to update user connection")
		}
	}
	if !found {
		conn := models.UserConnection{
			Type:        models.ConnectionTypeWargaming,
			UserID:      user.ID,
			ReferenceID: accountID,
			Metadata:    map[string]any{"verified": true, "default": true},
		}
		_, err := ctx.Database().UpsertConnection(ctx.Context, conn)
		if err != nil {
			return ctx.Error(err, "failed to update user connection")
		}
	}

	return ctx.Redirect("/linked?nickname="+ctx.Query("nickname"), http.StatusTemporaryRedirect)
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

	verifySession, err := ctx.Database().CreateSession(ctx.Context, verifySessionID, user.ID, time.Now().Add(time.Minute*5), map[string]string{"flow": "wargaming-redirect"})
	if err != nil {
		log.Err(err).Str("userId", user.ID).Msg("failed to create a new user session for wargaming auth flow")
		return ctx.Redirect("/error?message=failed to start a new session", http.StatusTemporaryRedirect)
	}

	redirectURL := fmt.Sprintf("%s?application_id=%s&redirect_uri=%s", baseWgURL, constants.WargamingPublicAppID, auth.NewWargamingAuthRedirectURL(verifySession.PublicID))
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
