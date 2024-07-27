package middleware

import (
	"runtime/debug"

	"github.com/cufee/aftermath/cmd/frontend/handler"
	"github.com/cufee/aftermath/internal/log"
)

var CatchPanic handler.Middleware = func(ctx *handler.Context, next func(ctx *handler.Context) error) func(ctx *handler.Context) error {
	return func(ctx *handler.Context) error {
		defer func() {
			if rec := recover(); rec != nil {
				log.Error().Str("stack", string(debug.Stack())).Msg("panic in http handler")
			}
		}()
		return next(ctx)
	}
}
