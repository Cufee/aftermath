// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.731
package widget

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"context"
	"fmt"
	"github.com/cufee/aftermath/cmd/frontend/components/widget"
	"github.com/cufee/aftermath/cmd/frontend/handler"
	"github.com/cufee/aftermath/cmd/frontend/layouts"
	"github.com/cufee/aftermath/internal/stats/client/v1"
	"github.com/pkg/errors"
	"golang.org/x/text/language"
	"time"
)

var PersonalWidget handler.Page = func(ctx *handler.Context) (handler.Layout, templ.Component, error) {
	userID := ctx.Path("userId")
	if userID == "" {
		return nil, nil, errors.New("invalid user id id")
	}

	account, err := ctx.Fetch().Account(ctx.Context, userID)
	if err != nil {
		return nil, nil, errors.New("invalid account id")
	}

	cards, _, err := ctx.Client.Stats(language.English).SessionCards(context.Background(), account.ID, time.Now(), client.WithWN8())
	if err != nil {
		return nil, nil, err
	}

	return layouts.Main, personalWidgetPage(widget.Widget(account, cards)), nil
}

func personalWidgetPage(widget templ.Component) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-row justify-center gap-4 flex-wrap min-w-max\"><div class=\"grow flex flex-col max-w-xl gap-2\"><div class=\"flex flex-col items-center justify-center gap-4\"><div class=\"form-control w-full flex gap-2\"><div class=\"flex flex-col bg-base-200 rounded-lg p-4\"><span class=\"text-lg\">Regular Battles</span> <label class=\"label cursor-pointer\"><span class=\"label-text\">Show Overview Card</span> <input type=\"checkbox\" class=\"toggle toggle-secondary\" checked=\"checked\" disabled></label> <label class=\"label cursor-pointer flex flex-col items-start gap-1\"><span class=\"label-text\">Vehicle Cards</span> <input type=\"range\" min=\"0\" max=\"10\" value=\"3\" class=\"range\" step=\"1\" disabled><div class=\"flex w-full justify-between px-2 text-xs\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for i := range 11 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col items-center\"><span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprint(i))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/frontend/routes/widget/personal.templ`, Line: 51, Col: 31}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></label></div><div class=\"flex flex-col bg-base-200 rounded-lg p-4\"><span class=\"text-lg\">Rating Battles</span> <label class=\"label cursor-pointer\"><span class=\"label-text\">Show Overview Card</span> <input type=\"checkbox\" class=\"toggle toggle-secondary\" checked=\"checked\" disabled></label></div></div><button class=\"btn btn-primary\">copy link</button></div></div><div class=\"flex flex-col gap-2\"><div class=\"max-m-max\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = widget.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
