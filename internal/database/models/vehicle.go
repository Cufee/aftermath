package models

import (
	"fmt"

	"github.com/cufee/aftermath/internal/json"

	assets "github.com/cufee/aftermath-assets/types"
	"github.com/cufee/aftermath/internal/database/gen/model"
	"golang.org/x/text/language"
)

type Vehicle assets.Vehicle

func (v Vehicle) Name(locale language.Tag) string {
	if n := v.LocalizedNames[locale]; n != "" {
		return n
	}
	if n := v.LocalizedNames[language.English]; n != "" {
		return n
	}
	return fmt.Sprintf("Secret Tank %s", v.ID)
}

func ToVehicle(record *model.Vehicle) Vehicle {
	v := Vehicle{
		ID:             record.ID,
		Tier:           int(record.Tier),
		LocalizedNames: make(map[language.Tag]string, 0),
	}
	json.Unmarshal([]byte(record.LocalizedNames), &v.LocalizedNames)
	return v
}
