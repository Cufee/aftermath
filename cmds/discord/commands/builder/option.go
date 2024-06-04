package builder

import (
	"github.com/bwmarrin/discordgo"
)

type Option struct {
	name        string
	discordType discordgo.ApplicationCommandOptionType

	// options    []commandOption

	// data discordgo.ApplicationCommand
}

func NewOption(name string) Option {
	return Option{}
}
