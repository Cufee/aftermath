package rest

import (
	"context"
	"net/url"

	"github.com/bwmarrin/discordgo"
)

func (c *Client) CreateMessage(ctx context.Context, channelID string, data discordgo.MessageSend, files []File) (discordgo.Message, error) {
	req, err := c.requestWithFiles("POST", discordgo.EndpointChannelMessages(channelID), data, files)
	if err != nil {
		return discordgo.Message{}, err
	}
	var m discordgo.Message
	return m, c.do(ctx, "create_message", req, &m)
}

func (c *Client) GetMessage(ctx context.Context, channelID string, messageID string) (discordgo.Message, error) {
	req, err := c.request("GET", discordgo.EndpointChannelMessage(channelID, messageID), nil)
	if err != nil {
		return discordgo.Message{}, err
	}
	var m discordgo.Message
	return m, c.do(ctx, "get_message", req, &m)
}

func (c *Client) UpdateMessage(ctx context.Context, channelID string, messageID string, data discordgo.MessageEdit, files []File) (discordgo.Message, error) {
	req, err := c.requestWithFiles("PATCH", discordgo.EndpointChannelMessage(channelID, messageID), data, files)
	if err != nil {
		return discordgo.Message{}, err
	}
	var m discordgo.Message
	return m, c.do(ctx, "update_message", req, &m)
}

func (c *Client) DeleteMessage(ctx context.Context, channelID string, messageID string) error {
	req, err := c.request("DELETE", discordgo.EndpointChannelMessage(channelID, messageID), nil)
	if err != nil {
		return err
	}
	return c.do(ctx, "delete_message", req, nil)
}

func (c *Client) CreateMessageReaction(ctx context.Context, channelID string, messageID string, emojiID string) error {
	req, err := c.request("PUT", discordgo.EndpointMessageReaction(channelID, messageID, url.QueryEscape(emojiID), "@me"), nil)
	if err != nil {
		return err
	}
	return c.do(ctx, "create_message_reaction", req, nil)
}

func (c *Client) DeleteOwnMessageReaction(ctx context.Context, channelID string, messageID string, emojiID string) error {
	req, err := c.request("DELETE", discordgo.EndpointMessageReaction(channelID, messageID, url.QueryEscape(emojiID), "@me"), nil)
	if err != nil {
		return err
	}
	return c.do(ctx, "delete_own_message_reaction", req, nil)
}
