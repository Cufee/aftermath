package discord

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type userResponse struct {
	User   discordgo.User `json:"user"`
	Scopes []string       `json:"scopes"`
}

func GetUserFromToken(token authToken) (discordgo.User, error) {
	url := discordgo.EndpointOAuth2 + "@me"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return discordgo.User{}, errors.Wrap(err, "failed to create a GET request")
	}
	req.Header.Add("Authorization", "Bearer "+token.Value)

	res, err := defaultClient.Do(req)
	if err != nil {
		return discordgo.User{}, errors.Wrap(err, "GET request failed")
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		body, _ := io.ReadAll(res.Body)
		log.Error().Str("body", string(body)).Msg("bad status code")
		return discordgo.User{}, errors.Errorf("received status code %d", res.StatusCode)
	}

	var data userResponse
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return discordgo.User{}, errors.Wrap(err, "failed to decode user data")
	}

	return data.User, nil
}
