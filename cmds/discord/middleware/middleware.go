package middleware

import (
	"github.com/cufee/aftermath/cmds/discord/common"
)

type MiddlewareFunc func(*common.Context, func(*common.Context)) func(*common.Context)
