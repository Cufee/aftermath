//go:build ignore

package main

import (
	"fmt"
	"image/png"
	"os"
	"path/filepath"

	"github.com/cufee/aftermath/cmd/frontend/assets"
	"github.com/cufee/aftermath/internal/stats/render/common/v1"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/rs/zerolog/log"
)

func main() {
	generateWN8Icons()
	generateLogoOptions()
	generateOGImages()
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

func generateLogoOptions() {
	log.Debug().Msg("generating logo options")

	for _, size := range []int{16, 32, 64, 128, 256, 512} {
		filename := fmt.Sprintf("icon-%d.png", size)

		opts := common.DefaultLogoOptions()
		opts.Gap *= 10
		opts.Jump *= 10
		opts.LineStep *= 10
		opts.LineWidth *= 10

		img := common.AftermathLogo(common.ColorAftermathRed, opts)
		f, err := os.Create(filepath.Join("../public", filename))
		if err != nil {
			panic(err)
		}
		err = png.Encode(f, imaging.Fit(img, size, size, imaging.Linear))
		if err != nil {
			panic(err)
		}
		f.Close()
	}
}

func generateOGImages() {
	log.Debug().Msg("generating og images")

	imageWidth := 512
	imageHeight := imageWidth * 2 / 3
	logoSize := imageHeight * 2 / 3
	borderWidth := 2

	{
		filename := "og-widget.jpg"
		opts := common.DefaultLogoOptions()
		opts.Gap *= 10
		opts.Jump *= 10
		opts.LineStep *= 10
		opts.LineWidth *= 10

		logo := common.AftermathLogo(common.ColorAftermathRed, opts)
		ctx := gg.NewContext(imageWidth, imageHeight)

		bg, err := imaging.Open("./bg-default.jpg")
		if err != nil {
			panic(err)
		}

		ctx.DrawImage(imaging.Blur(imaging.Fill(bg, imageWidth, imageHeight, imaging.Center, imaging.Lanczos), 30), 0, 0)
		ctx.DrawRoundedRectangle(float64(borderWidth), float64(borderWidth), float64(imageWidth-borderWidth*2), float64(imageHeight-borderWidth*2), 20)
		ctx.SetColor(common.DefaultCardColorNoAlpha)
		ctx.Fill()

		ctx.DrawImageAnchored(imaging.Fit(logo, logoSize, logoSize, imaging.Linear), imageWidth/2, imageHeight/2, 0.5, 0.5)

		f, err := os.Create(filepath.Join("../public", filename))
		if err != nil {
			panic(err)
		}

		err = imaging.Encode(f, ctx.Image(), imaging.JPEG)
		if err != nil {
			panic(err)
		}
		f.Close()
	}

	{
		filename := "og.jpg"
		opts := common.DefaultLogoOptions()
		opts.Gap *= 10
		opts.Jump *= 10
		opts.LineStep *= 10
		opts.LineWidth *= 10

		logo := common.AftermathLogo(common.ColorAftermathRed, opts)
		ctx := gg.NewContext(imageWidth, imageHeight)

		bg, err := imaging.Open("./bg-default.jpg")
		if err != nil {
			panic(err)
		}

		ctx.DrawImage(imaging.Blur(imaging.Fill(bg, imageWidth, imageHeight, imaging.Center, imaging.Lanczos), 30), 0, 0)
		ctx.DrawRoundedRectangle(float64(borderWidth), float64(borderWidth), float64(imageWidth-borderWidth*2), float64(imageHeight-borderWidth*2), 20)
		ctx.SetColor(common.DefaultCardColorNoAlpha)
		ctx.Fill()

		ctx.DrawImageAnchored(imaging.Fit(logo, logoSize, logoSize, imaging.Linear), imageWidth/2, imageHeight/2, 0.5, 0.5)

		f, err := os.Create(filepath.Join("../public", filename))
		if err != nil {
			panic(err)
		}

		err = imaging.Encode(f, ctx.Image(), imaging.JPEG)
		if err != nil {
			panic(err)
		}
		f.Close()
	}
}
