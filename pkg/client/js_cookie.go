package client

import (
	"bytes"
	"context"
	_ "embed"
	"net/url"
	"text/template"

	"github.com/PuerkitoBio/goquery"
)

//go:embed js_cookie.env.cjs
var jsCookieEnvJS string

var jsCookieTemplates = template.New("").Funcs(template.FuncMap{
	"backtickQuote": func(s string) string {
		return "`" + s + "`"
	},
})

func init() {

	jsCookieTemplates = template.Must(jsCookieTemplates.New("env").Parse(jsCookieEnvJS))
	jsCookieTemplates = template.Must(jsCookieTemplates.New("src").Parse(`
{{- template "env" . }}

{{ .Code }};

document.getElementsByTagName('head')[0].children[0].src;
`))

	jsCookieTemplates = template.Must(jsCookieTemplates.New("TODO").Parse(`
{{- template "env" . }}

{{ .Script1 }};
{{ .Script2 }};

window.onload();
[document.head, document, window];
`))

	jsCookieTemplates = template.Must(jsCookieTemplates.New("TODO2").Parse(`
{{- template "env" . }}

[document, window];
`))
}

func jsCookieSrc(ctx context.Context, rawURL string, doc goquery.Document) (script, src string, err error) {
	var scriptEl = doc.Find("script#_rspj")
	if scriptEl.Length() == 0 {
		return
	}
	script = scriptEl.Text()
	var jsEngine = ContextJavaScriptEngine(ctx)
	var b bytes.Buffer
	u, err := url.Parse(rawURL)
	if err != nil {
		return
	}
	err = jsCookieTemplates.Lookup("src").Execute(&b, struct {
		URL  *url.URL
		Code string
	}{u, script})
	if err != nil {
		return
	}
	src, err = jsEngine.Run(ctx, b.String())
	return
}

func jsCookieTODO(ctx context.Context, rawURL string, script1, script2 string) (script3 string, err error) {
	var jsEngine = ContextJavaScriptEngine(ctx)
	var b bytes.Buffer
	u, err := url.Parse(rawURL)
	if err != nil {
		return
	}
	err = jsCookieTemplates.Lookup("TODO").Execute(&b, struct {
		URL     *url.URL
		Script1 string
		Script2 string
	}{u, script1, script2})
	if err != nil {
		return
	}
	return jsEngine.Run(ctx, b.String())
}

func jsCookieTODO2(ctx context.Context, rawURL string, script1, script2, script3 string) (_ string, err error) {
	var jsEngine = ContextJavaScriptEngine(ctx)
	var b bytes.Buffer
	u, err := url.Parse(rawURL)
	if err != nil {
		return
	}
	err = jsCookieTemplates.Lookup("TODO2").Execute(&b, struct {
		URL     *url.URL
		Script1 string
		Script2 string
		Script3 string
	}{u, script1, script2, script3})
	if err != nil {
		return
	}
	return jsEngine.Run(ctx, b.String())
}
