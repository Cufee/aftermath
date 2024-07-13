package common

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/rest"
)

type reply struct {
	ctx *Context

	text       []string
	files      []rest.File
	components []discordgo.MessageComponent
	embeds     []*discordgo.MessageEmbed
}

func (r reply) Choices(data ...*discordgo.ApplicationCommandOptionChoice) error {
	return r.ctx.respond(discordgo.InteractionResponse{Type: discordgo.InteractionApplicationCommandAutocompleteResult, Data: &discordgo.InteractionResponseData{Choices: data}}, nil)
}

func (r reply) Text(message ...string) reply {
	r.text = append(r.text, message...)
	return r
}

func (r reply) Format(format string, args ...any) reply {
	r.text = append(r.text, fmt.Sprintf(r.ctx.Localize(format), args...))
	return r
}

func (r reply) File(data []byte, name string) reply {
	if data == nil {
		return r
	}
	r.files = append(r.files, rest.File{Data: data, Name: name})
	return r
}

func (r reply) Component(components ...discordgo.MessageComponent) reply {
	for _, c := range components {
		if c == nil {
			continue
		}
		r.components = append(r.components, c)
	}
	return r
}

func (r reply) Embed(embeds ...*discordgo.MessageEmbed) reply {
	for _, e := range embeds {
		if e == nil {
			continue
		}
		r.embeds = append(r.embeds, e)
	}
	return r
}

func (r reply) Send(content ...string) error {
	r.text = append(r.text, content...)
	return r.ctx.respond(r.data())
}

func (r reply) data() (discordgo.InteractionResponse, []rest.File) {
	var content []string
	for _, t := range r.text {
		content = append(content, r.ctx.Localize(t))
	}

	return discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content:    strings.Join(content, "\n"),
			Components: r.components,
			Embeds:     r.embeds,
		},
	}, r.files
}
