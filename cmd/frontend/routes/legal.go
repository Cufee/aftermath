package routes

import (
	"github.com/a-h/templ"
	"github.com/cufee/aftermath/cmd/frontend/components"
	"github.com/cufee/aftermath/cmd/frontend/handler"
	"github.com/cufee/aftermath/cmd/frontend/layouts"
)

var TermsOfService handler.Page = func(ctx *handler.Context) (handler.Layout, templ.Component, error) {
	return layouts.Main, components.RemoteContentPage("https://byvko-dev.github.io/legal/terms-of-service-partial"), nil
}

var PrivacyPolicy handler.Page = func(ctx *handler.Context) (handler.Layout, templ.Component, error) {
	return layouts.Main, components.RemoteContentPage("https://byvko-dev.github.io/legal/privacy-policy"), nil
}
