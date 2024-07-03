package assets

import (
	"path/filepath"

	"github.com/cufee/aftermath/internal/stats/render/common/v1"
)

func RatingIconPath(rating float32) string {
	return filepath.Join("/assets", "rating", common.GetRatingIconName(rating)+".png")
}
