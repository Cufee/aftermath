package main

import (
	"context"
	"testing"
	"time"

	common "github.com/cufee/aftermath/internal/stats/client/common"
	client "github.com/cufee/aftermath/internal/stats/client/v2"
	"github.com/cufee/aftermath/internal/stats/render/themes/spring2026"
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
		saveTestImage(t, "period", "full_small.png", image.PNG)
	})

	t.Run("render period image for large nickname", func(t *testing.T) {
		image, _, err := stats.PeriodImage(context.Background(), tests.DefaultAccountNA, time.Now(), common.WithBackgroundURL(bgImage, bgIsCustom), common.WithWN8())
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")
		saveTestImage(t, "period", "full_large.png", image.PNG)
	})

	t.Run("render period image with large name no highlights", func(t *testing.T) {
		image, _, err := stats.PeriodImage(context.Background(), tests.DefaultAccountNA, time.Now(), common.WithBackgroundURL(bgImage, bgIsCustom), common.WithVehicleIDs("0"), common.WithWN8())
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")
		saveTestImage(t, "period", "single_large.png", image.PNG)
	})

	t.Run("render period image with small name and no highlights", func(t *testing.T) {
		image, _, err := stats.PeriodImage(context.Background(), tests.DefaultAccountNAShort, time.Now(), common.WithBackgroundURL(bgImage, bgIsCustom), common.WithVehicleIDs("0"), common.WithWN8())
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")
		saveTestImage(t, "period", "single_small.png", image.PNG)
	})

	t.Run("render period image with no vehicles", func(t *testing.T) {
		image, _, err := stats.PeriodImage(context.Background(), tests.DefaultAccountNAShort, time.Now(), common.WithBackgroundURL(bgImage, bgIsCustom), common.WithVehicleIDs("-"), common.WithWN8())
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")
		saveTestImage(t, "period", "none_small.png", image.PNG)
	})
}

func TestRenderSessionV2(t *testing.T) {
	env.LoadTestEnv(t)
	stats := client.NewClient(tests.StaticTestingFetch(), tests.StaticTestingDatabase(), nil, language.English)

	t.Run("render session image for small nickname", func(t *testing.T) {
		image, _, err := stats.SessionImage(context.Background(), tests.DefaultAccountNAShort, time.Now(), common.WithWN8())
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")
		saveTestImage(t, "session", "full_small.png", image.PNG)
	})

	t.Run("render session image for large nickname", func(t *testing.T) {
		image, _, err := stats.SessionImage(context.Background(), tests.DefaultAccountNA, time.Now(), common.WithBackgroundURL(bgImage, bgIsCustom), common.WithWN8())
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")
		saveTestImage(t, "session", "full_large.png", image.PNG)
	})

	t.Run("render session image for large nickname and no vehicles", func(t *testing.T) {
		image, _, err := stats.SessionImage(context.Background(), tests.DefaultAccountNA, time.Now(), common.WithBackgroundURL(bgImage, bgIsCustom), common.WithVehicleIDs("-"), common.WithWN8())
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")
		saveTestImage(t, "session", "0_large.png", image.PNG)
	})

	t.Run("render session image for large nickname and 1 vehicle", func(t *testing.T) {
		image, _, err := stats.SessionImage(context.Background(), tests.DefaultAccountNA, time.Now(), common.WithBackgroundURL(bgImage, bgIsCustom), common.WithVehicleIDs("1"), common.WithWN8())
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")
		saveTestImage(t, "session", "1_large.png", image.PNG)
	})

	t.Run("render session image for large nickname and 3 vehicles", func(t *testing.T) {
		image, _, err := stats.SessionImage(context.Background(), tests.DefaultAccountNA, time.Now(), common.WithBackgroundURL(bgImage, bgIsCustom), common.WithVehicleIDs("1", "2", "3"), common.WithWN8())
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")
		saveTestImage(t, "session", "3_large.png", image.PNG)
	})

	t.Run("render session image for large nickname and 5 vehicles", func(t *testing.T) {
		image, _, err := stats.SessionImage(context.Background(), tests.DefaultAccountNA, time.Now(), common.WithBackgroundURL(bgImage, bgIsCustom), common.WithVehicleIDs("1", "2", "3", "4", "5"), common.WithWN8())
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")
		saveTestImage(t, "session", "5_large.png", image.PNG)
	})
}

func TestRenderSessionV2Spring2026(t *testing.T) {
	env.LoadTestEnv(t)
	stats := client.NewClient(tests.StaticTestingFetch(), tests.StaticTestingDatabase(), nil, language.English)

	theme := spring2026.Theme()

	t.Run("render session image for large nickname", func(t *testing.T) {
		image, _, err := stats.SessionImage(context.Background(), tests.DefaultAccountNA, time.Now(),
			common.WithTheme(theme), common.WithWN8(),
		)
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")
		saveTestImage(t, "session-spring2026", "full_large.png", image.PNG)
	})

	t.Run("render session image for small nickname", func(t *testing.T) {
		image, _, err := stats.SessionImage(context.Background(), tests.DefaultAccountNAShort, time.Now(),
			common.WithTheme(theme), common.WithWN8(),
		)
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")
		saveTestImage(t, "session-spring2026", "full_small.png", image.PNG)
	})
}

func TestRenderPeriodV2Spring2026(t *testing.T) {
	env.LoadTestEnv(t)
	stats := client.NewClient(tests.StaticTestingFetch(), tests.StaticTestingDatabase(), nil, language.English)

	theme := spring2026.Theme()

	t.Run("render period image for large nickname", func(t *testing.T) {
		image, _, err := stats.PeriodImage(context.Background(), tests.DefaultAccountNA, time.Now(),
			common.WithTheme(theme), common.WithWN8(),
		)
		assert.NoError(t, err, "failed to render a period image")
		assert.NotNil(t, image, "image is nil")
		saveTestImage(t, "period-spring2026", "full_large.png", image.PNG)
	})

	t.Run("render period image for small nickname", func(t *testing.T) {
		image, _, err := stats.PeriodImage(context.Background(), tests.DefaultAccountNAShort, time.Now(),
			common.WithTheme(theme), common.WithWN8(),
		)
		assert.NoError(t, err, "failed to render a period image")
		assert.NotNil(t, image, "image is nil")
		saveTestImage(t, "period-spring2026", "full_small.png", image.PNG)
	})
}
