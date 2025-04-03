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

var cacheLock sync.Mutex
var cacheTime time.Time
var vehiclesCache map[string]models.Vehicle
var vehicleNames map[language.Tag][]string
var vehicleNameToID map[language.Tag]map[string]string
var vehicleTierToIDs map[int][]string

func RefreshVehicleCache(ctx context.Context, db database.GlossaryClient) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	data, err := db.GetAllVehicles(ctx)
	if err != nil {
		return err
	}

	newVehicleTierToIDs := make(map[int][]string, 11)
	newVehicleNameToID := make(map[language.Tag]map[string]string)
	newVehiclesCache := make(map[string]models.Vehicle)
	newVehicleNames := make(map[language.Tag][]string)

	for id, vehicle := range data {
		newVehicleTierToIDs[vehicle.Tier] = append(newVehicleTierToIDs[vehicle.Tier], vehicle.ID)

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
	vehicleTierToIDs = newVehicleTierToIDs
	vehicleNameToID = newVehicleNameToID
	vehiclesCache = newVehiclesCache
	vehicleNames = newVehicleNames
	cacheLock.Unlock()

	return nil
}

func GetVehicleFromCache(locale language.Tag, id string) (models.Vehicle, bool) {
	v, ok := vehiclesCache[id]
	return v, ok
}
