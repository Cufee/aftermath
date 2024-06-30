package rest

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

func (c *Client) OverwriteGlobalApplicationCommands(ctx context.Context, data []discordgo.ApplicationCommand) error {
	req, err := c.request("PUT", discordgo.EndpointApplicationGlobalCommands(c.applicationID), data)
	if err != nil {
		return err
	}
	return c.do(ctx, req, nil)
}

func (c *Client) CreateGlobalApplicationCommand(ctx context.Context, data discordgo.ApplicationCommand) (discordgo.ApplicationCommand, error) {
	req, err := c.request("POST", discordgo.EndpointApplicationGlobalCommands(c.applicationID), data)
	if err != nil {
		return discordgo.ApplicationCommand{}, err
	}
	var command discordgo.ApplicationCommand
	return command, c.do(ctx, req, &command)
}

func (c *Client) GetGlobalApplicationCommands(ctx context.Context) ([]discordgo.ApplicationCommand, error) {
	req, err := c.request("GET", discordgo.EndpointApplicationGlobalCommands(c.applicationID), nil)
	if err != nil {
		return nil, err
	}

	var data []discordgo.ApplicationCommand
	return data, c.do(ctx, req, &data)
}

func (c *Client) UpdateGlobalApplicationCommand(ctx context.Context, id string, data discordgo.ApplicationCommand) (discordgo.ApplicationCommand, error) {
	req, err := c.request("PATCH", discordgo.EndpointApplicationGlobalCommand(c.applicationID, id), data)
	if err != nil {
		return discordgo.ApplicationCommand{}, err
	}
	var command discordgo.ApplicationCommand
	return command, c.do(ctx, req, &command)
}

func (c *Client) DeleteGlobalApplicationCommand(ctx context.Context, id string) error {
	req, err := c.request("DELETE", discordgo.EndpointApplicationGlobalCommand(c.applicationID, id), nil)
	if err != nil {
		return err
	}
	return c.do(ctx, req, nil)
}
