package builder

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmds/discord/common"
	"github.com/cufee/aftermath/cmds/discord/middleware"
)

type CommandHandler func(*common.Context) error

type Command struct {
	name   string
	params parameters

	match      func(string) bool
	handler    CommandHandler
	middleware []middleware.MiddlewareFunc

	options []Option
}

func NewCommand(name string) Command {
	return Command{name: name, match: func(s string) bool { return s == name }}
}

func (c Command) Build() (discordgo.ApplicationCommand, func(s string) bool, CommandHandler) {
	nameLocalized := common.LocalizeKey(c.nameKey())
	descLocalized := common.LocalizeKey(c.descKey())

	return discordgo.ApplicationCommand{
			Name:                     stringOr(nameLocalized[discordgo.EnglishUS], c.name),
			Description:              stringOr(descLocalized[discordgo.EnglishUS], c.name),
			NameLocalizations:        &nameLocalized,
			DescriptionLocalizations: &descLocalized,
			Type:                     discordgo.ChatApplicationCommand,
		},
		c.match,
		c.handler
}

func (c Command) Option(o Option) Command {
	c.options = append(c.options, o)
	return c
}

func (c Command) Handler(fn CommandHandler) Command {
	c.handler = fn
	return c
}

func (c Command) Middleware(mw ...middleware.MiddlewareFunc) Command {
	c.middleware = append(c.middleware, mw...)
	return c
}

func (c Command) nameKey() string {
	return stringOr(c.params.descKey, fmt.Sprintf("command_%s_name", c.name))
}

func (c Command) descKey() string {
	return stringOr(c.params.descKey, fmt.Sprintf("command_%s_description", c.name))
}
