package rest

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

func (c *Client) CreateDMChannel(ctx context.Context, userID string) (discordgo.Channel, error) {
	req, err := c.request("POST", discordgo.EndpointUserChannels("@me"), nil)
	if err != nil {
		return discordgo.Channel{}, err
	}
	var ch discordgo.Channel
	return ch, c.do(ctx, req, &ch)
}
