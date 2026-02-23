package public

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/cufee/aftermath/cmd/discord/commands"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/middleware"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/glossary"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/logic"
	"github.com/cufee/aftermath/internal/permissions"
	stats "github.com/cufee/aftermath/internal/stats/client/common"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
)

var (
	careerCommandMiddleware = []middleware.MiddlewareFunc{middleware.RequirePermissions(permissions.UseTextCommands, permissions.UseImageCommands)}
	careerCommandOptions    = commands.CareerStatsOptions
)

func careerCommandHandler(ctx common.Context) error {
	options := commands.GetDefaultStatsOptions(ctx.Options())
	message, valid := options.Validate(ctx)
	if !valid {
		return ctx.Reply().Send(message)
	}

	var accountID string
	var opts = []stats.RequestOption{stats.WithWN8()}
	if options.TankID != "" {
		opts = append(opts, stats.WithVehicleIDs(options.TankID))
	}
	if options.TankTier > 0 && options.TankTier <= 10 {
		ids, ok := glossary.TierVehicleIDs(options.TankTier)
		if ok {
			opts = append(opts, stats.WithVehicleIDs(ids...), stats.WithFooterText(ctx.Localize(fmt.Sprintf("common_label_tier_%d", options.TankTier))))
		}
	}

	ioptions := statsOptions{StatsOptions: options}

	switch {
	case options.UserID != "":
		// mentioned another user, check if the user has an account linked
		mentionedUser, _ := ctx.Core().Database().GetUserByID(ctx.Ctx(), options.UserID, database.WithConnections(), database.WithSubscriptions(), database.WithContent())
		defaultAccount, hasDefaultAccount := mentionedUser.Connection(models.ConnectionTypeWargaming, nil, new(true))
		if !hasDefaultAccount {
			return ctx.Reply().IsError(common.UserError).Send("stats_error_connection_not_found_vague")
		}
		accountID = defaultAccount.ReferenceID

		if img, content, err := logic.GetAccountBackgroundImage(ctx.Ctx(), ctx.Core().Database(), accountID); err == nil {
			opts = append(opts, stats.WithBackground(img, true))
			ioptions.BackgroundID = content.ID
		}

	case options.AccountID != "":
		// account selected from autocomplete
		accountID = options.AccountID

		if img, content, err := logic.GetAccountBackgroundImage(ctx.Ctx(), ctx.Core().Database(), accountID); err == nil {
			opts = append(opts, stats.WithBackground(img, true))
			ioptions.BackgroundID = content.ID
		}

	case options.NicknameSearch != "" && options.AccountID == "":
		// nickname provided, but user did not select an option from autocomplete
		accounts, err := accountsFromBadInput(ctx.Ctx(), ctx.Core().Fetch(), options.NicknameSearch)
		if err != nil && !errors.Is(err, fetch.ErrAccountNotFound) {
			return ctx.Err(err, common.ApplicationError)
		}
		if len(accounts) == 0 {
			return ctx.Reply().IsError(common.UserError).Send("stats_account_not_found")
		}

		realms := make(map[string]struct{})
		for _, a := range accounts {
			realms[a.Realm.String()] = struct{}{}
		}
		if len(realms) > 1 {
			reply, err := realmSelectButtons(ctx, ctx.ID(), accounts)
			if err != nil {
				return ctx.Err(err, common.ApplicationError)
			}
			return reply.Send()
		}

		// one or more options on the same server - just pick the first one
		message = "stats_bad_nickname_input_hint"
		accountID = fmt.Sprint(accounts[0].ID)

	default:
		defaultAccount, hasDefaultAccount := ctx.User().Connection(models.ConnectionTypeWargaming, nil, new(true))
		if !hasDefaultAccount {
			return ctx.Reply().Send("command_career_help_message")
		}
		// command used without options, but user has a default connection
		accountID = defaultAccount.ReferenceID

		if img, content, err := logic.GetAccountBackgroundImage(ctx.Ctx(), ctx.Core().Database(), accountID); err == nil {
			opts = append(opts, stats.WithBackground(img, true))
			ioptions.BackgroundID = content.ID
		} else {
			background, _ := ctx.User().Content(models.UserContentTypePersonalBackground)
			if img, err := logic.UserContentToImage(background); err == nil {
				opts = append(opts, stats.WithBackground(img, true))
				ioptions.BackgroundID = background.ID
			}
		}
	}

	image, meta, err := ctx.Core().Stats(ctx.Locale()).PeriodImage(context.Background(), accountID, options.PeriodStart, opts...)
	if err != nil {
		return ctx.Err(err, common.ApplicationError)
	}

	ioptions.AccountID = accountID
	button, saveErr := ioptions.refreshButton(ctx, ctx.ID())
	if saveErr != nil {
		// nil button will not cause an error and will be ignored
		log.Err(err).Str("interactionId", ctx.ID()).Str("command", "session").Msg("failed to save discord interaction")
	}

	var buf bytes.Buffer
	err = image.PNG(&buf)
	if err != nil {
		return ctx.Err(err, common.ApplicationError)
	}

	var timings []string
	if ctx.User().HasPermission(permissions.UseDebugFeatures) {
		timings = append(timings, "```")
		for name, duration := range meta.Timings {
			timings = append(timings, fmt.Sprintf("%s: %v", name, duration.Milliseconds()))
		}
		timings = append(timings, "```")
	}

	return ctx.Reply().WithAds().Hint(message).File(buf.Bytes(), "stats_command_by_aftermath.png").Component(button).Text(timings...).Send()
}

func init() {
	commands.LoadedPublic.Add(builder.NewCommand("stats").
		Middleware(careerCommandMiddleware...).
		Options(careerCommandOptions...).
		Params(builder.SetDescKey("command_career_description")).
		Handler(careerCommandHandler))
	commands.LoadedPublic.Add(builder.NewCommand("career").
		Middleware(careerCommandMiddleware...).
		Options(careerCommandOptions...).
		Params(builder.SetNameKey("command_career_name"), builder.SetDescKey("command_career_description")).
		Handler(careerCommandHandler))
}
