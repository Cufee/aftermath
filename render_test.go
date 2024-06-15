package main

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cufee/aftermath/internal/stats"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

func TestRenderSession(t *testing.T) {
	// Logger
	level, _ := zerolog.ParseLevel(os.Getenv("LOG_LEVEL"))
	zerolog.SetGlobalLevel(level)

	loadStaticAssets(static)
	coreClient := coreClientFromEnv()

	renderer := stats.NewRenderer(coreClient.Fetch(), coreClient.Database(), coreClient.Wargaming(), language.English)
	image, _, err := renderer.Session(context.Background(), "1013072123", time.Now())
	assert.NoError(t, err, "failed to render a session image")
	assert.NotNil(t, image, "image is nil")

	f, err := os.Open("render_test_session.png")
	assert.NoError(t, err, "failed to open a file")
	defer f.Close()

	err = image.PNG(f)
	assert.NoError(t, err, "failed to encode a png image")
}
