package client

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"text/template"

	"github.com/PuerkitoBio/goquery"
)

//go:embed js_cookie.env.cjs
var jsCookieEnvJS string

type jsCookieTemplateData struct {
	URL              *url.URL
	Script1          string
	Script1Attrs     map[string]string
	Script2          string
	HiddenInputName  string
	HiddenInputValue string
	UserAgent        string
}

var jsCookieTemplates = template.New("").Funcs(template.FuncMap{
	"toJSON": func(s any) (_ string, err error) {
		data, err := json.Marshal(s)
		if err != nil {
			return
		}
		return string(data), nil
	},
})

func init() {

	jsCookieTemplates = template.Must(jsCookieTemplates.New("env").Parse(jsCookieEnvJS))
	jsCookieTemplates = template.Must(jsCookieTemplates.New("src").Parse(`
{{- template "env" . }}

{{ .Script1 }};

document.getElementsByTagName('head')[0].children[0].src;
`))

	jsCookieTemplates = template.Must(jsCookieTemplates.New("value").Parse(`
{{- template "env" . }}

{{ .Script1 }};
{{ .Script2 }};

document.cookie;
`))
}

func newJSCookieTemplateData(ctx context.Context, rawURL string, doc goquery.Document) (data *jsCookieTemplateData, err error) {
	// spell-checker: word _rspj
	data = new(jsCookieTemplateData)
	data.UserAgent = ContextUserAgent(ctx)
	var scriptEl = doc.Find("script#_rspj")
	if scriptEl.Length() == 0 {
		return
	}
	data.URL, err = url.Parse(rawURL)
	if err != nil {
		return
	}
	data.Script1 = scriptEl.Text()
	data.Script1Attrs = make(map[string]string, len(scriptEl.Nodes[0].Attr))
	for _, i := range scriptEl.Nodes[0].Attr {
		data.Script1Attrs[i.Key] = i.Val
	}

	var hiddenInput = doc.Find("#__onload__")
	if hiddenInput.Length() == 0 {
		err = fmt.Errorf("missing '#__onload__' element")
		return
	}
	var ok bool
	data.HiddenInputName, ok = hiddenInput.Attr("name")
	if !ok {
		err = fmt.Errorf("missing '#__onload__' name attr")
		return
	}
	data.HiddenInputValue, ok = hiddenInput.Attr("value")
	if !ok {
		err = fmt.Errorf("missing '#__onload__' value attr")
		return
	}
	return
}

func jsCookieSrc(ctx context.Context, data jsCookieTemplateData) (src string, err error) {
	var jsEngine = ContextJavaScriptEngine(ctx)
	var b bytes.Buffer
	err = jsCookieTemplates.Lookup("src").Execute(&b, data)
	if err != nil {
		return
	}
	src, err = jsEngine.Run(ctx, b.String())
	return
}

func jsCookieValue(ctx context.Context, data jsCookieTemplateData) (cookie string, err error) {
	var jsEngine = ContextJavaScriptEngine(ctx)
	var b bytes.Buffer
	err = jsCookieTemplates.Lookup("value").Execute(&b, data)
	if err != nil {
		return
	}
	return jsEngine.Run(ctx, b.String())
}

func jsCookie(ctx context.Context, data jsCookieTemplateData) (cookie []*http.Cookie, err error) {
	src, err := jsCookieSrc(ctx, data)
	if err != nil {
		return
	}
	resp, err := For(ctx).Get(src)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	script2Data, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	data.Script2 = string(script2Data)
	value, err := jsCookieValue(ctx, data)
	if err != nil {
		return
	}
	return parseCookies(value), nil

}
