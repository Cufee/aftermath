package middleware

import (
	"github.com/cufee/aftermath/cmds/discord/common"
	"github.com/cufee/aftermath/internal/permissions"
)

type MiddlewareFunc func(*common.Context, func(*common.Context) error) func(*common.Context) error

func RequirePermissions(required permissions.Permissions) MiddlewareFunc {
	return func(ctx *common.Context, next func(*common.Context) error) func(*common.Context) error {
		if !ctx.User.Permissions.Has(required) {
			return func(ctx *common.Context) error {
				return ctx.Reply("common_error_command_missing_permissions")
			}
		}
		return next
	}
}
