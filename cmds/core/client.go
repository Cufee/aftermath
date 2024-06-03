package core

import (
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/stats"
	"github.com/cufee/aftermath/internal/stats/fetch"
	"golang.org/x/text/language"
)

type Client struct {
	Fetch fetch.Client
	DB    database.Client
}

func (c *Client) Render(locale language.Tag) stats.Renderer {
	return stats.NewRenderer(c.Fetch, locale)
}
