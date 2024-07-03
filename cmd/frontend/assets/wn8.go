package assets

import (
	"path/filepath"
	"strings"

	"github.com/cufee/aftermath/internal/stats/render/common/v1"
)

func WN8IconFilename(rating float32) string {
	name := strings.ReplaceAll(strings.ToLower(common.GetWN8TierName(rating)), " ", "_")
	if rating < 1 {
		name = "invalid"
	}
	return name + ".png"
}

func WN8IconPath(rating float32) string {
	return filepath.Join("/assets", "wn8", WN8IconFilename(rating))
}

func WN8IconPathSmall(rating float32) string {
	return filepath.Join("/assets", "wn8", "small_"+WN8IconFilename(rating))
}
