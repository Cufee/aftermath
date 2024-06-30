package rest

import (
	"bytes"
	"context"
	"errors"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

type File struct {
	Data []byte
	Name string
}

/*
Optimistically send an interaction response update request with fallback to interaction response send request
*/
func (c *Client) UpdateOrSendInteractionResponse(ctx context.Context, appID, interactionID, token string, data discordgo.InteractionResponse, files []File) error {
	err := c.UpdateInteractionResponse(ctx, appID, token, *data.Data, files)
	if err != nil {
		if errors.Is(err, ErrUnknownWebhook) {
			return c.SendInteractionResponse(ctx, interactionID, token, data, files)
		}
		return err
	}
	return nil
}

func (c *Client) SendInteractionResponse(ctx context.Context, interactionID, token string, data discordgo.InteractionResponse, files []File) error {
	req, err := c.interactionRequest("POST", discordgo.EndpointInteractionResponse(interactionID, token), data, files)
	if err != nil {
		return err
	}
	return c.do(ctx, req, nil)
}

func (c *Client) UpdateInteractionResponse(ctx context.Context, appID, token string, data discordgo.InteractionResponseData, files []File) error {
	req, err := c.interactionRequest("PATCH", discordgo.EndpointInteractionResponseActions(appID, token), data, files)
	if err != nil {
		return err
	}
	return c.do(ctx, req, nil)
}

func (c *Client) SendInteractionFollowup(ctx context.Context, appID, token string, data discordgo.InteractionResponse, files []File) error {
	req, err := c.interactionRequest("POST", discordgo.EndpointFollowupMessage(appID, token), data, files)
	if err != nil {
		return err
	}
	return c.do(ctx, req, nil)
}

func (c *Client) EditInteractionFollowup(ctx context.Context, appID, token string, data discordgo.InteractionResponseData, files []File) error {
	req, err := c.interactionRequest("PATCH", discordgo.EndpointFollowupMessage(appID, token), data, files)
	if err != nil {
		return err
	}
	return c.do(ctx, req, nil)
}

func (c *Client) interactionRequest(method string, url string, payload any, files []File) (*http.Request, error) {
	if len(files) > 0 {
		var df []*discordgo.File
		for _, f := range files {
			df = append(df, &discordgo.File{Name: f.Name, Reader: bytes.NewReader(f.Data)})
		}
		return c.requestWithFiles(method, url, payload, df)
	}
	return c.request(method, url, payload)
}
