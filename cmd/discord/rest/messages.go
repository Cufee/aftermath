package rest

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

func (c *Client) CreateMessage(ctx context.Context, channelID string, data discordgo.Message, files []File) (discordgo.Message, error) {
	req, err := c.requestWithFiles("POST", discordgo.EndpointChannelMessages(channelID), data, files)
	if err != nil {
		return discordgo.Message{}, err
	}
	var m discordgo.Message
	return m, c.do(ctx, req, &m)
}

func (c *Client) UpdateMessage(ctx context.Context, channelID string, messageID string, data discordgo.Message, files []File) (discordgo.Message, error) {
	req, err := c.requestWithFiles("PATCH", discordgo.EndpointChannelMessage(channelID, messageID), data, files)
	if err != nil {
		return discordgo.Message{}, err
	}
	var m discordgo.Message
	return m, c.do(ctx, req, &m)
}

func (c *Client) DeleteMessage(ctx context.Context, channelID string, messageID string) error {
	req, err := c.request("DELETE", discordgo.EndpointChannelMessage(channelID, messageID), nil)
	if err != nil {
		return err
	}
	return c.do(ctx, req, nil)
}
