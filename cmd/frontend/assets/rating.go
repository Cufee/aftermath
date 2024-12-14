package assets

import (
	"path/filepath"

	"github.com/cufee/aftermath/internal/render/v1"
)

func RatingIconPath(rating float32) string {
	return filepath.Join("/assets", "rating", render.GetRatingIconName(rating)+".png")
}
