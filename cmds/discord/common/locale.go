package common

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/internal/localization"

	"github.com/rs/zerolog/log"
	"golang.org/x/text/language"
)

func LocaleToLanguageTag(locale discordgo.Locale) language.Tag {
	// Some discord locale tags don't match the standard
	switch code := locale.String(); code {
	case "":
		return language.English
	case "en-GB":
		fallthrough
	case "en-US":
		return language.English
	case "es-419":
		return language.LatinAmericanSpanish
	case "es-ES":
		return language.Spanish

	default:
		tag, err := language.Parse(code)
		if err != nil {
			return language.English
		}
		return tag
	}
}

func LanguageToLocale(tag language.Tag) discordgo.Locale {
	// Some discord locale tags don't match the standard
	switch tag {
	case language.BritishEnglish:
		fallthrough
	case language.AmericanEnglish:
		fallthrough
	case language.English:
		return discordgo.EnglishUS

	case language.LatinAmericanSpanish:
		return discordgo.SpanishLATAM
	case language.Spanish:
		return discordgo.SpanishES

	default:
		return discordgo.EnglishUS
	}
}

func LocalizeKey(key string) map[discordgo.Locale]string {
	localized := make(map[discordgo.Locale]string)

	values, err := localization.ModuleKeyValues("discord", key)
	if err != nil {
		log.Err(err).Str("key", key).Msg("failed to get localizations for a key")
		return nil
	}

	for tag, value := range values {
		if value == "" {
			continue
		}
		localized[LanguageToLocale(tag)] = value
	}

	return localized
}
