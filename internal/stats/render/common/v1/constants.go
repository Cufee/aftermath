package common

import (
	"image/color"

	"github.com/golang/freetype/truetype"
	"github.com/pkg/errors"
	"golang.org/x/image/font"

	"github.com/cufee/aftermath/internal/stats/render/assets"
)

var DiscordBackgroundColor = color.RGBA{49, 51, 56, 255}

var (
	FontXL     Font
	Font2XL    Font
	FontLarge  Font
	FontMedium Font
	FontSmall  Font

	TextPrimary   = color.RGBA{255, 255, 255, 255}
	TextSecondary = color.RGBA{204, 204, 204, 255}
	TextAlt       = color.RGBA{150, 150, 150, 255}

	TextSubscriptionPlus    = color.RGBA{72, 167, 250, 255}
	TextSubscriptionPremium = color.RGBA{255, 223, 0, 255}

	DefaultCardColor        = color.RGBA{10, 10, 10, 180}
	DefaultCardColorNoAlpha = color.RGBA{10, 10, 10, 255}

	ColorAftermathRed  = color.RGBA{255, 0, 120, 255}
	ColorAftermathBlue = color.RGBA{90, 90, 255, 255}

	BorderRadiusLG = 25.0
	BorderRadiusMD = 20.0
	BorderRadiusSM = 15.0
	BorderRadiusXS = 10.0
)

type Font struct {
	size float64
	data []byte
}

func (f *Font) Valid() bool {
	return f.data != nil && f.size > 0
}

func (f *Font) Face() font.Face {
	ttf, _ := truetype.Parse(f.data)
	return truetype.NewFace(ttf, &truetype.Options{
		Size: f.size,
	})
}

// var fontCache map[float64]font.Face

func InitLoadedAssets() error {
	// var ok bool
	// fontCache, ok = assets.GetLoadedFontFaces("default", 36, 32, 24, 18, 14)
	fontData, ok := assets.GetLoadedFontFace("default")
	if !ok {
		return errors.New("default font not found")
	}

	FontXL = Font{32, fontData}
	Font2XL = Font{36, fontData}

	FontLarge = Font{24, fontData}
	FontMedium = Font{18, fontData}
	FontSmall = Font{14, fontData}

	return nil
}
