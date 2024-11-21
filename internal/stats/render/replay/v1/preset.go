package replay

import (
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/replay/v1"
	"github.com/cufee/aftermath/internal/stats/render/common/v1"
)

func statsBlockToBlock(stats prepare.StatsBlock[replay.BlockData, string], width float64) common.Block {
	value := common.NewTextContent(common.Style{
		Font:      common.FontLarge(),
		FontColor: common.TextPrimary,
	}, stats.Value().String())

	return common.NewBlocksContent(common.Style{Direction: common.DirectionVertical, AlignItems: common.AlignItemsCenter, Width: width},
		value,
		common.NewTextContent(common.Style{
			Font:      common.FontSmall(),
			FontColor: common.TextAlt,
		}, stats.Label))
}
