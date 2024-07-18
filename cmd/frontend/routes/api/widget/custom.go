package widget

import (
	"errors"
	"net/http"
	"net/url"
	"reflect"
	"strconv"

	"github.com/a-h/templ"
	"github.com/cufee/aftermath/cmd/frontend/components/widget"
	"github.com/cufee/aftermath/cmd/frontend/handler"
	"github.com/cufee/aftermath/cmd/frontend/routes/app/widgets"
	"github.com/cufee/aftermath/internal/constants"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
)

type widgetSettingsPayload struct {
	Title               *string `form:"widget_title"`
	AccountID           *string `form:"account_id"`
	Realm               *string `form:"account_realm"`
	Nickname            *string `form:"account_nickname"`
	ShowRatingOverview  *bool   `form:"rating_overview"`
	ShowUnratedOverview *bool   `form:"unrated_overview"`
	ShowVehicles        *bool   `form:"unrated_vehicles"`
	VehicleLimit        *int    `form:"vehicle_limit"`
}

func (p *widgetSettingsPayload) updateOptions(s *models.WidgetOptions) {
	if s == nil {
		return
	}
	if v := p.Title; v != nil {
		s.Title = *v
	}
	if v := p.AccountID; v != nil {
		s.AccountID = *v
	}
	if v := p.ShowRatingOverview; v != nil {
		s.Style.RatingOverview.Visible = *v
	}
	if v := p.ShowUnratedOverview; v != nil {
		s.Style.UnratedOverview.Visible = *v
	}
	if v := p.ShowVehicles; v != nil {
		s.Style.Vehicles.Visible = *v
	}
	if v := p.VehicleLimit; v != nil && *v >= 0 && *v <= 10 {
		s.Style.Vehicles.Limit = *v
	}
}

func (p *widgetSettingsPayload) parse(values url.Values) {
	v := reflect.ValueOf(p).Elem() // Get the value of the struct

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := v.Type().Field(i)
		formTag := fieldType.Tag.Get("form")

		if formTag == "" {
			continue
		}

		if formValues, ok := values[formTag]; ok && len(formValues) > 0 {
			switch field.Kind() {
			case reflect.Ptr:
				switch fieldType.Type.Elem().Kind() {
				case reflect.String:
					strVal := formValues[0]
					field.Set(reflect.ValueOf(&strVal))
				case reflect.Bool:
					boolVal, err := strconv.ParseBool(formValues[0])
					if err == nil {
						field.Set(reflect.ValueOf(&boolVal))
					}
				case reflect.Int:
					intVal, err := strconv.Atoi(formValues[0])
					if err == nil {
						field.Set(reflect.ValueOf(&intVal))
					}
				}
			}
		}
	}
}

var CreateCustomWidget handler.Partial = func(ctx *handler.Context) (templ.Component, error) {
	user, err := ctx.SessionUser()
	if err != nil {
		return nil, ctx.Redirect("/login", http.StatusTemporaryRedirect)
	}

	existing, err := ctx.Database().GetUserWidgetSettings(ctx.Context, user.ID, nil)
	if err != nil && !database.IsNotFound(err) {
		return nil, ctx.Error(err, "failed to get user widgets")
	}
	if len(existing) >= constants.WidgetAccountLimit {
		return nil, ctx.Error(errors.New("widget limit reached"))
	}

	form, err := ctx.FormValues()
	if err != nil {
		return nil, ctx.Error(err, "invalid form data")
	}

	var data widgetSettingsPayload
	data.parse(form)

	var settings models.WidgetOptions
	settings.Style = models.DefaultWidgetStyle
	settings.UserID = user.ID
	data.updateOptions(&settings)

	if settings.AccountID == "" {
		var account models.Account
		if data.Realm != nil {
			account.Realm = *data.Realm
		}
		if data.Nickname != nil {
			account.Nickname = *data.Nickname
		}
		if data.AccountID != nil {
			account.ID = *data.AccountID
		}
		return widgets.NewWidgetPage(widget.WidgetWithAccount{WidgetOptions: settings, Account: account}, map[string]string{"account_id": "You need to select a user from the list of search results."}), nil
	}

	created, err := ctx.Database().CreateWidgetSettings(ctx.Context, user.ID, settings)
	if err != nil {
		return nil, ctx.Error(err, "failed to create widget settings")
	}
	return nil, ctx.Redirect("/app/widgets/"+created.ID, http.StatusTemporaryRedirect)
}

var UpdateCustomWidget handler.Partial = func(ctx *handler.Context) (templ.Component, error) {
	user, err := ctx.SessionUser()
	if err != nil {
		return nil, ctx.Redirect("/login", http.StatusTemporaryRedirect)
	}

	widgetID := ctx.Path("widgetId")
	if widgetID == "" {
		return nil, ctx.Error(errors.New("invalid widget id"))
	}

	form, err := ctx.FormValues()
	if err != nil {
		return nil, ctx.Error(err, "invalid form data")
	}

	var data widgetSettingsPayload
	data.parse(form)

	settings, err := ctx.Database().GetWidgetSettings(ctx.Context, widgetID)
	if err != nil {
		if database.IsNotFound(err) {
			return nil, ctx.Error(errors.New("widget not found"))
		}
		return nil, ctx.Error(err, "failed to get widget settings")
	}
	if settings.UserID != user.ID {
		return nil, ctx.Redirect("/login", http.StatusTemporaryRedirect)
	}
	data.updateOptions(&settings)

	if settings.AccountID == "" {
		var account models.Account
		if data.Realm != nil {
			account.Realm = *data.Realm
		}
		if data.Nickname != nil {
			account.Nickname = *data.Nickname
		}
		if data.AccountID != nil {
			account.ID = *data.AccountID
		}
		return widgets.WidgetConfiguratorPage(widget.WidgetWithAccount{WidgetOptions: settings, Account: account}, map[string]string{"account_id": "You need to select a user from the list of search results."}), nil
	}

	updated, err := ctx.Database().UpdateWidgetSettings(ctx.Context, settings.ID, settings)
	if err != nil {
		return nil, ctx.Error(err, "failed to update widget settings")
	}

	account, err := ctx.Database().GetAccountByID(ctx.Context, updated.AccountID)
	if database.IsNotFound(err) {
		account, err = ctx.Fetch().Account(ctx.Context, updated.AccountID)
	}
	if err != nil {
		return nil, ctx.Error(err, "failed to get updated widget account")
	}

	return widgets.WidgetConfiguratorPage(widget.WidgetWithAccount{WidgetOptions: updated, Account: account}, nil), nil
}

var QuickAction handler.Partial = func(ctx *handler.Context) (templ.Component, error) {
	user, err := ctx.SessionUser()
	if err != nil {
		return nil, ctx.Redirect("/login", http.StatusTemporaryRedirect)
	}

	widgetID := ctx.Path("widgetId")
	if widgetID == "" {
		return nil, ctx.Error(errors.New("invalid widget id"))
	}

	form, err := ctx.FormValues()
	if err != nil {
		return nil, ctx.Error(err, "invalid form data")
	}

	var data widgetSettingsPayload
	data.parse(form)

	settings, err := ctx.Database().GetWidgetSettings(ctx.Context, widgetID)
	if err != nil {
		if database.IsNotFound(err) {
			return nil, ctx.Error(errors.New("widget not found"))
		}
		return nil, ctx.Error(err, "failed to get widget settings")
	}
	if settings.UserID != user.ID {
		return nil, ctx.Redirect("/login", http.StatusTemporaryRedirect)
	}
	data.updateOptions(&settings)

	if settings.AccountID == "" {
		return nil, ctx.Error(errors.New("widget has no account id set"))
	}

	updated, err := ctx.Database().UpdateWidgetSettings(ctx.Context, settings.ID, settings)
	if err != nil {
		return nil, ctx.Error(err, "failed to update widget settings")
	}

	return widget.QuickSettingButton(ctx.Query("field"), updated), nil
}
