package discord

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type authToken struct {
	Scope string `json:"scope"`
	Type  string `json:"token_type"`
	Value string `json:"access_token"`

	ExpiresAt    time.Time `json:"-"`
	ExpiresIn    int       `json:"expires_in"`
	RefreshToken string    `json:"refresh_token"`
}

var defaultClient = http.Client{
	Timeout: time.Second * 5,
}

const baseURL = "https://discord.com/oauth2"

var defaultClientID = os.Getenv("DISCORD_AUTH_CLIENT_ID")
var defaultClientSecret = os.Getenv("DISCORD_AUTH_CLIENT_SECRET")

var defaultScopes = os.Getenv("DISCORD_AUTH_DEFAULT_SCOPES")
var defaultRedirectURL = os.Getenv("DISCORD_AUTH_REDIRECT_URL")

type oAuthURL struct {
	scope        string
	state        string
	prompt       string
	clientID     string
	redirectURL  string
	responseType string
}

func NewOAuthURL(state string) oAuthURL {
	url := oAuthURL{
		prompt:       "none",
		state:        state,
		scope:        defaultScopes,
		clientID:     defaultClientID,
		redirectURL:  defaultRedirectURL,
		responseType: "code",
	}
	return url
}

func (u oAuthURL) String() string {
	// https://discord.com/developers/docs/topics/oauth2#authorization-code-grant
	base := baseURL + "/authorize"
	query := url.Values{}
	query.Set("scope", u.scope)
	query.Set("prompt", u.prompt)
	query.Set("state", u.state)
	query.Set("client_id", u.clientID)
	query.Set("redirect_uri", u.redirectURL)
	query.Set("response_type", u.responseType)
	return base + "?" + query.Encode()
}

func ExchangeOAuthCode(code, redirectURL string) (authToken, error) {
	base := discordgo.EndpointOAuth2 + "token"
	query := url.Values{}
	query.Set("code", code)
	query.Set("redirect_uri", defaultRedirectURL)
	query.Set("grant_type", "authorization_code")

	req, err := http.NewRequest("POST", base, strings.NewReader(query.Encode()))
	if err != nil {
		return authToken{}, errors.Wrap(err, "failed to create an exchange request")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(defaultClientID, defaultClientSecret)

	res, err := defaultClient.Do(req)
	if err != nil {
		return authToken{}, errors.Wrap(err, "failed to make a POST request")
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		body, _ := io.ReadAll(res.Body)
		log.Error().Str("body", string(body)).Msg("bad status code")
		return authToken{}, errors.Errorf("received status code %d", res.StatusCode)
	}

	var data authToken
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return authToken{}, errors.Wrap(err, "failed to decode response")
	}

	data.ExpiresAt = time.Now().Add(time.Second * time.Duration(data.ExpiresIn))
	return data, nil
}
