package constants

import "os"

var (
	DiscordAlertsEnabled           = os.Getenv("DISCORD_ERROR_REPORT_WEBHOOK_URL") != ""
	DiscordAlertsWebhookURL string = MustGetEnv("DISCORD_ERROR_REPORT_WEBHOOK_URL", func() bool { return !DiscordAlertsEnabled })
)
