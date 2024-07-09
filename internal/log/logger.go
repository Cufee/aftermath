package log

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/justinas/alice"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

var globalLogger = zerolog.New(os.Stdout).With().Timestamp().Logger()

func SetupGlobalLogger(setup func(zerolog.Logger) zerolog.Logger) {
	globalLogger = setup(globalLogger)
}

func Logger() zerolog.Logger { return globalLogger }

func Debug() *zerolog.Event        { return globalLogger.Debug() }
func Info() *zerolog.Event         { return globalLogger.Info() }
func Warn() *zerolog.Event         { return globalLogger.Warn() }
func Err(err error) *zerolog.Event { return globalLogger.Err(err) }
func Error() *zerolog.Event        { return globalLogger.Error() }
func Fatal() *zerolog.Event        { return globalLogger.Fatal() }

func NewMiddleware(logger zerolog.Logger) func(http.Handler) http.Handler {
	c := alice.New()

	c = c.Append(hlog.NewHandler(logger))
	c = c.Append(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		if strings.HasPrefix(r.URL.Path, "/assets/") {
			return
		}
		hlog.FromRequest(r).Info().
			Int("status", status).
			Str("method", r.Method).
			Str("duration", duration.String()).
			Stringer("url", r.URL).
			Msg("")
	}))

	return c.Then
}
