package assets

import (
	"context"
	"time"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/localization"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/session/v1"
	"github.com/cufee/aftermath/tests"
	"golang.org/x/text/language"
)

type mockWidget struct {
	Account models.Account
	Cards   session.Cards
}

var mockWidgetData *mockWidget

func MockWidgetData() mockWidget {
	if mockWidgetData == nil {
		loadWidgetData()
	}
	return *mockWidgetData
}

func loadWidgetData() {
	// TODO: need a better way to do this

	mockWidgetData = &mockWidget{}
	mockWidgetData.Account = models.Account{ID: "mock-account-id", Realm: "NA", Nickname: "@your_name", CreatedAt: time.Now(), LastBattleTime: time.Now(), ClanID: "mock-clan", ClanTag: "AMTH"}

	sess, career, err := tests.StaticTestingFetch().SessionStats(context.Background(), tests.DefaultAccountNA, time.Time{}, fetch.WithWN8())
	if err != nil {
		panic(err)
	}

	printer, err := localization.NewPrinter("stats", language.English)
	if err != nil {
		panic(err)
	}

	cards, err := session.NewCards(sess, career, nil, common.WithPrinter(printer, language.English))
	if err != nil {
		panic(err)
	}
	mockWidgetData.Cards = cards
}
