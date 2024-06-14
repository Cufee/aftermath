package rest

import "github.com/bwmarrin/discordgo"

func (c *Client) OverwriteGlobalApplicationCommands(data []discordgo.ApplicationCommand) error {
	req, err := c.request("PUT", discordgo.EndpointApplicationGlobalCommands(c.applicationID), data)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

func (c *Client) CreateGlobalApplicationCommand(data discordgo.ApplicationCommand) (discordgo.ApplicationCommand, error) {
	req, err := c.request("POST", discordgo.EndpointApplicationGlobalCommands(c.applicationID), data)
	if err != nil {
		return discordgo.ApplicationCommand{}, err
	}
	var command discordgo.ApplicationCommand
	return command, c.do(req, &command)
}

func (c *Client) GetGlobalApplicationCommands() ([]discordgo.ApplicationCommand, error) {
	req, err := c.request("GET", discordgo.EndpointApplicationGlobalCommands(c.applicationID), nil)
	if err != nil {
		return nil, err
	}

	var data []discordgo.ApplicationCommand
	return data, c.do(req, &data)
}

func (c *Client) UpdateGlobalApplicationCommand(id string, data discordgo.ApplicationCommand) (discordgo.ApplicationCommand, error) {
	req, err := c.request("PATCH", discordgo.EndpointApplicationGlobalCommand(c.applicationID, id), data)
	if err != nil {
		return discordgo.ApplicationCommand{}, err
	}
	var command discordgo.ApplicationCommand
	return command, c.do(req, &command)
}

func (c *Client) DeleteGlobalApplicationCommand(id string) error {
	req, err := c.request("DELETE", discordgo.EndpointApplicationGlobalCommand(c.applicationID, id), nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}
