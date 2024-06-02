package core

import (
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/stats/fetch"
)

type Client struct {
	Fetch fetch.Client
	DB    database.Client
}
