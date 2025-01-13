//go:build ignore

package main

import (
	"fmt"
	"image/color"
	"image/png"
	"math"
	"os"
	"path/filepath"

	"github.com/cufee/aftermath/internal/localization"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/render/v1"
	common "github.com/cufee/aftermath/internal/render/v1"
	"github.com/fogleman/gg"
	"github.com/nao1215/imaging"
	"golang.org/x/text/language"

	"github.com/joho/godotenv"
)

var outDirPath = "../../static"
var brandColor color.NRGBA

var discordColorLight = color.NRGBA{54, 56, 61, 255}
var discordColorMedium = color.NRGBA{46, 47, 52, 255}
var discordColorDark = color.NRGBA{30, 31, 34, 255}
var discordColorText = color.NRGBA{151, 155, 162, 255}

func main() {
	godotenv.Load("../../.env")

	switch os.Getenv("BRAND_FLAVOR") {
	default:
		brandColor = common.ColorAftermathRed
	case "blue":
		brandColor = common.ColorAftermathBlue
	case "gold":
		brandColor = common.ColorAftermathYellow
	}

	localization.LoadAssets(os.DirFS("../../static/localization"), ".")
	printer, _ := localization.NewPrinterWithFallback("discord", language.English)

	generateDiscordHelpImage(printer)

	generateDiscordLogo("red", common.ColorAftermathRed)
	generateDiscordLogo("blue", common.ColorAftermathBlue)
	generateDiscordLogo("yellow", common.ColorAftermathYellow)

	generateRatingIcons()
}

func generateDiscordHelpImage(printer func(string) string) {
	log.Debug().Msg("generating discord help image")

	imageWidth := 550
	imageHeight := imageWidth * 3 / 5
	inputHeight := 50
	iconSize := 25
	padding := 15
	gap := 5

	{
		filename := "images/discord/discord-help.png"

		ctx := gg.NewContext(imageWidth, imageHeight)
		ctx.SetColor(discordColorMedium)
		ctx.Clear()

		{ // draw Discord UI
			{
				// text input
				tctx := gg.NewContext(imageWidth-padding*2, inputHeight)
				tctx.DrawRoundedRectangle(0, 0, float64(tctx.Width()), float64(tctx.Height()), common.BorderRadiusXS)
				tctx.SetColor(discordColorLight)
				tctx.Fill()

				input := ctx.Image()
				ctx.DrawImage(input, padding, imageHeight-padding-tctx.Height())

				sctx := gg.NewContext(imageWidth, tctx.Height()+padding*2)
				sctx.DrawRoundedRectangle(float64(padding), float64(padding), float64(tctx.Width()), float64(tctx.Height()), common.BorderRadiusXS)
				sctx.SetColor(color.NRGBA{0, 0, 0, 100})
				sctx.Fill()

				ctx.DrawImage(imaging.Blur(sctx.Image(), 2), 0, imageHeight-sctx.Height())
				ctx.DrawImage(tctx.Image(), padding, imageHeight-padding-tctx.Height())

				circleIconR := float64(iconSize / 2)
				circleIconX := float64(padding + 50 - (inputHeight / 2))
				circleIconY := float64(imageHeight - padding - (inputHeight / 2))

				ctx.DrawCircle(circleIconX, circleIconY, circleIconR)
				ctx.SetColor(discordColorText)
				ctx.Fill()

				ctx.SetLineWidth(2)
				ctx.SetLineCapRound()
				ctx.SetColor(discordColorMedium)
				ctx.DrawLine(circleIconX, circleIconY-circleIconR+7, circleIconX, circleIconY+circleIconR-7)
				ctx.DrawLine(circleIconX-circleIconR+7, circleIconY, circleIconX+circleIconR-7, circleIconY)
				ctx.Stroke()

				ctx.SetColor(discordColorText)
				ctx.LoadFontFace("../../static/fonts/default.ttf", 30)
				ctx.DrawString("/", float64(padding+inputHeight), float64(circleIconY+10))
			}

			// commands drawer
			{
				dctx := gg.NewContext(imageWidth-padding*2, imageHeight-padding*2-gap-inputHeight)

				dctx.DrawRoundedRectangle(0, 0, float64(dctx.Width()), float64(dctx.Height()), common.BorderRadiusXS)
				dctx.Clip()

				dctx.DrawRectangle(0, 0, float64(dctx.Width()), float64(dctx.Height()))
				dctx.SetColor(discordColorDark)
				dctx.Fill()

				dctx.DrawRectangle(float64(inputHeight), 0, float64(dctx.Width()-inputHeight), float64(dctx.Height()))
				dctx.SetColor(discordColorMedium)
				dctx.Fill()

				fontSize := 20.0
				commands := []string{"help", "links", "stats", "session"}
				for i, name := range commands {
					drawY := float64((padding + int(fontSize)) + i*(int(fontSize*2)+padding))
					dctx.SetColor(color.White)
					dctx.LoadFontFace("../../static/fonts/default.ttf", fontSize)
					dctx.DrawString("/"+printer(fmt.Sprintf("command_%s_name", name)), float64(padding+inputHeight+gap), drawY)

					dctx.LoadFontFace("../../static/fonts/default.ttf", fontSize*0.8)
					dctx.SetColor(discordColorText)
					dctx.DrawString(printer(fmt.Sprintf("command_%s_description", name)), float64(padding+inputHeight+gap), drawY+fontSize+float64(gap/2))
					dctx.DrawString("Aftermath", float64(dctx.Width()-padding-gap*2-70), drawY+fontSize/2+float64(gap/2))
				}

				sctx := gg.NewContext(imageWidth,
					dctx.Height()+padding*2)
				sctx.DrawRoundedRectangle(float64(padding), float64(padding), float64(dctx.Width()), float64(dctx.Height()), common.BorderRadiusXS)
				sctx.SetColor(color.NRGBA{0, 0, 0, 100})
				sctx.Fill()

				ctx.DrawImage(imaging.Blur(sctx.Image(), 5), 0, 0)
				ctx.DrawImage(dctx.Image(), padding, padding)

			}

			// commands
			{

			}
		}

		logo := common.AftermathLogo(brandColor, common.DefaultLogoOptions())
		ctx.DrawImage(imaging.Fit(logo, iconSize, iconSize, imaging.Linear), padding+((inputHeight-iconSize)/2), padding*2)

		f, err := os.Create(filepath.Join(outDirPath, filename))
		if err != nil {
			panic(err)
		}

		err = imaging.Encode(f, ctx.Image(), imaging.PNG)
		if err != nil {
			panic(err)
		}
		f.Close()
	}

}

type point struct {
	x int
	y int
}

func generateDiscordLogo(suffix string, logoColor color.Color) {
	log.Debug().Msg("generating discord logo image")
	{
		filename := "images/discord/logo_" + suffix + ".png"

		opts := common.LargeLogoOptions()
		padding := 80
		img := imaging.Fit(common.AftermathLogo(logoColor, opts), 256, 256, imaging.Linear)
		nctx := gg.NewContext(256+padding, 256+padding)
		nctx.SetColor(color.NRGBA{30, 31, 34, 255})
		nctx.Clear()
		nctx.DrawImage(img, padding/2+(256-img.Bounds().Dx())/2, padding/6+(256-img.Bounds().Dy())/2)

		f, err := os.Create(filepath.Join(outDirPath, filename))
		if err != nil {
			panic(err)
		}
		err = png.Encode(f, nctx.Image())
		if err != nil {
			panic(err)
		}
		f.Close()
	}
	{
		filename := "images/discord/logo_" + suffix + "_centered.png"

		opts := common.LargeLogoOptions()
		padding := 80
		img := imaging.Fit(common.AftermathLogo(logoColor, opts), 256, 256, imaging.Linear)
		nctx := gg.NewContext(256+padding, 256+padding)
		nctx.SetColor(color.NRGBA{30, 31, 34, 255})
		nctx.Clear()
		nctx.DrawImageAnchored(img, nctx.Width()/2, nctx.Height()/2, 0.5, 0.5)

		f, err := os.Create(filepath.Join(outDirPath, filename))
		if err != nil {
			panic(err)
		}
		err = png.Encode(f, nctx.Image())
		if err != nil {
			panic(err)
		}
		f.Close()
	}
	{
		filename := "images/discord/logo_" + suffix + "_centered_alpha.png"

		opts := common.LargeLogoOptions()
		padding := 80
		img := imaging.Fit(common.AftermathLogo(logoColor, opts), 256, 256, imaging.Linear)
		nctx := gg.NewContext(256+padding, 256+padding)
		nctx.DrawImageAnchored(img, nctx.Width()/2, nctx.Height()/2, 0.5, 0.5)

		f, err := os.Create(filepath.Join(outDirPath, filename))
		if err != nil {
			panic(err)
		}
		err = png.Encode(f, nctx.Image())
		if err != nil {
			panic(err)
		}
		f.Close()
	}
}

type ratingIcon struct {
	name  string
	color color.Color
	fill  [][]int
}

var ratingIcons = []ratingIcon{
	{name: "calibration", color: common.TextAlt, fill: [][]int{{0, 0, 1, 0, 0}, {0, 1, 1, 0, 1, 1, 0}, {1, 1, 1, 1, 0, 1, 1, 1, 1}, {0, 1, 1, 0, 1, 1, 0}, {0, 0, 1, 0, 0}}},
	{name: "bronze", color: render.GetRatingColors(1).Background, fill: [][]int{{0, 0, 0, 0, 0}, {0, 0, 1, 1, 1, 0, 0}, {0, 0, 1, 1, 1, 1, 1, 0, 0}, {0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0}}},
	{name: "silver", color: render.GetRatingColors(2001).Background, fill: [][]int{{0, 0, 0, 0, 0}, {0, 0, 1, 1, 1, 0, 0}, {0, 0, 1, 1, 0, 1, 1, 0, 0}, {0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0}}},
	{name: "gold", color: render.GetRatingColors(3001).Background, fill: [][]int{{0, 0, 0, 0, 0}, {0, 1, 1, 1, 1, 1, 0}, {0, 1, 1, 0, 1, 0, 1, 1, 0}, {0, 1, 1, 1, 1, 1, 0}, {0, 0, 0, 0, 0}}},
	{name: "platinum", color: render.GetRatingColors(4001).Background, fill: [][]int{{0, 0, 1, 0, 0}, {0, 1, 1, 1, 1, 1, 0}, {1, 1, 1, 1, 0, 1, 1, 1, 1}, {0, 1, 1, 1, 1, 1, 0}, {0, 0, 1, 0, 0}}},
	{name: "diamond", color: render.GetRatingColors(5001).Background, fill: [][]int{{0, 0, 1, 0, 0}, {0, 1, 1, 0, 1, 1, 0}, {1, 1, 1, 1, 0, 1, 1, 1, 1}, {0, 1, 1, 0, 1, 1, 0}, {0, 0, 1, 0, 0}}},
}

func generateRatingIcons() {
	log.Debug().Msg("generating rating icons")

	var ratingIconLineWidth = 8
	var ratingIconBackgroundColor = color.Transparent

	for _, icon := range ratingIcons {
		filename := "images/game/rating-" + icon.name + ".png"

		centerIndex := len(icon.fill) / 2
		iconWidth := (centerIndex * 2) * (ratingIconLineWidth * 2)
		iconHeight := 0
		for _, items := range icon.fill {
			iconHeight = max(iconHeight, len(items)*(ratingIconLineWidth))
		}

		ctx := gg.NewContext(iconWidth, iconHeight)
		offsetX := ratingIconLineWidth / 2
		for _, items := range icon.fill {
			colHeight := len(items) * (ratingIconLineWidth)
			offsetY := (iconHeight - colHeight) / 2
			ctx.DrawRoundedRectangle(float64(offsetX), float64(offsetY), float64(ratingIconLineWidth), float64(colHeight), (float64(ratingIconLineWidth)/2)-1)
			ctx.SetColor(ratingIconBackgroundColor)
			ctx.Fill()

			ctx.SetColor(icon.color)
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

		f, err := os.Create(filepath.Join(outDirPath, filename))
		if err != nil {
			panic(err)
		}
		err = png.Encode(f, ctx.Image())
		if err != nil {
			panic(err)
		}
		f.Close()

	}
}
