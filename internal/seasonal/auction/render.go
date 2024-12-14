package auction

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"sync"
	"time"

	"github.com/cufee/aftermath/internal/logic"
	"github.com/cufee/aftermath/internal/render/assets"
	"github.com/cufee/aftermath/internal/render/v1"
	"github.com/nao1215/imaging"
	"golang.org/x/sync/errgroup"
	"golang.org/x/text/language"
)

var (
	timeIconSize = 16.0

	goldIconSmallSize = 20.0
	goldIconLargeSize = 32.0

	tankImageSize = 300.0

	cardStyle = render.Style{
		Width:           tankImageSize + 120,
		BackgroundColor: render.DefaultCardColor,
		BorderRadius:    render.BorderRadiusLG,
		PaddingX:        60,
		PaddingY:        50,
		Direction:       render.DirectionVertical,
		JustifyContent:  render.JustifyContentCenter,
		Gap:             2,
	}
)

var timerIcon image.Image

type styledText struct {
	value string
	style render.Style
	size  render.StringSize
}

func centerText(target float64, text ...*styledText) float64 {
	var widthMax float64 = target

	for _, line := range text {
		line.size = render.MeasureString(line.value, line.style.Font)
		widthMax = max(widthMax, line.size.TotalWidth)
	}

	for _, line := range text {
		line.style.PaddingX += (widthMax - line.size.TotalWidth) / 2
	}

	return widthMax
}

func headerCard() (image.Image, error) {
	bg, _ := assets.GetLoadedImage("auction-splash")
	return imaging.Fill(bg, 1000, 50, imaging.Top, imaging.Linear), nil
}

func vehicleCard(style render.Style, locale language.Tag, vehicle Vehicle) (image.Image, error) {
	nameColor := render.TextPrimary
	if vehicle.Premium {
		nameColor = render.TextSubscriptionPremium
	}
	if vehicle.Collectible {
		nameColor = render.TextSubscriptionPlus
	}

	name := styledText{
		value: fmt.Sprintf("%s %s", logic.IntToRoman(vehicle.Tier), vehicle.Name(locale)),
		style: render.Style{Font: render.FontXL(), FontColor: nameColor},
	}

	remaining := time.Until(vehicle.NextPriceDrop).Round(time.Second)
	hours := int(remaining.Hours())
	minutes := int(remaining.Minutes()) % 60
	seconds := int(remaining.Seconds()) % 60
	nextDropString := fmt.Sprintf("%d:%02d:%02d", hours, minutes, seconds)

	availableWithTime := styledText{
		value: fmt.Sprintf("Available: %d %s", vehicle.Available, nextDropString),
		style: render.Style{Font: render.FontLarge(), FontColor: render.TextSecondary, PaddingX: (timeIconSize / 2) + 4},
	}

	widthMax := centerText(style.Width-(style.PaddingX*2), &name, &availableWithTime)
	style.Width = max(widthMax+(style.PaddingX*2), style.Width)
	contentWidth := style.Width - style.PaddingX*2

	var blocks []render.Block
	blocks = append(blocks, render.NewTextContent(name.style, name.value))
	blocks = append(blocks, render.NewBlocksContent(render.Style{Direction: render.DirectionHorizontal, AlignItems: render.AlignItemsCenter, JustifyContent: render.JustifyContentSpaceBetween, Width: contentWidth},
		render.NewTextContent(render.Style{Font: render.FontLarge(), FontColor: render.TextSecondary}, fmt.Sprintf("Available: %d", vehicle.Available)),
		render.NewBlocksContent(render.Style{Direction: render.DirectionHorizontal, AlignItems: render.AlignItemsCenter, Gap: 4},
			render.NewImageContent(render.Style{BackgroundColor: render.TextSecondary}, timerIcon),
			render.NewTextContent(render.Style{Font: render.FontLarge(), FontColor: render.TextSecondary}, nextDropString),
		),
	))

	if vehicle.Image == nil {
		blocks = append(blocks, render.NewEmptyContent(1, tankImageSize))
	} else {
		iconPadding := (contentWidth - tankImageSize) / 2
		blocks = append(blocks, render.NewBlocksContent(render.Style{PaddingX: iconPadding}, render.NewImageContent(render.Style{Width: tankImageSize, Height: tankImageSize}, vehicle.Image)))
	}

	icon, _ := assets.GetLoadedImage("currency-gold")
	blocks = append(blocks, render.NewBlocksContent(render.Style{Direction: render.DirectionHorizontal, Gap: 2, JustifyContent: render.JustifyContentCenter, AlignItems: render.AlignItemsCenter, Width: contentWidth},
		render.NewTextContent(render.Style{Font: render.FontXL(), FontColor: color.White}, fmt.Sprint(vehicle.Price.Current.Value)),
		render.NewImageContent(render.Style{Width: goldIconLargeSize, Height: goldIconLargeSize}, icon),
	))
	blocks = append(blocks, render.NewBlocksContent(render.Style{Direction: render.DirectionHorizontal, Gap: 4, JustifyContent: render.JustifyContentCenter, AlignItems: render.AlignItemsCenter, Width: contentWidth},
		render.NewTextContent(render.Style{Font: render.FontSmall(), FontColor: render.TextAlt}, "Next Price:"),
		render.NewBlocksContent(render.Style{Direction: render.DirectionHorizontal, AlignItems: render.AlignItemsCenter},
			render.NewTextContent(render.Style{Font: render.FontMedium(), FontColor: render.TextSecondary}, fmt.Sprint(vehicle.Price.Upcoming.Value)),
			render.NewImageContent(render.Style{Width: goldIconSmallSize, Height: goldIconSmallSize}, icon),
		),
	))

	card := render.NewBlocksContent(style, blocks...)

	return card.Render()
}

func AuctionCards(data AuctionVehicles, locale language.Tag) ([]image.Image, error) {
	if timerIcon == nil {
		img, ok := assets.GetLoadedImage("timer")
		if !ok {
			return nil, errors.New("failed to get a timer icon")
		}
		timerIcon = imaging.Resize(img, int(timeIconSize), int(timeIconSize), imaging.Linear)
	}

	var cardsMx sync.Mutex
	cards := make([]image.Image, len(data.Vehicles)+1)

	var group errgroup.Group
	group.Go(func() error {
		header, err := headerCard()
		cards[0] = header
		return err
	})

	for i, vehicle := range data.Vehicles {
		group.Go(func() error {
			card, err := vehicleCard(cardStyle, locale, vehicle)
			if err != nil {
				return err
			}
			cardsMx.Lock()
			defer cardsMx.Unlock()
			cards[i+1] = card
			return nil
		})
	}
	err := group.Wait()
	if err != nil {
		return nil, err
	}

	return cards, nil
}
