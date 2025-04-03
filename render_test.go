package main

import (
	"bytes"
	"context"
	"os"
	"slices"
	"testing"

	"github.com/cufee/aftermath/internal/localization"
	rc "github.com/cufee/aftermath/internal/render/common"

	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/common/v1"

	prepare "github.com/cufee/aftermath/internal/stats/prepare/replay/v1"
	render "github.com/cufee/aftermath/internal/stats/render/replay/v1"

	"github.com/cufee/aftermath/tests"
	"github.com/cufee/aftermath/tests/env"
	"github.com/matryer/is"
	"github.com/nao1215/imaging"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"

	_ "github.com/joho/godotenv/autoload"
)

var bgImage = "static://bg-default"
var bgIsCustom = false

func init() {
	loadStaticAssets(static)
	level, _ := zerolog.ParseLevel(os.Getenv("LOG_LEVEL"))
	zerolog.SetGlobalLevel(level)
}

func TestRenderReplay(t *testing.T) {
	is := is.New(t)

	env.LoadTestEnv(t)
	db := tests.StaticTestingDatabase()
	wg, _ := wargamingClientsFromEnv()

	printer, err := localization.NewPrinterWithFallback("stats", language.English)
	is.NoErr(err)

	fetch, err := fetch.NewMultiSourceClient(wg, nil, db)
	is.NoErr(err)

	replayFiles := []string{"replay_defeat.wotbreplay", "replay_victory.wotbreplay", "replay_draw.wotbreplay", "replay_rating.wotbreplay"}
	for _, name := range replayFiles {
		t.Run("render replay image from "+name, func(t *testing.T) {
			is := is.New(t)

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

			{
				cards, err := prepare.NewCards(replay, glossary, gameModeNames, common.WithPrinter(printer, language.English))
				is.NoErr(err)

				image, err := render.CardsToImage(replay, cards, rc.WithBackgroundURL(bgImage, bgIsCustom), rc.WithPrinter(printer))
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
