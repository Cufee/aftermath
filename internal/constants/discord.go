package constants

import (
	"os"
	"time"
)

var (
	DiscordBotInviteURL          = MustGetEnv("BOT_INVITE_LINK")
	DiscordPrimaryGuildInviteURL = MustGetEnv("PRIMARY_GUILD_INVITE_LINK")
)

var (
	DiscordPrimaryToken     = MustGetEnv("DISCORD_TOKEN")
	DiscordPrimaryPublicKey = MustGetEnv("DISCORD_PUBLIC_KEY")

	DiscordPrivateBotEnabled = os.Getenv("INTERNAL_DISCORD_TOKEN") != ""
	DiscordPrivateToken      = MustGetEnv("INTERNAL_DISCORD_TOKEN", func() bool { return !DiscordPrivateBotEnabled })
	DiscordPrivatePublicKey  = MustGetEnv("INTERNAL_DISCORD_PUBLIC_KEY", func() bool { return !DiscordPrivateBotEnabled })
)

var (
	DiscordAuthClientID      = MustGetEnv("DISCORD_AUTH_CLIENT_ID")
	DiscordAuthClientSecret  = MustGetEnv("DISCORD_AUTH_CLIENT_SECRET")
	DiscordAuthRedirectURL   = MustGetEnv("DISCORD_AUTH_REDIRECT_URL")
	DiscordAuthDefaultScopes = MustGetEnv("DISCORD_AUTH_DEFAULT_SCOPES")
)

var (
	DiscordContentModerationChannelID = MustGetEnv("DISCORD_CONTENT_MODERATION_CHANNEL_ID")
)

var (
	ImageUploadMaxSize  int64 = 1_000 * 1_000 * 2 // ~2MB - assuming 1000x1000 16Bit image at 2 bytes per pixel
	ReplayUploadMaxSize int64 = 1_000_000 * 3     // ~3MB - replays are typically significantly smaller than this
)

var (
	DefaultFeatureBanDuration time.Duration = time.Hour * 24 * 180 // 180 days
)

var (
	DiscordEmojiYellowID = MustGetEnv("EMOJI_AFTERMATH_YELLOW_ID")
	DiscordEmojiBlueID   = MustGetEnv("EMOJI_AFTERMATH_BLUE_ID")
	DiscordEmojiRedID    = MustGetEnv("EMOJI_AFTERMATH_RED_ID")
)
