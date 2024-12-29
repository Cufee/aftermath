package emoji

import (
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/internal/constants"
)

var aftermathLogo = map[string]discordgo.ComponentEmoji{
	"red":    {Name: "red", ID: constants.DiscordEmojiRedID},
	"blue":   {Name: "blue", ID: constants.DiscordEmojiBlueID},
	"yellow": {Name: "yellow", ID: constants.DiscordEmojiYellowID},
}

func AftermathLogoColored(color string) *discordgo.ComponentEmoji {
	e, ok := aftermathLogo[color]
	if !ok {
		e = aftermathLogo["red"]
	}
	return &e
}

func AftermathLogoDefault() *discordgo.ComponentEmoji {
	e, ok := aftermathLogo[os.Getenv("BRAND_FLAVOR")]
	if !ok {
		e = aftermathLogo["red"]
	}
	return &e
}
