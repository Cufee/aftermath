package main

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cufee/aftermath/internal/stats/render/common/v1"
	stats "github.com/cufee/aftermath/internal/stats/renderer/v1"
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

	db, err := newDatabaseClientFromEnv()
	assert.NoError(t, err)

	coreClient, _ := coreClientsFromEnv(db)
	defer coreClient.Database().Disconnect()

	renderer := stats.NewRenderer(coreClient.Fetch(), coreClient.Database(), coreClient.Wargaming(), language.English)
	image, _, err := renderer.Session(context.Background(), "1013072123", time.Now(), common.WithBackground(""))
	assert.NoError(t, err, "failed to render a session image")
	assert.NotNil(t, image, "image is nil")

	f, err := os.Create("tmp/render_test_session.png")
	assert.NoError(t, err, "failed to create a file")
	defer f.Close()

	err = image.PNG(f)
	assert.NoError(t, err, "failed to encode a png image")
}
