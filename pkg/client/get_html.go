package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var getHTMLMu sync.Mutex
var CaptchaDelay = 1 * time.Minute

func rawGetHTML(ctx context.Context, url string, options ...GetHTMLOption) (res GetHTMLResult, err error) {
	var opts = newGetHTMLOptions(options...)
	req, err := newRequest(ctx, "GET", url, nil)
	if err != nil {
		return
	}
	if opts.visitRequest != nil {
		opts.visitRequest(req)
	}
	resp, err := For(ctx).Do(req)
	if err != nil {
		return
	}
	return makeGetHTMLResult(ctx, resp, func() (_ GetHTMLResult, err error) {
		return rawGetHTML(ctx, url, options...)
	})
}

func GetHTML(ctx context.Context, url string, options ...GetHTMLOption) (res GetHTMLResult, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("qidian: client.GetHTML('%s'): %w", url, err)
		}
	}()

	getHTMLMu.Lock()
	defer getHTMLMu.Unlock()

	return rawGetHTML(ctx, url, options...)
}

func makeGetHTMLResult(ctx context.Context, resp *http.Response, retry func() (_ GetHTMLResult, err error)) (res GetHTMLResult, err error) {
	res.retry = retry
	err = func() (err error) {
		defer resp.Body.Close()
		res.response = resp
		res.body, err = io.ReadAll(resp.Body)
		if err != nil {
			return
		}
		if resp.StatusCode >= 400 {
			err = fmt.Errorf("response status %d\n\n%s", resp.StatusCode, res.body)
			return
		}
		return
	}()
	if err != nil {
		return
	}
	err = handleStatusForbidden(ctx, &res)
	if err != nil {
		return
	}
	err = handleCaptcha(ctx, &res)
	if err != nil {
		return
	}
	err = handleJSProtect(ctx, &res)
	if err != nil {
		return
	}
	err = handleAccessDeny(ctx, &res)
	if err != nil {
		return
	}
	return
}

func handleJSProtect(ctx context.Context, res *GetHTMLResult) (err error) {
	if !bytes.Contains(res.Body(), []byte("_rspj")) {
		return
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(res.Body()))
	if err != nil {
		return
	}
	data, err := newJSCookieTemplateData(ctx, res.Request().URL.String(), *doc)
	if err != nil {
		return
	}
	if data == nil {
		return nil
	}
	var c = For(ctx)
	if c.Jar == nil {
		err = fmt.Errorf("nil cookie jar")
		return
	}
	cookie, err := jsCookie(ctx, *data)
	if err != nil {
		return
	}
	c.Jar.SetCookies(res.Request().URL, []*http.Cookie{cookie})
	*res, err = res.retry()
	return
}

func handleStatusForbidden(ctx context.Context, res *GetHTMLResult) (err error) {
	if res.response.StatusCode != http.StatusForbidden {
		return
	}
	time.Sleep(CaptchaDelay)
	*res, err = res.retry()
	return
}

func handleCaptcha(ctx context.Context, res *GetHTMLResult) (err error) {
	if !bytes.Contains(res.Body(), []byte("/TCaptcha.js\"")) {
		return
	}
	if !bytes.Contains(res.Body(), []byte("<body></body>")) && !bytes.HasPrefix(res.Body(), []byte("<script>")) {
		return
	}
	time.Sleep(CaptchaDelay)
	*res, err = res.retry()
	return
}

func handleAccessDeny(ctx context.Context, res *GetHTMLResult) (err error) {
	if !bytes.Contains(res.Body(), []byte("<title>AccessDeny</title>")) {
		return
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(res.Body()))
	if err != nil {
		return
	}
	if v := doc.Find("h1").Text(); v != "" {
		return fmt.Errorf("access deny: %s", v)
	}
	return
}

type GetHTMLOptions struct {
	visitRequest func(req *http.Request)
}

func newGetHTMLOptions(options ...GetHTMLOption) *GetHTMLOptions {
	var opts = new(GetHTMLOptions)
	for _, i := range options {
		i(opts)
	}
	return opts
}

type GetHTMLOption func(opts *GetHTMLOptions)

type GetHTMLResult struct {
	response *http.Response
	body     []byte
	retry    func() (_ GetHTMLResult, err error)
}

func (obj GetHTMLResult) Request() *http.Request {
	if obj.response == nil {
		return nil
	}
	return obj.response.Request
}

func (obj GetHTMLResult) Response() *http.Response {
	return obj.response
}

func (obj GetHTMLResult) Body() []byte {
	return obj.body
}

func GetHTMLOptionVisitRequest(visitor func(req *http.Request)) GetHTMLOption {
	return func(opts *GetHTMLOptions) {
		opts.visitRequest = visitor
	}
}
