package emoji

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

var aftermathLogo = map[string]discordgo.ComponentEmoji{
	"red":    {Name: "aftermath_red", ID: "1315494620167409744"},
	"blue":   {Name: "aftermath_blue", ID: "1315494694649598023"},
	"yellow": {Name: "aftermath_yellow", ID: "1315494518140833865"},
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
