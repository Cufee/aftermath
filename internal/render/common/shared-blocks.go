package common

import (
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/facepaint"
	"github.com/cufee/facepaint/style"
)

func NewPlayerNameBlock(account models.Account, theme Theme) *facepaint.Block {
	var blocks []*facepaint.Block

	var clanTagBlock *facepaint.Block
	if account.ClanTag != "" {
		textStl := ClanTagTextStyle(theme)
		cardStl := ClanTagCardStyle(theme)
		clanTagBlock = facepaint.NewBlocksContent(cardStl.Options(), facepaint.MustNewTextContent(textStl.Options(), account.ClanTag))
		blocks = append(blocks, clanTagBlock)
	}

	nameStl := PlayerNameTextStyle(theme)
	blocks = append(blocks, facepaint.NewBlocksContent(PlayerNameCardLayout.Options(),
		facepaint.MustNewTextContent(nameStl.Options(), account.Nickname),
	))

	if clanTagBlock != nil {
		size := clanTagBlock.Dimensions()
		stl := style.Style{
			Width:  float64(size.Width),
			Height: 1,
		}
		blocks = append(blocks, facepaint.NewEmptyContent(stl.Options()))
	}

	wrapperStl := PlayerNameWrapperStyle(theme)
	return facepaint.NewBlocksContent(wrapperStl.Options(), blocks...)
}

func NewFooterBlock(stats fetch.AccountStatsOverPeriod, opts Options) *facepaint.Block {
	stl := FooterPillStyle(opts.Theme)
	var footer []*facepaint.Block
	for _, text := range opts.FooterText {
		footer = append(footer, facepaint.MustNewTextContent(stl.Options(), text))
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

	return facepaint.NewBlocksContent(FooterWrapperLayout.Options(), footer...)
}
