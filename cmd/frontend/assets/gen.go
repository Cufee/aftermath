//go:build ignore

package main

import (
	"image/png"
	"os"
	"path/filepath"

	"github.com/cufee/aftermath/cmd/frontend/assets"
	"github.com/cufee/aftermath/internal/stats/render/common/v1"
	"github.com/rs/zerolog/log"
)

func main() {
	generateWN8Icons()
}

var wn8Tiers = []int{0, 1, 301, 451, 651, 901, 1201, 1601, 2001, 2451, 2901}

func generateWN8Icons() {
	log.Debug().Msg("generating wn8 image assets")

	for _, tier := range wn8Tiers {
		color := common.GetWN8Colors(float32(tier)).Background
		if tier < 1 {
			color = common.TextAlt
		}
		{
			filename := assets.WN8IconFilename(float32(tier))
			img := common.AftermathLogo(color, common.DefaultLogoOptions())
			f, err := os.Create(filepath.Join("../public", "wn8", filename))
			if err != nil {
				panic(err)
			}
			err = png.Encode(f, img)
			if err != nil {
				panic(err)
			}
			f.Close()
		}
		{
			filename := "small_" + assets.WN8IconFilename(float32(tier))
			img := common.AftermathLogo(color, common.SmallLogoOptions())
			f, err := os.Create(filepath.Join("../public", "wn8", filename))
			if err != nil {
				panic(err)
			}
			err = png.Encode(f, img)
			if err != nil {
				panic(err)
			}
			f.Close()
		}
	}
}
