package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/internal/logic"
	"github.com/cufee/am-wg-proxy-next/v2/types"
)

var DaysOption = builder.NewOption("days", discordgo.ApplicationCommandOptionInteger).
	Min(1).
	Max(90).
	Params(
		builder.SetDescKey("common_option_stats_days_description"),
	)

var VehicleOption = builder.NewOption("tank", discordgo.ApplicationCommandOptionString).
	Autocomplete().
	Min(3).
	Max(64).
	Params(
		builder.SetDescKey("common_option_stats_tank_description"),
	)

var VehicleTierOption = builder.NewOption("tier", discordgo.ApplicationCommandOptionInteger).
	Params(
		builder.SetDescKey("common_option_stats_tank_tier_description"),
	).
	Choices(buildTierChoices()...)

var NicknameOption = builder.NewOption("nickname", discordgo.ApplicationCommandOptionString).
	Autocomplete().
	Min(3).
	Max(24).
	Params(
		builder.SetDescKey("common_option_stats_nickname_description"),
	)

var UserOption = builder.NewOption("user", discordgo.ApplicationCommandOptionUser).
	Params(
		builder.SetDescKey("common_option_stats_user_description"),
	)

var SessionStatsOptions = []builder.Option{
	DaysOption,
	NicknameOption,
	VehicleOption,
	VehicleTierOption,
	UserOption,
}

var CareerStatsOptions = []builder.Option{
	NicknameOption,
	VehicleOption,
	VehicleTierOption,
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
	TankTier       int
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

	tier, ok := common.GetOption[float64](data, "tier")
	if ok && options.TankID == "" {
		options.TankTier = int(tier)
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

func buildTierChoices() []builder.OptionChoice {
	var opts []builder.OptionChoice
	for i := range 10 {
		tier := 10 - i
		roman := logic.IntToRoman(tier)
		opts = append(opts, builder.NewChoice(fmt.Sprintf("tier_%s", roman), tier).Params(builder.SetNameKey(fmt.Sprintf("common_label_tier_%d", tier))))
	}
	return opts
}
