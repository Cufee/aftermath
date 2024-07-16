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
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/stats/client/v1"
	"golang.org/x/text/language"
	"net/http"
	"slices"
	"strconv"
	"time"
)

var AccountWidget handler.Partial = func(ctx *handler.Context) (templ.Component, error) {
	accountID := ctx.Path("accountId")
	if accountID == "" {
		return nil, ctx.Redirect("/widget", http.StatusTemporaryRedirect)
	}

	account, err := ctx.Fetch().Account(ctx.Context, accountID)
	if err != nil {
		return nil, ctx.Redirect("/widget", http.StatusTemporaryRedirect)
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
		return nil, err
	}

	var wopts []widget.WidgetOption
	if v, err := strconv.Atoi(ctx.Query("vl")); err == nil && v >= 0 && v <= 10 {
		wopts = append(wopts, widget.WithVehicleLimit(int(v)))
	}
	if v := ctx.Query("or"); v != "" {
		wopts = append(wopts, widget.WithRatingOverview(v == "1"))
	}
	if v := ctx.Query("ou"); v != "" {
		wopts = append(wopts, widget.WithUnratedOverview(v == "1"))
	}

	return accountWidget(widget.Widget(account, cards, wopts...)), nil
}

func accountWidget(widget templ.Component) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"account-widget\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = widget.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
