package common

import (
	"image/color"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/log"
)

type TitleCardStyle struct {
	Container Style
	Nickname  Style
	ClanTag   Style
}

func (style TitleCardStyle) TotalPaddingAndGaps() float64 {
	return style.Container.PaddingX*2 + style.Container.Gap*2 + style.Nickname.PaddingX*2 + style.ClanTag.PaddingX*4
}

func DefaultPlayerTitleStyle(name string, containerStyle Style) TitleCardStyle {
	containerStyle.AlignItems = AlignItemsCenter
	containerStyle.Direction = DirectionHorizontal
	// containerStyle.Debug = true

	nameFontSize := FontLarge()
	if len(name) > 10 {
		nameFontSize = FontMedium()
	}
	return TitleCardStyle{
		Container: containerStyle,
		Nickname:  Style{Font: nameFontSize, FontColor: TextPrimary},
		ClanTag:   Style{Font: FontMedium(), FontColor: TextSecondary, PaddingX: 10, PaddingY: 5, BackgroundColor: ClanTagBackgroundColor, BorderRadius: BorderRadiusSM},
	}
}

func NewPlayerTitleCard(style TitleCardStyle, nickname, clanTag string, subscriptions []models.UserSubscription) Block {
	clanTagBlock, hasClanTagBlock := newClanTagBlock(style.ClanTag, clanTag, subscriptions)
	if !hasClanTagBlock {
		return NewBlocksContent(style.Container, NewTextContent(style.Nickname, nickname))
	}

	content := make([]Block, 0, 3)
	style.Container.JustifyContent = JustifyContentSpaceBetween

	clanTagImage, err := clanTagBlock.Render()
	if err != nil {
		log.Warn().Err(err).Msg("failed to render clan tag")
		// This error is not fatal, we can just render the name
		return NewBlocksContent(style.Container, NewTextContent(style.Nickname, nickname))
	}
	content = append(content, NewImageContent(Style{Width: float64(clanTagImage.Bounds().Dx()), Height: float64(clanTagImage.Bounds().Dy())}, clanTagImage))

	// Nickname
	content = append(content, NewTextContent(style.Nickname, nickname))

	// Invisible tag to offset the nickname
	clanTagOffsetBlock := NewBlocksContent(Style{
		Width:          float64(clanTagImage.Bounds().Dx()),
		JustifyContent: JustifyContentEnd,
	}, NewTextContent(Style{Font: style.ClanTag.Font, FontColor: color.Transparent}, "-"))
	content = append(content, clanTagOffsetBlock)

	return NewBlocksContent(style.Container, content...)

}

func newClanTagBlock(style Style, clanTag string, subs []models.UserSubscription) (Block, bool) {
	if clanTag == "" {
		return Block{}, false
	}

	var blocks []Block
	blocks = append(blocks, NewTextContent(Style{Font: FontMedium(), FontColor: TextSecondary}, clanTag))
	if sub := ClanSubscriptionsBadges(subs); sub != nil {
		iconBlock, err := sub.Block()
		if err == nil {
			blocks = append(blocks, iconBlock)
		}
	}

	return NewBlocksContent(style, blocks...), true
}
