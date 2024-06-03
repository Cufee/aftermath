package commands

import (
	"fmt"

	"github.com/cufee/aftermath/cmds/discord/common"
	"github.com/cufee/aftermath/cmds/discord/middleware"
	"github.com/disgoorg/disgo/discord"
)

func stringOr(options ...string) string {
	for _, value := range options {
		if value != "" {
			return value
		}
	}
	return ""
}

type commandBuilder struct {
	name string

	options    []commandOption
	handler    Handler
	middleware []middleware.MiddlewareFunc

	params parameters
}

type commandOption struct {
	Name                     string
	NameLocalizations        map[discord.Locale]string
	Description              string
	DescriptionLocalizations map[discord.Locale]string

	discord.ApplicationCommandOption

	params parameters
}

type parameters struct {
	nameKey        string
	descriptionKey string
}

type Param func(p *parameters)

func SetNameKey(key string) Param {
	return func(p *parameters) {
		p.nameKey = key
	}
}

func SetDescKey(key string) Param {
	return func(p *parameters) {
		p.descriptionKey = key
	}
}

func cmd(params ...Param) commandBuilder {
	p := parameters{}
	for _, apply := range params {
		apply(&p)
	}
	return commandBuilder{params: p}
}

func (b commandBuilder) Name(name string) commandBuilder {
	b.name = name
	b.params.nameKey = stringOr(b.params.nameKey, fmt.Sprintf("command_%s_name", b.name))
	b.params.descriptionKey = stringOr(b.params.descriptionKey, fmt.Sprintf("command_%s_description", b.name))
	return b
}

func (b commandBuilder) Middleware(mw ...middleware.MiddlewareFunc) commandBuilder {
	b.middleware = append(b.middleware, mw...)
	return b
}

func (b commandBuilder) Option(name string, data discord.ApplicationCommandOption, params ...Param) commandBuilder {
	p := parameters{
		nameKey:        fmt.Sprintf("command_%s_option_%s_name", b.name, name),
		descriptionKey: fmt.Sprintf("command_%s_option_%s_description", b.name, name),
	}
	for _, apply := range params {
		apply(&p)
	}

	b.options = append(b.options, commandOption{
		name,
		nil,
		name,
		nil,
		data,
		p,
	})
	return b
}

func (b commandBuilder) Handler(f func(ctx *ctx) error) commandBuilder {
	b.handler = f
	return b
}

func (b commandBuilder) Build() Command {
	nameLocalized := common.LocalizeKey(b.params.nameKey)
	descriptionLocalized := common.LocalizeKey(b.params.descriptionKey)

	var options []discord.ApplicationCommandOption
	for _, opt := range b.options {
		opt.NameLocalizations = common.LocalizeKey(opt.params.nameKey)
		opt.Name = stringOr(opt.NameLocalizations[discord.LocaleEnglishUS], opt.Name)

		opt.DescriptionLocalizations = common.LocalizeKey(opt.params.descriptionKey)
		opt.Name = stringOr(opt.DescriptionLocalizations[discord.LocaleEnglishUS], opt.Name)

		options = append(options, (discord.ApplicationCommandOption)(opt))
	}

	return Command{
		discord.SlashCommandCreate{
			Name:                     stringOr(nameLocalized[discord.LocaleEnglishUS], b.name),
			NameLocalizations:        nameLocalized,
			Description:              stringOr(descriptionLocalized[discord.LocaleEnglishUS], "/"+b.name),
			DescriptionLocalizations: descriptionLocalized,
			Options:                  options,
		},
		b.handler,
		b.middleware,
	}
}
