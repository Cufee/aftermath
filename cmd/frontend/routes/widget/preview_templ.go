// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package widget

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/cufee/aftermath/cmd/frontend/components"
	cwidget "github.com/cufee/aftermath/cmd/frontend/components/widget"
	"github.com/cufee/aftermath/cmd/frontend/handler"
	"github.com/cufee/aftermath/cmd/frontend/layouts"
	"github.com/cufee/aftermath/cmd/frontend/routes/api/widget"
	"strconv"
)

var WidgetPreview handler.Page = func(ctx *handler.Context) (handler.Layout, templ.Component, error) {
	widget, err := widget.AccountWidget(ctx)
	if err != nil {
		return layouts.Main, nil, ctx.Error(err, "failed to generate a widget preview")
	}
	if widget == nil {
		return nil, nil, nil
	}

	var withUnrated = ctx.Query("ou") != "0"
	var withRating = ctx.Query("or") != "0"
	var vehicles int = 3
	if v, err := strconv.Atoi(ctx.Query("vl")); err == nil && v >= 0 && v <= 10 {
		vehicles = v
	}

	return layouts.Main, widgetPreview(ctx.Path("accountId"), widget, withRating, withUnrated, vehicles), nil
}

func widgetPreview(accountID string, widget templ.Component, or, ou bool, vl int) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-row flex-wrap gap-4\"><div class=\"flex flex-col gap-4 basis-1/2 grow\"><div class=\"flex flex-col gap-2 text-center\"><div class=\"text-3xl font-bold\">Aftermath Streaming Widget</div><p>Level up your stream with a real-time stats widget!</p></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = cwidget.Settings(handlePreview(accountID), or, ou, vl).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, copyButtonAction())
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button type=\"button\" id=\"copy-widget-link\" class=\"btn btn-primary w-full transition-all duration-250 ease-in-out\" onclick=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 templ.ComponentScript = copyButtonAction()
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var2.Call)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">Copy OBS Link</button></div><div class=\"flex flex-col items-center justify-center grow overflow-auto\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var3 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"min-h-96 w-full overflow-auto p-2 relative\"><div class=\"min-w-[400px] w-max mx-auto h-1/2\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = widget.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = components.OBSMockup("/assets/widget-background.jpg").Render(templ.WithChildren(ctx, templ_7745c5c3_Var3), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func copyButtonAction() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_copyButtonAction_e532`,
		Function: `function __templ_copyButtonAction_e532(){const url = window.location.protocol + "//" + window.location.host + window.location.pathname + "live" + window.location.search
	navigator.clipboard.writeText(url);
	
	const btn = document.getElementById("copy-widget-link")
	const oldText = btn.textContent
	btn.textContent = "Copied to Clipboard!";
	btn.classList.add("btn-success");
	setTimeout(()=> {
		btn.textContent = oldText;
		btn.classList.remove("btn-success");
	}, 2000)
}`,
		Call:       templ.SafeScript(`__templ_copyButtonAction_e532`),
		CallInline: templ.SafeScriptInline(`__templ_copyButtonAction_e532`),
	}
}

func handlePreview(id string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_handlePreview_5e47`,
		Function: `function __templ_handlePreview_5e47(id){const ouEl = document.getElementById("widget-settings-ou")
	const orEl = document.getElementById("widget-settings-or")
	const vlEl = document.getElementById("widget-settings-vl")
	const button = document.getElementById("copy-widget-link")
	
	const ou = ouEl.checked ? "1" : "0"
	const or = orEl.checked ? "1" : "0"
	const vl = vlEl.value
	const newQuery = ` + "`" + `?or=${or}&ou=${ou}&vl=${vl}` + "`" + `
	if (newQuery != window.location.search) {
		ouEl.disabled = true
		orEl.disabled = true
		vlEl.disabled = true
		button.disabled = true

		fetch("/api/p/widget/"+id+newQuery).then((r) => r.text()).then((html) => {
			document.getElementById("account-widget").outerHTML = html
			const url = window.location.protocol + "//" + window.location.host + window.location.pathname + newQuery;
			window.history?.pushState({path:url},'',url);
		}).catch(e => console.log(e)).finally(() => {
			ouEl.disabled = false
			orEl.disabled = false
			vlEl.disabled = false
			button.disabled = false
		})
	}
}`,
		Call:       templ.SafeScript(`__templ_handlePreview_5e47`, id),
		CallInline: templ.SafeScriptInline(`__templ_handlePreview_5e47`, id),
	}
}
