//go:build ignore

package main

import (
	"fmt"
	"image/color"
	"os"
	"path/filepath"

	"github.com/cufee/aftermath/internal/localization"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/stats/render/common/v1"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"golang.org/x/text/language"

	"github.com/joho/godotenv"
)

var outDirPath = "../../static"
var brandColor color.RGBA

var discordColorLight = color.RGBA{54, 56, 61, 255}
var discordColorMedium = color.RGBA{46, 47, 52, 255}
var discordColorDark = color.RGBA{30, 31, 34, 255}
var discordColorText = color.RGBA{151, 155, 162, 255}

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
	printer, _ := localization.NewPrinter("discord", language.English)

	generateDiscordHelpImage(printer)
}

func generateDiscordHelpImage(printer func(string) string) {
	log.Debug().Msg("generating discord help image")

	imageWidth := 512
	imageHeight := imageWidth * 2 / 4
	inputHeight := 50
	iconSize := 25
	gap := 5

	{
		filename := "images/discord/discord-help.jpg"

		ctx := gg.NewContext(imageWidth, imageHeight)
		ctx.SetColor(discordColorMedium)
		ctx.Clear()

		{ // draw Discord UI
			{
				// text input
				tctx := gg.NewContext(imageWidth-gap*2, inputHeight)
				tctx.DrawRoundedRectangle(0, 0, float64(tctx.Width()), float64(tctx.Height()), common.BorderRadiusXS)
				tctx.SetColor(discordColorLight)
				tctx.Fill()

				input := ctx.Image()
				ctx.DrawImage(input, gap, imageHeight-gap-tctx.Height())

				sctx := gg.NewContext(imageWidth, tctx.Height()+gap*2)
				sctx.DrawRoundedRectangle(float64(gap), float64(gap), float64(tctx.Width()), float64(tctx.Height()), common.BorderRadiusXS)
				sctx.SetColor(color.RGBA{0, 0, 0, 100})
				sctx.Fill()

				ctx.DrawImage(imaging.Blur(sctx.Image(), 1.5), 0, imageHeight-sctx.Height())
				ctx.DrawImage(tctx.Image(), gap, imageHeight-gap-tctx.Height())

				circleIconR := float64(iconSize / 2)
				circleIconX := float64(gap + 50 - (inputHeight / 2))
				circleIconY := float64(imageHeight - gap - (inputHeight / 2))

				ctx.DrawCircle(circleIconX, circleIconY, circleIconR)
				ctx.SetColor(discordColorText)
				ctx.Fill()

				ctx.SetLineWidth(3)
				ctx.SetLineCapRound()
				ctx.SetColor(discordColorMedium)
				ctx.DrawLine(circleIconX, circleIconY-circleIconR+7, circleIconX, circleIconY+circleIconR-7)
				ctx.DrawLine(circleIconX-circleIconR+7, circleIconY, circleIconX+circleIconR-7, circleIconY)
				ctx.Stroke()

				ctx.SetColor(discordColorText)
				ctx.SetLineCapSquare()
				ctx.DrawLine(float64(gap+inputHeight), float64(circleIconY+circleIconR), float64(gap+inputHeight)+5, float64(circleIconY-circleIconR))
				ctx.Stroke()
			}

			// commands drawer
			{
				dctx := gg.NewContext(imageWidth-gap*2, imageHeight-gap*3-inputHeight)

				dctx.DrawRoundedRectangle(0, 0, float64(dctx.Width()), float64(dctx.Height()), common.BorderRadiusXS)
				dctx.Clip()

				dctx.DrawRectangle(0, 0, float64(dctx.Width()), float64(dctx.Height()))
				dctx.SetColor(discordColorDark)
				dctx.Fill()

				dctx.DrawRectangle(float64(inputHeight), 0, float64(dctx.Width()-inputHeight), float64(dctx.Height()))
				dctx.SetColor(discordColorMedium)
				dctx.Fill()

				fontSize := 20.0
				commands := []string{"help", "links", "stats"}
				for i, name := range commands {
					drawY := float64((gap*2 + int(fontSize)) + i*(int(fontSize*2)+gap*3))
					dctx.SetColor(color.White)
					dctx.LoadFontFace("../../static/fonts/default.ttf", fontSize)
					dctx.DrawString("/"+printer(fmt.Sprintf("command_%s_name", name)), float64(gap+inputHeight+gap*2), drawY)

					dctx.LoadFontFace("../../static/fonts/default.ttf", fontSize*0.8)
					dctx.SetColor(discordColorText)
					dctx.DrawString(printer(fmt.Sprintf("command_%s_description", name)), float64(gap+inputHeight+gap*2), drawY+fontSize+float64(gap/2))
					dctx.DrawString("Aftermath", float64(dctx.Width()-gap*3-70), drawY+fontSize/2+float64(gap/2))
				}

				sctx := gg.NewContext(imageWidth, dctx.Height()+gap*2)
				sctx.DrawRoundedRectangle(float64(gap), float64(gap), float64(dctx.Width()), float64(dctx.Height()), common.BorderRadiusXS)
				sctx.SetColor(color.RGBA{0, 0, 0, 100})
				sctx.Fill()

				ctx.DrawImage(imaging.Blur(sctx.Image(), 2), 0, 0)
				ctx.DrawImage(dctx.Image(), gap, gap)

			}

			// commands
			{

			}
		}

		opts := common.DefaultLogoOptions()
		logo := common.AftermathLogo(brandColor, opts)

		ctx.DrawImage(imaging.Fit(logo, iconSize, iconSize, imaging.Linear), gap+((inputHeight-iconSize)/2), gap+((inputHeight-iconSize)/2))

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
