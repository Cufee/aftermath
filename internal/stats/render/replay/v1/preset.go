package replay

import (
	"image/color"

	common "github.com/cufee/aftermath/internal/render/v1"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/replay/v1"
)

func statsBlockToBlock(stats prepare.StatsBlock[replay.BlockData, string], width float64) common.Block {
	var fontColor color.Color = common.TextPrimary

	label := []common.Block{
		common.NewTextContent(common.Style{
			Font:      common.FontSmall(),
			FontColor: common.TextAlt,
		}, stats.Label),
	}
	if stats.Tag == prepare.TagWinrate {
		label = append(label, playerWinrateIndicator(stats.V))
	}

	return common.NewBlocksContent(common.Style{Direction: common.DirectionVertical, AlignItems: common.AlignItemsCenter, Width: width},
		common.NewTextContent(common.Style{Font: common.FontLarge(), FontColor: fontColor}, stats.Value().String()),
		common.NewBlocksContent(common.Style{AlignItems: common.AlignItemsCenter, JustifyContent: common.JustifyContentCenter, Gap: playerWinrateIndicatorSize}, label...),
	)
}
