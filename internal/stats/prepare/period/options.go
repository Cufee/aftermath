package period

var defaultOptions = options{localePrinter: func(s string) string { return s }}

type options struct {
	localePrinter func(string) string
}

type Option func(*options)

func WithPrinter(printer func(string) string) func(*options) {
	return func(o *options) { o.localePrinter = printer }
}
