package common

import (
	"image/color"

	"github.com/cufee/aftermath/internal/database"
	"github.com/rs/zerolog/log"
)

type TitleCardStyle struct {
	Container Style
	Nickname  Style
	ClanTag   Style
}

func (style TitleCardStyle) TotalPaddingAndGaps() float64 {
	return style.Container.PaddingX*2 + style.Container.Gap + style.Nickname.PaddingX*2 + style.ClanTag.PaddingX*2
}

func DefaultPlayerTitleStyle(containerStyle Style) TitleCardStyle {
	containerStyle.AlignItems = AlignItemsCenter
	containerStyle.Direction = DirectionHorizontal

	clanTagBackgroundColor := DefaultCardColor
	clanTagBackgroundColor.R += 10
	clanTagBackgroundColor.G += 10
	clanTagBackgroundColor.B += 10

	return TitleCardStyle{
		Container: containerStyle,
		Nickname:  Style{Font: &FontLarge, FontColor: TextPrimary},
		ClanTag:   Style{Font: &FontMedium, FontColor: TextSecondary, PaddingX: 10, PaddingY: 5, BackgroundColor: clanTagBackgroundColor, BorderRadius: 10},
	}
}

func NewPlayerTitleCard(style TitleCardStyle, nickname, clanTag string, subscriptions []database.UserSubscription) Block {
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

func newClanTagBlock(style Style, clanTag string, subs []database.UserSubscription) (Block, bool) {
	if clanTag == "" {
		return Block{}, false
	}

	var blocks []Block
	blocks = append(blocks, NewTextContent(Style{Font: &FontMedium, FontColor: TextSecondary}, clanTag))
	if sub := ClanSubscriptionsBadges(subs); sub != nil {
		iconBlock, err := sub.Block()
		if err == nil {
			blocks = append(blocks, iconBlock)
		}
	}

	return NewBlocksContent(style, blocks...), true
}
