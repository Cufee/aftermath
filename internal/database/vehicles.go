package database

import (
	"context"
	"encoding/json"

	"github.com/cufee/aftermath/internal/database/prisma/db"
	"github.com/rs/zerolog/log"
	"github.com/steebchen/prisma-client-go/runtime/transaction"
	"golang.org/x/text/language"
)

type Vehicle struct {
	ID     string
	Tier   int
	Type   string
	Class  string
	Nation string

	LocalizedNames map[language.Tag]string
}

func (v Vehicle) FromModel(model db.VehicleModel) Vehicle {
	v.ID = model.ID
	v.Tier = model.Tier
	v.Type = model.Type
	v.Nation = model.Nation
	v.LocalizedNames = make(map[language.Tag]string)
	if model.LocalizedNamesEncoded != "" {
		_ = json.Unmarshal([]byte(model.LocalizedNamesEncoded), &v.LocalizedNames)
	}
	return v
}

func (v Vehicle) Name(locale language.Tag) string {
	if n := v.LocalizedNames[locale]; n != "" {
		return n
	}
	if n := v.LocalizedNames[language.English]; n != "" {
		return n
	}
	return "Secret Tank"
}

func (c *client) UpsertVehicles(ctx context.Context, vehicles map[string]Vehicle) error {
	if len(vehicles) < 1 {
		return nil
	}

	var transactions []transaction.Transaction
	for id, data := range vehicles {
		encoded, err := json.Marshal(data.LocalizedNames)
		if err != nil {
			log.Err(err).Str("id", id).Msg("failed to encode a stats frame for vehicle averages")
			continue
		}

		transactions = append(transactions, c.Raw.Vehicle.
			UpsertOne(db.Vehicle.ID.Equals(id)).
			Create(
				db.Vehicle.ID.Set(id),
				db.Vehicle.Tier.Set(data.Tier),
				db.Vehicle.Type.Set(data.Type),
				db.Vehicle.Class.Set(data.Class),
				db.Vehicle.Nation.Set(data.Nation),
				db.Vehicle.LocalizedNamesEncoded.Set(string(encoded)),
			).
			Update(
				db.Vehicle.Tier.Set(data.Tier),
				db.Vehicle.Type.Set(data.Type),
				db.Vehicle.Class.Set(data.Class),
				db.Vehicle.Nation.Set(data.Nation),
				db.Vehicle.LocalizedNamesEncoded.Set(string(encoded)),
			).Tx(),
		)
	}

	return c.Raw.Prisma.Transaction(transactions...).Exec(ctx)
}

func (c *client) GetVehicles(ctx context.Context, ids []string) (map[string]Vehicle, error) {
	if len(ids) < 1 {
		return nil, nil
	}

	models, err := c.Raw.Vehicle.FindMany(db.Vehicle.ID.In(ids)).Exec(ctx)
	if err != nil {
		return nil, err
	}

	vehicles := make(map[string]Vehicle)
	for _, model := range models {
		vehicles[model.ID] = vehicles[model.ID].FromModel(model)
	}

	return vehicles, nil
}
