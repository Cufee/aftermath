package period

import (
	"github.com/cufee/aftermath/internal/database/models"
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
