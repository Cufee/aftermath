package cta

import (
	"math/rand/v2"
	"slices"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/internal/constants"
	"github.com/cufee/aftermath/internal/localization"
	"golang.org/x/text/language"
)

type CallToActionMessage struct {
	headKey string
	bodyKey string
	buttons []discordgo.Button

	Color         int
	TagsBlacklist []string
	TagsWhitelist []string
}

func (m CallToActionMessage) Verify(printer localization.Printer) bool {
	localizedButtons := m.Buttons(printer)
	if len(localizedButtons) != len(m.buttons) {
		return false
	}
	for i, template := range m.buttons {
		if localizedButtons[i].Label == template.Label {
			return false
		}
	}
	return printer(m.headKey) != m.headKey && printer(m.bodyKey) != m.bodyKey
}

func (m CallToActionMessage) HeadText(printer localization.Printer) string {
	return printer(m.headKey)
}

func (m CallToActionMessage) BodyText(printer localization.Printer) string {
	return printer(m.bodyKey)
}

func (m CallToActionMessage) Buttons(printer localization.Printer) []discordgo.Button {
	var localized []discordgo.Button
	for _, template := range m.buttons {
		instance := template
		instance.Label = printer(template.Label)
		localized = append(localized, instance)
	}
	return localized
}

var (
	ctaCommandReplay = CallToActionMessage{
		TagsBlacklist: []string{"command_replay"},
		headKey:       "cta_command_replay_head",
		bodyKey:       "cta_command_replay_body",
	}
	ctaCommandLinksAdd = CallToActionMessage{
		TagsBlacklist: []string{"command_links"},
		headKey:       "cta_command_links_add_head",
		bodyKey:       "cta_command_links_add_body",
	}
	ctaCommandLinksVerify = CallToActionMessage{
		TagsBlacklist: []string{"command_links"},
		headKey:       "cta_command_links_verify_head",
		bodyKey:       "cta_command_links_verify_body",
	}
	ctaCommandHelp = CallToActionMessage{
		TagsBlacklist: []string{"command_help"},
		headKey:       "cta_command_help_head",
		bodyKey:       "cta_command_help_body",
		buttons:       []discordgo.Button{common.ButtonJoinPrimaryGuild("cta_command_help_button")},
	}
	ctaReportIssues = CallToActionMessage{
		// TagsBlacklist: []string{"command_help"},
		headKey: "cta_report_issues_head",
		bodyKey: "cta_report_issues_body",
		buttons: []discordgo.Button{common.ButtonJoinPrimaryGuild("cta_report_issues_button")},
	}
	ctaPersonalInstall = CallToActionMessage{
		TagsBlacklist: []string{"command_help"},
		headKey:       "cta_personal_install_head",
		bodyKey:       "cta_personal_install_body",
		buttons:       []discordgo.Button{common.ButtonInviteAftermath("cta_personal_install_button")},
	}
	ctaGuildInstall = CallToActionMessage{
		TagsBlacklist: []string{"command_help"},
		headKey:       "cta_guild_install_head",
		bodyKey:       "cta_guild_install_body",
		buttons:       []discordgo.Button{common.ButtonInviteAftermath("cta_guild_install_button")},
	}

	ctaAbandonedPositiveGrowth = CallToActionMessage{
		TagsWhitelist: []string{"growth", "growth_positive"},
		headKey:       "cta_abandoned_positive_growth_head",
		bodyKey:       "cta_abandoned_positive_growth_body",
	}
	ctaAbandonedNegativeGrowth = CallToActionMessage{
		TagsWhitelist: []string{"growth", "growth_negative"},
		headKey:       "cta_abandoned_negative_growth_head",
		bodyKey:       "cta_abandoned_negative_growth_body",
	}
	ctaAbandonedNeutralGrowth = CallToActionMessage{
		TagsWhitelist: []string{"growth", "growth_neutral"},
		headKey:       "cta_abandoned_neutral_growth_head",
		bodyKey:       "cta_abandoned_neutral_growth_body",
	}
)

type MessageCollection []CallToActionMessage

var defaultCTACollection = MessageCollection{
	ctaPersonalInstall,
	ctaGuildInstall,
	ctaReportIssues,

	ctaCommandHelp,
	ctaCommandReplay,
	ctaCommandLinksAdd,
	ctaCommandLinksVerify,
}

func (c MessageCollection) NewEmbed(locale language.Tag, tags []string) (discordgo.MessageEmbed, []discordgo.MessageComponent, bool) {
	var applicable []CallToActionMessage
	printer, err := localization.NewPrinter("cta", locale)
	if err != nil {
		// if the requested locale is not supported, we should not send any ads
		return discordgo.MessageEmbed{}, nil, false
	}

optionsLoop:
	for _, option := range c {
		if !option.Verify(printer) {
			continue
		}

		for tag := range slices.Values(tags) {
			if slices.Contains(option.TagsBlacklist, tag) {
				continue optionsLoop
			}
			if option.TagsWhitelist != nil && !slices.Contains(option.TagsWhitelist, tag) {
				continue optionsLoop
			}
		}

		applicable = append(applicable, option)
	}
	if len(applicable) == 0 {
		return discordgo.MessageEmbed{}, nil, false
	}

	selected := applicable[0]
	if len(applicable) > 1 {
		selected = applicable[rand.IntN(len(applicable)-1)]
	}

	// body
	embed := discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			URL:     constants.FrontendURL,
			Name:    constants.FrontendAppName,
			IconURL: constants.FrontendURL + "/assets/icon/64.png",
		},
		Title:       selected.HeadText(printer),
		Description: selected.BodyText(printer),
	}

	// buttons
	var row discordgo.ActionsRow
	for _, btn := range selected.Buttons(printer) {
		row.Components = append(row.Components, btn)
	}
	if len(row.Components) > 0 {
		return embed, []discordgo.MessageComponent{row}, true
	}
	return embed, nil, true
}
