package models

import "golang.org/x/text/language"

type Vehicle struct {
	ID     string
	Tier   int
	Type   string
	Class  string
	Nation string

	LocalizedNames map[string]string
}

func (v Vehicle) Name(locale language.Tag) string {
	if n := v.LocalizedNames[locale.String()]; n != "" {
		return n
	}
	if n := v.LocalizedNames[language.English.String()]; n != "" {
		return n
	}
	return "Secret Tank"
}
