package render

import (
	"image/color"

	"github.com/cufee/aftermath/internal/render/assets"
	"github.com/cufee/aftermath/internal/stats/frame"
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

func GetRatingTierName(rating float32) string {
	switch {
	case rating > 5000:
		return "Diamond"
	case rating > 4000:
		return "Platinum"
	case rating > 3000:
		return "Gold"
	case rating > 2000:
		return "Silver"
	case rating > 0:
		return "Bronze"
	default:
		return ""
	}
}

func GetRatingColors(rating float32) ratingColors {
	switch {
	case rating > 5000:
		return ratingColors{color.NRGBA{181, 106, 181, 255}, color.Black}
	case rating > 4000:
		return ratingColors{color.NRGBA{154, 197, 219, 255}, color.Black}
	case rating > 3000:
		return ratingColors{color.NRGBA{255, 215, 0, 255}, color.Black}
	case rating > 2000:
		return ratingColors{color.NRGBA{192, 192, 192, 255}, color.Black}
	case rating > 0:
		return ratingColors{color.NRGBA{192, 105, 105, 255}, color.Black}
	default:
		return ratingColors{color.Transparent, color.Transparent}
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
