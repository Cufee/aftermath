package builder

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/middleware"
)

type commandType byte

const (
	CommandTypeChat      commandType = iota
	CommandTypeComponent commandType = iota
)

type CommandHandler func(common.Context) error

type Command struct {
	discordgo.ApplicationCommand

	Type       commandType
	Match      func(string) bool
	Handler    CommandHandler
	Middleware []middleware.MiddlewareFunc
	Ephemeral  bool
}

type Builder struct {
	name   string
	kind   commandType
	params parameters

	guildOnly bool

	match      func(string) bool
	handler    CommandHandler
	middleware []middleware.MiddlewareFunc

	ephemeral bool
	options   []Option
}

func NewCommand(name string) Builder {
	return Builder{name: name, match: func(s string) bool { return s == name }, kind: CommandTypeChat}
}

func (c Builder) Build() Command {
	if c.handler == nil {
		panic("command " + c.name + " is missing a handler")
	}
	if c.match == nil {
		panic("command " + c.name + " is missing a match function")
	}

	var nameLocalized map[discordgo.Locale]string
	if c.nameKey() != "" {
		nameLocalized = common.LocalizeKey(c.nameKey())
	}

	var descLocalized map[discordgo.Locale]string
	if c.descKey() != "" {
		descLocalized = common.LocalizeKey(c.descKey())
	}

	var options []*discordgo.ApplicationCommandOption
	for _, o := range c.options {
		built := o.Build(c.name)
		options = append(options, &built)
	}

	return Command{
		discordgo.ApplicationCommand{
			Name:                     strings.ToLower(stringOr(nameLocalized[discordgo.EnglishUS], c.name)),
			Description:              stringOr(descLocalized[discordgo.EnglishUS], c.name),
			NameLocalizations:        &nameLocalized,
			DescriptionLocalizations: &descLocalized,
			Options:                  options,
			DMPermission:             common.Pointer(!c.guildOnly),
			Type:                     discordgo.ChatApplicationCommand,
		},
		c.kind,
		c.match,
		c.handler,
		c.middleware,
		c.ephemeral,
	}
}

func (c Builder) Options(o ...Option) Builder {
	c.options = append(c.options, o...)
	return c
}

func (c Builder) Params(params ...Param) Builder {
	for _, apply := range params {
		apply(&c.params)
	}
	return c
}

func (c Builder) Ephemeral() Builder {
	c.ephemeral = true
	return c
}

func (c Builder) GuildOnly() Builder {
	c.guildOnly = true
	return c
}

func (c Builder) ComponentType(matchFn func(string) bool) Builder {
	c.kind = CommandTypeComponent
	c.match = matchFn
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
	return stringOr(c.params.nameKey, fmt.Sprintf("command_%s_name", c.name))
}

func (c Builder) descKey() string {
	return stringOr(c.params.descKey, fmt.Sprintf("command_%s_description", c.name))
}
