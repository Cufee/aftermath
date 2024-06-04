package commands

// import (
// 	c "context"
// 	"io"

// 	"github.com/cufee/aftermath/cmds/core"
// 	"github.com/cufee/aftermath/cmds/discord/common"
// 	"github.com/cufee/aftermath/cmds/discord/middleware"
// 	"github.com/disgoorg/disgo/discord"
// 	"github.com/disgoorg/disgo/events"
// )

// type commandInit []commandBuilder

// func (c *commandInit) add(cmd commandBuilder) {
// 	*c = append(*c, cmd)
// }

// func (c *commandInit) Build() []Command {
// 	var commands []Command
// 	for _, builder := range *c {
// 		commands = append(commands, builder.Build())
// 	}
// 	return commands
// }

// var Loaded commandInit

// type Handler func(*ctx) error

// type Command struct {
// 	discord.ApplicationCommandCreate
// 	Handler
// 	Middleware []middleware.MiddlewareFunc
// }

// type ctx struct {
// 	*common.Context
// 	event *events.ApplicationCommandInteractionCreate
// }

// func ContextFrom(c c.Context, client core.Client, event *events.ApplicationCommandInteractionCreate) *ctx {
// 	return &ctx{
// 		common.NewContext(c, client),
// 		event,
// 	}
// }

// func (c *ctx) Options() discord.SlashCommandInteractionData {
// 	return c.event.SlashCommandInteractionData()
// }

// func (c *ctx) Reply(msg discord.MessageCreate) error {
// 	return c.event.CreateMessage(msg)
// }

// func (c *ctx) File(reader io.Reader, name string) error {
// 	f := []*discord.File{{Name: name, Reader: reader}}
// 	return c.event.CreateMessage(discord.MessageCreate{Files: f})
// }
