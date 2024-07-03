// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.696
package widget

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/cufee/aftermath/cmd/frontend/handler"
import "github.com/pkg/errors"
import "golang.org/x/text/language"
import "time"
import "github.com/cufee/aftermath/cmd/frontend/layouts"
import "github.com/cufee/aftermath/cmd/frontend/components/widget"
import "github.com/cufee/aftermath/internal/database/models"
import "github.com/cufee/aftermath/internal/stats/client/v1"
import "slices"

var LiveWidget handler.Page = func(ctx *handler.Context) (handler.Layout, templ.Component, error) {
	accountID := ctx.Req().PathValue("accountId")
	if accountID == "" {
		return nil, nil, errors.New("invalid account id")
	}

	account, err := ctx.Fetch().Account(ctx.Context, accountID)
	if err != nil {
		return nil, nil, errors.New("invalid account id")
	}

	var opts = []client.RequestOption{client.WithWN8()}
	if ref := ctx.Req().URL.Query().Get("ref"); ref != "" {
		opts = append(opts, client.WithReferenceID(ref))
	}
	if t := ctx.Req().URL.Query().Get("type"); t != "" && slices.Contains(models.SnapshotType("").Values(), t) {
		opts = append(opts, client.WithType(models.SnapshotType(t)))
	}

	cards, _, err := ctx.Client.Stats(language.English).SessionCards(context.Background(), account.ID, time.Now(), opts...)
	if err != nil {
		return nil, nil, err
	}

	return layouts.Main, liveWidget(widget.Widget(account, cards, widget.WithAutoReload())), nil
}

func liveWidget(widget templ.Component) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = widget.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>\n    :root { background-color: rgba(0, 0, 0, 0); white-space: nowrap; }\n  </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
