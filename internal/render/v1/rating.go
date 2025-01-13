package render

import (
	"image/color"
	"math"

	"github.com/cufee/aftermath/internal/render/assets"
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/fogleman/gg"
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
		return ratingColors{color.NRGBA{181, 106, 181, 255}, color.White}
	case rating > 4000:
		return ratingColors{color.NRGBA{154, 197, 219, 255}, color.Black}
	case rating > 3000:
		return ratingColors{color.NRGBA{255, 215, 0, 255}, color.Black}
	case rating > 2000:
		return ratingColors{color.NRGBA{192, 192, 192, 255}, color.Black}
	case rating > 0:
		return ratingColors{color.NRGBA{192, 105, 105, 255}, color.White}
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

type ratingIcon struct {
	Name  string
	Color color.Color
	Fill  [][]int
}

var RatingIconSettings = map[string]ratingIcon{
	"bronze": {Name: "bronze", Color: GetRatingColors(1).Background, Fill: [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}},
	"silver": {Name: "silver", Color: GetRatingColors(2001).Background, Fill: [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}},
	"gold": {Name: "gold", Color: GetRatingColors(3001).Background, Fill: [][]int{
		{0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0},
		{1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1},
		{0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
	}},
	"platinum": {Name: "platinum", Color: GetRatingColors(4001).Background, Fill: [][]int{
		{0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
	}},
	"diamond": {Name: "diamond", Color: GetRatingColors(5001).Background, Fill: [][]int{
		{0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 0, 1, 1, 1, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 1, 1, 1, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
	}},
}

func init() {
	RatingIconSettings["calibration"] = ratingIcon{
		Color: TextAlt,
		Name:  "calibration",
		Fill:  RatingIconSettings["diamond"].Fill,
	}
}

func NewRatingIcon(rating frame.Value) (Block, bool) {
	settings, ok := RatingIconSettings[GetRatingTierName(rating.Float())]
	if !ok {
		return NewEmptyContent(1, 1), false
	}
	return RenderRatingIcon(settings)
}

func RenderRatingIcon(settings ratingIcon) (Block, bool) {
	var ratingIconLineWidth = 8
	var ratingIconBackgroundColor = color.Transparent

	centerIndex := len(settings.Fill) / 2
	iconHeight := 0
	for _, items := range settings.Fill {
		iconHeight = max(iconHeight, len(items)*(ratingIconLineWidth))
	}
	iconWidth := (centerIndex*2)*(ratingIconLineWidth*2) - (len(settings.Fill)%2)*ratingIconLineWidth

	ctx := gg.NewContext(iconWidth, iconHeight)
	offsetX := 0
	for _, items := range settings.Fill {
		colHeight := len(items) * (ratingIconLineWidth)
		offsetY := (iconHeight - colHeight) / 2
		ctx.DrawRoundedRectangle(float64(offsetX), float64(offsetY), float64(ratingIconLineWidth), float64(colHeight), (float64(ratingIconLineWidth)/2)-1)
		ctx.SetColor(ratingIconBackgroundColor)
		ctx.Fill()

		ctx.SetColor(settings.Color)
		for i, section := range items {
			sectionOffsetY := float64(offsetY + (i * ratingIconLineWidth))
			if section == 0 {
				continue
			}

			var topRounded bool = true
			var bottomRounded bool = true
			if i-1 >= 0 {
				topRounded = items[i-1] == 0
			}
			if i+1 < len(items) {
				bottomRounded = items[i+1] == 0
			}

			positionY := sectionOffsetY

			if !topRounded && !bottomRounded {
				ctx.DrawRectangle(float64(offsetX), positionY, float64(ratingIconLineWidth), float64(ratingIconLineWidth))
				ctx.Fill()
			}

			// draw top part
			if topRounded {
				ctx.DrawArc(float64(offsetX)+float64(ratingIconLineWidth/2), positionY+float64(ratingIconLineWidth)/2, float64(ratingIconLineWidth)/2, -math.Pi, 0)
				ctx.Fill()
			} else {
				ctx.DrawRectangle(float64(offsetX), positionY, float64(ratingIconLineWidth), float64(ratingIconLineWidth)/2)
				ctx.Fill()
			}

			// draw bottom part
			if bottomRounded {

				ctx.DrawArc(float64(offsetX)+float64(ratingIconLineWidth/2), positionY+float64(ratingIconLineWidth/2), float64(ratingIconLineWidth)/2, math.Pi, 0)
				ctx.Fill()
			} else {
				ctx.DrawRectangle(float64(offsetX), positionY+float64(ratingIconLineWidth/2), float64(ratingIconLineWidth), float64(ratingIconLineWidth)/2)
				ctx.Fill()
			}
		}

		offsetX += ratingIconLineWidth * 3 / 2
	}

	return NewImageContent(Style{}, ctx.Image()), true
}
