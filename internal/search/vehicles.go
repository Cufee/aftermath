package search

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/logic"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"golang.org/x/text/language"
)

var cacheLock sync.Mutex
var cacheTime time.Time
var vehiclesCache map[string]models.Vehicle
var vehicleNames map[language.Tag][]string
var vehicleNameToID map[language.Tag]map[string]string

func RefreshVehicleCacheIfStale(ctx context.Context, db database.GlossaryClient) error {
	cacheLock.Lock()
	t := cacheTime
	cacheLock.Unlock()
	if time.Since(t) > time.Hour*1 {
		return RefreshVehicleCache(ctx, db)
	}
	return nil
}

func RefreshVehicleCache(ctx context.Context, db database.GlossaryClient) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	data, err := db.GetAllVehicles(ctx)
	if err != nil {
		return err
	}

	newVehicleNameToID := make(map[language.Tag]map[string]string)
	newVehiclesCache := make(map[string]models.Vehicle)
	newVehicleNames := make(map[language.Tag][]string)

	for id, vehicle := range data {
		for tag, value := range vehicle.LocalizedNames {
			value = strings.ToLower(value)
			newVehiclesCache[id] = vehicle

			localeNameToID := newVehicleNameToID[tag]
			if localeNameToID == nil {
				localeNameToID = make(map[string]string)
			}

			nameWithTier := strings.ToLower(fmt.Sprintf("%s %s", logic.IntToRoman(vehicle.Tier), value))
			localeNameToID[nameWithTier] = id
			newVehicleNameToID[tag] = localeNameToID
			newVehicleNames[tag] = append(newVehicleNames[tag], nameWithTier)
		}
	}

	cacheLock.Lock()
	cacheTime = time.Now()
	vehicleNameToID = newVehicleNameToID
	vehiclesCache = newVehiclesCache
	vehicleNames = newVehicleNames
	cacheLock.Unlock()

	return nil
}

func GetVehicleFromCache(locale language.Tag, id string) (models.Vehicle, bool) {
	cacheLock.Lock()
	v, ok := vehiclesCache[id]
	cacheLock.Unlock()
	return v, ok
}

func SearchVehicles(locale language.Tag, query string, limit int) ([]models.Vehicle, bool) {
	if len(query) < 3 {
		return nil, false
	}

	cacheLock.Lock()
	terms, ok := vehicleNames[locale]
	cacheLock.Unlock()
	if !ok {
		return nil, false
	}

	var vehicles []models.Vehicle
	result := fuzzy.RankFind(strings.ToLower(query), terms)
	for i, entry := range result {
		if i >= limit {
			break
		}
		cacheLock.Lock()
		id, ok := vehicleNameToID[locale][entry.Target]
		cacheLock.Unlock()
		if !ok {
			continue
		}
		cacheLock.Lock()
		v, ok := vehiclesCache[id]
		cacheLock.Unlock()
		vehicles = append(vehicles, v)
	}

	if len(vehicles) < 1 {
		return nil, false
	}
	return vehicles, true
}
