package main

import (
	"context"
	"os"
	"testing"
	"time"

	common "github.com/cufee/aftermath/internal/stats/client/common"
	client "github.com/cufee/aftermath/internal/stats/client/v2"
	"github.com/cufee/aftermath/tests"
	"github.com/cufee/aftermath/tests/env"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"

	_ "github.com/joho/godotenv/autoload"
)

func TestRenderPeriodV2(t *testing.T) {
	env.LoadTestEnv(t)
	stats := client.NewClient(tests.StaticTestingFetch(), tests.StaticTestingDatabase(), nil, language.English)

	t.Run("render period image for small nickname", func(t *testing.T) {
		image, _, err := stats.PeriodImage(context.Background(), tests.DefaultAccountNAShort, time.Now(), common.WithWN8())
		assert.NoError(t, err, "failed to render a period image")
		assert.NotNil(t, image, "image is nil")

		f, err := os.Create("tmp/render_test_period_v2_full_small.png")
		assert.NoError(t, err, "failed to create a file")
		defer f.Close()

		err = image.PNG(f)
		assert.NoError(t, err, "failed to encode a png image")
	})

	t.Run("render period image for large nickname", func(t *testing.T) {
		image, _, err := stats.PeriodImage(context.Background(), tests.DefaultAccountNA, time.Now(), common.WithBackgroundURL(bgImage, bgIsCustom), common.WithWN8())
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")

		f, err := os.Create("tmp/render_test_period_v2_full_large.png")
		assert.NoError(t, err, "failed to create a file")
		defer f.Close()

		err = image.PNG(f)
		assert.NoError(t, err, "failed to encode a png image")
	})

	t.Run("render period image with large name no highlights", func(t *testing.T) {
		image, _, err := stats.PeriodImage(context.Background(), tests.DefaultAccountNA, time.Now(), common.WithBackgroundURL(bgImage, bgIsCustom), common.WithVehicleID("0"), common.WithWN8())
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")

		f, err := os.Create("tmp/render_test_period_v2_single_large.png")
		assert.NoError(t, err, "failed to create a file")
		defer f.Close()

		err = image.PNG(f)
		assert.NoError(t, err, "failed to encode a png image")
	})

	t.Run("render period image with small name and no highlights", func(t *testing.T) {
		image, _, err := stats.PeriodImage(context.Background(), tests.DefaultAccountNAShort, time.Now(), common.WithBackgroundURL(bgImage, bgIsCustom), common.WithVehicleID("0"), common.WithWN8())
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")

		f, err := os.Create("tmp/render_test_period_v2_single_small.png")
		assert.NoError(t, err, "failed to create a file")
		defer f.Close()

		err = image.PNG(f)
		assert.NoError(t, err, "failed to encode a png image")
	})
}
