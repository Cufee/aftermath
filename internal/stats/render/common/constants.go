package common

import (
	"image/color"

	"github.com/pkg/errors"

	"github.com/cufee/aftermath/internal/stats/render/assets"
	"golang.org/x/image/font"
)

var DiscordBackgroundColor = color.RGBA{49, 51, 56, 255}

var (
	FontXL     font.Face
	Font2XL    font.Face
	FontLarge  font.Face
	FontMedium font.Face
	FontSmall  font.Face

	TextPrimary   = color.RGBA{255, 255, 255, 255}
	TextSecondary = color.RGBA{204, 204, 204, 255}
	TextAlt       = color.RGBA{150, 150, 150, 255}

	TextSubscriptionPlus    = color.RGBA{72, 167, 250, 255}
	TextSubscriptionPremium = color.RGBA{255, 223, 0, 255}

	DefaultCardColor        = color.RGBA{10, 10, 10, 180}
	DefaultCardColorNoAlpha = color.RGBA{10, 10, 10, 255}

	ColorAftermathRed  = color.RGBA{255, 0, 120, 255}
	ColorAftermathBlue = color.RGBA{90, 90, 255, 255}
)

var fontCache map[float64]font.Face

func InitLoadedAssets() error {
	var ok bool
	fontCache, ok = assets.GetLoadedFontFaces("default", 36, 32, 24, 18, 14)
	if !ok {
		return errors.New("default font not found")
	}
	FontXL = fontCache[32]
	Font2XL = fontCache[36]
	FontLarge = fontCache[24]
	FontMedium = fontCache[18]
	FontSmall = fontCache[14]

	return nil
}

func GetCustomFont(size float64) (font.Face, bool) {
	if f, ok := fontCache[size]; ok {
		return f, true
	}

	newCache, ok := assets.GetLoadedFontFaces("default", size)
	if !ok {
		return nil, false
	}
	for size, font := range newCache {
		fontCache[size] = font
	}

	return newCache[size], true
}
