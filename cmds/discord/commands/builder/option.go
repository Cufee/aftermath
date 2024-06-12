package builder

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmds/discord/common"
)

type Option struct {
	name     string
	required bool

	params parameters
	kind   discordgo.ApplicationCommandOptionType

	// Minimal value of number/integer option.
	minValue *float64
	// Maximum value of number/integer option.
	maxValue float64
	// Minimum length of string option.
	minLength *int
	// Maximum length of string option.
	maxLength int

	choices []OptionChoice

	options []Option
}

func NewOption(name string, kind discordgo.ApplicationCommandOptionType) Option {
	return Option{name: name, kind: kind}
}

func (o Option) Required() Option {
	o.required = true
	return o
}

func (o Option) Min(value float64) Option {
	if o.kind == discordgo.ApplicationCommandOptionString {
		i := int(value)
		o.minLength = &i
		return o
	}
	if o.kind == discordgo.ApplicationCommandOptionInteger || o.kind == discordgo.ApplicationCommandOptionNumber {
		o.minValue = &value
		return o
	}
	panic("invalid .Min call on option of type " + o.kind.String())
}

func (o Option) Max(value float64) Option {
	if o.kind == discordgo.ApplicationCommandOptionString {
		i := int(value)
		o.maxLength = i
		return o
	}
	if o.kind == discordgo.ApplicationCommandOptionInteger || o.kind == discordgo.ApplicationCommandOptionNumber {
		o.maxValue = value
		return o
	}
	panic("invalid .Max call on option of type " + o.kind.String())
}

func (o Option) Params(params ...Param) Option {
	for _, apply := range params {
		apply(&o.params)
	}
	return o
}

func (o Option) Options(options ...Option) Option {
	o.options = append(o.options, options...)
	return o
}

func (o Option) Choices(choices ...OptionChoice) Option {
	o.choices = append(o.choices, choices...)
	return o
}

func (o Option) Build(command string) discordgo.ApplicationCommandOption {
	if o.kind == 0 {
		panic("option type is not set")
	}

	nameLocalized := common.LocalizeKey(o.nameKey(command))
	descLocalized := common.LocalizeKey(o.descKey(command))

	var choices []*discordgo.ApplicationCommandOptionChoice
	for _, choice := range o.choices {
		c := choice.Build(command, o.name)
		choices = append(choices, &c)
	}

	return discordgo.ApplicationCommandOption{
		Name:                     strings.ToLower(stringOr(nameLocalized[discordgo.EnglishUS], o.name)),
		Description:              stringOr(descLocalized[discordgo.EnglishUS], o.name),
		NameLocalizations:        nameLocalized,
		DescriptionLocalizations: descLocalized,
		MinLength:                o.minLength,
		MaxLength:                o.maxLength,
		MinValue:                 o.minValue,
		MaxValue:                 o.maxValue,
		Required:                 o.required,
		Choices:                  choices,
		Type:                     o.kind,
	}
}

func (o Option) nameKey(command string) string {
	return stringOr(o.params.nameKey, fmt.Sprintf("command_%s_option_%s_name", command, o.name))
}

func (o Option) descKey(command string) string {
	return stringOr(o.params.descKey, fmt.Sprintf("command_%s_option_%s_description", command, o.name))
}
