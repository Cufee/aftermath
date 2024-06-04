package builder

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmds/discord/common"
	"github.com/cufee/aftermath/cmds/discord/middleware"
)

type CommandHandler func(*common.Context) error

type Command struct {
	discordgo.ApplicationCommand

	Match      func(string) bool
	Handler    CommandHandler
	Middleware []middleware.MiddlewareFunc
	Ephemeral  bool
}

type Builder struct {
	name   string
	params parameters

	match      func(string) bool
	handler    CommandHandler
	middleware []middleware.MiddlewareFunc

	ephemeral bool
	options   []Option
}

func NewCommand(name string) Builder {
	return Builder{name: name, match: func(s string) bool { return s == name }}
}

func (c Builder) Build() Command {
	if c.handler == nil {
		panic("command " + c.name + " is missing a handler")
	}
	if c.match == nil {
		panic("command " + c.name + " is missing a match function")
	}

	nameLocalized := common.LocalizeKey(c.nameKey())
	descLocalized := common.LocalizeKey(c.descKey())

	return Command{
		discordgo.ApplicationCommand{
			Name:                     stringOr(nameLocalized[discordgo.EnglishUS], c.name),
			Description:              stringOr(descLocalized[discordgo.EnglishUS], c.name),
			NameLocalizations:        &nameLocalized,
			DescriptionLocalizations: &descLocalized,
			Type:                     discordgo.ChatApplicationCommand,
		},
		c.match,
		c.handler,
		c.middleware,
		c.ephemeral,
	}
}

func (c Builder) Option(o Option) Builder {
	c.options = append(c.options, o)
	return c
}

func (c Builder) IsEphemeral() Builder {
	c.ephemeral = true
	return c
}

func (c Builder) Handler(fn CommandHandler) Builder {
	c.handler = fn
	return c
}

func (c Builder) Middleware(mw ...middleware.MiddlewareFunc) Builder {
	c.middleware = append(c.middleware, mw...)
	return c
}

func (c Builder) nameKey() string {
	return stringOr(c.params.descKey, fmt.Sprintf("command_%s_name", c.name))
}

func (c Builder) descKey() string {
	return stringOr(c.params.descKey, fmt.Sprintf("command_%s_description", c.name))
}
