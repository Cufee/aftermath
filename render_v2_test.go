package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	common "github.com/cufee/aftermath/internal/stats/client/common"
	options "github.com/cufee/aftermath/internal/stats/client/common"
	client "github.com/cufee/aftermath/internal/stats/client/v2"
	"github.com/cufee/aftermath/tests"
	"github.com/cufee/aftermath/tests/env"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"

	_ "github.com/joho/godotenv/autoload"
)

func sessionRenderOutputDir(t *testing.T) string {
	t.Helper()
	// e.g. TestRenderSessionV2/full_small -> tmp/render_test_session/TestRenderSessionV2__full_small
	name := strings.ReplaceAll(t.Name(), "/", "__")
	return filepath.Join("tmp", "render_test_session", name)
}

func writeSessionImagePages(t *testing.T, outDir string, images []common.Image) {
	t.Helper()
	assert.NotEmpty(t, images, "expected at least one session page")
	assert.NoError(t, os.MkdirAll(outDir, 0o755))
	for i, img := range images {
		var name string
		if len(images) == 1 {
			name = filepath.Join(outDir, "page.png")
		} else {
			name = filepath.Join(outDir, fmt.Sprintf("page_%d.png", i))
		}
		f, err := os.Create(name)
		assert.NoError(t, err, "failed to create output file")
		err = img.PNG(f)
		_ = f.Close()
		assert.NoError(t, err, "failed to encode a png image")
	}
}

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
		image, _, err := stats.PeriodImage(context.Background(), tests.DefaultAccountNA, time.Now(), common.WithBackgroundURL(bgImage, bgIsCustom), common.WithVehicleIDs("0"), common.WithWN8())
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")

		f, err := os.Create("tmp/render_test_period_v2_single_large.png")
		assert.NoError(t, err, "failed to create a file")
		defer f.Close()

		err = image.PNG(f)
		assert.NoError(t, err, "failed to encode a png image")
	})

	t.Run("render period image with small name and no highlights", func(t *testing.T) {
		image, _, err := stats.PeriodImage(context.Background(), tests.DefaultAccountNAShort, time.Now(), common.WithBackgroundURL(bgImage, bgIsCustom), common.WithVehicleIDs("0"), common.WithWN8())
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")

		f, err := os.Create("tmp/render_test_period_v2_single_small.png")
		assert.NoError(t, err, "failed to create a file")
		defer f.Close()

		err = image.PNG(f)
		assert.NoError(t, err, "failed to encode a png image")
	})

	t.Run("render period image with no vehicles", func(t *testing.T) {
		image, _, err := stats.PeriodImage(context.Background(), tests.DefaultAccountNAShort, time.Now(), common.WithBackgroundURL(bgImage, bgIsCustom), common.WithVehicleIDs("-"), common.WithWN8())
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")

		f, err := os.Create("tmp/render_test_period_v2_none_small.png")
		assert.NoError(t, err, "failed to create a file")
		defer f.Close()

		err = image.PNG(f)
		assert.NoError(t, err, "failed to encode a png image")
	})
}

func TestRenderSessionV2(t *testing.T) {
	env.LoadTestEnv(t)
	stats := client.NewClient(tests.StaticTestingFetch(), tests.StaticTestingDatabase(), nil, language.English)

	t.Run("render session image for small nickname", func(t *testing.T) {
		images, _, err := stats.SessionImage(context.Background(), tests.DefaultAccountNAShort, time.Now(), options.WithWN8())
		assert.NoError(t, err, "failed to render a session image")
		writeSessionImagePages(t, sessionRenderOutputDir(t), images)
	})

	t.Run("render session image for large nickname", func(t *testing.T) {
		// Static data includes regular + rating battles (2 overviews), so highlights + vehicles go to continuation pages.
		images, _, err := stats.SessionImage(context.Background(), tests.DefaultAccountNA, time.Now(), options.WithBackgroundURL(bgImage, bgIsCustom), options.WithWN8())
		assert.NoError(t, err, "failed to render a session image")
		assert.GreaterOrEqual(t, len(images), 2, "expected multi-page session with deferred highlights")
		writeSessionImagePages(t, sessionRenderOutputDir(t), images)
	})

	t.Run("render session image for large nickname and no vehicles", func(t *testing.T) {
		images, _, err := stats.SessionImage(context.Background(), tests.DefaultAccountNA, time.Now(), options.WithBackgroundURL(bgImage, bgIsCustom), options.WithVehicleIDs("-"), options.WithWN8())
		assert.NoError(t, err, "failed to render a session image")
		writeSessionImagePages(t, sessionRenderOutputDir(t), images)
	})

	t.Run("render session image for large nickname and 1 vehicle", func(t *testing.T) {
		images, _, err := stats.SessionImage(context.Background(), tests.DefaultAccountNA, time.Now(), options.WithBackgroundURL(bgImage, bgIsCustom), options.WithVehicleIDs("1"), options.WithWN8())
		assert.NoError(t, err, "failed to render a session image")
		writeSessionImagePages(t, sessionRenderOutputDir(t), images)
	})

	t.Run("render session image for large nickname and 3 vehicles", func(t *testing.T) {
		images, _, err := stats.SessionImage(context.Background(), tests.DefaultAccountNA, time.Now(), options.WithBackgroundURL(bgImage, bgIsCustom), options.WithVehicleIDs("1", "2", "3"), options.WithWN8())
		assert.NoError(t, err, "failed to render a session image")
		writeSessionImagePages(t, sessionRenderOutputDir(t), images)
	})

	t.Run("render session image for large nickname and 5 vehicles", func(t *testing.T) {
		images, _, err := stats.SessionImage(context.Background(), tests.DefaultAccountNA, time.Now(), options.WithBackgroundURL(bgImage, bgIsCustom), options.WithVehicleIDs("1", "2", "3", "4", "5"), options.WithWN8())
		assert.NoError(t, err, "failed to render a session image")
		writeSessionImagePages(t, sessionRenderOutputDir(t), images)
	})

	t.Run("render session with many vehicle ids (multi-page when data allows)", func(t *testing.T) {
		manyIDs := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}
		images, _, err := stats.SessionImage(context.Background(), tests.DefaultAccountNA, time.Now(),
			options.WithBackgroundURL(bgImage, bgIsCustom), options.WithVehicleIDs(manyIDs...), options.WithWN8())
		assert.NoError(t, err, "failed to render session images")
		writeSessionImagePages(t, sessionRenderOutputDir(t), images)
	})

	t.Run("highlights on page 0 when no rating battles", func(t *testing.T) {
		images, _, err := stats.SessionImage(context.Background(), tests.DefaultAccountNASessionNoRating, time.Now(),
			options.WithBackgroundURL(bgImage, bgIsCustom), options.WithWN8())
		assert.NoError(t, err, "failed to render session images")
		writeSessionImagePages(t, sessionRenderOutputDir(t), images)
	})

	t.Run("max pages (10)", func(t *testing.T) {
		var ids []string
		for i := range 70 {
			ids = append(ids, fmt.Sprintf("%d", i+1))
		}
		images, _, err := stats.SessionImage(context.Background(), tests.DefaultAccountNA, time.Now(),
			options.WithBackgroundURL(bgImage, bgIsCustom), options.WithVehicleIDs(ids...), options.WithWN8())
		assert.NoError(t, err, "failed to render session images")
		assert.Equal(t, 10, len(images), "expected exactly 10 pages (max)")
		writeSessionImagePages(t, sessionRenderOutputDir(t), images)
	})

}
