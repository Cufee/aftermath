package common

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/cufee/aftermath/cmd/discord/rest"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/retry"
)

type Reply struct {
	ctx Context

	includeAds bool
	hint       string
	text       []string
	files      []rest.File
	components []discordgo.MessageComponent
	embeds     []*discordgo.MessageEmbed
}

func WithRetry(fn func() (discordgo.Message, error), tries ...int) (discordgo.Message, error) {
	var triesCnt = 5
	if len(tries) > 0 && tries[0] > 0 {
		triesCnt = tries[0]
	}
	res := retry.Retry(fn, triesCnt, time.Second)
	return res.Data, res.Err
}

func ContextReply(ctx Context) Reply {
	return Reply{ctx: ctx}
}

func (r Reply) Choices(data ...*discordgo.ApplicationCommandOptionChoice) error {
	_, err := r.ctx.InteractionResponse(discordgo.InteractionResponseData{Choices: data}, nil)
	return err
}

func (r Reply) Hint(text string) Reply {
	r.hint = text
	return r
}

func (r Reply) Text(message ...string) Reply {
	r.text = append(r.text, message...)
	return r
}

func (r Reply) Format(format string, args ...any) Reply {
	r.text = append(r.text, fmt.Sprintf(r.ctx.Localize(format), args...))
	return r
}

func (r Reply) File(data []byte, name string) Reply {
	if data == nil {
		return r
	}
	r.files = append(r.files, rest.File{Data: data, Name: name})
	return r
}

func (r Reply) Component(components ...discordgo.MessageComponent) Reply {
	for _, c := range components {
		if c == nil {
			continue
		}
		r.components = append(r.components, c)
	}
	return r
}

func (r Reply) Embed(embeds ...*discordgo.MessageEmbed) Reply {
	for _, e := range embeds {
		if e == nil {
			continue
		}
		r.embeds = append(r.embeds, e)
	}
	return r
}

func (r Reply) WithAds() Reply {
	r.includeAds = true
	return r
}

func (r Reply) Send(content ...string) error {
	_, err := r.Message(content...)
	return err
}

func (r Reply) Message(content ...string) (discordgo.Message, error) {
	if r.includeAds {
		defer func() {
			data, send := r.newMessageAd()
			if !send {
				return
			}

			_, err := r.ctx.InteractionFollowUp(data, nil)
			if err != nil {
				log.Err(err).Msg("failed to send an interaction ad followup")
			}
		}()
	}

	r.text = append(r.text, content...)
	return r.ctx.InteractionResponse(r.data())
}

func (r Reply) data() (discordgo.InteractionResponseData, []rest.File) {
	var content []string
	for _, t := range r.text {
		content = append(content, r.ctx.Localize(t))
	}
	if r.hint != "" {
		content = append(content, "-# "+r.ctx.Localize(r.hint))
	}

	var cleatAttachments []*discordgo.MessageAttachment
	return discordgo.InteractionResponseData{
		Content:     strings.Join(content, "\n"),
		Components:  r.components,
		Embeds:      r.embeds,
		Attachments: &cleatAttachments,
	}, r.files
}
