// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.731
package widget

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/cufee/aftermath/cmd/frontend/logic"
import "fmt"
import "github.com/cufee/aftermath/internal/database/models"
import prepare "github.com/cufee/aftermath/internal/stats/prepare/session/v1"

type WidgetFlavor string

const (
	// WidgetFlavorTicker WidgetFlavor = "ticker"
	WidgetFlavorDefault WidgetFlavor = "default"
)

type WidgetOption func(*widget)

func WithAutoReload() WidgetOption {
	return func(w *widget) { w.autoReload = true }
}
func WithVehicleLimit(limit int) WidgetOption {
	return func(w *widget) { w.unratedVehiclesLimit = limit }
}
func WithRatingOverview(shown bool) WidgetOption {
	return func(w *widget) { w.showRatingOverview = shown }
}
func WithUnratedOverview(shown bool) WidgetOption {
	return func(w *widget) { w.showUnratedOverview = shown }
}
func WithFlavor(flavor WidgetFlavor) WidgetOption {
	return func(w *widget) { w.flavor = flavor }
}

func Widget(account models.Account, cards prepare.Cards, options ...WidgetOption) templ.Component {
	widget := widget{
		cards:         cards,
		account:       account,
		flavor:        WidgetFlavorDefault,
		autoReload:    false,
		vehicleStyle:  styleOptions{showTitle: true},
		overviewStyle: styleOptions{showLabel: true},

		showRatingOverview:   true,
		showUnratedOverview:  true,
		unratedVehiclesLimit: 3,
	}
	for _, apply := range options {
		apply(&widget)
	}

	return widget.Render()
}

type styleOptions struct {
	showTitle  bool
	showCareer bool
	showLabel  bool
}

type widget struct {
	cards   prepare.Cards
	account models.Account

	vehicleStyle  styleOptions
	overviewStyle styleOptions

	showRatingOverview   bool
	showUnratedOverview  bool
	unratedVehiclesLimit int

	flavor     WidgetFlavor
	autoReload bool
}

func (w widget) Render() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<head><title>Aftermath - ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(w.account.Nickname)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/frontend/components/widget/widget.templ`, Line: 76, Col: 41}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><meta property=\"og:type\" content=\"website\"><meta property=\"og:title\" content=\"Aftermath - Streaming Widget\"><meta property=\"og:image\" content=\"https://amth.one/assets/og-widget.jpg\"><meta property=\"og:description\" content=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("Aftermath streaming widget for %s [%s]", w.account.Nickname, w.account.Realm))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/frontend/components/widget/widget.templ`, Line: 80, Col: 134}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><meta name=\"twitter:card\" content=\"summary_large_image\"><meta name=\"twitter:title\" content=\"Aftermath - Streaming Widget\"><meta name=\"twitter:image:alt\" content=\"Aftermath Streaming Widget\"><meta name=\"twitter:image\" content=\"https://amth.one/assets/og-widget.jpg\"><meta name=\"twitter:description\" content=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("Aftermath streaming widget for %s [%s]", w.account.Nickname, w.account.Realm))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/frontend/components/widget/widget.templ`, Line: 85, Col: 135}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></head><div class=\"text-nowrap whitespace-nowrap min-w-max\" id=\"widget-container\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		switch w.flavor {
		default:
			templ_7745c5c3_Err = w.defaultWidget().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if w.autoReload {
			templ_7745c5c3_Err = logic.EmbedMinifiedScript(widgetRefresh(w.account.Realm, w.account.ID, w.account.LastBattleTime.Unix()), w.account.Realm, w.account.ID, w.account.LastBattleTime.Unix()).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func widgetRefresh(realm string, accountID string, lastBattle int64) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_widgetRefresh_731f`,
		Function: `function __templ_widgetRefresh_731f(realm, accountID, lastBattle){let apiHost = ""
  switch (realm.toLowerCase()) {
    case 'na':
      apiHost = "wotblitz.com"
      break;
    case 'eu':
      apiHost = "wotblitz.eu"
      break;
    case 'as':
      apiHost = "wotblitz.asia"
      break;
    default:
      throw new Error("Unknown realm: " + realm)
  }
  const refresh = () => {
    fetch(` + "`" + `https://api.${apiHost}/wotb/account/info/?application_id=f44aa6f863c9327c63ba26be3db0d07f&account_id=${accountID}&fields=last_battle_time` + "`" + `)
      .then(response => response.json())
      .then(data => {
        if (data.data[accountID.toString()].last_battle_time > lastBattle) {
          location.reload()
        }
      })
      .catch(error => { console.error(error); setTimeout(()=>location.reload(), 10000) })
  }
  setInterval(refresh, 5000)
}`,
		Call:       templ.SafeScript(`__templ_widgetRefresh_731f`, realm, accountID, lastBattle),
		CallInline: templ.SafeScriptInline(`__templ_widgetRefresh_731f`, realm, accountID, lastBattle),
	}
}
