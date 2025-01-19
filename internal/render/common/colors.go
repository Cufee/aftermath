package common

import "image/color"

var DiscordBackgroundColor = color.NRGBA{49, 51, 56, 255}

var DefaultLogoColorOptions = []color.Color{color.NRGBA{50, 50, 50, 180}, color.NRGBA{200, 200, 200, 180}}

var (
	TextPrimary   = color.NRGBA{255, 255, 255, 255}
	TextSecondary = color.NRGBA{204, 204, 204, 255}
	TextAlt       = color.NRGBA{150, 150, 150, 255}

	TextSubscriptionPlus    = color.NRGBA{72, 167, 250, 255}
	TextSubscriptionPremium = color.NRGBA{255, 223, 0, 255}

	DefaultCardColor        = color.NRGBA{10, 10, 10, 150}
	DefaultCardColorNoAlpha = color.NRGBA{10, 10, 10, 255}
	ClanTagBackgroundColor  = color.NRGBA{10, 10, 10, 100}

	ColorAftermathRed    = color.NRGBA{255, 0, 120, 255}
	ColorAftermathBlue   = color.NRGBA{72, 167, 250, 255}
	ColorAftermathYellow = color.NRGBA{255, 223, 0, 255}
)
