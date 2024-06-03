package localization

import "golang.org/x/text/language"

type Printer func(string) string

func NewPrinter(module string, lang language.Tag) (Printer, error) {
	return loadedLocalizationStrings.Printer(module, lang)
}

func ModuleKeyValues(module string, key string) (map[language.Tag]string, error) {
	return loadedLocalizationStrings.AllLanguages(module, key)
}
