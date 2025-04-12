package localization

import "golang.org/x/text/language"

type Printer func(string) string

func NewPrinter(module string, lang language.Tag) (Printer, error) {
	return loadedLocalizationStrings.Printer(module, lang)
}

func NewPrinterWithFallback(module string, lang, fallback language.Tag) (Printer, error) {
	return loadedLocalizationStrings.PrinterWithFallback(module, lang, fallback)
}

func ModuleKeyValues(module string, key string) (map[language.Tag]string, error) {
	return loadedLocalizationStrings.AllLanguages(module, key)
}
