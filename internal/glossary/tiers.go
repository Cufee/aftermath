package glossary

func TierVehicleIDs(tier int) ([]string, bool) {
	cacheLock.RLock()
	defer cacheLock.RUnlock()
	ids, ok := vehicleTierToIDs[tier]
	return ids, ok
}
