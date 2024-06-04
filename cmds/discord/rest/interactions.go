package rest

import (
	"github.com/bwmarrin/discordgo"
)

func (c *Client) SendInteractionResponse(id, token string, data discordgo.InteractionResponse) error {
	req, err := c.request("POST", discordgo.EndpointInteractionResponse(id, token), data)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

func (c *Client) UpdateInteractionResponse(id, token string, data discordgo.InteractionResponseData) error {
	req, err := c.request("PATCH", discordgo.EndpointInteractionResponseActions(id, token), data)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

func (c *Client) SendInteractionFollowup(id, token string, data discordgo.InteractionResponse) error {
	req, err := c.request("POST", discordgo.EndpointFollowupMessage(id, token), data)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

func (c *Client) EditInteractionFollowup(id, token string, data discordgo.InteractionResponseData) error {
	req, err := c.request("PATCH", discordgo.EndpointFollowupMessage(id, token), data)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}
