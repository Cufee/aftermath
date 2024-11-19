package middleware

import (
	"github.com/cufee/aftermath/cmd/frontend/handler"
)

func AllowOrigin(value string) handler.Middleware {
	return func(ctx *handler.Context, next func(ctx *handler.Context) error) func(ctx *handler.Context) error {
		ctx.SetHeader("Access-Control-Allow-Origin", value)
		return next
	}
}
