package core

import (
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/realtime"
	stats "github.com/cufee/aftermath/internal/stats/client/v2"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"golang.org/x/text/language"
)

var _ Client = &client{}

type Client interface {
	Stats(locale language.Tag) stats.Client

	Wargaming() wargaming.Client
	Database() database.Client
	PubSub() realtime.Client
	Fetch() fetch.Client
}

type client struct {
	fetch     fetch.Client
	wargaming wargaming.Client
	pubsub    realtime.Client
	db        database.Client
}

func (c *client) Wargaming() wargaming.Client {
	return c.wargaming
}

func (c *client) Database() database.Client {
	return c.db
}

func (c *client) PubSub() realtime.Client {
	return c.pubsub
}

func (c *client) Fetch() fetch.Client {
	return c.fetch
}

func (c *client) Stats(locale language.Tag) stats.Client {
	return stats.NewClient(c.fetch, c.db, c.wargaming, locale)
}

func NewClient(fetch fetch.Client, wargaming wargaming.Client, database database.Client, pubsub realtime.Client) *client {
	return &client{fetch: fetch, db: database, wargaming: wargaming, pubsub: pubsub}
}
