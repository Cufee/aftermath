package common

import (
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/aftermath/internal/stats/render/assets"
)

var iconsCache = make(map[string]Block, 6)

func GetRatingIconName(rating float32) string {
	switch {
	case rating > 5000:
		return "diamond"
	case rating > 4000:
		return "platinum"
	case rating > 3000:
		return "gold"
	case rating > 2000:
		return "silver"
	case rating > 0:
		return "bronze"
	default:
		return "calibration"
	}
}

func GetRatingIcon(rating frame.Value, size float64) (Block, bool) {
	style := Style{Width: size, Height: size}
	if rating.Float() < 0 {
		style.BackgroundColor = TextAlt
	}
	name := "rating-" + GetRatingIconName(rating.Float())

	if b, ok := iconsCache[name]; ok {
		return b, true
	}

	img, ok := assets.GetLoadedImage(name)
	if !ok {
		return Block{}, false
	}

	iconsCache[name] = NewImageContent(style, img)
	return iconsCache[name], true
}
