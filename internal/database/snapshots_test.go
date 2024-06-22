package database

// /*
// DATABASE_URL needs to be set and migrations need to be applied
// */
// func TestGetVehicleSnapshots(t *testing.T) {
// 	client, err := NewClient()
// 	assert.NoError(t, err, "new client should not error")
// 	defer client.prisma.Disconnect()

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
// 	defer cancel()

// 	client.prisma.VehicleSnapshot.FindMany().Delete().Exec(ctx)
// 	defer client.prisma.VehicleSnapshot.FindMany().Delete().Exec(ctx)

// 	createdAtVehicle1 := time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC)
// 	createdAtVehicle2 := time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC)
// 	createdAtVehicle3 := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)

// 	createdAtVehicle4 := time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC)
// 	createdAtVehicle5 := time.Date(2023, 9, 2, 0, 0, 0, 0, time.UTC)

// 	vehicle1 := VehicleSnapshot{
// 		VehicleID:   "v1",
// 		AccountID:   "a1",
// 		ReferenceID: "r1",

// 		Type:           SnapshotTypeDaily,
// 		CreatedAt:      createdAtVehicle1,
// 		LastBattleTime: createdAtVehicle1,
// 	}
// 	vehicle2 := VehicleSnapshot{
// 		VehicleID:   "v1",
// 		AccountID:   "a1",
// 		ReferenceID: "r1",

// 		Type:           SnapshotTypeDaily,
// 		CreatedAt:      createdAtVehicle2,
// 		LastBattleTime: createdAtVehicle2,
// 	}
// 	vehicle3 := VehicleSnapshot{
// 		VehicleID:   "v1",
// 		AccountID:   "a1",
// 		ReferenceID: "r1",

// 		Type:           SnapshotTypeDaily,
// 		CreatedAt:      createdAtVehicle3,
// 		LastBattleTime: createdAtVehicle3,
// 	}
// 	vehicle4 := VehicleSnapshot{
// 		VehicleID:   "v4",
// 		AccountID:   "a1",
// 		ReferenceID: "r2",

// 		Type:           SnapshotTypeDaily,
// 		CreatedAt:      createdAtVehicle4,
// 		LastBattleTime: createdAtVehicle4,
// 	}
// 	vehicle5 := VehicleSnapshot{
// 		VehicleID:   "v5",
// 		AccountID:   "a1",
// 		ReferenceID: "r2",

// 		Type:           SnapshotTypeDaily,
// 		CreatedAt:      createdAtVehicle5,
// 		LastBattleTime: createdAtVehicle5,
// 	}
// 	vehicle6 := VehicleSnapshot{
// 		VehicleID:   "v5",
// 		AccountID:   "a1",
// 		ReferenceID: "r2",

// 		Type:           SnapshotTypeDaily,
// 		CreatedAt:      createdAtVehicle5,
// 		LastBattleTime: createdAtVehicle5,
// 	}

// 	{ // create snapshots
// 		snaphots := []VehicleSnapshot{vehicle1, vehicle2, vehicle3, vehicle4, vehicle5, vehicle6}
// 		err = client.CreateVehicleSnapshots(ctx, snaphots...)
// 		assert.NoError(t, err, "create vehicle snapshot should not error")
// 	}
// 	{ // when we check created after, vehicles need to be ordered by createdAt ASC, so we expect to get vehicle2 back
// 		vehicles, err := client.GetVehicleSnapshots(ctx, "a1", "r1", SnapshotTypeDaily, WithCreatedAfter(createdAtVehicle1))
// 		assert.NoError(t, err, "get vehicle snapshot error")
// 		assert.Len(t, vehicles, 1, "should return exactly 1 snapshots")
// 		assert.True(t, vehicles[0].CreatedAt.Equal(createdAtVehicle2), "wrong vehicle snapshot returned", vehicles)
// 	}
// 	{ // when we check created before, vehicles need to be ordered by createdAt DESC, so we expect to get vehicle2 back
// 		vehicles, err := client.GetVehicleSnapshots(ctx, "a1", "r1", SnapshotTypeDaily, WithCreatedBefore(createdAtVehicle3))
// 		assert.NoError(t, err, "get vehicle snapshot error")
// 		assert.Len(t, vehicles, 1, "should return exactly 1 snapshots")
// 		assert.True(t, vehicles[0].CreatedAt.Equal(createdAtVehicle2), "wrong vehicle snapshot returned", vehicles)
// 	}
// 	{ // make sure only 1 vehicle is returned per ID
// 		vehicles, err := client.GetVehicleSnapshots(ctx, "a1", "r2", SnapshotTypeDaily, WithCreatedBefore(createdAtVehicle5.Add(time.Hour)))
// 		assert.NoError(t, err, "get vehicle snapshot error")
// 		assert.Len(t, vehicles, 2, "should return exactly 2 snapshots")
// 		assert.NotEqual(t, vehicles[0].ID, vehicles[1].ID, "each vehicle id should only be returned once", vehicles)
// 	}
// 	{ // get a cehicle with a specific id
// 		vehicles, err := client.GetVehicleSnapshots(ctx, "a1", "r2", SnapshotTypeDaily, WithVehicleIDs([]string{"v5"}))
// 		assert.NoError(t, err, "get vehicle snapshot error")
// 		assert.Len(t, vehicles, 1, "should return exactly 1 snapshots")
// 		assert.NotEqual(t, vehicles[0].ID, "v5", "incorrect vehicle returned", vehicles)
// 	}
// 	{ // this should return no result
// 		vehicles, err := client.GetVehicleSnapshots(ctx, "a1", "r1", SnapshotTypeDaily, WithCreatedBefore(createdAtVehicle1))
// 		assert.NoError(t, err, "no results from a raw query does not trigger an error")
// 		assert.Len(t, vehicles, 0, "return should have no results", vehicles)
// 	}
// }
