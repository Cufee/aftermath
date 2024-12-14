package assets

import (
	"path/filepath"
	"strings"

	"github.com/cufee/aftermath/internal/render/v1"
)

func wn8IconFilename(rating float32) string {
	name := strings.ReplaceAll(strings.ToLower(render.GetWN8TierName(rating)), " ", "_")
	if rating < 1 {
		name = "invalid"
	}
	return name + ".png"
}

func WN8IconPath(rating float32) string {
	return filepath.Join("/assets", "wn8", wn8IconFilename(rating))
}

func WN8IconPathSmall(rating float32) string {
	return filepath.Join("/assets", "wn8", "small_"+wn8IconFilename(rating))
}
