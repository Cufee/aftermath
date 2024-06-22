package database

import (
	"context"

	"github.com/cufee/aftermath/internal/database/models"
)

// func (v Vehicle) FromModel(model db.VehicleModel) Vehicle {
// 	v.ID = model.ID
// 	v.Tier = model.Tier
// 	v.Type = model.Type
// 	v.Nation = model.Nation
// 	v.LocalizedNames = make(map[string]string)
// 	if model.LocalizedNamesEncoded != nil {
// 		err := encoding.DecodeGob(model.LocalizedNamesEncoded, &v.LocalizedNames)
// 		if err != nil {
// 			log.Err(err).Str("id", v.ID).Msg("failed to decode vehicle localized names")
// 		}
// 	}
// 	return v
// }

// func (v Vehicle) EncodeNames() []byte {
// 	encoded, _ := encoding.EncodeGob(v.LocalizedNames)
// 	return encoded
// }

func (c *libsqlClient) UpsertVehicles(ctx context.Context, vehicles map[string]models.Vehicle) error {
	// if len(vehicles) < 1 {
	// 	return nil
	// }

	// var transactions []transaction.Transaction
	// for id, data := range vehicles {
	// 	transactions = append(transactions, c.prisma.Vehicle.
	// 		UpsertOne(db.Vehicle.ID.Equals(id)).
	// 		Create(
	// 			db.Vehicle.ID.Set(id),
	// 			db.Vehicle.Tier.Set(data.Tier),
	// 			db.Vehicle.Type.Set(data.Type),
	// 			db.Vehicle.Class.Set(data.Class),
	// 			db.Vehicle.Nation.Set(data.Nation),
	// 			db.Vehicle.LocalizedNamesEncoded.Set(data.EncodeNames()),
	// 		).
	// 		Update(
	// 			db.Vehicle.Tier.Set(data.Tier),
	// 			db.Vehicle.Type.Set(data.Type),
	// 			db.Vehicle.Class.Set(data.Class),
	// 			db.Vehicle.Nation.Set(data.Nation),
	// 			db.Vehicle.LocalizedNamesEncoded.Set(data.EncodeNames()),
	// 		).Tx(),
	// 	)
	// }

	// return c.prisma.Prisma.Transaction(transactions...).Exec(ctx)
	return nil
}

func (c *libsqlClient) GetVehicles(ctx context.Context, ids []string) (map[string]models.Vehicle, error) {
	// 	if len(ids) < 1 {
	// 		return nil, nil
	// 	}

	// 	models, err := c.prisma.Vehicle.FindMany(db.Vehicle.ID.In(ids)).Exec(ctx)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	vehicles := make(map[string]models.Vehicle)
	// 	for _, model := range models {
	// 		vehicles[model.ID] = vehicles[model.ID].FromModel(model)
	// 	}

	return vehicles, nil
}
