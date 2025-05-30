//go:build ignore

package main

import (
	"fmt"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/cufee/aftermath/internal/log"
	common "github.com/cufee/aftermath/internal/render/v1"
	"github.com/fogleman/gg"
	"github.com/nao1215/imaging"

	ico "github.com/Kodeworks/golang-image-ico"
	"github.com/joho/godotenv"
)

var outDirPath = "../../public"
var brandColor color.NRGBA

var cardColor = color.NRGBA{7, 7, 7, 200}

func main() {
	godotenv.Load("../../../../.env")

	switch os.Getenv("BRAND_FLAVOR") {
	default:
		brandColor = common.ColorAftermathRed
	case "blue":
		brandColor = common.ColorAftermathBlue
	case "gold":
		brandColor = common.TextSubscriptionPremium
	}

	generateWN8Icons()
	generateRatingIcons()
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
			filename := wn8IconFilename(float32(tier))
			img := common.AftermathLogo(color, common.LargeLogoOptions())
			f, err := os.Create(filepath.Join(outDirPath, "wn8", filename))
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
			filename := "small_" + wn8IconFilename(float32(tier))
			img := common.AftermathLogo(color, common.SmallLogoOptions())
			f, err := os.Create(filepath.Join(outDirPath, "wn8", filename))
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
		opts := common.LargeLogoOptions()
		img := common.AftermathLogo(brandColor, opts)
		{
			filename := fmt.Sprintf("icon/%d.png", size)
			f, err := os.Create(filepath.Join(outDirPath, filename))
			if err != nil {
				panic(err)
			}
			err = png.Encode(f, imaging.Fit(img, size, size, imaging.Linear))
			if err != nil {
				panic(err)
			}
			f.Close()
		}

		if size == 16 {
			f, err := os.Create(filepath.Join(outDirPath, "favicon.ico"))
			if err != nil {
				panic(err)
			}
			err = ico.Encode(f, imaging.Fit(img, size, size, imaging.Linear))
			if err != nil {
				panic(err)
			}
			f.Close()
		}
	}
}

func generateOGImages() {
	log.Debug().Msg("generating og images")

	imageWidth := 512
	imageHeight := imageWidth * 2 / 4
	logoSize := imageHeight * 1 / 2
	borderWidth := 2

	bg, err := imaging.Open("./bg-default.jpg")
	if err != nil {
		panic(err)
	}

	{
		filename := "og/widget.jpg"
		opts := common.LargeLogoOptions()
		logo := common.AftermathLogo(brandColor, opts)
		ctx := gg.NewContext(imageWidth, imageHeight)

		obsBg, err := imaging.Open("./obs-splash.png")
		if err != nil {
			panic(err)
		}

		ctx.DrawImage(imaging.Fill(obsBg, imageWidth, imageHeight, imaging.Center, imaging.Lanczos), 0, 0)
		ctx.DrawRectangle(0, 0, float64(imageWidth), float64(imageHeight))
		ctx.SetColor(color.NRGBA{7, 7, 7, 200})
		ctx.Fill()

		ctx.DrawImageAnchored(imaging.Fit(logo, logoSize, logoSize, imaging.Linear), imageWidth/2, imageHeight/2, 0.5, 0.5)

		f, err := os.Create(filepath.Join(outDirPath, filename))
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
		filename := "og/default.jpg"
		opts := common.LargeLogoOptions()
		logo := common.AftermathLogo(brandColor, opts)
		ctx := gg.NewContext(imageWidth, imageHeight)

		ctx.DrawImage(imaging.Blur(imaging.Fill(bg, imageWidth, imageHeight, imaging.Center, imaging.Lanczos), 30), 0, 0)
		ctx.DrawRectangle(0, 0, float64(imageWidth), float64(imageHeight))
		ctx.SetColor(cardColor)
		ctx.Fill()

		ctx.DrawImageAnchored(imaging.Fit(logo, logoSize, logoSize, imaging.Linear), imageWidth/2, imageHeight/2, 0.5, 0.5)

		f, err := os.Create(filepath.Join(outDirPath, filename))
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
		filename := "og/verify.jpg"
		opts := common.LargeLogoOptions()
		logo := common.AftermathLogo(brandColor, opts)
		ctx := gg.NewContext(imageWidth, imageHeight)

		// each logo should take 1/3 of the total height available
		heightAvailable := imageHeight - borderWidth*2
		widthAvailable := imageWidth - borderWidth*2
		singleLogoSize := heightAvailable / 3
		// link icon should be centered, other icons should be moved a specific distance away from it
		linkIconSize := 32 // px
		imageGap := 32     // px
		padding := (widthAvailable-linkIconSize-singleLogoSize*2-imageGap*2)/2 + borderWidth*2
		// we can now draw padding - logo - gap - link - gap - logo

		linkIcon, err := imaging.Open("./link.png")
		if err != nil {
			panic(err)
		}
		blitzLogo, err := imaging.Open("./blitz-logo.png")
		if err != nil {
			panic(err)
		}
		linkBlock := common.NewImageContent(common.Style{Width: float64(linkIconSize), Height: float64(linkIconSize), BackgroundColor: common.TextAlt}, linkIcon)
		linkColored, err := linkBlock.Render()
		if err != nil {
			panic(err)
		}

		ctx.DrawImage(imaging.Blur(imaging.Fill(bg, imageWidth, imageHeight, imaging.Center, imaging.Lanczos), 30), 0, 0)
		ctx.DrawRectangle(0, 0, float64(imageWidth), float64(imageHeight))
		ctx.SetColor(cardColor)
		ctx.Fill()

		ctx.DrawImage(imaging.Fill(logo, singleLogoSize, singleLogoSize, imaging.Center, imaging.Lanczos), padding, imageHeight/2-singleLogoSize/2)
		ctx.DrawImage(imaging.Fill(linkColored, linkIconSize, linkIconSize, imaging.Center, imaging.Lanczos), padding+singleLogoSize+imageGap, imageHeight/2-linkIconSize/2)
		ctx.DrawImage(imaging.Fill(blitzLogo, singleLogoSize, singleLogoSize, imaging.Center, imaging.Lanczos), padding+singleLogoSize+imageGap+linkIconSize+imageGap, imageHeight/2-singleLogoSize/2)

		f, err := os.Create(filepath.Join(outDirPath, filename))
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

func wn8IconFilename(rating float32) string {
	name := strings.ReplaceAll(strings.ToLower(common.GetWN8TierName(rating)), " ", "_")
	if rating < 1 {
		name = "invalid"
	}
	return name + ".png"
}

func generateRatingIcons() {
	log.Debug().Msg("generating rating image assets")

	for name, icon := range common.RatingIconSettings {
		block, ok := common.RenderRatingIcon(icon)
		if !ok {
			panic("failed to render rating icon " + name)
		}

		img, err := block.Render()
		if err != nil {
			panic(err)
		}

		f, err := os.Create(filepath.Join(outDirPath, "rating", name+".png"))
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
