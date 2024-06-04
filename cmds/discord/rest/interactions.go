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

func (c *Client) UpdateInteractionMessage(token string, data discordgo.InteractionResponse) error {
	req, err := c.request("PATCH", discordgo.EndpointInteractionResponseActions(c.applicationID, token), data)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}
