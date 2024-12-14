package middleware

import (
	"time"

	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/internal/permissions"
)

type MiddlewareFunc func(common.Context, func(common.Context) error) func(common.Context) error

func RequirePermissions(required ...permissions.Permissions) MiddlewareFunc {
	return func(ctx common.Context, next func(common.Context) error) func(common.Context) error {
		if !ctx.User().HasPermission(required...) {
			return func(ctx common.Context) error {
				return ctx.Reply().Send("common_error_command_missing_permissions")
			}
		}
		return next
	}
}

func ExpireAfter(expiration time.Time) MiddlewareFunc {
	return func(ctx common.Context, next func(common.Context) error) func(common.Context) error {
		if time.Now().After(expiration) {
			return func(ctx common.Context) error {
				return ctx.Reply().Format("common_error_command_expired_fmt", expiration.Unix()).Send()
			}
		}
		return next
	}
}
