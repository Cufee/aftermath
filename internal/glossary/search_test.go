package glossary

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"testing"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"golang.org/x/text/language"
)

type glossaryStub struct {
	database.GlossaryClient
	vehicles map[string]models.Vehicle
	err      error
}

func (s glossaryStub) GetAllVehicles(context.Context) (map[string]models.Vehicle, error) {
	return s.vehicles, s.err
}

func preserveCache(t *testing.T) {
	t.Helper()
	cacheLock.RLock()
	oldTime, oldVehicles, oldAliases, oldTiers := cacheTime, vehiclesCache, vehicleAliases, vehicleTierToIDs
	cacheLock.RUnlock()
	t.Cleanup(func() {
		cacheLock.Lock()
		cacheTime, vehiclesCache, vehicleAliases, vehicleTierToIDs = oldTime, oldVehicles, oldAliases, oldTiers
		cacheLock.Unlock()
	})
}

func loadVehicles(t *testing.T, vehicles ...models.Vehicle) {
	t.Helper()
	data := make(map[string]models.Vehicle, len(vehicles))
	for _, vehicle := range vehicles {
		data[vehicle.ID] = vehicle
	}
	if err := RefreshVehicleCache(context.Background(), glossaryStub{vehicles: data}); err != nil {
		t.Fatal(err)
	}
}

func assertSearch(t *testing.T, locale language.Tag, query string, limit int, want ...string) {
	t.Helper()
	vehicles, ok := SearchVehicles(locale, query, limit)
	var got []string
	for _, vehicle := range vehicles {
		got = append(got, vehicle.ID)
	}
	if !reflect.DeepEqual(got, want) || ok != (len(want) > 0) {
		t.Fatalf("SearchVehicles(%s, %q, %d) = %v, %v; want %v, %v", locale, query, limit, got, ok, want, len(want) > 0)
	}
}

func TestSearchVehiclesAllLanguages(t *testing.T) {
	preserveCache(t)
	loadVehicles(t,
		models.Vehicle{ID: "maus", Tier: 10, LocalizedNames: map[language.Tag]string{language.English: "Maus"}},
		models.Vehicle{ID: "tiger", Tier: 7, LocalizedNames: map[language.Tag]string{language.English: "Tiger I"}},
		models.Vehicle{ID: "e100", Tier: 10, LocalizedNames: map[language.Tag]string{language.English: "E 100"}},
		models.Vehicle{ID: "is", Tier: 7, LocalizedNames: map[language.Tag]string{language.English: "IS-1", language.Russian: "ИС-1"}},
		models.Vehicle{ID: "chinese", Tier: 6, LocalizedNames: map[language.Tag]string{language.Chinese: "虎"}},
	)
	for _, locale := range []language.Tag{language.English, language.Russian, language.Japanese} {
		t.Run(locale.String(), func(t *testing.T) {
			assertSearch(t, locale, "MAUS", 5, "maus")
			assertSearch(t, locale, "Tiger I", 5, "tiger")
			assertSearch(t, locale, "E 100", 5, "e100")
			assertSearch(t, locale, "IS-1", 5, "is")
			assertSearch(t, locale, "иС-1", 5, "is")
			assertSearch(t, locale, "vii ИС-1", 5, "is")
			assertSearch(t, locale, "X Maus", 5, "maus")
			assertSearch(t, locale, "虎", 5, "chinese") // Preserve the three-byte minimum.
		})
	}
	vehicle, ok := GetVehicleFromCache(language.Russian, "maus")
	if !ok || vehicle.Name(language.Russian) != "Maus" {
		t.Fatal("English display fallback was lost")
	}
	vehicle, ok = GetVehicleFromCache(language.Russian, "is")
	if !ok || vehicle.Name(language.Russian) != "ИС-1" {
		t.Fatal("localized display name was lost")
	}
}

func TestSearchVehiclesDeduplicatesIDsNotNames(t *testing.T) {
	preserveCache(t)
	loadVehicles(t,
		models.Vehicle{ID: "10", Tier: 6, LocalizedNames: map[language.Tag]string{
			language.English: "T-34-85", language.German: "T-34-85", language.French: "t-34-85", language.Russian: "Т-34-85",
		}},
		models.Vehicle{ID: "2", Tier: 6, LocalizedNames: map[language.Tag]string{language.English: "T-34-85"}},
	)
	assertSearch(t, language.Russian, "34-85", 2, "10", "2")
	assertSearch(t, language.English, "T-34-85", 5, "10", "2")
	assertSearch(t, language.Russian, "VI T-34-85", 5, "10", "2")
	assertSearch(t, language.Russian, "T-34-85", 1, "10")
}

func TestSearchVehiclesRanksBeforeLimiting(t *testing.T) {
	preserveCache(t)
	vehicles := []models.Vehicle{{ID: "z-exact", Tier: 10, LocalizedNames: map[language.Tag]string{
		language.English: "Maus", language.French: "Maus Extra Long",
	}}}
	for i := 1; i <= 6; i++ {
		vehicles = append(vehicles, models.Vehicle{
			ID: fmt.Sprintf("%d", i), Tier: 10,
			LocalizedNames: map[language.Tag]string{language.English: "Maus " + strings.Repeat("x", i)},
		})
	}
	for iteration := 0; iteration < 20; iteration++ {
		// Reverse insertion order as well as exercising Go's map iteration order.
		for i, j := 0, len(vehicles)-1; i < j; i, j = i+1, j-1 {
			vehicles[i], vehicles[j] = vehicles[j], vehicles[i]
		}
		loadVehicles(t, vehicles...)
		assertSearch(t, language.Russian, "Maus", 5, "z-exact", "1", "2", "3", "4")
	}
}

func TestSearchVehiclesEmptyAndInvalidQueries(t *testing.T) {
	preserveCache(t)
	loadVehicles(t,
		models.Vehicle{ID: "named", Tier: 10, LocalizedNames: map[language.Tag]string{language.English: "Maus"}},
		models.Vehicle{ID: "empty", Tier: 8, LocalizedNames: map[language.Tag]string{language.English: ""}},
		models.Vehicle{ID: "unnamed", Tier: 9},
	)
	for _, query := range []string{"", "M", "Ma", "И", "missing", "viii "} {
		assertSearch(t, language.English, query, 5)
	}
	for _, limit := range []int{-1, 0} {
		assertSearch(t, language.English, "Maus", limit)
	}
	for _, id := range []string{"empty", "unnamed"} {
		if vehicle, ok := GetVehicleFromCache(language.Russian, id); !ok || vehicle.ID != id {
			t.Fatalf("unnamed vehicle %q is missing from the ID cache", id)
		}
	}
	if _, ok := GetVehicleFromCache(language.English, "missing"); ok {
		t.Fatal("unexpected vehicle for missing ID")
	}
	if ids, ok := TierVehicleIDs(9); !ok || !reflect.DeepEqual(ids, []string{"unnamed"}) {
		t.Fatalf("unnamed tier lookup = %v, %v", ids, ok)
	}
	loadVehicles(t)
	assertSearch(t, language.English, "Maus", 5)
	if _, ok := TierVehicleIDs(9); ok {
		t.Fatal("empty refresh retained old tier entries")
	}
}

func TestRefreshVehicleCacheFailurePreservesSnapshot(t *testing.T) {
	preserveCache(t)
	loadVehicles(t, models.Vehicle{ID: "maus", Tier: 10, LocalizedNames: map[language.Tag]string{language.English: "Maus"}})
	cacheLock.RLock()
	before := cacheTime
	cacheLock.RUnlock()
	wantErr := errors.New("database unavailable")
	if err := RefreshVehicleCache(context.Background(), glossaryStub{err: wantErr}); !errors.Is(err, wantErr) {
		t.Fatalf("refresh error = %v; want %v", err, wantErr)
	}
	assertSearch(t, language.Russian, "Maus", 5, "maus")
	if _, ok := GetVehicleFromCache(language.English, "maus"); !ok {
		t.Fatal("failed refresh discarded the vehicle cache")
	}
	if ids, ok := TierVehicleIDs(10); !ok || !reflect.DeepEqual(ids, []string{"maus"}) {
		t.Fatalf("failed refresh changed tier lookup: %v, %v", ids, ok)
	}
	cacheLock.RLock()
	after := cacheTime
	cacheLock.RUnlock()
	if !before.Equal(after) {
		t.Fatal("failed refresh changed the cache timestamp")
	}
}

func TestVehicleCacheConcurrentRefreshAndReads(t *testing.T) {
	preserveCache(t)
	clients := []glossaryStub{
		{vehicles: map[string]models.Vehicle{"a": {ID: "a", Tier: 10, LocalizedNames: map[language.Tag]string{language.English: "Maus Alpha"}}}},
		{vehicles: map[string]models.Vehicle{"b": {ID: "b", Tier: 10, LocalizedNames: map[language.Tag]string{language.English: "Maus Beta"}}}},
	}
	if err := RefreshVehicleCache(context.Background(), clients[0]); err != nil {
		t.Fatal(err)
	}
	var wg sync.WaitGroup
	for worker := 0; worker < 6; worker++ {
		wg.Go(func() {
			for iteration := 0; iteration < 100; iteration++ {
				if worker < 2 {
					if err := RefreshVehicleCache(context.Background(), clients[iteration%2]); err != nil {
						t.Error(err)
					}
					continue
				}
				vehicles, ok := SearchVehicles(language.Russian, "Maus", 5)
				if !ok || len(vehicles) != 1 || (vehicles[0].ID != "a" && vehicles[0].ID != "b") || vehicles[0].Name(language.English) == "" {
					t.Errorf("inconsistent search snapshot: %v, %v", vehicles, ok)
				}
				if vehicle, ok := GetVehicleFromCache(language.English, "a"); ok && vehicle.ID != "a" {
					t.Errorf("inconsistent ID lookup: %v", vehicle)
				}
				if ids, ok := TierVehicleIDs(10); !ok || len(ids) != 1 || (ids[0] != "a" && ids[0] != "b") {
					t.Errorf("inconsistent tier lookup: %v, %v", ids, ok)
				}
			}
		})
	}
	wg.Wait()
}
