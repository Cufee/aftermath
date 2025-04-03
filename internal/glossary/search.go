package glossary

import (
	"strings"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"golang.org/x/text/language"
)

func SearchVehicles(locale language.Tag, query string, limit int) ([]models.Vehicle, bool) {
	if len(query) < 3 {
		return nil, false
	}

	terms, ok := vehicleNames[locale]
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
