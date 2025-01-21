package sakura

import (
	"strings"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/facepaint"
)

func newPlayerNameCard(account models.Account) *facepaint.Block {
	var blocks []*facepaint.Block

	// nickname
	blocks = append(blocks, facepaint.NewBlocksContent(nameStyle.nicknameContainer(),
		facepaint.MustNewTextContent(nameStyle.nickname(), strings.ReplaceAll(account.Nickname, "_", " ")),
	))

	// clan tag
	var clanTagBlock *facepaint.Block
	if account.ClanTag != "" {
		clanTagBlock = facepaint.NewBlocksContent(nameStyle.tagContainer(), facepaint.MustNewTextContent(nameStyle.tag(), strings.ReplaceAll(account.ClanTag, "_", " ")))
		blocks = append(blocks, clanTagBlock)
	}

	return facepaint.NewBlocksContent(nameStyle.container(), blocks...)
}
