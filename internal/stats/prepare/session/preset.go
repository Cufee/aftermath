package session

import (
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/aftermath/internal/stats/prepare/common"
)

func presetToBlock(preset common.Tag, session, career frame.StatsFrame) (common.StatsBlock[BlockData], error) {
	// create blocks from stats
	// this module has no special blocks, so we can just use the common.Block#FillValue
	sessionBlock := common.StatsBlock[BlockData](common.NewBlock(preset, BlockData{}))
	err := sessionBlock.FillValue(session)
	if err != nil {
		return common.StatsBlock[BlockData]{}, nil
	}
	careerBlock := common.StatsBlock[BlockData](common.NewBlock(preset, BlockData{}))
	err = careerBlock.FillValue(career)
	if err != nil {
		return common.StatsBlock[BlockData]{}, nil
	}

	block := common.StatsBlock[BlockData](common.NewBlock(preset, BlockData{
		Session: sessionBlock.Value,
		Career:  careerBlock.Value,
	}))
	block.Value = sessionBlock.Value
	return block, nil
}
