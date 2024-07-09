package log

import (
	"net/http"
	"os"
	"time"

	"github.com/justinas/alice"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
)

var globalLogger = zerolog.New(os.Stdout).With().Timestamp().Logger()

func SetupGlobalLogger(setup func(zerolog.Logger) zerolog.Logger) {
	globalLogger = setup(globalLogger)
}

func Debug() *zerolog.Event        { return log.Debug() }
func Info() *zerolog.Event         { return log.Info() }
func Warn() *zerolog.Event         { return log.Warn() }
func Err(err error) *zerolog.Event { return log.Err(err) }
func Error() *zerolog.Event        { return log.Error() }
func Fatal() *zerolog.Event        { return log.Fatal() }

func HTTPHandler(logger zerolog.Logger) http.Handler {
	c := alice.New()

	// Install the logger handler with default output on the console
	c = c.Append(hlog.NewHandler(logger))

	// Install some provided extra handler to set some request's context fields.
	// Thanks to that handler, all our logs will come with some prepopulated fields.
	c = c.Append(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Stringer("url", r.URL).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("")
	}))
	c = c.Append(hlog.UserAgentHandler("user_agent"))
	c = c.Append(hlog.RefererHandler("referer"))

	return c.Then(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the logger from the request's context. You can safely assume it
		// will be always there: if the handler is removed, hlog.FromRequest
		// will return a no-op logger.
		hlog.FromRequest(r).Info().Msg("served request")
	}))
}
