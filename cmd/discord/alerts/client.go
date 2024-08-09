package alerts

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
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

func (c *alertClient) Error(ctx context.Context, message, codeBlock string) error {
	content := fmt.Sprintf("%s\n```%s```", message, codeBlock)
	if len(content) > 1999 {
		content = content[:1999]
	}
	return c.rest.PostWebhookMessage(ctx, c.webhookURL, discordgo.WebhookParams{
		Content:  content,
		Username: constants.FrontendAppName + " Error Report",
	}, nil)
}

func (c *alertClient) Reader(r io.ReadCloser, levels ...zerolog.Level) {
	var levelSlice []string
	for _, l := range levels {
		levelSlice = append(levelSlice, l.String())
	}
	if len(levelSlice) < 1 {
		return
	}

	go func() {
		dec := json.NewDecoder(r)

		for {
			var e = make(map[string]any)
			err := dec.Decode(&e)
			if err == io.EOF {
				return
			}
			if err != nil {
				fmt.Fprintf(os.Stderr, "unmarshaling log failed: %v\n", err)
				continue
			}

			level, ok := e["level"].(string)
			if !ok || !slices.Contains(levelSlice, level) {
				continue
			}
			data, err := json.MarshalIndent(e, "", "  ")
			if err != nil {
				fmt.Fprintf(os.Stderr, "marshaling log failed: %v\n", err)
				continue
			}
			message, _ := e["message"].(string)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			err = c.Error(ctx, "**["+strings.ToUpper(level)+"]**: "+message, string(data))
			cancel()
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to report an error: %v\n", err)
				continue
			}
		}
	}()
}
