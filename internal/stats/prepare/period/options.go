package period

import "golang.org/x/text/language"

var defaultOptions = options{localePrinter: func(s string) string { return s }, locale: language.English}

type options struct {
	localePrinter func(string) string
	locale        language.Tag
}

type Option func(*options)

func WithPrinter(printer func(string) string, locale language.Tag) func(*options) {
	return func(o *options) { o.localePrinter = printer; o.locale = locale }
}
