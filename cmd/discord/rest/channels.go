package rest

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

func (c *Client) CreateDMChannel(ctx context.Context, userID string) (discordgo.Channel, error) {
	data := struct {
		RecipientID string `json:"recipient_id"`
	}{userID}
	req, err := c.request("POST", discordgo.EndpointUserChannels("@me"), data)
	if err != nil {
		return discordgo.Channel{}, err
	}
	var ch discordgo.Channel
	return ch, c.do(ctx, "create_dm_channel", req, &ch)
}
