package rest

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

func (c *Client) DeleteMessage(ctx context.Context, channelID string, messageID string) error {
	req, err := c.request("DELETE", discordgo.EndpointChannelMessage(channelID, messageID), nil)
	if err != nil {
		return err
	}
	return c.do(ctx, req, nil)
}
