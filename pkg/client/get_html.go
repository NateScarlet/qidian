package client

import (
	"context"
	"io"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type transport struct {
}

// var jsProtectTemplate = template.Must(template.New("").Parse(`
// {{- /* */ -}}

// const vm = require('vm');

// const code = {{ . | json }};

// (async () => {
// 	console.log(await vm.runInNewContext(code));
// })().catch((err) => {
// 	console.error(err);
// 	process.exit(1);
// })
// `))

func handleJSProtect(ctx context.Context, doc *goquery.Document) (ret *goquery.Document, err error) {
	var scriptEl = doc.Find("script#_rspj").Get(0)
	if scriptEl == nil {
		return doc, nil
	}
	var jsEngine = ContextJavaScriptEngine(ctx)
	var script = scriptEl.Data
	jsEngine.Run(ctx, ``+script+`
	
	`)
	return

}

func GetHTML(ctx context.Context, url string, body io.Reader) (doc *goquery.Document, err error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return
	}
	resp, err := For(ctx).Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	doc, err = goquery.NewDocumentFromReader(resp.Request.Body)
	if err != nil {
		return
	}

	return
}
