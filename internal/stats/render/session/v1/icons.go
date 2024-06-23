package session

import (
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/aftermath/internal/stats/render/assets"
	"github.com/cufee/aftermath/internal/stats/render/common/v1"
)

var iconsCache = make(map[string]common.Block, 6)

func getRatingIcon(rating frame.Value) (common.Block, bool) {
	style := common.Style{Width: specialRatingIconSize, Height: specialRatingIconSize}
	name := "rating-"
	switch {
	case rating.Float() > 5000:
		name += "diamond"
	case rating.Float() > 4000:
		name += "platinum"
	case rating.Float() > 3000:
		name += "gold"
	case rating.Float() > 2000:
		name += "silver"
	case rating.Float() > 0:
		name += "bronze"
	default:
		name += "calibration"
		style.BackgroundColor = common.TextAlt
	}
	if b, ok := iconsCache[name]; ok {
		return b, true
	}

	img, ok := assets.GetLoadedImage(name)
	if !ok {
		return common.Block{}, false
	}

	iconsCache[name] = common.NewImageContent(style, img)
	return iconsCache[name], true
}
