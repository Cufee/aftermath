package public

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/json"
	"github.com/cufee/am-wg-proxy-next/v2/types"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
)

type statsOptions struct {
	commands.StatsOptions
	BackgroundID string
	ReferenceID  string
}

func (o statsOptions) fromInteraction(data models.DiscordInteraction) (statsOptions, error) {
	if data.Type != models.InteractionTypeComponent || data.Meta == nil || data.Meta["stats_options"] == nil {
		return statsOptions{}, errors.New("interactions contains no stats options")
	}

	raw, ok := data.Meta["stats_options"].(string)
	if !ok {
		return statsOptions{}, errors.New("invalid stats options data type")
	}

	return o, json.Unmarshal([]byte(raw), &o)
}

func (o statsOptions) refreshButton(ctx common.Context, id string) (discordgo.MessageComponent, error) {
	encoded, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}

	interaction := models.DiscordInteraction{
		Result:    "generated-refresh-button",
		EventID:   id,
		Locale:    ctx.Locale(),
		UserID:    ctx.User().ID,
		GuildID:   ctx.RawInteraction().GuildID,
		ChannelID: ctx.RawInteraction().ChannelID,
		MessageID: "not-available",
		Type:      models.InteractionTypeComponent,
		Meta:      map[string]any{"stats_options": string(encoded)},
	}
	if ctx.RawInteraction().Message != nil {
		interaction.MessageID = ctx.RawInteraction().Message.ID
	}

	interaction, err = ctx.Core().Database().CreateDiscordInteraction(ctx.Ctx(), interaction)
	if err != nil {
		return nil, err
	}

	return newStatsRefreshButton(interaction), nil
}

var validNicknameSeparatorChars = []string{"!", "@", "#", "$", "%", "^", "&", "*", "-", "=", "+"}
var validFuzzyServers = map[string]types.Realm{
	"na":            types.RealmNorthAmerica,
	"north america": types.RealmNorthAmerica,
	"north-america": types.RealmNorthAmerica,
	"northamerica":  types.RealmNorthAmerica,
	"america":       types.RealmNorthAmerica,
	"com":           types.RealmNorthAmerica,
	"eu":            types.RealmEurope,
	"europe":        types.RealmEurope,
	"as":            types.RealmAsia,
	"asia":          types.RealmAsia,
}

func accountsFromBadInput(ctx context.Context, client fetch.Client, input string) ([]fetch.AccountWithRealm, error) {
	// input is most likely a valid account nickname
	if wargaming.ValidatePlayerNickname(input) {
		return client.BroadSearch(ctx, input, 2)
	}

	// input contained some extra characters that cannot possibly be in a player name
	for _, sep := range validNicknameSeparatorChars {
		split := strings.Split(input, sep)
		if len(split) != 2 {
			continue
		}

		for value, realm := range validFuzzyServers {
			var search string
			if strings.ToLower(split[0]) == value {
				search = split[1]
			}
			if strings.ToLower(split[1]) == value {
				search = split[0]
			}

			if search == "" {
				continue
			}
			if !wargaming.ValidatePlayerNickname(search) {
				return nil, nil // nothing found
			}

			account, err := client.Search(ctx, search, realm, 3)
			if err != nil && errors.Is(err, fetch.ErrAccountNotFound) {
				return nil, err
			}
			if account.ID == 0 {
				return nil, nil // nothing found
			}
			return []fetch.AccountWithRealm{{Account: account, Realm: realm}}, nil
		}
	}

	return nil, nil
}

func realmSelectButtons(ctx common.Context, id string, accounts []fetch.AccountWithRealm) (common.Reply, error) {
	realmAccounts := make(map[string]fetch.AccountWithRealm)
	for _, a := range accounts {
		if _, ok := realmAccounts[a.Realm.String()]; !ok {
			realmAccounts[a.Realm.String()] = a
		}
	}

	var message = []string{ctx.Localize("stats_multiple_accounts_found")}
	row := discordgo.ActionsRow{}
	for _, a := range realmAccounts {
		message = append(message, fmt.Sprintf("**[%s]** %s", a.Realm, a.Nickname))

		encoded, err := json.Marshal(statsOptions{StatsOptions: commands.StatsOptions{AccountID: fmt.Sprint(a.ID), Realm: a.Realm}})
		if err != nil {
			return ctx.Reply(), err
		}

		interaction := models.DiscordInteraction{
			Result:    "generated-realm-select-button",
			EventID:   id,
			Locale:    ctx.Locale(),
			UserID:    ctx.User().ID,
			GuildID:   ctx.RawInteraction().GuildID,
			ChannelID: ctx.RawInteraction().ChannelID,
			MessageID: "not-available",
			Type:      models.InteractionTypeComponent,
			Meta:      map[string]any{"stats_options": string(encoded)},
		}
		if ctx.RawInteraction().Message != nil {
			interaction.MessageID = ctx.RawInteraction().Message.ID
		}

		interaction, err = ctx.Core().Database().CreateDiscordInteraction(ctx.Ctx(), interaction)
		if err != nil {
			return ctx.Reply(), err
		}

		row.Components = append(row.Components, discordgo.Button{
			Label:    a.Realm.String(),
			Style:    discordgo.SecondaryButton,
			CustomID: fmt.Sprintf("refresh_stats_from_button#%s", interaction.ID),
		})
	}
	return ctx.Reply().Hint("stats_bad_nickname_input_hint").Component(row).Text(strings.Join(message, "\n")), nil
}
