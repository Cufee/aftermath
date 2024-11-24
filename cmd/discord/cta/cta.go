package cta

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/internal/localization"
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
		TagsBlacklist: []string{"command_help"},
		headKey:       "cta_report_issues_head",
		bodyKey:       "cta_report_issues_body",
		buttons:       []discordgo.Button{common.ButtonJoinPrimaryGuild("cta_report_issues_button")},
	}
	ctaPersonalInstall = CallToActionMessage{
		TagsBlacklist: []string{"command_help"},
		headKey:       "cta_personal_install_head",
		bodyKey:       "cta_personal_install_body",
		buttons:       []discordgo.Button{common.ButtonInviteAftermath("cta_personal_install_button")},
	}

	ctaAbandonedPositiveGrowth = CallToActionMessage{
		headKey: "cta_abandoned_positive_growth_head",
		bodyKey: "cta_abandoned_positive_growth_body",
	}
	ctaAbandonedNegativeGrowth = CallToActionMessage{
		headKey: "cta_abandoned_negative_growth_head",
		bodyKey: "cta_abandoned_negative_growth_body",
	}
	ctaAbandonedNeutralGrowth = CallToActionMessage{
		headKey: "cta_abandoned_neutral_growth_head",
		bodyKey: "cta_abandoned_neutral_growth_body",
	}
)

func DefaultCTACollection(printer localization.Printer) ([]CallToActionMessage, bool) {
	col := []CallToActionMessage{
		ctaPersonalInstall,
		ctaReportIssues,

		ctaCommandHelp,
		ctaCommandReplay,
		ctaCommandLinksAdd,
		ctaCommandLinksVerify,
	}
	for _, m := range col {
		if !m.Verify(printer) {
			return nil, false
		}
	}
	return col, true
}
