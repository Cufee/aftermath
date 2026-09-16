package glossary

import (
	"sort"
	"strings"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"golang.org/x/text/language"
)

// SearchVehicles matches all localized names. Locale is retained for caller compatibility
// but does not restrict matching; callers can use Vehicle.Name to localize display names.
func SearchVehicles(locale language.Tag, query string, limit int) ([]models.Vehicle, bool) {
	if len(query) < 3 || limit <= 0 {
		return nil, false
	}

	// Both values belong to the same refresh and remain immutable after publication.
	cacheLock.RLock()
	aliases, cachedVehicles := vehicleAliases, vehiclesCache
	cacheLock.RUnlock()

	query = strings.ToLower(query)
	bestScores := make(map[string]int)
	for _, alias := range aliases {
		score := fuzzy.RankMatch(query, alias.term)
		if score < 0 {
			continue
		}
		if previous, ok := bestScores[alias.vehicleID]; !ok || score < previous {
			bestScores[alias.vehicleID] = score
		}
	}

	if len(bestScores) == 0 {
		return nil, false
	}
	ids := make([]string, 0, len(bestScores))
	for id := range bestScores {
		ids = append(ids, id)
	}
	sort.Slice(ids, func(i, j int) bool {
		if bestScores[ids[i]] == bestScores[ids[j]] {
			return ids[i] < ids[j]
		}
		return bestScores[ids[i]] < bestScores[ids[j]]
	})
	if len(ids) > limit {
		ids = ids[:limit]
	}
	vehicles := make([]models.Vehicle, 0, len(ids))
	for _, id := range ids {
		vehicles = append(vehicles, cachedVehicles[id])
	}
	return vehicles, true
}
