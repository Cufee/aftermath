package localization

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/cufee/aftermath/internal/files"
	"golang.org/x/text/language"
)

type localizationString struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Notes   string `json:"notes"`
	Context string `json:"context"`
}

type localizationStrings struct {
	// format [module_name][lang_tag][key]value
	data map[string]map[language.Tag]map[string]string
}

func (l *localizationStrings) AddStrings(module string, lang language.Tag, newStrings []localizationString) error {
	if l.data == nil {
		l.data = make(map[string]map[language.Tag]map[string]string)
	}
	if l.data[module] == nil {
		l.data[module] = make(map[language.Tag]map[string]string)
	}
	if l.data[module][lang] == nil {
		l.data[module][lang] = make(map[string]string)
	}
	for _, data := range newStrings {
		if strings.HasPrefix("_skip", data.Key) {
			continue
		}
		if _, ok := l.data[module][lang][data.Key]; ok {
			return fmt.Errorf("%s/%s/%s already registered", module, lang.String(), data.Key)
		}
		l.data[module][lang][data.Key] = data.Value
	}

	return nil
}

func (l *localizationStrings) Printer(module string, lang language.Tag) (Printer, error) {
	if l.data[module] == nil {
		return nil, fmt.Errorf("module %s not registered", module)
	}
	if l.data[module][lang] == nil {
		return nil, fmt.Errorf("language %s not registered", lang.String())
	}
	return func(s string) string {
		return l.data[module][lang][s]
	}, nil
}

var loadedLocalizationStrings localizationStrings

func LoadAssets(assets fs.FS, directory string) error {
	dirFiles, err := files.ReadDirFiles(assets, directory)
	if err != nil {
		return err
	}

	for name, data := range dirFiles {
		if !strings.HasSuffix(name, ".json") {
			continue
		}

		langPath, err := filepath.Rel(directory, name)
		if err != nil {
			return fmt.Errorf("failed to trim the locale directory: %w", err)
		}

		nameSlice := strings.Split(langPath, "/")
		if len(nameSlice) != 2 {
			return errors.New("bad localization file structure found, expected {locale}/{module_name}.json")
		}

		tag, err := language.Parse(nameSlice[0])
		if err != nil {
			return fmt.Errorf("failed to parse language tag: %w", err)
		}

		var localeStrings []localizationString
		err = json.Unmarshal(data, &localeStrings)
		if err != nil {
			return fmt.Errorf("failed to unmarshal a locale file: %w", err)
		}

		module := strings.Split(nameSlice[1], ".")[0]
		err = loadedLocalizationStrings.AddStrings(module, tag, localeStrings)
		if err != nil {
			return err
		}
	}

	return nil
}
