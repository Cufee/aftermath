package rest

import (
	"context"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

func (c *Client) SendInteractionResponse(ctx context.Context, id, token string, data discordgo.InteractionResponse) error {
	var files []*discordgo.File
	if data.Data != nil {
		files = data.Data.Files
	}
	req, err := c.interactionRequest("POST", discordgo.EndpointInteractionResponse(id, token), data, files)
	if err != nil {
		return err
	}
	return c.do(ctx, req, nil)
}

func (c *Client) UpdateInteractionResponse(ctx context.Context, id, token string, data discordgo.InteractionResponseData) error {
	req, err := c.interactionRequest("PATCH", discordgo.EndpointInteractionResponseActions(id, token), data, data.Files)
	if err != nil {
		return err
	}
	return c.do(ctx, req, nil)
}

func (c *Client) SendInteractionFollowup(ctx context.Context, id, token string, data discordgo.InteractionResponse) error {
	req, err := c.interactionRequest("POST", discordgo.EndpointFollowupMessage(id, token), data, data.Data.Files)
	if err != nil {
		return err
	}
	return c.do(ctx, req, nil)
}

func (c *Client) EditInteractionFollowup(ctx context.Context, id, token string, data discordgo.InteractionResponseData) error {
	req, err := c.interactionRequest("PATCH", discordgo.EndpointFollowupMessage(id, token), data, data.Files)
	if err != nil {
		return err
	}
	return c.do(ctx, req, nil)
}

func (c *Client) interactionRequest(method string, url string, payload any, files []*discordgo.File) (*http.Request, error) {
	if len(files) > 0 {
		return c.requestWithFiles(method, url, payload, files)
	}
	return c.request(method, url, payload)
}
