package widget

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"time"

	"github.com/a-h/templ"
	"github.com/cufee/aftermath/cmd/frontend/components/widget"
	"github.com/cufee/aftermath/cmd/frontend/handler"
	"github.com/cufee/aftermath/cmd/frontend/routes/app/widgets"
	"github.com/cufee/aftermath/internal/constants"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/logic"
	"github.com/cufee/aftermath/internal/realtime"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/pkg/errors"
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
		return nil, ctx.Err(err, "failed to get user widgets")
	}
	if len(existing) >= constants.WidgetAccountLimit {
		return nil, ctx.Error("widget limit reached")
	}

	form, err := ctx.FormValues()
	if err != nil {
		return nil, ctx.Err(err, "invalid form data")
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
			account.Realm = types.Realm(*data.Realm)
		}
		if data.Nickname != nil {
			account.Nickname = *data.Nickname
		}
		if data.AccountID != nil {
			account.ID = *data.AccountID
		}
		return widgets.NewWidgetPage(widget.WidgetWithAccount{WidgetOptions: settings, Account: account}, map[string]string{"account_id": "You need to select a user from the list of search results."}), nil
	}

	account, err := ctx.Database().GetAccountByID(ctx.Context, settings.AccountID)
	if database.IsNotFound(err) {
		account, err = ctx.Fetch().Account(ctx.Context, settings.AccountID)
	}
	if err != nil || account.Private {
		return widgets.NewWidgetPage(widget.WidgetWithAccount{WidgetOptions: settings, Account: account}, map[string]string{"account_id": "You have selected a private account - stats are not available for private accounts."}), nil
	}

	created, err := ctx.Database().CreateWidgetSettings(ctx.Context, user.ID, settings)
	if err != nil {
		return nil, ctx.Err(err, "failed to create widget settings")
	}

	c, cancel := context.WithTimeout(ctx.Context, time.Second*5)
	defer cancel()

	realm, ok := ctx.Wargaming().RealmFromID(created.AccountID)
	if !ok {
		return nil, ctx.Error("failed to create widget settings")
	}

	_, err = logic.RecordAccountSnapshots(c, ctx.Wargaming(), ctx.Database(), realm, logic.WithReference(created.AccountID, created.ID), logic.WithType(models.SnapshotTypeWidget))
	if err != nil {
		return nil, ctx.Err(err, "failed to create a new widget session")
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
		return nil, ctx.Error("invalid widget id")
	}

	form, err := ctx.FormValues()
	if err != nil {
		return nil, ctx.Err(err, "invalid form data")
	}

	var data widgetSettingsPayload
	data.parse(form)

	settings, err := ctx.Database().GetWidgetSettings(ctx.Context, widgetID)
	if err != nil {
		if database.IsNotFound(err) {
			return nil, ctx.Error("widget not found")
		}
		return nil, ctx.Err(err, "failed to get widget settings")
	}
	if settings.UserID != user.ID {
		return nil, ctx.Redirect("/login", http.StatusTemporaryRedirect)
	}
	data.updateOptions(&settings)

	if settings.AccountID == "" {
		var account models.Account
		if data.Realm != nil {
			account.Realm = types.Realm(*data.Realm)
		}
		if data.Nickname != nil {
			account.Nickname = *data.Nickname
		}
		if data.AccountID != nil {
			account.ID = *data.AccountID
		}
		return widgets.WidgetConfiguratorPage(widget.WidgetWithAccount{WidgetOptions: settings, Account: account}, map[string]string{"account_id": "You need to select a user from the list of search results."}), nil
	}

	account, err := ctx.Database().GetAccountByID(ctx.Context, settings.AccountID)
	if database.IsNotFound(err) {
		account, err = ctx.Fetch().Account(ctx.Context, settings.AccountID)
	}
	if err != nil || account.Private {
		return widgets.NewWidgetPage(widget.WidgetWithAccount{WidgetOptions: settings, Account: account}, map[string]string{"account_id": "You have selected a private account - stats are not available for private accounts."}), nil
	}

	updated, err := ctx.Database().UpdateWidgetSettings(ctx.Context, settings.ID, settings)
	if err != nil {
		return nil, ctx.Err(err, "failed to update widget settings")
	}

	go func(widgetID string) {
		topicID := fmt.Sprintf("widget-settings-%s", widgetID)
		pctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()
		err = ctx.PubSub().Send(pctx, realtime.Message{Topic: topicID, Strategy: realtime.RouteToAll, Data: "reload"})
		if err != nil && !errors.Is(err, realtime.ErrInvalidTopic) && !errors.Is(err, realtime.ErrTopicHasNoListeners) {
			log.Err(err).Str("topic", topicID).Str("widget", widgetID).Msg("failed to update the live widget through pubsub")
		}
	}(settings.ID)

	return widgets.WidgetConfiguratorPage(widget.WidgetWithAccount{WidgetOptions: updated, Account: account}, nil), nil
}

var QuickAction handler.Partial = func(ctx *handler.Context) (templ.Component, error) {
	user, err := ctx.SessionUser()
	if err != nil {
		return nil, ctx.Redirect("/login", http.StatusTemporaryRedirect)
	}

	widgetID := ctx.Path("widgetId")
	if widgetID == "" {
		return nil, ctx.Error("invalid widget id")
	}

	form, err := ctx.FormValues()
	if err != nil {
		return nil, ctx.Err(err, "invalid form data")
	}

	var data widgetSettingsPayload
	data.parse(form)

	settings, err := ctx.Database().GetWidgetSettings(ctx.Context, widgetID)
	if err != nil {
		if database.IsNotFound(err) {
			return nil, ctx.Error("widget not found")
		}
		return nil, ctx.Err(err, "failed to get widget settings")
	}
	if settings.UserID != user.ID {
		return nil, ctx.Redirect("/login", http.StatusTemporaryRedirect)
	}
	data.updateOptions(&settings)

	if settings.AccountID == "" {
		return nil, ctx.Error("widget has no account id set")
	}

	updated, err := ctx.Database().UpdateWidgetSettings(ctx.Context, settings.ID, settings)
	if err != nil {
		return nil, ctx.Err(err, "failed to update widget settings")
	}

	go func(widgetID string) {
		topicID := fmt.Sprintf("widget-settings-%s", widgetID)
		pctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()
		err = ctx.PubSub().Send(pctx, realtime.Message{Topic: topicID, Strategy: realtime.RouteToAll, Data: "reload"})
		if err != nil && !errors.Is(err, realtime.ErrInvalidTopic) && !errors.Is(err, realtime.ErrTopicHasNoListeners) {
			log.Err(err).Str("topic", topicID).Str("widget", widgetID).Msg("failed to update the live widget through pubsub")
		}
	}(settings.ID)

	return widget.QuickSettingButton(ctx.Query("field"), updated), nil
}

var ResetSession handler.Partial = func(ctx *handler.Context) (templ.Component, error) {
	user, err := ctx.SessionUser()
	if err != nil {
		return nil, ctx.Redirect("/login", http.StatusTemporaryRedirect)
	}

	widgetID := ctx.Path("widgetId")
	if widgetID == "" {
		return nil, ctx.Error("invalid widget id")
	}

	settings, err := ctx.Database().GetWidgetSettings(ctx.Context, widgetID)
	if err != nil {
		if database.IsNotFound(err) {
			return nil, ctx.Error("widget not found")
		}
		return nil, ctx.Err(err, "failed to get widget settings")
	}
	if settings.UserID != user.ID {
		return nil, ctx.Redirect("/login", http.StatusTemporaryRedirect)
	}

	if time.Since(settings.SessionFrom) < constants.WidgetSessionResetTimeout {
		return nil, ctx.Error("too many requests")
	}
	if settings.AccountID == "" {
		return nil, ctx.Error("widget has no account id set")
	}

	settings.SessionFrom = time.Now()
	updated, err := ctx.Database().UpdateWidgetSettings(ctx.Context, settings.ID, settings)
	if err != nil {
		return nil, ctx.Err(err, "failed to update widget settings")
	}

	go func(widgetID, accountID string) {
		c, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()

		realm, ok := ctx.Wargaming().RealmFromID(accountID)
		if !ok {
			log.Error().Str("widgetId", widgetID).Str("accountId", accountID).Msg("failed to parse realm")
			return
		}

		_, err = logic.RecordAccountSnapshots(c, ctx.Wargaming(), ctx.Database(), realm, logic.WithReference(accountID, widgetID), logic.WithType(models.SnapshotTypeWidget))
		if err != nil {
			log.Err(err).Str("widgetId", widgetID).Str("accountId", accountID).Msg("failed to refresh a session for widget")
		}

		topicID := fmt.Sprintf("widget-settings-%s", widgetID)
		pctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()
		err = ctx.PubSub().Send(pctx, realtime.Message{Topic: topicID, Strategy: realtime.RouteToAll, Data: "reload"})
		if err != nil && !errors.Is(err, realtime.ErrInvalidTopic) && !errors.Is(err, realtime.ErrTopicHasNoListeners) {
			log.Err(err).Str("topic", topicID).Str("widget", widgetID).Msg("failed to update the live widget through pubsub")
		}
	}(updated.ID, updated.AccountID)

	return widget.SessionResetButton(updated.ID, constants.WidgetSessionResetTimeout.Seconds()), nil
}
