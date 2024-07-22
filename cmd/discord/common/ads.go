package common

import (
	"github.com/bwmarrin/discordgo"
)

func (r Reply) newMessageAd() (discordgo.InteractionResponseData, bool) {

	return discordgo.InteractionResponseData{}, false
}
