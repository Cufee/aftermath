package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cufee/aftermath/cmd/frontend/handler"
	"github.com/cufee/aftermath/cmd/frontend/logic/auth"
	"github.com/cufee/aftermath/internal/constants"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/json"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/logic"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
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
	realm, ok := ctx.Wargaming().RealmFromID(accountID)
	if !ok {
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

	err = verifyAccountToken(ctx.Context, constants.WargamingPublicAppID, realm, accountID, accessToken)
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

	go func(fetch fetch.Client, id string) {
		// make sure this account exists in the database
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		_, err := fetch.Account(ctx, id)
		if err != nil {
			log.Err(err).Msg("failed to cache an account after login")
		}
	}(ctx.Fetch(), accountID)

	return ctx.Redirect("/linked?nickname="+ctx.Query("nickname"), http.StatusTemporaryRedirect)
}

var WargamingBegin handler.Endpoint = func(ctx *handler.Context) error {
	user, err := ctx.SessionUser()
	if err != nil {
		return ctx.Redirect("/login", http.StatusTemporaryRedirect)
	}

	realm, ok := ctx.Wargaming().ParseRealm(ctx.Path("realm"))
	if !ok {
		return ctx.Redirect("/error?message=invalid realm", http.StatusTemporaryRedirect)
	}

	domain, ok := realm.DomainWorldOfTanksAPI()
	if !ok {
		return ctx.Redirect("/error?message=invalid realm", http.StatusTemporaryRedirect)
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
	redirectURL := fmt.Sprintf("https://%s/wot/auth/login/?application_id=%s&expires_at=%s&redirect_uri=%s", domain, constants.WargamingPublicAppID, tokenExpiration, auth.NewWargamingAuthRedirectURL(verifySession.PublicID))
	return ctx.Redirect(redirectURL, http.StatusTemporaryRedirect)
}

func verifyAccountToken(ctx context.Context, appID string, realm types.Realm, accountID string, token string) error {
	domain, ok := realm.DomainBlitzAPI()
	if !ok {
		return wargaming.ErrRealmNotSupported
	}

	url := fmt.Sprintf("https://%s/wotb/account/info/?application_id=%s&account_id=%s&access_token=%s", domain, appID, accountID, token)
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

	if res.StatusCode != http.StatusOK {
		return errors.New("bad status code")
	}

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
