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

func (record *ApplicationCommand) Model() model.ApplicationCommand {
	return model.ApplicationCommand{
		ID:          record.ID,
		CreatedAt:   TimeToString(time.Now()),
		UpdatedAt:   TimeToString(time.Now()),
		Name:        record.Name,
		OptionsHash: record.Hash,
		Version:     record.Version,
	}
}
