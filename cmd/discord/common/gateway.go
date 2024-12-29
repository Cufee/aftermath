package common

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/internal/database"
)

type EventHandler[T any] struct {
	Match  func(database.Client, *discordgo.Session, *T) bool
	Handle func(Context, *T) error
}
