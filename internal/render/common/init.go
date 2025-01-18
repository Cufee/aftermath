package render

import (
	"image/color"
)

var DiscordBackgroundColor = color.NRGBA{49, 51, 56, 255}

var (
	TextPrimary   = color.NRGBA{255, 255, 255, 255}
	TextSecondary = color.NRGBA{204, 204, 204, 255}
	TextAlt       = color.NRGBA{150, 150, 150, 255}

	TextSubscriptionPlus    = color.NRGBA{72, 167, 250, 255}
	TextSubscriptionPremium = color.NRGBA{255, 223, 0, 255}

	DefaultCardColor        = color.NRGBA{10, 10, 10, 180}
	DefaultCardColorNoAlpha = color.NRGBA{10, 10, 10, 255}
	ClanTagBackgroundColor  = color.NRGBA{10, 10, 10, 120}

	ColorAftermathRed    = color.NRGBA{255, 0, 120, 255}
	ColorAftermathBlue   = color.NRGBA{72, 167, 250, 255}
	ColorAftermathYellow = color.NRGBA{255, 223, 0, 255}

	BorderRadiusXL = 30.0
	BorderRadiusLG = 25.0
	BorderRadiusMD = 20.0
	BorderRadiusSM = 15.0
	BorderRadiusXS = 10.0
)
