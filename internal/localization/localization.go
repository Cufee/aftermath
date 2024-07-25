package localization

import "golang.org/x/text/language"

type Printer func(string) string

func NewPrinter(module string, lang, fallback language.Tag) (Printer, error) {
	p, err := loadedLocalizationStrings.Printer(module, lang)
	if err != nil {
		return loadedLocalizationStrings.Printer(module, fallback)
	}
	return p, nil
}

func NewPrinterWithFallback(module string, lang language.Tag) (Printer, error) {
	p, err := loadedLocalizationStrings.Printer(module, lang)
	if err != nil {
		return loadedLocalizationStrings.Printer(module, language.English)
	}
	return p, nil
}

func ModuleKeyValues(module string, key string) (map[language.Tag]string, error) {
	return loadedLocalizationStrings.AllLanguages(module, key)
}
