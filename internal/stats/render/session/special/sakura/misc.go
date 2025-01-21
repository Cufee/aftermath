package sakura

import (
	"github.com/cufee/aftermath/internal/render/common"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/session/v1"
	"github.com/cufee/facepaint"
)

func newFooterCard(stats fetch.AccountStatsOverPeriod, cards session.Cards, opts common.Options) *facepaint.Block {
	pillStyle := styledFooter.pill()
	var footer []*facepaint.Block
	if opts.VehicleID != "" {
		footer = append(footer, facepaint.MustNewTextContent(pillStyle, cards.Unrated.Overview.Title))
	}

	sessionTo := stats.PeriodEnd.Format("Jan 2, 2006")
	sessionFromFormat := "Jan 2, 2006"
	if stats.PeriodStart.Year() == stats.PeriodEnd.Year() {
		sessionFromFormat = "Jan 2"
	}
	sessionFrom := stats.PeriodStart.Format(sessionFromFormat)
	if stats.PeriodStart.IsZero() || sessionFrom == sessionTo {
		footer = append(footer, facepaint.MustNewTextContent(pillStyle, sessionTo))
	} else {
		footer = append(footer, facepaint.MustNewTextContent(pillStyle, sessionFrom+" - "+sessionTo))
	}

	return facepaint.NewBlocksContent(styledFooter.container, footer...)
}
