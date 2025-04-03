package glossary

func TierVehicleIDs(tier int) ([]string, bool) {
	ids, ok := vehicleTierToIDs[tier]
	return ids, ok
}
