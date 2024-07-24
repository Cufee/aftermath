package common

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/emoji"
)

func ButtonInviteAftermath(label string) discordgo.Button {
	return discordgo.Button{
		Style: discordgo.LinkButton,
		Label: label,
		Emoji: emoji.AftermathLogoDefault(),
		URL:   "https://amth.one/invite",
	}
}
func ButtonJoinPrimaryGuild(label string) discordgo.Button {
	return discordgo.Button{
		Style: discordgo.LinkButton,
		Label: label,
		Emoji: emoji.AftermathLogoColored("yellow"),
		URL:   "https://amth.one/join",
	}
}
