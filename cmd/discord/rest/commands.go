package rest

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/internal/discord"
)

func (c *Client) OverwriteGlobalApplicationCommands(ctx context.Context, data []discord.ApplicationCommand) error {
	req, err := c.request("PUT", discordgo.EndpointApplicationGlobalCommands(c.applicationID), data)
	if err != nil {
		return err
	}
	return c.do(ctx, req, nil)
}

func (c *Client) CreateGlobalApplicationCommand(ctx context.Context, data discord.ApplicationCommand) (discord.ApplicationCommand, error) {
	req, err := c.request("POST", discordgo.EndpointApplicationGlobalCommands(c.applicationID), data)
	if err != nil {
		return discord.ApplicationCommand{}, err
	}
	var command discord.ApplicationCommand
	return command, c.do(ctx, req, &command)
}

func (c *Client) GetGlobalApplicationCommands(ctx context.Context) ([]discord.ApplicationCommand, error) {
	req, err := c.request("GET", discordgo.EndpointApplicationGlobalCommands(c.applicationID), nil)
	if err != nil {
		return nil, err
	}

	var data []discord.ApplicationCommand
	return data, c.do(ctx, req, &data)
}

func (c *Client) UpdateGlobalApplicationCommand(ctx context.Context, id string, data discord.ApplicationCommand) (discord.ApplicationCommand, error) {
	req, err := c.request("PATCH", discordgo.EndpointApplicationGlobalCommand(c.applicationID, id), data)
	if err != nil {
		return discord.ApplicationCommand{}, err
	}
	var command discord.ApplicationCommand
	return command, c.do(ctx, req, &command)
}

func (c *Client) DeleteGlobalApplicationCommand(ctx context.Context, id string) error {
	req, err := c.request("DELETE", discordgo.EndpointApplicationGlobalCommand(c.applicationID, id), nil)
	if err != nil {
		return err
	}
	return c.do(ctx, req, nil)
}
