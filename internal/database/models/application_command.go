package models

import (
	"time"

	"github.com/cufee/aftermath/internal/database/gen/model"
)

type ApplicationCommand struct {
	ID      string
	Hash    string
	Name    string
	Version string
}

func ToApplicationCommand(record *model.ApplicationCommand) ApplicationCommand {
	return ApplicationCommand{
		ID:      record.ID,
		Name:    record.Name,
		Hash:    record.OptionsHash,
		Version: record.Version,
	}
}

func FromApplicationCommand(record *ApplicationCommand) model.ApplicationCommand {
	return model.ApplicationCommand{
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		ID:          record.ID,
		Name:        record.Name,
		OptionsHash: record.Hash,
		Version:     record.Version,
	}
}
