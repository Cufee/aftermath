package rest

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

func (c *Client) PostWebhookMessage(ctx context.Context, webhookURL string, data discordgo.WebhookParams, files []File) error {
	req, err := c.requestWithFiles("POST", webhookURL, data, files)
	if err != nil {
		return err
	}
	return c.do(ctx, req, nil)
}
