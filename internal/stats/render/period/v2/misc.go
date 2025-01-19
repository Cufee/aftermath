package period

import (
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/render/common"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/period/v1"
	"github.com/cufee/facepaint"
	"github.com/cufee/facepaint/style"
)

func newPlayerNameCard(account models.Account) *facepaint.Block {
	var blocks []*facepaint.Block

	// clan tag
	var clanTagBlock *facepaint.Block
	if account.ClanTag != "" {
		stl := styledPlayerClanTag()
		clanTagBlock = facepaint.NewBlocksContent(styledPlayerClanTagCard.Options(), facepaint.MustNewTextContent(stl.Options(), account.ClanTag))
		blocks = append(blocks, clanTagBlock)
	}

	// nickname
	stl := styledPlayerName()
	blocks = append(blocks, facepaint.NewBlocksContent(styledPlayerNameCard.Options(),
		facepaint.MustNewTextContent(stl.Options(), account.Nickname),
	))

	// spacer
	if clanTagBlock != nil {
		size := clanTagBlock.Dimensions()
		stl := style.Style{
			Width:  float64(size.Width),
			Height: 1,
		}
		blocks = append(blocks, facepaint.NewEmptyContent(stl.Options()))
	}

	return facepaint.NewBlocksContent(styledPlayerNameWrapper.Options(), blocks...)
}

func newFooterCard(stats fetch.AccountStatsOverPeriod, cards period.Cards, opts common.Options) *facepaint.Block {
	stl := styledFooterCard()
	var footer []*facepaint.Block
	if opts.VehicleID != "" {
		footer = append(footer, facepaint.MustNewTextContent(stl.Options(), cards.Overview.Title))
	}

	sessionTo := stats.PeriodEnd.Format("Jan 2, 2006")
	sessionFromFormat := "Jan 2, 2006"
	if stats.PeriodStart.Year() == stats.PeriodEnd.Year() {
		sessionFromFormat = "Jan 2"
	}
	sessionFrom := stats.PeriodStart.Format(sessionFromFormat)
	if stats.PeriodStart.IsZero() || sessionFrom == sessionTo {
		footer = append(footer, facepaint.MustNewTextContent(stl.Options(), sessionTo))
	} else {
		footer = append(footer, facepaint.MustNewTextContent(stl.Options(), sessionFrom+" - "+sessionTo))
	}

	return facepaint.NewBlocksContent(styledFooterWrapper.Options(), footer...)
}
