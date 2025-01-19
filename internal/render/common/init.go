package common

import (
	"errors"

	"github.com/cufee/aftermath/internal/render/assets"
)

func InitLoadedAssets() error {
	fontData, ok := assets.GetLoadedFontFace("default")
	if !ok {
		return errors.New("default font not found")
	}
	defaultFont = fontData

	return nil
}
