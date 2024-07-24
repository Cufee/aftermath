package rest

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

type File struct {
	Data []byte
	Name string
}

func (c *Client) SendInteractionResponse(ctx context.Context, interactionID, token string, data discordgo.InteractionResponse, files []File) (discordgo.Message, error) {
	req, err := c.requestWithFiles("POST", discordgo.EndpointInteractionResponse(interactionID, token), data, files)
	if err != nil {
		return discordgo.Message{}, err
	}
	var m discordgo.Message
	return m, c.do(ctx, req, &m)
}

func (c *Client) UpdateInteractionResponse(ctx context.Context, appID, token string, data discordgo.InteractionResponseData, files []File) (discordgo.Message, error) {
	req, err := c.requestWithFiles("PATCH", discordgo.EndpointInteractionResponseActions(appID, token), data, files)
	if err != nil {
		return discordgo.Message{}, err
	}
	var m discordgo.Message
	return m, c.do(ctx, req, &m)
}

func (c *Client) DeleteInteractionResponse(ctx context.Context, appID, token string) error {
	req, err := c.request("DELETE", discordgo.EndpointWebhookMessage(appID, token, "@original"), nil)
	if err != nil {
		return err
	}
	return c.do(ctx, req, nil)
}

func (c *Client) SendInteractionFollowup(ctx context.Context, appID, token string, data discordgo.InteractionResponseData, files []File) (discordgo.Message, error) {
	req, err := c.requestWithFiles("POST", discordgo.EndpointFollowupMessage(appID, token), data, files)
	if err != nil {
		return discordgo.Message{}, err
	}
	var m discordgo.Message
	return m, c.do(ctx, req, &m)
}

func (c *Client) EditInteractionFollowup(ctx context.Context, appID, token string, data discordgo.InteractionResponseData, files []File) (discordgo.Message, error) {
	req, err := c.requestWithFiles("PATCH", discordgo.EndpointFollowupMessage(appID, token), data, files)
	if err != nil {
		return discordgo.Message{}, err
	}
	var m discordgo.Message
	return m, c.do(ctx, req, &m)
}
