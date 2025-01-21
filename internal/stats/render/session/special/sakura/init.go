package sakura

import (
	"image"

	"github.com/cufee/aftermath/internal/stats/render/session/special/sakura/assets"
	"github.com/cufee/facepaint/style"
	"github.com/pkg/errors"
)

var initComplete bool

// lazy load assets, as this package is not guaranteed to be used
func lazyLoadAssets() error {
	if !initComplete {
		err := assets.Load()
		if err != nil {
			return err
		}
		initComplete = true
	}

	ffs, ok := assets.FontFace("zyukiharu")
	if !ok {
		return errors.New("missing ungai font")
	}
	fontData = ffs

	backgroundImage, ok = assets.Image("background")
	if !ok {
		return errors.New("missing background image")
	}

	return nil
}

var backgroundImage image.Image

var fontData []byte
var fonts = make(map[float64]style.Font)

func Font2XL() style.Font {
	return getFont(36)
}
func FontXL() style.Font {
	return getFont(32)
}
func FontLarge() style.Font {
	return getFont(24)
}
func FontMedium() style.Font {
	return getFont(18)
}
func FontSmall() style.Font {
	return getFont(14)
}

func getFont(size float64) style.Font {
	if fonts[size] == nil {
		fonts[size], _ = style.NewFont(fontData, size)
	}
	return fonts[size]
}
