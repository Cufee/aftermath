package alerts

import (
	"context"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/rest"
	"github.com/cufee/aftermath/internal/constants"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type DiscordAlertClient interface {
	Error(ctx context.Context, header, codeBlock string) error
}

type alertClient struct {
	rest       *rest.Client
	webhookURL string
}

func NewClient(token string, webhookURL string) (*alertClient, error) {
	client, err := rest.NewClient(token)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create a rest client")
	}
	return &alertClient{rest: client, webhookURL: webhookURL}, nil
}

func (c *alertClient) Error(ctx context.Context, header, codeBlock string) error {
	content := fmt.Sprintf("### An error was logged through zerolog\n%s\n```%s```", header, codeBlock)
	return c.rest.PostWebhookMessage(ctx, c.webhookURL, discordgo.WebhookParams{
		Content:  content,
		Username: constants.FrontendAppName + " Error Report",
	}, nil)

}

type zerologHook struct {
	client DiscordAlertClient
}

func (h *zerologHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	switch level {
	case zerolog.ErrorLevel:
		h.client.Error(ctx, msg, fmt.Sprintf("%v", e.GetCtx().Err()))
	}
}

func NewHook(c DiscordAlertClient) *zerologHook {
	return &zerologHook{client: c}
}
