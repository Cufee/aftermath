package models

import (
	"fmt"

	assets "github.com/cufee/aftermath-assets/types"
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
