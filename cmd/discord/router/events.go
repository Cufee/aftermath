package router

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/rest"
	"github.com/cufee/aftermath/internal/constants"
	"github.com/cufee/aftermath/internal/log"
)

type eventsClient struct {
	rest *rest.Client

	mx            sync.Mutex
	messageLength int

	webhookURL    string
	pushFrequency time.Duration
	queue         chan string
}

func newEventsClient(rest *rest.Client, webhook string, pushFrequency time.Duration) *eventsClient {
	return &eventsClient{
		rest:          rest,
		webhookURL:    webhook,
		pushFrequency: pushFrequency,
		queue:         make(chan string, 10),
		mx:            sync.Mutex{},
	}
}

func (c *eventsClient) start() {
	ticker := time.NewTicker(c.pushFrequency)
	for range ticker.C {
		if len(c.queue) < 1 {
			continue
		}

		c.mx.Lock()
		messages := c.flush()
		c.mx.Unlock()

		c.send(messages)
	}
}

func (c *eventsClient) flush() []string {
	var messages []string
outer:
	for {
		select {
		case message := <-c.queue:
			messages = append(messages, message)
		default:
			break outer
		}
	}
	c.messageLength = 0
	return messages
}

func (c *eventsClient) send(messages []string) {
	log.Debug().Msgf("sending %d messages", len(messages))
	if len(messages) < 1 {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err := c.rest.PostWebhookMessage(ctx, c.webhookURL, discordgo.WebhookParams{
		Username: constants.FrontendAppName + "Event Firehose",
		Content:  strings.Join(messages, "\n"),
	}, nil)
	if err != nil {
		log.Err(err).Msg("failed to send messages")
	}

}

func (c *eventsClient) push(message string) {
	c.mx.Lock()
	defer c.mx.Unlock()

	if c.messageLength+len(message)+2 >= 2000 {
		messages := c.flush()
		go c.send(messages)
	}

	c.messageLength += len(message) + 2
	c.queue <- message
}
