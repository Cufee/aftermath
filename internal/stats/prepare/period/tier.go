package period

// func calculateAvgTier(vehicles map[string]fetch.VehicleStatsFrame, glossary map[string]struct{}) fetch.Value {
// 	var weightedTierTotal float32

// 	for _, vehicle := range vehicles {
// 		if data, ok := glossary[vehicle.VehicleID]; ok && data.Tier > 0 {
// 			battlesTotal += vehicle.Battles
// 			weightedTierTotal += float32(vehicle.Battles * data.Tier)
// 		}
// 	}
// 	if battlesTotal == 0 {
// 		return fetch.InvalidValue
// 	}

// 	return weightedTierTotal / float32(battlesTotal)
// }
