package core

import (
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/stats"
	"github.com/cufee/aftermath/internal/stats/fetch"
	"golang.org/x/text/language"
)

var _ Client = &client{}

type Client interface {
	Render(locale language.Tag) stats.Renderer

	Database() database.Client
	Fetch() fetch.Client
}

type client struct {
	fetch fetch.Client
	db    database.Client
}

func (c *client) Database() database.Client {
	return c.db
}

func (c *client) Fetch() fetch.Client {
	return c.fetch
}

func (c *client) Render(locale language.Tag) stats.Renderer {
	return stats.NewRenderer(c.fetch, locale)
}

func NewClient(fetch fetch.Client, database database.Client) *client {
	return &client{fetch: fetch, db: database}
}
