package common

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/cufee/aftermath/cmd/discord/rest"
	"github.com/cufee/aftermath/internal/retry"
)

type Reply struct {
	ctx Context

	internal replyInternal
}

type replyInternal struct {
	Hint       string
	Text       []string
	Files      []rest.File
	Embeds     []*discordgo.MessageEmbed
	Components []discordgo.MessageComponent
	Choices    []*discordgo.ApplicationCommandOptionChoice

	Reference *discordgo.MessageReference

	eventMetadata map[string]any
}

type ResponseData discordgo.InteractionResponseData

func (d ResponseData) Interaction() discordgo.InteractionResponseData {
	return discordgo.InteractionResponseData(d)
}

func (d ResponseData) Message() discordgo.MessageSend {
	return discordgo.MessageSend{
		Content:         d.Content,
		Embeds:          d.Embeds,
		TTS:             d.TTS,
		Components:      d.Components,
		AllowedMentions: d.AllowedMentions,
		Flags:           d.Flags,
	}
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

func (r Reply) Choices(data ...*discordgo.ApplicationCommandOptionChoice) Reply {
	r.internal.Choices = append(r.internal.Choices, data...)
	return r
}

func (r Reply) Metadata() map[string]any {
	if r.internal.eventMetadata != nil {
		return r.internal.eventMetadata
	}
	return make(map[string]any)
}

func (r Reply) WithMeta(data map[string]any) Reply {
	meta := r.Metadata()
	for key, value := range data {
		meta[key] = value
	}
	return r
}

func (r Reply) Reference(messageID, channelID, guildID string) Reply {
	r.internal.Reference = &discordgo.MessageReference{
		MessageID: messageID,
		ChannelID: channelID,
		GuildID:   guildID,
	}
	return r
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
		r.internal.Files = make([]rest.File, 0)
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
	r.internal.Embeds = append(r.internal.Embeds, embeds...)
	return r
}

func (r Reply) WithAds() Reply {
	r.ctx.WithAds(true)
	return r
}

func (r Reply) Send(content ...string) error {
	_, err := r.Message(content...)
	return err
}

func (r Reply) IsError(e Error) Reply {
	r.ctx.SetError(e)
	return r
}

func (r Reply) Message(content ...string) (discordgo.Message, error) {
	r.internal.Text = append(r.internal.Text, content...)
	return r.ctx.InteractionResponse(r)
}

func (r Reply) Followup(ctx context.Context, content ...string) (discordgo.Message, error) {
	r.internal.Text = append(r.internal.Text, content...)
	return r.ctx.InteractionFollowUp(ctx, r)
}

func (r replyInternal) Build(localize func(string) string) (ResponseData, []rest.File) {
	var content []string
	for _, t := range r.Text {
		content = append(content, localize(t))
	}
	if r.Hint != "" {
		content = append(content, "-# "+localize(r.Hint))
	}

	data := discordgo.InteractionResponseData{
		Content:    strings.Join(content, "\n"),
		Components: r.Components,
		Choices:    r.Choices,
		Embeds:     r.Embeds,
	}
	if r.Files != nil {
		data.Attachments = &[]*discordgo.MessageAttachment{} // clear existing attachments
	}
	return ResponseData(data), r.Files
}
