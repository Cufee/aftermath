package common

import (
	"fmt"
	"io"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type reply struct {
	ctx *Context

	text       []string
	files      []*discordgo.File
	components []discordgo.MessageComponent
	embeds     []*discordgo.MessageEmbed
}

func (r reply) Text(message ...string) reply {
	r.text = append(r.text, message...)
	return r
}

func (r reply) Fmt(format string, args ...any) reply {
	r.text = append(r.text, fmt.Sprintf(format, args...))
	return r
}

func (r reply) File(reader io.Reader, name string) reply {
	if reader == nil {
		return r
	}
	r.files = append(r.files, &discordgo.File{Reader: reader, Name: name})
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
	return r.ctx.respond(r.data(r.ctx.Localize))
}

func (r reply) data(localePrinter func(string) string) discordgo.InteractionResponseData {
	var content []string
	for _, t := range r.text {
		content = append(content, localePrinter(t))
	}
	return discordgo.InteractionResponseData{
		Content:    strings.Join(content, "\n"),
		Components: r.components,
		Embeds:     r.embeds,
		Files:      r.files,
	}
}
