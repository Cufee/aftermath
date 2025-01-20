package session

import (
	"image"
	"image/color"
	"slices"
	"time"

	"github.com/cufee/aftermath/internal/render/common"
	"github.com/cufee/aftermath/internal/stats/frame"
)

func addBackgroundBranding(background image.Image, vehicles map[string]frame.VehicleStatsFrame, patternSeed int) image.Image {
	var values []vehicleWN8
	for _, vehicle := range vehicles {
		if wn8 := vehicle.WN8(); !frame.InvalidValue.Equals(wn8) {
			values = append(values, vehicleWN8{vehicle.VehicleID, wn8, int(vehicle.LastBattleTime.Unix())})
		}
	}
	slices.SortFunc(values, func(a, b vehicleWN8) int { return b.sortKey - a.sortKey })
	if len(values) >= 10 {
		values = values[:9]
	}

	var accentColors []color.Color
	for _, value := range values {
		c := common.GetWN8Colors(value.wn8.Float()).Background
		if _, _, _, a := c.RGBA(); a > 0 {
			accentColors = append(accentColors, c)
		}
	}

	if patternSeed == 0 {
		patternSeed = int(time.Now().Unix())
	}

	return common.AddDefaultBrandedOverlay(background, accentColors, patternSeed, 0.5)
}
