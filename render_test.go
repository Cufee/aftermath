package main

import (
	"bytes"
	"context"
	"image/png"
	"os"
	"slices"
	"testing"
	"time"

	"github.com/cufee/aftermath/internal/localization"
	client "github.com/cufee/aftermath/internal/stats/client/v1"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/replay/v1"
	rc "github.com/cufee/aftermath/internal/stats/render/common/v1"
	render "github.com/cufee/aftermath/internal/stats/render/replay/v1"
	session "github.com/cufee/aftermath/internal/stats/render/session/v1"
	"github.com/cufee/aftermath/tests"
	"github.com/matryer/is"
	"github.com/nao1215/imaging"
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

	bgImage := "static://bg-default"

	stats := client.NewClient(tests.StaticTestingFetch(), tests.StaticTestingDatabase(), nil, language.English)
	t.Run("generate content mask before generating image", func(t *testing.T) {
		cards, meta, err := stats.SessionCards(context.Background(), tests.DefaultAccountNAShort, time.Now(), client.WithWN8())
		assert.NoError(t, err, "failed to generate session cards")

		segments, err := session.CardsToSegments(meta.Stats["session"], meta.Stats["career"], cards, nil)
		assert.NoError(t, err, "failed to render a session segments")
		assert.NotNil(t, segments, "segments is nil")

		start := time.Now()
		mask, err := segments.ContentMask()
		assert.NoError(t, err, "failed to generate content mask")
		println("mask", time.Since(start).Milliseconds())

		f, err := os.Create("tmp/render_test_session_full_small_mask.png")
		assert.NoError(t, err, "failed to create a file")
		defer f.Close()

		err = png.Encode(f, mask)
		assert.NoError(t, err, "failed to encode a png image")

		start = time.Now()
		_, err = segments.Render()
		println("image after mask", time.Since(start).Milliseconds())
		assert.NoError(t, err, "failed to render segments")
	})

	t.Run("render session image for small nickname", func(t *testing.T) {
		image, _, err := stats.SessionImage(context.Background(), tests.DefaultAccountNAShort, time.Now(), client.WithWN8())
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")

		f, err := os.Create("tmp/render_test_session_full_small.png")
		assert.NoError(t, err, "failed to create a file")
		defer f.Close()

		err = image.PNG(f)
		assert.NoError(t, err, "failed to encode a png image")
	})

	t.Run("render session image for large nickname", func(t *testing.T) {
		image, _, err := stats.SessionImage(context.Background(), tests.DefaultAccountNA, time.Now(), client.WithBackgroundURL(bgImage, false), client.WithWN8())
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")

		f, err := os.Create("tmp/render_test_session_full_large.png")
		assert.NoError(t, err, "failed to create a file")
		defer f.Close()

		err = image.PNG(f)
		assert.NoError(t, err, "failed to encode a png image")
	})

	t.Run("render session image for large nickname and no vehicles", func(t *testing.T) {
		image, _, err := stats.SessionImage(context.Background(), tests.DefaultAccountNA, time.Now(), client.WithBackgroundURL(bgImage, false), client.WithVehicleID("0"), client.WithWN8())
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")

		f, err := os.Create("tmp/render_test_session_single_large.png")
		assert.NoError(t, err, "failed to create a file")
		defer f.Close()

		err = image.PNG(f)
		assert.NoError(t, err, "failed to encode a png image")
	})
}

func TestRenderPeriod(t *testing.T) {
	// Logger
	level, _ := zerolog.ParseLevel(os.Getenv("LOG_LEVEL"))
	zerolog.SetGlobalLevel(level)

	loadStaticAssets(static)

	bgImage := "static://bg-default"

	stats := client.NewClient(tests.StaticTestingFetch(), tests.StaticTestingDatabase(), nil, language.English)

	t.Run("render period image for small nickname", func(t *testing.T) {
		image, _, err := stats.PeriodImage(context.Background(), tests.DefaultAccountNAShort, time.Now(), client.WithWN8())
		assert.NoError(t, err, "failed to render a period image")
		assert.NotNil(t, image, "image is nil")

		f, err := os.Create("tmp/render_test_period_full_small.png")
		assert.NoError(t, err, "failed to create a file")
		defer f.Close()

		err = image.PNG(f)
		assert.NoError(t, err, "failed to encode a png image")
	})

	t.Run("render period image for large nickname", func(t *testing.T) {
		image, _, err := stats.PeriodImage(context.Background(), tests.DefaultAccountNA, time.Now(), client.WithBackgroundURL(bgImage, false), client.WithWN8())
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")

		f, err := os.Create("tmp/render_test_period_full_large.png")
		assert.NoError(t, err, "failed to create a file")
		defer f.Close()

		err = image.PNG(f)
		assert.NoError(t, err, "failed to encode a png image")
	})

	t.Run("render period image with large name no highlights", func(t *testing.T) {
		image, _, err := stats.PeriodImage(context.Background(), tests.DefaultAccountNA, time.Now(), client.WithBackgroundURL(bgImage, false), client.WithVehicleID("0"), client.WithWN8())
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")

		f, err := os.Create("tmp/render_test_period_single_large.png")
		assert.NoError(t, err, "failed to create a file")
		defer f.Close()

		err = image.PNG(f)
		assert.NoError(t, err, "failed to encode a png image")
	})

	t.Run("render period image with small name and no highlights", func(t *testing.T) {
		image, _, err := stats.PeriodImage(context.Background(), tests.DefaultAccountNAShort, time.Now(), client.WithBackgroundURL(bgImage, false), client.WithVehicleID("0"), client.WithWN8())
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")

		f, err := os.Create("tmp/render_test_period_single_small.png")
		assert.NoError(t, err, "failed to create a file")
		defer f.Close()

		err = image.PNG(f)
		assert.NoError(t, err, "failed to encode a png image")
	})
}

func TestRenderReplay(t *testing.T) {
	is := is.New(t)

	// Logger
	level, _ := zerolog.ParseLevel(os.Getenv("LOG_LEVEL"))
	zerolog.SetGlobalLevel(level)

	loadStaticAssets(static)

	db := tests.StaticTestingDatabase()
	wg, _ := wargamingClientsFromEnv()

	printer, err := localization.NewPrinterWithFallback("stats", language.English)
	is.NoErr(err)

	fetch, err := fetch.NewMultiSourceClient(wg, nil, db)
	is.NoErr(err)

	replayFiles := []string{"replay_defeat.wotbreplay", "replay_victory.wotbreplay", "replay_draw.wotbreplay"}
	for _, name := range replayFiles {
		t.Run("render replay image from "+name, func(t *testing.T) {
			file, err := os.ReadFile("tests/" + name)
			is.NoErr(err)

			replay, err := fetch.Replay(context.Background(), bytes.NewReader(file), int64(len(file)))
			is.NoErr(err)

			var vehicles []string
			for _, player := range append(replay.Teams.Allies, replay.Teams.Enemies...) {
				if id := player.VehicleID; !slices.Contains(vehicles, id) {
					vehicles = append(vehicles, id)
				}
			}

			glossary, err := db.GetVehicles(context.Background(), vehicles)
			is.NoErr(err)

			gameModeNames, err := db.GetGameModeNames(context.Background(), replay.GameMode.String())
			is.NoErr(err)

			bgImage := "static://bg-default"

			{
				cards, err := prepare.NewCards(replay, glossary, gameModeNames, common.WithPrinter(printer, language.English))
				is.NoErr(err)

				image, err := render.CardsToImage(replay, cards, rc.WithBackgroundURL(bgImage, false), rc.WithPrinter(printer))
				assert.NoError(t, err, "failed to render a replay image")
				assert.NotNil(t, image, "image is nil")

				out := "tmp/render_test_" + name + ".png"
				f, err := os.Create(out)
				assert.NoError(t, err, "failed to create a file")
				defer f.Close()

				err = imaging.Save(image, out)
				assert.NoError(t, err, "failed to encode a png image")
			}
		})
	}
}
