// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.731
package redirect

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/cufee/aftermath/cmd/frontend/handler"
	"github.com/cufee/aftermath/cmd/frontend/logic"
	"github.com/cufee/aftermath/internal/constants"
	"net/http"
	"slices"
	"strings"
)

var VerifyFromDiscord handler.Partial = func(ctx *handler.Context) (templ.Component, error) {
	realm := strings.ToLower(ctx.Path("realm"))
	if !slices.Contains([]string{"na", "eu", "as"}, realm) {
		return nil, ctx.Redirect("/", http.StatusMovedPermanently)
	}
	return discordVerifyRedirect(realm), nil
}

func discordVerifyRedirect(realm string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<head><title>Aftermath 🔗 Wargaming - ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(strings.ToUpper(realm))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/frontend/routes/redirect/discord_verify.templ`, Line: 22, Col: 60}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><meta property=\"og:type\" content=\"website\"><meta property=\"og:title\" content=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs("Aftermath 🔗 Wargaming -" + strings.ToUpper(realm))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/frontend/routes/redirect/discord_verify.templ`, Line: 24, Col: 91}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><meta property=\"og:image\" content=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(constants.FrontendURL + "/assets/og-verify.jpg")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/frontend/routes/redirect/discord_verify.templ`, Line: 25, Col: 85}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><meta property=\"og:description\" content=\"Link your Blitz account to Aftermath!\"><meta name=\"twitter:card\" content=\"summary_large_image\"><meta name=\"twitter:title\" content=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs("Aftermath 🔗 Wargaming -" + strings.ToUpper(realm))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/frontend/routes/redirect/discord_verify.templ`, Line: 28, Col: 92}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><meta name=\"twitter:image:alt\" content=\"Link your Blitz account to Aftermath\"><meta name=\"twitter:image\" content=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(constants.FrontendURL + "/assets/og-verify.jpg")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/frontend/routes/redirect/discord_verify.templ`, Line: 30, Col: 86}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><meta name=\"twitter:description\" content=\"Link your Blitz account to Aftermath!\"></head>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = logic.EmbedScript(discordVerifyRedirectScript(realm)).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func discordVerifyRedirectScript(realm string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_discordVerifyRedirectScript_fac0`,
		Function: `function __templ_discordVerifyRedirectScript_fac0(realm){window.location.href="/login?from=/api/auth/wargaming/login/"+realm;
}`,
		Call:       templ.SafeScript(`__templ_discordVerifyRedirectScript_fac0`, realm),
		CallInline: templ.SafeScriptInline(`__templ_discordVerifyRedirectScript_fac0`, realm),
	}
}
