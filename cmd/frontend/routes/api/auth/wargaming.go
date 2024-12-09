package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/cufee/aftermath/cmd/frontend/handler"
	"github.com/cufee/aftermath/cmd/frontend/logic/auth"
	"github.com/cufee/aftermath/internal/constants"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/json"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/logic"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/pkg/errors"
)

var WargamingRedirect handler.Endpoint = func(ctx *handler.Context) error {
	if ctx.Query("status") == "error" {
		return ctx.Redirect("/error?message="+ctx.Query("message"), http.StatusTemporaryRedirect)
	}

	accountID := ctx.Query("account_id")
	if accountID == "" {
		return ctx.Redirect("/error?message=invalid response received from wargaming", http.StatusTemporaryRedirect)
	}
	accessToken := ctx.Query("access_token")
	if accessToken == "" {
		return ctx.Redirect("/error?message=invalid response received from wargaming", http.StatusTemporaryRedirect)
	}
	realm, err := ctx.Wargaming().RealmFromID(accountID)
	if err != nil {
		return ctx.Redirect("/error?message=invalid response received from wargaming", http.StatusTemporaryRedirect)
	}

	token := ctx.Path("token")
	if token == "" {
		log.Debug().Msg("missing token")
		return ctx.Redirect("/error?message=this verification link has expired", http.StatusTemporaryRedirect)
	}

	session, err := ctx.Database().FindSession(ctx.Context, token)
	if err != nil {
		log.Warn().Err(err).Msg("failed to find session")
		return ctx.Redirect("/error?message=this verification link has expired", http.StatusTemporaryRedirect)
	}
	if session.Meta["flow"] != "wargaming-redirect" {
		return ctx.Redirect("/error?message=this verification link has expired", http.StatusTemporaryRedirect)
	}

	err = verifyAccountToken(ctx.Context, constants.WargamingPublicAppID, realm.String(), accountID, accessToken)
	if err != nil {
		log.Warn().Err(err).Msg("failed to verify access token")

		err := ctx.Database().SetSessionExpiresAt(ctx.Context, session.ID, time.Time{})
		if err != nil {
			log.Err(err).Msg("failed to set session expiration")
		}
		return ctx.Redirect("/error?message=this verification link has expired", http.StatusTemporaryRedirect)
	}

	user, err := ctx.Database().GetUserByID(ctx.Context, session.UserID, database.WithConnections())
	if err != nil {
		log.Err(err).Msg("failed to find user")
		return ctx.Redirect("/error?message=this verification link has expired", http.StatusTemporaryRedirect)
	}

	err = ctx.Database().SetSessionExpiresAt(ctx.Context, session.ID, time.Time{})
	if err != nil {
		log.Err(err).Msg("failed to set session expiration")
		return ctx.Redirect("/error?message=this verification link has expired", http.StatusTemporaryRedirect)
	}

	// Mark all other connection to this account as not verified
	err = ctx.Database().SetReferenceConnectionsUnverified(ctx.Context, accountID)
	if err != nil && !database.IsNotFound(err) {
		log.Err(err).Msg("failed to claim connection verification status")
		return ctx.Err(err, "failed to update user connection")
	}

	var found bool
	for _, conn := range user.Connections {
		if conn.Type != models.ConnectionTypeWargaming {
			continue
		}

		conn.Selected = conn.ReferenceID == accountID
		if conn.ReferenceID == accountID {
			conn.Verified = true
			found = true
		}

		_, err := ctx.Database().UpdateUserConnection(ctx.Context, conn.ID, conn)
		if err != nil {
			return ctx.Err(err, "failed to update user connection")
		}
	}
	if !found {
		conn := models.UserConnection{
			Type:        models.ConnectionTypeWargaming,
			UserID:      user.ID,
			ReferenceID: accountID,
			Verified:    true,
			Selected:    true,
		}
		_, err := ctx.Database().UpsertUserConnection(ctx.Context, conn)
		if err != nil {
			return ctx.Err(err, "failed to update user connection")
		}
	}

	return ctx.Redirect("/linked?nickname="+ctx.Query("nickname"), http.StatusTemporaryRedirect)
}

var WargamingBegin handler.Endpoint = func(ctx *handler.Context) error {
	user, err := ctx.SessionUser()
	if err != nil {
		return ctx.Redirect("/login", http.StatusTemporaryRedirect)
	}

	baseWgURL, err := loginUriFromRealm(ctx.Path("realm"))
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

	tokenExpiration := "300" // seconds
	redirectURL := fmt.Sprintf("%s/auth/login/?application_id=%s&expires_at=%s&redirect_uri=%s", baseWgURL, constants.WargamingPublicAppID, tokenExpiration, auth.NewWargamingAuthRedirectURL(verifySession.PublicID))
	return ctx.Redirect(redirectURL, http.StatusTemporaryRedirect)
}

func verifyAccountToken(ctx context.Context, appID string, realm string, accountID string, token string) error {
	baseWgURL, err := verifyBaseUriFromRealm(realm)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/account/info/?application_id=%s&account_id=%s&access_token=%s", baseWgURL, appID, accountID, token)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var response types.WgResponse[map[string]types.Account]
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return err
	}
	if response.Error.Code != 0 {
		return errors.New(response.Error.Message)
	}
	if fmt.Sprint(response.Data[accountID].ID) != accountID {
		return errors.New("invalid token for account")
	}
	return nil
}

func loginUriFromRealm(realm string) (string, error) {
	switch strings.ToUpper(realm) {
	default:
		return "", errors.New("unsupported realm")

	case types.RealmEurope.String():
		return "https://api.worldoftanks.eu/wot", nil
	case types.RealmNorthAmerica.String():
		return "https://api.worldoftanks.com/wot", nil
	case types.RealmAsia.String():
		return "https://api.worldoftanks.asia/wot", nil
	}
}

func verifyBaseUriFromRealm(realm string) (string, error) {
	switch strings.ToUpper(realm) {
	default:
		return "", errors.New("unsupported realm")

	case types.RealmEurope.String():
		return "https://api.wotblitz.eu/wotb", nil
	case types.RealmNorthAmerica.String():
		return "https://api.wotblitz.com/wotb", nil
	case types.RealmAsia.String():
		return "https://api.wotblitz.asia/wotb", nil
	}
}
