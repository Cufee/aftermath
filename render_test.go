package main

import (
	"bytes"
	"context"
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
	"github.com/cufee/aftermath/tests"
	"github.com/disintegration/imaging"
	"github.com/matryer/is"
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
		image, _, err := stats.SessionImage(context.Background(), tests.DefaultAccountNA, time.Now(), client.WithBackgroundURL(bgImage))
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")

		f, err := os.Create("tmp/render_test_session_full_large.png")
		assert.NoError(t, err, "failed to create a file")
		defer f.Close()

		err = image.PNG(f)
		assert.NoError(t, err, "failed to encode a png image")
	}
	{
		image, _, err := stats.SessionImage(context.Background(), tests.DefaultAccountNA, time.Now(), client.WithBackgroundURL(bgImage), client.WithVehicleID("0"))
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")

		f, err := os.Create("tmp/render_test_session_single_large.png")
		assert.NoError(t, err, "failed to create a file")
		defer f.Close()

		err = image.PNG(f)
		assert.NoError(t, err, "failed to encode a png image")
	}
	{
		image, _, err := stats.SessionImage(context.Background(), tests.DefaultAccountNA, time.Now(), client.WithBackgroundURL(bgImage), client.WithVehicleID("0"))
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")

		f, err := os.Create("tmp/render_test_session_single_large.png")
		assert.NoError(t, err, "failed to create a file")
		defer f.Close()

		err = image.PNG(f)
		assert.NoError(t, err, "failed to encode a png image")
	}
}

func TestRenderPeriod(t *testing.T) {
	// Logger
	level, _ := zerolog.ParseLevel(os.Getenv("LOG_LEVEL"))
	zerolog.SetGlobalLevel(level)

	loadStaticAssets(static)

	bgImage := "static://bg-default"

	stats := client.NewClient(tests.StaticTestingFetch(), tests.StaticTestingDatabase(), nil, language.English)

	{
		image, _, err := stats.PeriodImage(context.Background(), tests.DefaultAccountNAShort, time.Now())
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")

		f, err := os.Create("tmp/render_test_period_full_small.png")
		assert.NoError(t, err, "failed to create a file")
		defer f.Close()

		err = image.PNG(f)
		assert.NoError(t, err, "failed to encode a png image")
	}
	{
		image, _, err := stats.PeriodImage(context.Background(), tests.DefaultAccountNA, time.Now(), client.WithBackgroundURL(bgImage))
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")

		f, err := os.Create("tmp/render_test_period_full_large.png")
		assert.NoError(t, err, "failed to create a file")
		defer f.Close()

		err = image.PNG(f)
		assert.NoError(t, err, "failed to encode a png image")
	}
	{
		image, _, err := stats.PeriodImage(context.Background(), tests.DefaultAccountNA, time.Now(), client.WithBackgroundURL(bgImage), client.WithVehicleID("0"))
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")

		f, err := os.Create("tmp/render_test_period_single_large.png")
		assert.NoError(t, err, "failed to create a file")
		defer f.Close()

		err = image.PNG(f)
		assert.NoError(t, err, "failed to encode a png image")
	}
	{
		image, _, err := stats.PeriodImage(context.Background(), tests.DefaultAccountNAShort, time.Now(), client.WithBackgroundURL(bgImage), client.WithVehicleID("0"))
		assert.NoError(t, err, "failed to render a session image")
		assert.NotNil(t, image, "image is nil")

		f, err := os.Create("tmp/render_test_period_single_small.png")
		assert.NoError(t, err, "failed to create a file")
		defer f.Close()

		err = image.PNG(f)
		assert.NoError(t, err, "failed to encode a png image")
	}
}

func TestRenderReplay(t *testing.T) {
	is := is.New(t)

	// Logger
	level, _ := zerolog.ParseLevel(os.Getenv("LOG_LEVEL"))
	zerolog.SetGlobalLevel(level)

	loadStaticAssets(static)

	db, err := newDatabaseClientFromEnv()
	is.NoErr(err)

	wg, _ := wargamingClientsFromEnv()

	printer, err := localization.NewPrinter("stats", language.English)
	is.NoErr(err)

	file, err := os.ReadFile("tests/replay_2.wotbreplay")
	is.NoErr(err)

	fetch, err := fetch.NewMultiSourceClient(wg, nil, db)
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

	{
		cards, err := prepare.NewCards(replay, glossary, gameModeNames, common.WithPrinter(printer, language.English))
		is.NoErr(err)

		image, err := render.CardsToImage(replay, cards, rc.WithBackground(""))
		assert.NoError(t, err, "failed to render a replay image")
		assert.NotNil(t, image, "image is nil")

		f, err := os.Create("tmp/render_test_replay.png")
		assert.NoError(t, err, "failed to create a file")
		defer f.Close()

		err = imaging.Save(image, "tmp/render_test_replay.png")
		assert.NoError(t, err, "failed to encode a png image")
	}
}
