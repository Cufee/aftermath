package render

import (
	"image/color"
	"math"
	"strings"

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
	Fill  string
}

var RatingIconSettings = map[string]ratingIcon{
	"bronze": {Name: "bronze", Color: GetRatingColors(1).Background, Fill: `
_______
_______
___x___
___x___
__xxx__
__xxx__
__xxx__
__xxx__
___x___
___x___
_______
_______
`},
	"silver": {Name: "silver", Color: GetRatingColors(2001).Background, Fill: `
_______
_______
___x___
___x___
_x_x_x_
_x___x_
_xx_xx_
__xxx__
__xxx__
___x___
___x___
_______
`},
	"gold": {Name: "gold", Color: GetRatingColors(3001).Background, Fill: `
_______
___x___
__xxx__
_xxxxx_
_x_x_x_
_x___x_
_xx_xx_
__xxx__
__xxx__
___x___
___x___
_______
`},
	"platinum": {Name: "platinum", Color: GetRatingColors(4001).Background, Fill: `
_______
___x___
x_xxx_x
xxxxxxx
xx_x_xx
_x___x_
_xx_xx_
__xxx__
__xxx__
___x___
___x___
_______
`},
	"diamond": {Name: "diamond", Color: GetRatingColors(5001).Background, Fill: `
___x___
x__x__x
x_xxx_x
xxxxxxx
xx_x_xx
_x___x_
_xx_xx_
__xxx__
__xxx__
___x___
___x___
_______
`},
}

func init() {
	RatingIconSettings["calibration"] = ratingIcon{
		Color: TextAlt,
		Name:  "calibration",
		Fill:  RatingIconSettings["bronze"].Fill,
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

	rows := strings.Split(strings.TrimSpace(settings.Fill), "\n")

	iconHeight := len(rows) * (ratingIconLineWidth)
	iconWidth := (len(rows[0]) * ratingIconLineWidth) + ((len(rows[0]) - 1) * (ratingIconLineWidth / 2))

	ctx := gg.NewContext(iconWidth, iconHeight)
	// ctx.SetColor(color.Black)
	// ctx.Clear()
	ctx.SetColor(settings.Color)

	for rowI, row := range rows {
		positionY := float64(rowI * ratingIconLineWidth)
		for itemI, item := range strings.Split(row, "") {
			if strings.ToLower(item) != "x" {
				continue
			}

			positionX := itemI*ratingIconLineWidth + max(0, (itemI)*(ratingIconLineWidth/2))

			var topRounded bool = true
			var bottomRounded bool = true
			if rowI-1 >= 0 {
				topRounded = strings.ToLower(strings.Split(rows[rowI-1], "")[itemI]) != "x"
			}
			if rowI+1 < len(rows) {
				bottomRounded = strings.ToLower(strings.Split(rows[rowI+1], "")[itemI]) != "x"
			}

			if !topRounded && !bottomRounded {
				ctx.DrawRectangle(float64(positionX), positionY, float64(ratingIconLineWidth), float64(ratingIconLineWidth))
				ctx.Fill()
			}

			// draw top part
			if topRounded {
				ctx.DrawArc(float64(positionX)+float64(ratingIconLineWidth/2), positionY+float64(ratingIconLineWidth)/2, float64(ratingIconLineWidth)/2, -math.Pi, 0)
				ctx.Fill()
			} else {
				ctx.DrawRectangle(float64(positionX), positionY, float64(ratingIconLineWidth), float64(ratingIconLineWidth)/2)
				ctx.Fill()
			}

			// draw bottom part
			if bottomRounded {

				ctx.DrawArc(float64(positionX)+float64(ratingIconLineWidth/2), positionY+float64(ratingIconLineWidth/2), float64(ratingIconLineWidth)/2, math.Pi, 0)
				ctx.Fill()
			} else {
				ctx.DrawRectangle(float64(positionX), positionY+float64(ratingIconLineWidth/2), float64(ratingIconLineWidth), float64(ratingIconLineWidth)/2)
				ctx.Fill()
			}
		}
	}

	return NewImageContent(Style{}, ctx.Image()), true
}
