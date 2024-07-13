package rest

import (
	"context"
	"errors"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/internal/retry"
)

type File struct {
	Data []byte
	Name string
}

/*
Optimistically send an interaction response update request with fallback to interaction response send request
*/
func (c *Client) UpdateOrSendInteractionResponse(ctx context.Context, appID, interactionID, token string, data discordgo.InteractionResponse, files []File) error {
	res := retry.Retry(func() (struct{}, error) {
		ctx, cancel := context.WithTimeout(ctx, time.Second*3)
		defer cancel()

		err := c.UpdateInteractionResponse(ctx, appID, token, *data.Data, files)
		if err != nil {
			if errors.Is(err, ErrUnknownWebhook) || errors.Is(err, ErrUnknownInteraction) {
				return struct{}{}, c.SendInteractionResponse(ctx, interactionID, token, data, files)
			}
			return struct{}{}, err
		}
		return struct{}{}, nil
	}, 3, time.Millisecond*250)
	return res.Err
}

func (c *Client) SendInteractionResponse(ctx context.Context, interactionID, token string, data discordgo.InteractionResponse, files []File) error {
	req, err := c.requestWithFiles("POST", discordgo.EndpointInteractionResponse(interactionID, token), data, files)
	if err != nil {
		return err
	}
	return c.do(ctx, req, nil)
}

func (c *Client) UpdateInteractionResponse(ctx context.Context, appID, token string, data discordgo.InteractionResponseData, files []File) error {
	req, err := c.requestWithFiles("PATCH", discordgo.EndpointInteractionResponseActions(appID, token), data, files)
	if err != nil {
		return err
	}
	return c.do(ctx, req, nil)
}

func (c *Client) SendInteractionFollowup(ctx context.Context, appID, token string, data discordgo.InteractionResponse, files []File) error {
	req, err := c.requestWithFiles("POST", discordgo.EndpointFollowupMessage(appID, token), data, files)
	if err != nil {
		return err
	}
	return c.do(ctx, req, nil)
}

func (c *Client) EditInteractionFollowup(ctx context.Context, appID, token string, data discordgo.InteractionResponseData, files []File) error {
	req, err := c.requestWithFiles("PATCH", discordgo.EndpointFollowupMessage(appID, token), data, files)
	if err != nil {
		return err
	}
	return c.do(ctx, req, nil)
}
