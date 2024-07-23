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

	internal replyInternal
}

type replyInternal struct {
	IncludeAds bool
	Hint       string
	Text       []string
	Files      []rest.File
	Components []discordgo.MessageComponent
	Embeds     []*discordgo.MessageEmbed
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

func (r Reply) Peek() replyInternal {
	return r.internal
}

func (r Reply) Choices(data ...*discordgo.ApplicationCommandOptionChoice) error {
	_, err := r.ctx.InteractionResponse(discordgo.InteractionResponseData{Choices: data}, nil)
	return err
}

func (r Reply) Hint(text string) Reply {
	r.internal.Hint = text
	return r
}

func (r Reply) Text(message ...string) Reply {
	r.internal.Text = append(r.internal.Text, message...)
	return r
}

func (r Reply) Format(format string, args ...any) Reply {
	r.internal.Text = append(r.internal.Text, fmt.Sprintf(r.ctx.Localize(format), args...))
	return r
}

func (r Reply) File(data []byte, name string) Reply {
	if data == nil {
		return r
	}
	r.internal.Files = append(r.internal.Files, rest.File{Data: data, Name: name})
	return r
}

func (r Reply) Component(components ...discordgo.MessageComponent) Reply {
	if len(components) == 1 && components[0] == nil {
		r.internal.Components = make([]discordgo.MessageComponent, 0)
		return r
	}

	r.internal.Components = append(r.internal.Components, components...)
	return r
}

func (r Reply) Embed(embeds ...*discordgo.MessageEmbed) Reply {
	if len(embeds) == 1 && embeds[0] == nil {
		r.internal.Embeds = make([]*discordgo.MessageEmbed, 0)
		return r
	}

	r.internal.Embeds = append(r.internal.Embeds, embeds...)
	return r
}

func (r Reply) WithAds() Reply {
	r.internal.IncludeAds = true
	return r
}

func (r Reply) Send(content ...string) error {
	_, err := r.Message(content...)
	return err
}

func (r Reply) Message(content ...string) (discordgo.Message, error) {
	if r.internal.IncludeAds {
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

	r.internal.Text = append(r.internal.Text, content...)
	return r.ctx.InteractionResponse(r.internal.Data(r.ctx.Localize))
}

func (r replyInternal) Data(localize func(string) string) (discordgo.InteractionResponseData, []rest.File) {
	var content []string
	for _, t := range r.Text {
		content = append(content, localize(t))
	}
	if r.Hint != "" {
		content = append(content, "-# "+localize(r.Hint))
	}

	var clearAttachments []*discordgo.MessageAttachment
	return discordgo.InteractionResponseData{
		Content:     strings.Join(content, "\n"),
		Components:  r.Components,
		Embeds:      r.Embeds,
		Attachments: &clearAttachments,
	}, r.Files
}
