package common

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/internal/localization"

	"github.com/cufee/aftermath/internal/log"
	"golang.org/x/text/language"
)

func LocaleToLanguageTag(locale discordgo.Locale) language.Tag {
	// Some discord locale tags don't match the standard
	switch locale {
	case discordgo.EnglishUS:
		fallthrough
	case discordgo.EnglishGB:
		return language.English

	case "es-419":
		fallthrough
	case discordgo.SpanishES:
		return language.Spanish

	case discordgo.PortugueseBR:
		return language.BrazilianPortuguese

	case "":
		return language.English

	default:
		tag, err := language.Parse(string(locale))
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

	case language.Spanish:
		fallthrough
	case language.LatinAmericanSpanish:
		return discordgo.SpanishES

	case language.BrazilianPortuguese:
		fallthrough
	case language.EuropeanPortuguese:
		fallthrough
	case language.Portuguese:
		return discordgo.PortugueseBR

	default:
		locale := discordgo.Locale(tag.String())
		_, ok := discordgo.Locales[locale]
		if !ok {
			return discordgo.EnglishUS
		}
		return locale
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
