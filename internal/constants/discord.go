package constants

import (
	"os"
	"time"
)

var (
	DiscordBotInviteURL          = mustGetEnv("BOT_INVITE_LINK")
	DiscordPrimaryGuildInviteURL = mustGetEnv("PRIMARY_GUILD_INVITE_LINK")
)

var (
	DiscordPrimaryToken     = mustGetEnv("DISCORD_TOKEN")
	DiscordPrimaryPublicKey = mustGetEnv("DISCORD_PUBLIC_KEY")

	DiscordPrivateBotEnabled = os.Getenv("INTERNAL_DISCORD_TOKEN") != ""
	DiscordPrivateToken      = mustGetEnv("INTERNAL_DISCORD_TOKEN", func() bool { return !DiscordPrivateBotEnabled })
	DiscordPrivatePublicKey  = mustGetEnv("INTERNAL_DISCORD_PUBLIC_KEY", func() bool { return !DiscordPrivateBotEnabled })
)

var (
	DiscordAuthClientID      = mustGetEnv("DISCORD_AUTH_CLIENT_ID")
	DiscordAuthClientSecret  = mustGetEnv("DISCORD_AUTH_CLIENT_SECRET")
	DiscordAuthRedirectURL   = mustGetEnv("DISCORD_AUTH_REDIRECT_URL")
	DiscordAuthDefaultScopes = mustGetEnv("DISCORD_AUTH_DEFAULT_SCOPES")
)

var (
	DiscordContentModerationChannelID = mustGetEnv("DISCORD_CONTENT_MODERATION_CHANNEL_ID")
)

var (
	ImageUploadMaxSize  int64 = 1_000 * 1_000 * 2 // ~2MB - assuming 1000x1000 16Bit image at 2 bytes per pixel
	ReplayUploadMaxSize int64 = 1_000_000         // ~1MB - replays are typically significantly smaller than this
)

var (
	DefaultFeatureBanDuration time.Duration = time.Hour * 24 * 180 // 180 days
)
