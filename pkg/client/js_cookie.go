package client

import (
	"bytes"
	"context"
	"net/url"
	"text/template"

	"github.com/PuerkitoBio/goquery"
)

var jsCookieTemplates = template.New("")

func init() {
	jsCookieTemplates = template.Must(jsCookieTemplates.New("window").Parse(`
{{- /* */ -}}
const window = {
	location: {
		protocol: "{{.URL.Scheme}}:",
		host: "{{.URL.Host}}",
	},
};
`))
	jsCookieTemplates = template.Must(jsCookieTemplates.New("src").Parse(`
{{- /* */ -}}
{{ template "window" . }}
const document = {
	createElement(tag) {
		return {};
	},
	getElementsByTagName(name) {
		if (name == "head") {
			return [
				{
					appendChild(v) {
						document.ret = v.src;
					},
				},
			];
		}
		return {};
	},
};

{{ .Code }};

document.ret;
`))
}

func jsCookieSrc(ctx context.Context, rawURL string, doc goquery.Document) (src string, err error) {
	var scriptEl = doc.Find("script#_rspj")
	if scriptEl.Length() == 0 {
		return "", nil
	}
	var jsEngine = ContextJavaScriptEngine(ctx)
	var b bytes.Buffer
	u, err := url.Parse(rawURL)
	if err != nil {
		return
	}
	err = jsCookieTemplates.Lookup("src").Execute(&b, struct {
		URL  url.URL
		Code string
	}{*u, scriptEl.Text()})
	if err != nil {
		return
	}
	return jsEngine.Run(ctx, b.String())
}

// func jsCookieValue(ctx context.Context, rawURL string, script string) (src string, err error) {
// 	var scriptEl = doc.Find("script#_rspj")
// 	if scriptEl.Length() == 0 {
// 		return "", nil
// 	}
// 	var jsEngine = ContextJavaScriptEngine(ctx)
// 	var b bytes.Buffer
// 	u, err := url.Parse(rawURL)
// 	if err != nil {
// 		return
// 	}
// 	err = jsCookieTemplates.Execute(&b, struct {
// 		URL  url.URL
// 		Code string
// 	}{*u, scriptEl.Text()})
// 	if err != nil {
// 		return
// 	}
// 	return jsEngine.Run(ctx, b.String())
// }
