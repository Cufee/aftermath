package commands

import (
	"regexp"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/am-wg-proxy-next/v2/types"
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
	Autocomplete().
	Min(3).
	Max(64).
	Params(
		builder.SetNameKey("common_option_stats_tank_name"),
		builder.SetDescKey("common_option_stats_tank_description"),
	)

var NicknameOption = builder.NewOption("nickname", discordgo.ApplicationCommandOptionString).
	Autocomplete().
	Min(3).
	Max(64).
	Params(
		builder.SetNameKey("common_option_stats_nickname_name"),
		builder.SetDescKey("common_option_stats_nickname_description"),
	)

var UserOption = builder.NewOption("user", discordgo.ApplicationCommandOptionUser).
	Params(
		builder.SetNameKey("common_option_stats_user_name"),
		builder.SetDescKey("common_option_stats_user_description"),
	)

var DefaultStatsOptions = []builder.Option{
	DaysOption,
	NicknameOption,
	VehicleOption,
	UserOption,
}

type StatsOptions struct {
	PeriodStart    time.Time
	Days           int
	NicknameSearch string
	AccountID      string
	Realm          types.Realm
	UserID         string
	TankSearch     string
	TankID         string
}

func (o StatsOptions) Validate(ctx common.Context) (string, bool) {
	if o.UserID != "" && o.UserID == ctx.User().ID {
		// mentioning self is redundant - this should not prevent the command from working though
		return "stats_error_mentioned_self_non_blocking", true
	}
	if o.UserID != "" && o.NicknameSearch != "" {
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
	options.NicknameSearch, _ = common.GetOption[string](data, "nickname")
	if strings.HasPrefix(options.NicknameSearch, "valid#account#") {
		data := strings.Split(strings.TrimPrefix(options.NicknameSearch, "valid#account#"), "#")
		if len(data) == 2 {
			options.Realm = types.Realm(data[1])
			options.AccountID = data[0]
		}
	}

	options.UserID, _ = common.GetOption[string](data, "user")

	if days, _ := common.GetOption[float64](data, "days"); days > 0 {
		options.Days = int(days)
		options.PeriodStart = time.Now().Add(time.Hour * 24 * time.Duration(days) * -1)
	}

	return options
}

func ValidatePlayerName(name string) bool {
	return len(name) > 3 && len(name) < 24 && strings.HasPrefix(name, "valid#account#") || !validNameRegex.MatchString(name)
}
