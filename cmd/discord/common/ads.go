package common

import (
	"github.com/bwmarrin/discordgo"
)

func (r reply) newMessageAd() (discordgo.InteractionResponseData, bool) {

	return discordgo.InteractionResponseData{}, false
}
