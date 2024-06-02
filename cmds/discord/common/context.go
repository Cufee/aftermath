package common

import (
	"context"

	"github.com/cufee/aftermath/cmds/core"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/localization"
	"golang.org/x/text/language"
)

type contextKey int

const (
	ContextKeyUser contextKey = iota
)

type Context struct {
	context.Context
	User database.User

	Locale   language.Tag
	Localize localization.Printer

	Core core.Client
}

func NewContext(ctx context.Context, client core.Client, user database.User, locale string) *Context {
	return &Context{Context: ctx, User: user, Core: client}
}
