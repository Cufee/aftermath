package builder

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmds/discord/common"
)

type OptionChoice struct {
	name  string
	value any

	params parameters
}

func NewChoice(name string, value any) OptionChoice {
	return OptionChoice{name: name, value: value}
}

func (c OptionChoice) Params(params ...Param) OptionChoice {
	for _, apply := range params {
		apply(&c.params)
	}
	return c
}

func (c OptionChoice) Build(command, option string) discordgo.ApplicationCommandOptionChoice {
	if c.value == nil {
		panic("option value cannot be nil")
	}

	nameLocalized := common.LocalizeKey(c.nameKey(command, option))

	return discordgo.ApplicationCommandOptionChoice{
		Name:              stringOr(nameLocalized[discordgo.EnglishUS], c.name),
		NameLocalizations: nameLocalized,
		Value:             c.value,
	}
}

func (c OptionChoice) nameKey(command, option string) string {
	return stringOr(c.params.nameKey, fmt.Sprintf("command_%s_option_%s_choice_%s_name", command, option, c.name))
}
