package main

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cufee/aftermath/internal/stats/render/common/v1"
	stats "github.com/cufee/aftermath/internal/stats/renderer/v1"
	"github.com/cufee/aftermath/tests"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"

	_ "github.com/joho/godotenv/autoload"
)

func TestRenderSession(t *testing.T) {
	// Logger
	level, _ := zerolog.ParseLevel(os.Getenv("LOG_LEVEL"))
	zerolog.SetGlobalLevel(level)

	loadStaticAssets(static)

	renderer := stats.NewRenderer(tests.StaticTestingFetch(), tests.StaticTestingDatabase(), nil, language.English)

	{
		image, _, err := renderer.Session(context.Background(), tests.DefaultAccountNAShort, time.Now(), common.WithBackground(""))
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")

		f, err := os.Create("tmp/render_test_session_full_small.png")
		assert.NoError(t, err, "failed to create a file")
		defer f.Close()

		err = image.PNG(f)
		assert.NoError(t, err, "failed to encode a png image")
	}
	{
		image, _, err := renderer.Session(context.Background(), tests.DefaultAccountNA, time.Now(), common.WithBackground(""))
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")

		f, err := os.Create("tmp/render_test_session_full_large.png")
		assert.NoError(t, err, "failed to create a file")
		defer f.Close()

		err = image.PNG(f)
		assert.NoError(t, err, "failed to encode a png image")
	}
}
