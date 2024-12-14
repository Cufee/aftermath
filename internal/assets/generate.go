//go:build ignore

package main

import (
	"fmt"
	"image/color"
	"image/png"
	"os"
	"path/filepath"

	"github.com/cufee/aftermath/internal/localization"
	common "github.com/cufee/aftermath/internal/render/v1"
	"github.com/fogleman/gg"
	"github.com/nao1215/imaging"
	"github.com/rs/zerolog/log"
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
		brandColor = common.TextSubscriptionPremium
	}

	localization.LoadAssets(os.DirFS("../../static/localization"), ".")
	printer, _ := localization.NewPrinterWithFallback("discord", language.English)

	generateDiscordHelpImage(printer)
	generateDiscordLogo()
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

func generateDiscordLogo() {
	log.Debug().Msg("generating discord logo image")
	{
		filename := "images/discord/logo.png"

		opts := common.LargeLogoOptions()
		padding := 80
		img := imaging.Fit(common.AftermathLogo(brandColor, opts), 256, 256, imaging.Linear)
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
		filename := "images/discord/logo_centered.png"

		opts := common.LargeLogoOptions()
		padding := 80
		img := imaging.Fit(common.AftermathLogo(brandColor, opts), 256, 256, imaging.Linear)
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
		filename := "images/discord/logo_centered_alpha.png"

		opts := common.LargeLogoOptions()
		padding := 80
		img := imaging.Fit(common.AftermathLogo(brandColor, opts), 256, 256, imaging.Linear)
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
