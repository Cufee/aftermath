package rest

import (
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

var (
	ErrUnknownWebhook          = errors.New("discord api: unknown webhook")
	ErrUnknownInteraction      = errors.New("discord api: unknown interaction")
	ErrInteractionAlreadyAcked = errors.New("discord api: interaction already acked")
)

func knownError(code int) error {
	switch code {
	default:
		return nil
	case discordgo.ErrCodeUnknownWebhook:
		return ErrUnknownWebhook
	case discordgo.ErrCodeUnknownInteraction:
		return ErrUnknownInteraction
	case discordgo.ErrCodeInteractionHasAlreadyBeenAcknowledged:
		return ErrInteractionAlreadyAcked
	}
}
