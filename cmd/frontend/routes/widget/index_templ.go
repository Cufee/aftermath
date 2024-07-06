// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.731
package widget

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/cufee/aftermath/cmd/frontend/components"
	cwidget "github.com/cufee/aftermath/cmd/frontend/components/widget"
	"github.com/cufee/aftermath/cmd/frontend/handler"
	"github.com/cufee/aftermath/cmd/frontend/layouts"
	"github.com/cufee/aftermath/cmd/frontend/logic"
	"github.com/cufee/aftermath/cmd/frontend/routes/api/widget"
	"github.com/cufee/aftermath/internal/constants"
	"strconv"
)

var WidgetHome handler.Page = func(ctx *handler.Context) (handler.Layout, templ.Component, error) {
	widget, err := widget.MockWidget(ctx)
	if err != nil {
		return layouts.Main, nil, ctx.Error(err, "failed to generate a widget preview")
	}

	var withUnrated = ctx.Query("ou") != "0"
	var withRating = ctx.Query("or") != "0"
	var vehicles int = 3
	if v, err := strconv.Atoi(ctx.Query("vl")); err == nil && v >= 0 && v <= 10 {
		vehicles = v
	}

	return layouts.Main, widgetHome(widget, withRating, withUnrated, vehicles), nil
}

func widgetHome(widget templ.Component, or, ou bool, vl int) templ.Component {
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
		templ_7745c5c3_Err = cwidget.Settings(handlePreviewOnHome(), or, ou, vl).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-row flex-wrap gap-2\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, onRealmSelect())
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<select id=\"player-search-realm\" class=\"select select-bordered grow\" onchange=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 templ.ComponentScript = onRealmSelect()
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var2.Call)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><option disabled selected class=\"text-center\">Select a Server</option> <option value=\"NA\">North America</option> <option value=\"EU\">Europe</option> <option value=\"AS\">Asia</option></select><div class=\"dropdown dropdown-top grow\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, onNicknameInput())
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input id=\"player-search-nickname\" type=\"text\" placeholder=\"Nickname\" class=\"join-item input input-bordered placeholder:text-center w-full\" disabled oninput=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 templ.ComponentScript = onNicknameInput()
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var3.Call)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><ul hx-boost=\"true\" id=\"player-search-results\" tabindex=\"0\" class=\"dropdown-content z-[1] menu p-2 shadow bg-base-300 rounded-t-box flex-nowrap overflow-auto w-full\"><span class=\"text-xs text-center cursor-default\">Start typing to search</span></ul></div></div></div><div class=\"flex flex-col items-center justify-center grow overflow-auto\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var4 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
		templ_7745c5c3_Err = components.OBSMockup("/assets/widget-background.jpg").Render(templ.WithChildren(ctx, templ_7745c5c3_Var4), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = logic.EmbedMinifiedScript(searchEventHandler(constants.WargamingPublicAppID), constants.WargamingPublicAppID).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func onRealmSelect() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_onRealmSelect_da95`,
		Function: `function __templ_onRealmSelect_da95(){const results = document.getElementById("player-search-results");
	const select = document.getElementById("player-search-realm");
	const realm = select.value;
	// enable/disable field
	const nickname = document.getElementById("player-search-nickname");
	if (realm) {
		if (nickname.value.length >= 5) {
			// if the input already exists - send a request
			results.innerHTML = '<li><span class="loading loading-dots loading-xs m-auto"></span></li>';
			const event = new CustomEvent("player-search", { detail: {query: nickname.value, realm} });
			document.dispatchEvent(event);
			return
		}
		nickname.disabled = false;
	} else {
		nickname.disabled = true;
	}
	// clear results
	results.innerHTML = '<span class="text-xs text-center cursor-default">Start typing to search</span>';
}`,
		Call:       templ.SafeScript(`__templ_onRealmSelect_da95`),
		CallInline: templ.SafeScriptInline(`__templ_onRealmSelect_da95`),
	}
}

func onNicknameInput() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_onNicknameInput_042c`,
		Function: `function __templ_onNicknameInput_042c(){const nickname = document.getElementById("player-search-nickname");
	const results = document.getElementById("player-search-results");
	const select = document.getElementById("player-search-realm");
	const realm = select.value;
	if (!realm) {
		results.innerHTML = '<span class="text-xs text-center cursor-default">Start typing to search</span>';
		nickname.disabled = true;
		nickname.value = '';
		return;
	}
	if (!nickname.value) {
		results.innerHTML = '<span class="text-xs text-center cursor-default">Start typing to search</span>';
		return;
	}
	if (nickname.value.length < 5) {
		results.innerHTML = '<span class="text-xs text-center cursor-default">Continue typing to search</span>';
		return;
	}
	const event = new CustomEvent("player-search", { detail: {query: nickname.value, realm} });
	document.dispatchEvent(event);
}`,
		Call:       templ.SafeScript(`__templ_onNicknameInput_042c`),
		CallInline: templ.SafeScriptInline(`__templ_onNicknameInput_042c`),
	}
}

func searchEventHandler(appId string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_searchEventHandler_3147`,
		Function: `function __templ_searchEventHandler_3147(appId){const results = document.getElementById("player-search-results");
	const nickname = document.getElementById("player-search-nickname");

	const setResultsLoading = () => {
		results.innerHTML = '<li><span class="loading loading-dots loading-xs m-auto"></span></li>';
	}
	window.setResultsLoading = setResultsLoading

	const getApiUrl = (realm, query) => {
		let baseUrl = "";
		switch (realm) {
			case "NA":
				baseUrl = "https://api.wotblitz.com/wotb";
				break;
			case "EU":
				baseUrl = "https://api.wotblitz.eu/wotb";
				break;
			case "AS":
				baseUrl = "https://api.wotblitz.asia/wotb";
				break;
			default:
				return "";
		}
		return baseUrl + "/account/list/" + "?application_id=" + appId + "&limit=5" + "&search=" + query;
	}

	let debounceTimeout;
	const debounceDelay = 500;

	document.addEventListener(
		"player-search",
		(evt) => {
			clearTimeout(debounceTimeout);
			debounceTimeout = setTimeout(() => {
				setResultsLoading()
				
				const { query, realm } = evt.detail;
				const url = getApiUrl(realm, query);
				if (!url) return;

				fetch(url).then(res => res.json()).then(data => {
					if (data.status != "ok") {
						results.innerHTML = ` + "`" + `<span class="text-xs text-center cursor-default">${data?.error?.message.toLocaleLowerCase().replaceAll("_", " ") || "Failed to get accounts"}</span>` + "`" + `;
						nickname.disabled = false;
						return;
					}
					const elements = []
					for (const account of data.data?.reverse() || []) {
						if (!account.account_id || !account.nickname) continue;
						elements.push(` + "`" + `<li><a onclick="window.setResultsLoading();htmx.trigger('body','htmx:beforeSend');" href="/widget/${account.account_id}${window.location.search}">${account.nickname}</a></li>` + "`" + `);
					}
					if (elements.length == 0) {
						results.innerHTML = '<span class="text-xs text-center cursor-default">No players found</span>';
						nickname.disabled = false;
						return;
					}
					results.innerHTML = elements.join("");
					return;
				}).catch(e => {
					console.log("failed to search for accounts",e);
					results.innerHTML = '<span class="text-xs text-center cursor-default">No players found</span>';
					nickname.disabled = false;
				});
			}, debounceDelay);
		},
		false,
	);
}`,
		Call:       templ.SafeScript(`__templ_searchEventHandler_3147`, appId),
		CallInline: templ.SafeScriptInline(`__templ_searchEventHandler_3147`, appId),
	}
}

func handlePreviewOnHome() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_handlePreviewOnHome_2c46`,
		Function: `function __templ_handlePreviewOnHome_2c46(){const ouEl = document.getElementById("widget-settings-ou");
	const orEl = document.getElementById("widget-settings-or");
	const vlEl = document.getElementById("widget-settings-vl");

	const ou = ouEl.checked ? "1" : "0";
	const or = orEl.checked ? "1" : "0";
	const vl = vlEl.value;
	const newQuery = ` + "`" + `?or=${or}&ou=${ou}&vl=${vl}` + "`" + `;
	if (newQuery != window.location.search) {
		ouEl.disabled = true;
		orEl.disabled = true;
		vlEl.disabled = true;
		fetch("/api/widget/mock"+newQuery).then((r) => r.text()).then((html) => {
			document.getElementById("mock-widget").outerHTML = html;
			const url = window.location.protocol + "//" + window.location.host + window.location.pathname + newQuery;
			window.history?.pushState({path:url},'',url);
		}).catch(e => console.log(e)).finally(() => {
			setTimeout(() => {
				ouEl.disabled = false;
				orEl.disabled = false;
				vlEl.disabled = false;
			}, 500);
		});
	}
}`,
		Call:       templ.SafeScript(`__templ_handlePreviewOnHome_2c46`),
		CallInline: templ.SafeScriptInline(`__templ_handlePreviewOnHome_2c46`),
	}
}
