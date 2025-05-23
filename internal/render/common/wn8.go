package common

import (
	"image/color"
)

type ratingColors struct {
	Background color.Color
	Content    color.Color
}

func GetWN8Colors(r float32) ratingColors {
	if r > 0 && r < 301 {
		return ratingColors{color.NRGBA{255, 0, 0, 255}, color.White}
	}
	if r > 300 && r < 451 {
		return ratingColors{color.NRGBA{251, 83, 83, 255}, color.White}
	}
	if r > 450 && r < 651 {
		return ratingColors{color.NRGBA{255, 160, 49, 255}, color.White}
	}
	if r > 650 && r < 901 {
		return ratingColors{color.NRGBA{255, 244, 65, 255}, color.Black}
	}
	if r > 900 && r < 1201 {
		return ratingColors{color.NRGBA{149, 245, 62, 255}, color.Black}
	}
	if r > 1200 && r < 1601 {
		return ratingColors{color.NRGBA{103, 190, 51, 255}, color.Black}
	}
	if r > 1600 && r < 2001 {
		return ratingColors{color.NRGBA{106, 236, 255, 255}, color.Black}
	}
	if r > 2000 && r < 2451 {
		return ratingColors{color.NRGBA{46, 174, 193, 255}, color.White}
	}
	if r > 2450 && r < 2901 {
		return ratingColors{color.NRGBA{208, 108, 255, 255}, color.White}
	}
	if r > 2900 {
		return ratingColors{color.NRGBA{160, 87, 193, 255}, color.Black}
	}
	return ratingColors{color.Transparent, color.Transparent}
}

func GetWN8TierName(r float32) string {
	if r > 0 && r < 301 {
		return "Very Bad"
	}
	if r > 300 && r < 451 {
		return "Bad"
	}
	if r > 450 && r < 651 {
		return "Below Average"
	}
	if r > 650 && r < 901 {
		return "Average"
	}
	if r > 900 && r < 1201 {
		return "Above Average"
	}
	if r > 1200 && r < 1601 {
		return "Good"
	}
	if r > 1600 && r < 2001 {
		return "Very Good"
	}
	if r > 2000 && r < 2451 {
		return "Great"
	}
	if r > 2450 && r < 2901 {
		return "Unicum"
	}
	if r > 2900 {
		return "Super Unicum"
	}
	return ""
}
