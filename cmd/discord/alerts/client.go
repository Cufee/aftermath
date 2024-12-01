package alerts

import (
	"context"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/cufee/aftermath/internal/json"

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
		var decoded = make(map[string]any)
		var marshaled = make([]byte, 1000)

		for {
			decoded = map[string]any{}

			err := dec.Decode(&decoded)
			if err == io.EOF {
				return
			}
			if err != nil {
				fmt.Fprintf(os.Stderr, "unmarshaling log failed: %v\n", err)
				continue
			}

			level, ok := decoded["level"].(string)
			if !ok || !slices.Contains(levelSlice, level) {
				continue
			}
			marshaled, err = json.MarshalIndent(decoded, "", "  ")
			if err != nil {
				fmt.Fprintf(os.Stderr, "marshaling log failed: %v\n", err)
				continue
			}
			message, _ := decoded["message"].(string)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			err = c.Error(ctx, "**["+strings.ToUpper(level)+"]**: "+message, string(marshaled))
			cancel()
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to report an error: %v\n", err)
				continue
			}
		}
	}()
}
