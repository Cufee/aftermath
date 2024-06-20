package main

import (
	"context"
	"image/png"
	"os"
	"testing"
	"time"

	"github.com/cufee/aftermath/internal/stats"
	"github.com/cufee/aftermath/internal/stats/render"
	"github.com/cufee/aftermath/internal/stats/render/assets"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

func TestRenderSession(t *testing.T) {
	// Logger
	level, _ := zerolog.ParseLevel(os.Getenv("LOG_LEVEL"))
	zerolog.SetGlobalLevel(level)

	loadStaticAssets(static)
	coreClient, _ := coreClientsFromEnv()
	defer coreClient.Database().Disconnect()

	rating, _ := assets.GetLoadedImage("rating-calibration")
	c := imaging.PasteCenter(gg.NewContext(int(float64(rating.Bounds().Dx())*1.5), int(float64(rating.Bounds().Dy())*1.5)).Image(), rating)

	fc, err := os.Create("tmp/rating-calibration.png")
	assert.NoError(t, err, "failed to create a file")
	defer fc.Close()

	png.Encode(fc, c)

	bgImage, ok := assets.GetLoadedImage("bg-default")
	assert.True(t, ok, "failed to load a background image")

	renderer := stats.NewRenderer(coreClient.Fetch(), coreClient.Database(), coreClient.Wargaming(), language.English)
	image, _, err := renderer.Session(context.Background(), "1013072123", time.Now(), render.WithBackground(bgImage))
	assert.NoError(t, err, "failed to render a session image")
	assert.NotNil(t, image, "image is nil")

	f, err := os.Create("tmp/render_test_session.png")
	assert.NoError(t, err, "failed to create a file")
	defer f.Close()

	err = image.PNG(f)
	assert.NoError(t, err, "failed to encode a png image")
}
