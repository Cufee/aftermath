// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.731
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/cufee/aftermath/cmd/frontend/logic"
	"github.com/cufee/aftermath/internal/constants"
)

func Head(additional ...templ.Component) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<head><meta charset=\"utf-8\"><meta name=\"color-scheme\" content=\"light\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"><link href=\"https://cdn.jsdelivr.net/npm/daisyui@4.11.1/dist/full.min.css\" rel=\"stylesheet\" type=\"text/css\"><script src=\"https://unpkg.com/htmx.org@1.9.12\"></script><script src=\"https://unpkg.com/htmx.org/dist/ext/head-support.js\"></script><script src=\"https://unpkg.com/htmx.org@1.9.12/dist/ext/multi-swap.js\"></script><script src=\"https://cdn.tailwindcss.com\"></script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = logic.EmbedMinifiedScript(tailwindConfig()).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<title>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(constants.FrontendAppName)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/frontend/components/head.templ`, Line: 19, Col: 36}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><link rel=\"icon\" type=\"image/x-icon\" href=\"/assets/favicon.ico\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, component := range additional {
			templ_7745c5c3_Err = component.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</head>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func tailwindConfig() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_tailwindConfig_bf3d`,
		Function: `function __templ_tailwindConfig_bf3d(){tailwind.config = {
		theme: {
			borderRadius: {
				"none": "0px",
				"sm": "10px",
				"md": "15px",
				"lg": "20px",
				"xl": "25px",
				'full': '9999px',
			}
		}
	}
}`,
		Call:       templ.SafeScript(`__templ_tailwindConfig_bf3d`),
		CallInline: templ.SafeScriptInline(`__templ_tailwindConfig_bf3d`),
	}
}
