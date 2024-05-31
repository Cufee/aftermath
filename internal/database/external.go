package database

import (
	"github.com/cufee/aftermath/internal/database/prisma/db"
	"github.com/cufee/aftermath/internal/localization"
)

type GlossaryVehicle struct {
	db.VehicleModel
}

func (v GlossaryVehicle) Name(printer localization.Printer) string {
	return printer("name")
}
