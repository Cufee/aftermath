// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package widget

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"context"
	"github.com/cufee/aftermath/cmd/frontend/components/widget"
	"github.com/cufee/aftermath/cmd/frontend/handler"
	"github.com/cufee/aftermath/cmd/frontend/layouts"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/stats/client/v1"
	"github.com/pkg/errors"
	"golang.org/x/text/language"
	"slices"
	"strconv"
	"time"
)

var PersonalLiveWidget handler.Page = func(ctx *handler.Context) (handler.Layout, templ.Component, error) {
	userID := ctx.Path("userId")
	if userID == "" {
		return nil, nil, errors.New("invalid account id")
	}

	account, err := ctx.Fetch().Account(ctx.Context, userID)
	if err != nil {
		return nil, nil, errors.New("invalid account id")
	}

	var opts = []client.RequestOption{client.WithWN8()}
	if ref := ctx.Query("ref"); ref != "" {
		opts = append(opts, client.WithReferenceID(ref))
	}
	if t := ctx.Query("type"); t != "" && slices.Contains(models.SnapshotType("").Values(), t) {
		opts = append(opts, client.WithType(models.SnapshotType(t)))
	}

	cards, _, err := ctx.Client.Stats(language.English).SessionCards(context.Background(), account.ID, time.Now(), opts...)
	if err != nil {
		return nil, nil, err
	}

	var wopts = []widget.WidgetOption{widget.WithAutoReload()}
	if v, err := strconv.Atoi(ctx.Query("vl")); err == nil && v >= 0 && v <= 10 {
		wopts = append(wopts, widget.WithVehicleLimit(int(v)))
	}
	if v := ctx.Query("or"); v != "" {
		wopts = append(wopts, widget.WithRatingOverview(v == "1"))
	}
	if v := ctx.Query("ou"); v != "" {
		wopts = append(wopts, widget.WithUnratedOverview(v == "1"))
	}

	return layouts.Main, liveWidget(widget.Widget(account, cards, wopts...)), nil
}

var LiveWidget handler.Page = func(ctx *handler.Context) (handler.Layout, templ.Component, error) {
	accountID := ctx.Path("accountId")
	if accountID == "" {
		return nil, nil, errors.New("invalid account id")
	}

	account, err := ctx.Fetch().Account(ctx.Context, accountID)
	if err != nil {
		return nil, nil, errors.New("invalid account id")
	}

	var opts = []client.RequestOption{client.WithWN8()}
	if ref := ctx.Query("ref"); ref != "" {
		opts = append(opts, client.WithReferenceID(ref))
	}
	if t := ctx.Query("type"); t != "" && slices.Contains(models.SnapshotType("").Values(), t) {
		opts = append(opts, client.WithType(models.SnapshotType(t)))
	}

	cards, _, err := ctx.Client.Stats(language.English).SessionCards(context.Background(), account.ID, time.Now(), opts...)
	if err != nil {
		return nil, nil, err
	}

	var wopts = []widget.WidgetOption{widget.WithAutoReload()}
	if v, err := strconv.Atoi(ctx.Query("vl")); err == nil && v >= 0 && v <= 10 {
		wopts = append(wopts, widget.WithVehicleLimit(int(v)))
	}
	if v := ctx.Query("or"); v != "" {
		wopts = append(wopts, widget.WithRatingOverview(v == "1"))
	}
	if v := ctx.Query("ou"); v != "" {
		wopts = append(wopts, widget.WithUnratedOverview(v == "1"))
	}

	return layouts.StyleOnly, liveWidget(widget.Widget(account, cards, wopts...)), nil
}

func liveWidget(widget templ.Component) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"min-w-max\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = widget.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><style>\n    :root { background-color: rgba(0, 0, 0, 0); white-space: nowrap; }\n  </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
