package period

import (
	"image"
	"strconv"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/render/common"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/aftermath/internal/stats/prepare/period/v1"
	"github.com/cufee/facepaint"
	"github.com/cufee/facepaint/style"
)

type vehicleWN8 struct {
	id      string
	wn8     frame.Value
	sortKey int
}

func CardsToImage(stats fetch.AccountStatsOverPeriod, cards period.Cards, subs []models.UserSubscription, opts ...common.Option) (image.Image, error) {
	o := common.DefaultOptions()
	for _, apply := range opts {
		apply(&o)
	}

	// Generate cards
	cardsBlock, err := generateCards(stats, cards, subs, o)
	if err != nil {
		return nil, err
	}

	if o.Background == nil {
		return cardsBlock.Render()
	}

	if !o.BackgroundIsCustom {
		seed, _ := strconv.Atoi(stats.Account.ID)
		o.Background = addBackgroundBranding(o.Background, stats.RegularBattles.Vehicles, seed)
	}

	contentSize := cardsBlock.Dimensions()
	withBackground := facepaint.NewBlocksContent(style.NewStyle(),
		facepaint.MustNewImageContent(
			style.NewStyle(
				style.SetWidth(float64(contentSize.Width)),
				style.SetHeight(float64(contentSize.Height)),
				style.SetBlur(common.DefaultBackgroundBlur),
				style.SetPosition(style.PositionAbsolute),
				style.SetZIndex(-99),
			), o.Background),
		cardsBlock,
	)
	return withBackground.Render()

}
