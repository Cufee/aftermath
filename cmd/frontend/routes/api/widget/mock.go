package widget

import (
	"github.com/a-h/templ"
	"github.com/cufee/aftermath/cmd/frontend/handler"
)

var MockWidget handler.Partial = func(ctx *handler.Context) (templ.Component, error) {

	return nil, nil
}
