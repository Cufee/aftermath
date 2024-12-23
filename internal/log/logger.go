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

type DokployTypeHook struct{}

func (h DokployTypeHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	e.Str("type", "["+level.String()+"]")
}

var globalLogger = zerolog.New(os.Stdout).Hook(DokployTypeHook{}).With().Timestamp().Logger()

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

func NewMiddleware(logger zerolog.Logger, ignorePath ...string) func(http.Handler) http.Handler {
	c := alice.New()

	c = c.Append(hlog.NewHandler(logger))
	c = c.Append(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		if strings.HasPrefix(r.URL.Path, "/assets/") {
			return
		}
		for _, path := range ignorePath {
			if path == r.URL.Path {
				return
			}
		}
		hlog.FromRequest(r).Info().
			Int("status", status).
			Str("method", r.Method).
			Str("duration", duration.String()).
			Str("url", r.URL.Path).
			Msg("")
	}))

	return c.Then
}
