package auction

import (
	"fmt"
	"image"
	"image/color"

	"github.com/cufee/aftermath/internal/render/assets"
	"github.com/cufee/aftermath/internal/render/v1"
	"golang.org/x/sync/errgroup"
	"golang.org/x/text/language"
)

var (
	cardStyle = render.Style{
		BackgroundColor: render.DefaultCardColor,
		BorderRadius:    render.BorderRadiusLG,
		PaddingX:        25,
		PaddingY:        25,
		Direction:       render.DirectionVertical,
		JustifyContent:  render.JustifyContentCenter,
		Gap:             2,
	}
)

func vehicleCard(style render.Style, locale language.Tag, vehicle Vehicle) (image.Image, error) {
	name := render.NewTextContent(render.Style{Font: render.FontLarge(), FontColor: render.TextPrimary}, vehicle.Name(locale))
	leftCount := render.NewTextContent(render.Style{Font: render.FontMedium(), FontColor: render.TextSecondary}, fmt.Sprintf("Vehicles left: %d", vehicle.Available))

	var vehicleIcon render.Block
	if vehicle.Image == nil {
		vehicleIcon = render.NewEmptyContent(250, 250)
	} else {
		vehicleIcon = render.NewImageContent(render.Style{Width: 250, Height: 250}, vehicle.Image)
	}

	icon, _ := assets.GetLoadedImage("currency-gold")
	currentPrice := render.NewBlocksContent(render.Style{Direction: render.DirectionHorizontal, Gap: 1, JustifyContent: render.JustifyContentCenter},
		render.NewTextContent(render.Style{Font: render.FontSmall(), FontColor: render.TextAlt}, "CURRENT PRICE"),
		render.NewTextContent(render.Style{Font: render.FontLarge(), FontColor: color.White}, fmt.Sprint(vehicle.Price.Current.Value)),
		render.NewImageContent(render.Style{Width: 26, Height: 26}, icon),
	)
	nextPrice := render.NewBlocksContent(render.Style{Direction: render.DirectionHorizontal, Gap: 1, JustifyContent: render.JustifyContentCenter},
		render.NewTextContent(render.Style{Font: render.FontSmall(), FontColor: render.TextAlt}, "UPCOMING"),
		render.NewTextContent(render.Style{Font: render.FontMedium(), FontColor: color.White}, fmt.Sprint(vehicle.Price.Upcoming.Value)),
		render.NewImageContent(render.Style{Width: 20, Height: 20}, icon),
	)

	card := render.NewBlocksContent(style, name, leftCount, vehicleIcon, currentPrice, nextPrice)

	return card.Render()
}

func AuctionCards(data AuctionVehicles, locale language.Tag) ([]image.Image, error) {
	vehicleCardsCh := make(chan image.Image, len(data.Vehicles))

	var group errgroup.Group
	for _, vehicle := range data.Vehicles {
		group.Go(func() error {
			card, err := vehicleCard(cardStyle, locale, vehicle)
			if err != nil {
				return err
			}
			vehicleCardsCh <- card
			return nil
		})
	}
	err := group.Wait()
	close(vehicleCardsCh)
	if err != nil {
		return nil, err
	}

	var cards []image.Image
	for c := range vehicleCardsCh {
		cards = append(cards, c)
	}

	return cards, nil
}
