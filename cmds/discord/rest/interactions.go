package rest

import (
	"net/http"

	"github.com/bwmarrin/discordgo"
)

func (c *Client) SendInteractionResponse(id, token string, data discordgo.InteractionResponse) error {
	req, err := c.interactionRequest("POST", discordgo.EndpointInteractionResponse(id, token), data, data.Data.Files)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

func (c *Client) UpdateInteractionResponse(id, token string, data discordgo.InteractionResponseData) error {
	req, err := c.interactionRequest("PATCH", discordgo.EndpointInteractionResponseActions(id, token), data, data.Files)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

func (c *Client) SendInteractionFollowup(id, token string, data discordgo.InteractionResponse) error {
	req, err := c.interactionRequest("POST", discordgo.EndpointFollowupMessage(id, token), data, data.Data.Files)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

func (c *Client) EditInteractionFollowup(id, token string, data discordgo.InteractionResponseData) error {
	req, err := c.interactionRequest("PATCH", discordgo.EndpointFollowupMessage(id, token), data, data.Files)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

func (c *Client) interactionRequest(method string, url string, payload any, files []*discordgo.File) (*http.Request, error) {
	if len(files) > 0 {
		return c.requestWithFiles(method, url, payload, files)
	}
	return c.request(method, url, payload)
}
