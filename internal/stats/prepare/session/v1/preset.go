package session

import (
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/aftermath/internal/stats/prepare/common/v1"
)

func (b *cardBuilder) presetToBlock(preset common.Tag, session, career frame.StatsFrame) (common.StatsBlock[BlockData, string], error) {
	var err error
	// create blocks from stats
	// this module has no special blocks, so we can just use the common.Block#FillValue
	sessionBlock := common.StatsBlock[BlockData, string](common.NewBlock[BlockData, string](preset, BlockData{}))
	if preset == common.TagAvgTier {
		err = sessionBlock.FillValue(session, b.session.RegularBattles.Vehicles, b.glossary)
	} else {
		err = sessionBlock.FillValue(session)
	}
	if err != nil {
		return common.StatsBlock[BlockData, string]{}, err
	}
	careerBlock := common.StatsBlock[BlockData, string](common.NewBlock[BlockData, string](preset, BlockData{}))
	if preset == common.TagAvgTier {
		err = careerBlock.FillValue(career, b.career.RegularBattles.Vehicles, b.glossary)
	} else {
		err = careerBlock.FillValue(career)
	}
	if err != nil {
		return common.StatsBlock[BlockData, string]{}, err
	}

	block := common.StatsBlock[BlockData, string](common.NewBlock[BlockData, string](preset, BlockData{
		S: sessionBlock.Value(),
		C: careerBlock.Value(),
	}))
	block.SetValue(sessionBlock.Value())
	return block, nil
}
