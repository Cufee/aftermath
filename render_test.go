package main

import (
	"context"
	"os"
	"testing"
	"time"

	stats "github.com/cufee/aftermath/internal/stats/client/v1"
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

	stats := stats.NewClient(tests.StaticTestingFetch(), tests.StaticTestingDatabase(), nil, language.English)

	{
		image, _, err := stats.SessionImage(context.Background(), tests.DefaultAccountNAShort, time.Now())
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")

		f, err := os.Create("tmp/render_test_session_full_small.png")
		assert.NoError(t, err, "failed to create a file")
		defer f.Close()

		err = image.PNG(f)
		assert.NoError(t, err, "failed to encode a png image")
	}
	{
		image, _, err := stats.SessionImage(context.Background(), tests.DefaultAccountNA, time.Now())
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")

		f, err := os.Create("tmp/render_test_session_full_large.png")
		assert.NoError(t, err, "failed to create a file")
		defer f.Close()

		err = image.PNG(f)
		assert.NoError(t, err, "failed to encode a png image")
	}
}
