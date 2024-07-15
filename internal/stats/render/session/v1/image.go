package session

import (
	"image"
	"image/color"
	"strconv"
	"time"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/session/v1"
	"github.com/cufee/aftermath/internal/stats/render/common/v1"
	"github.com/fogleman/gg"
)

var logoColorOptions = []color.Color{color.NRGBA{150, 150, 150, 25}, color.NRGBA{204, 204, 204, 25}, color.NRGBA{255, 255, 255, 25}}

func CardsToImage(session, career fetch.AccountStatsOverPeriod, cards session.Cards, subs []models.UserSubscription, opts ...common.Option) (image.Image, error) {
	o := common.DefaultOptions()
	for _, apply := range opts {
		apply(&o)
	}

	segments, err := cardsToSegments(session, career, cards, subs, o)
	if err != nil {
		return nil, err
	}

	content, err := segments.Render()
	if err != nil {
		return nil, err
	}

	patternSeed, _ := strconv.Atoi(session.Account.ID)
	if patternSeed == 0 {
		patternSeed = int(time.Now().Unix())
	}

	bg := common.NewBrandedBackground(content.Bounds().Dx(), content.Bounds().Dy(), logoColorOptions, patternSeed)
	contentWithBg := gg.NewContext(content.Bounds().Dx(), content.Bounds().Dy())
	contentWithBg.SetColor(color.NRGBA{24, 26, 27, 255})
	contentWithBg.Clear()
	contentWithBg.DrawImage(bg, 0, 0)
	contentWithBg.DrawImage(content, 0, 0)

	f := common.FontCustom(common.Max(float64(content.Bounds().Dx())/8, 50))
	nameSize := common.MeasureString(session.Account.Nickname, f)

	frameWidth := nameSize.LineHeight / 4
	frame := common.NewFrameContext(contentWithBg.Image(),
		common.WithFrameBackground(o.Background),
		common.WithFrameWidth(frameWidth),
		common.WithBorderRadius(40),
		common.WithShadow(f.Size()/15),
	)

	{ // draw player name
		letterSize := common.MeasureString(string(session.Account.Nickname[0]), f)
		frame.Ctx().SetFontFace(f.Face())
		frame.Ctx().SetColor(common.TextPrimary)
		frame.Ctx().DrawString(session.Account.Nickname, frame.Width/2-letterSize.TotalWidth/2, nameSize.LineHeight/2.75)
	}

	return frame.Image(), nil
}
