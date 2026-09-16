package glossary

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/logic"
	"golang.org/x/text/language"
)

var cacheLock sync.RWMutex
var cacheTime time.Time
var vehiclesCache map[string]models.Vehicle
var vehicleAliases []vehicleAlias
var vehicleTierToIDs map[int][]string

type vehicleAlias struct {
	vehicleID string
	term      string
}

func RefreshVehicleCache(ctx context.Context, db database.GlossaryClient) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	data, err := db.GetAllVehicles(ctx)
	if err != nil {
		return err
	}

	newVehicleTierToIDs := make(map[int][]string, 11)
	newVehiclesCache := make(map[string]models.Vehicle)
	var newVehicleAliases []vehicleAlias
	seenAliases := make(map[vehicleAlias]struct{})

	for id, vehicle := range data {
		newVehiclesCache[id] = vehicle
		newVehicleTierToIDs[vehicle.Tier] = append(newVehicleTierToIDs[vehicle.Tier], vehicle.ID)

		for _, value := range vehicle.LocalizedNames {
			if value == "" {
				continue
			}
			value = strings.ToLower(value)
			nameWithTier := strings.ToLower(fmt.Sprintf("%s %s", logic.IntToRoman(vehicle.Tier), value))
			for _, term := range []string{value, nameWithTier} {
				alias := vehicleAlias{vehicleID: id, term: term}
				if _, exists := seenAliases[alias]; exists {
					continue
				}
				seenAliases[alias] = struct{}{}
				newVehicleAliases = append(newVehicleAliases, alias)
			}
		}
	}

	cacheLock.Lock()
	cacheTime = time.Now()
	vehicleTierToIDs = newVehicleTierToIDs
	vehiclesCache = newVehiclesCache
	vehicleAliases = newVehicleAliases
	cacheLock.Unlock()

	return nil
}

func GetVehicleFromCache(locale language.Tag, id string) (models.Vehicle, bool) {
	cacheLock.RLock()
	defer cacheLock.RUnlock()
	v, ok := vehiclesCache[id]
	return v, ok
}
