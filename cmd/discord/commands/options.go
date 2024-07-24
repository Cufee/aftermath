package commands

import (
	"regexp"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
)

var validNameRegex = regexp.MustCompile(`[^\w\_]`)

var DaysOption = builder.NewOption("days", discordgo.ApplicationCommandOptionInteger).
	Min(1).
	Max(90).
	Params(
		builder.SetNameKey("common_option_stats_days_name"),
		builder.SetDescKey("common_option_stats_days_description"),
	)

var VehicleOption = builder.NewOption("tank", discordgo.ApplicationCommandOptionString).
	Min(3).
	Max(32).
	Autocomplete().
	Params(
		builder.SetNameKey("common_option_stats_tank_name"),
		builder.SetDescKey("common_option_stats_tank_description"),
	)

var NicknameAndServerOptions = []builder.Option{
	builder.NewOption("nickname", discordgo.ApplicationCommandOptionString).
		Min(5).
		Max(30).
		Params(
			builder.SetNameKey("common_option_stats_nickname_name"),
			builder.SetDescKey("common_option_stats_nickname_description"),
		),
	builder.NewOption("server", discordgo.ApplicationCommandOptionString).
		Params(
			builder.SetNameKey("common_option_stats_realm_name"),
			builder.SetDescKey("common_option_stats_realm_description"),
		).
		Choices(
			builder.NewChoice("realm_na", "na").Params(builder.SetNameKey("common_label_realm_na")),
			builder.NewChoice("realm_eu", "eu").Params(builder.SetNameKey("common_label_realm_eu")),
			builder.NewChoice("realm_as", "as").Params(builder.SetNameKey("common_label_realm_as")),
		),
}

var UserOption = builder.NewOption("user", discordgo.ApplicationCommandOptionUser).
	Params(
		builder.SetNameKey("common_option_stats_user_name"),
		builder.SetDescKey("common_option_stats_user_description"),
	)

var DefaultStatsOptions = append([]builder.Option{
	DaysOption,
	VehicleOption,
	UserOption,
}, NicknameAndServerOptions...)

type StatsOptions struct {
	PeriodStart time.Time
	Days        int
	Server      string
	Nickname    string
	UserID      string
	TankSearch  string
	TankID      string
}

func (o StatsOptions) Validate(ctx common.Context) (string, bool) {
	// check if the name is valid
	if o.UserID == "" && o.Nickname != "" && !ValidatePlayerName(o.Nickname) {
		return "errors_generic_nickname_invalid", false
	}
	if o.UserID == "" && o.Nickname != "" && o.Server == "" {
		// entering nickname also requires to enter the server
		return "errors_generic_nickname_requires_server", false
	}

	if o.UserID != "" && o.UserID == ctx.User().ID {
		// mentioning self is redundant - this should not prevent the command from working though
		return "stats_error_mentioned_self_non_blocking", true
	}
	if o.UserID != "" && o.Nickname != "" {
		// mentioning a user and providing a nickname is redundant
		return "stats_error_too_many_arguments_non_blocking", true
	}
	return "", true
}

func GetDefaultStatsOptions(data []*discordgo.ApplicationCommandInteractionDataOption) StatsOptions {
	var options StatsOptions

	options.TankSearch, _ = common.GetOption[string](data, "tank")
	if strings.HasPrefix(options.TankSearch, "valid#vehicle#") {
		options.TankID = strings.TrimPrefix(options.TankSearch, "valid#vehicle#")
	}
	options.Nickname, _ = common.GetOption[string](data, "nickname")
	options.Server, _ = common.GetOption[string](data, "server")
	options.UserID, _ = common.GetOption[string](data, "user")

	if days, _ := common.GetOption[float64](data, "days"); days > 0 {
		options.Days = int(days)
		options.PeriodStart = time.Now().Add(time.Hour * 24 * time.Duration(days) * -1)
	}

	return options
}

func ValidatePlayerName(name string) bool {
	return !validNameRegex.MatchString(name)
}
