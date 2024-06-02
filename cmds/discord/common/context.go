package common

import (
	"context"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/localization"
	"golang.org/x/text/language"
)

type Context struct {
	context.Context
	User database.User

	Locale   language.Tag
	Localize localization.Printer
}

func NewContext(ctx context.Context, user database.User, locale string) *Context {
	return &Context{Context: ctx, User: user}
}
